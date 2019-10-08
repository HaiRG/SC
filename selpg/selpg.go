package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"bufio"
	"github.com/spf13/pflag"
)
const INT_MAX = int(^uint(0) >> 1)
type selpg_args struct{
	start_page int
	end_page int
	in_filename string
    page_len int
    page_type string
    print_dest string
} 
var proname string
func main(){
	proname = os.Args[0]
	var sa selpg_args
	pflag.IntVarP(&sa.start_page,"start_page","s",-1,"start_page")
	pflag.IntVarP(&sa.end_page,"end_page","e",-1,"end_page")
	pflag.IntVarP(&sa.page_len,"page_len","l",72,"page_len")
	pflag.StringVarP(&sa.page_type,"page_type","f","l","page_type")
	pflag.StringVarP(&sa.print_dest,"print_dest","d","","print_dest")
	
	pflag.Parse()
	file_name := pflag.Args()
	if len(file_name) > 0{
		_, err := os.Stat(file_name[0])
		if err != nil{
			panic(err)
		}
		sa.in_filename =  file_name[0]
	}else{
		sa.in_filename = ""
	}
	process_args(&sa)
	process_input(&sa)
}
func process_args(sa *selpg_args){
	if len(os.Args) < 3{
		fmt.Fprintf(os.Stderr,"%s: not enough arguments\n",proname)
		os.Exit(0)
	}
	if os.Args[1] != "-s" {
		fmt.Fprintf(os.Stderr,"%s: first arg should be -sstartpage\n",proname)
		os.Exit(0)
	}
	if sa.start_page == -1 {
		fmt.Fprintf(os.Stderr,"%s: you shoule input start_page\n",proname)
		os.Exit(0)
	}
	if sa.start_page < 1 || sa.start_page > (INT_MAX - 1) {
		fmt.Fprintf(os.Stderr,"%s: invalid start_page\n",proname)
		os.Exit(0)
	}
	if os.Args[3] != "-e"{
		fmt.Fprintf(os.Stderr,"%s: 2nd arg should be -eendpage",proname)
		os.Exit(0)
	} 
	if sa.end_page == -1 {
		fmt.Fprintf(os.Stderr,"%s: you shoule input end_page\n",proname)
		os.Exit(0)
	}

	if sa.end_page < 1 || sa.end_page > (INT_MAX - 1) || sa.end_page < sa.start_page{
		fmt.Fprintf(os.Stderr,"%s: invalid end_page\n",proname)
		os.Exit(0)
	} 
	if sa.page_type == "l" && (sa.page_len < 1 || sa.page_len > (INT_MAX - 1)){
		fmt.Fprintf(os.Stderr,"%s: invalid page_len\n",proname)
		os.Exit(0)
	}

}
func process_input(sa *selpg_args){
	var fin *os.File
	var fout io.WriteCloser
	is_read := false
	is_read2 := false
	if sa.in_filename == ""{
		fin = os.Stdin
	}else{
		var err error
		fin,err = os.Open(sa.in_filename)
		if err != nil{
			panic(err)
		}
		is_read = true; 
	}
	buf := bufio.NewReader(fin)
	cmd := &exec.Cmd{}
	if len(sa.print_dest) == 0 {
		fout = os.Stdout
	} else {
		cmd = exec.Command("")
		var err error
		cmd.Stdout,err = os.OpenFile(sa.print_dest, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
		fout,err = cmd.StdinPipe()
		if err != nil{
			fmt.Fprintf(os.Stderr,"error\n")
		}
		cmd.Start()
		is_read2 = false
	}

	var page_ctr int
	if sa.page_type == "l"{
		line_ctr := 0
		page_ctr = 1
		for {
			line, crc := buf.ReadString('\n')
			if crc == io.EOF{

				break
			}
			line_ctr++
			if line_ctr > sa.page_len{
				page_ctr++
				line_ctr = 1
			}
			if page_ctr >= sa.start_page && page_ctr <= sa.end_page{
				fout.Write([]byte(line))
			}
		}
	}else{
		page_ctr = 1
		for{
			line, crc := buf.ReadString('\n')
			if crc == io.EOF{
				break
			}
			if page_ctr >= sa.start_page && page_ctr <= sa.end_page{
				fout.Write([]byte(line))
			}
			page_ctr++
		}
	}
	if page_ctr < sa.start_page{
		fmt.Fprintf(os.Stderr,"%s: start_page %d greater tan total pages %d\n",proname,sa.start_page,page_ctr)
	}else if page_ctr < sa.end_page{
		fmt.Fprintf(os.Stderr,"%s: end_page %d greater tan total pages %d\n",proname,sa.end_page,page_ctr)
	}
	if is_read{
		fin.Close()
	}
	if is_read2{
		fout.Close()
	}
}
