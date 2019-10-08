# 中山大学数据科学与计算机学院本科生实验报告
| 课程名称 | 服务计算 |   任课老师   | 潘茂林 |
| :------: | :------------------: | :----------: | :----: |
|   年级   |      2017                | 专业（方向） |     软件工程   |
|   学号   |  17343033                    |     姓名     |    郭海锐    |

## 实验名称
CLI 命令行实用程序开发基础

## 实验要求

使用 golang 开发 开发 Linux 命令行实用程序 中的 selpg

提示：

1. 请按文档 使用 selpg 章节要求测试你的程序
2. 请使用 pflag 替代 goflag 以满足 Unix 命令行规范， 参考：Golang之使用Flag和Pflag
3. golang 文件读写、读环境变量，请自己查 os 包
4. “-dXXX” 实现，请自己查 os/exec 库，例如案例 Command，管理子进程的标准输入和输出通常使用 io.Pipe，具体案例见 Pipe

## 实验过程

改写的主要流程，对于原c代码，主要需要改写以下部分

命令结构：
```c
struct selpg_args
{
	int start_page;
	int end_page;
	char in_filename[BUFSIZ];
	int page_len; /* default value, can be overriden by "-l number" on command line */
	int page_type; /* 'l' for lines-delimited, 'f' for form-feed-delimited */
					/* default is 'l' */
	char print_dest[BUFSIZ];
};
```

将这部分改写为以下部分，使用go语言的方式重新定义
```go
type selpg_args struct{
	start_page int
	end_page int
	in_filename string
    page_len int
    page_type string
    print_dest string
}
```

源代码除了go结构体外，还包括了参数的处理及判断函数`process_args(ac, av, &sa);`以及输出处理函数`process_input(sa);
`以下分别进行改写

首先是`process_args(ac, av, &sa);`

对于这个函数，源代码是对输入字符串进行一一判断，取值，而我们可以采用pflag，直接将命令行参数绑定到相应的值，因此，只需要判断所赋的值是否正确。

`pflag`如下：

```go
var sa selpg_args
pflag.IntVarP(&sa.start_page,"start_page","s",-1,"start_page")
pflag.IntVarP(&sa.end_page,"end_page","e",-1,"end_page")
pflag.IntVarP(&sa.page_len,"page_len","l",72,"page_len")
pflag.StringVarP(&sa.page_type,"page_type","f","l","page_type")
pflag.StringVarP(&sa.print_dest,"print_dest","d","","print_dest")
```
上面的语句包括了shorthand以及默认赋值，默认赋值可以用来判断用户输入的参数是否正确

参数判断如下：
```go
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
```

主要是判断是否有输入参数，以及参数的逻辑是否正确，比如开始页面不能大于终止页面。

然后是`process_input(sa);
`输出函数

确定输入源
```c
FILE *fin; /* input stream */
	FILE *fout; /* output stream */
	char s1[BUFSIZ]; /* temp string var */
	char *crc; /* for char ptr return code */
	int c; /* to read 1 char */
	char line[BUFSIZ];
	int line_ctr; /* line counter */
	int page_ctr; /* page counter */
	char inbuf[INBUFSIZ]; /* for better performance on input stream */

	/* set the input source */
	if (sa.in_filename[0] == '\0')
	{
		fin = stdin;
	}
	else
	{
		fin = fopen(sa.in_filename, "r");
		if (fin == NULL)
		{
			fprintf(stderr, "%s: could not open input file \"%s\"\n",
			progname, sa.in_filename);
			perror("");
			exit(12);
		}
	}
```
改为

```go
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
```
确定输出目的地

```c
setvbuf(fin, inbuf, _IOFBF, INBUFSIZ);

	/* set the output destination */
	if (sa.print_dest[0] == '\0')
	{
		fout = stdout;
	}
	else
	{
		fflush(stdout);
		sprintf(s1, "lp -d%s", sa.print_dest);
		fout = popen(s1, "w");
		if (fout == NULL)
		{
			fprintf(stderr, "%s: could not open pipe to \"%s\"\n",
			progname, s1);
			perror("");
			exit(13);
		}
	}
```
改为
```go
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
```
输出，根据选择的格式与输出方式
```c
if (sa.page_type == 'l')
	{
		line_ctr = 0;
		page_ctr = 1;

		while (1)
		{
			crc = fgets(line, BUFSIZ, fin);
			if (crc == NULL) /* error or EOF */
				break;
			line_ctr++;
			if (line_ctr > sa.page_len)
			{
				page_ctr++;
				line_ctr = 1;
			}
			if ( (page_ctr >= sa.start_page) && (page_ctr <= sa.end_page) )
			{
				fprintf(fout, "%s", line);
			}
		}
	}
	else
	{
		page_ctr = 1;
		while (1)
		{
			c = getc(fin);
			if (c == EOF) /* error or EOF */
				break;
			if (c == '\f') /* form feed */
				page_ctr++;
			if ( (page_ctr >= sa.start_page) && (page_ctr <= sa.end_page) )
			{
				putc(c, fout);
			}
		}
	}
```
改为以下代码（为了方便测试，将换页符改为换行符）
```go
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
```
最后是错误判断以及文件输入输出流关闭
```c
if (page_ctr < sa.start_page)
	{
		fprintf(stderr,
		"%s: start_page (%d) greater than total pages (%d),"
		" no output written\n", progname, sa.start_page, page_ctr);
	}
	else if (page_ctr < sa.end_page)
	{
		fprintf(stderr,"%s: end_page (%d) greater than total pages (%d),"
		" less output than expected\n", progname, sa.end_page, page_ctr);
	}
	if (ferror(fin)) /* fgets()/getc() encountered an error on stream fin */
	{
		strcpy(s1, strerror(errno)); /* !!! PBO */
		fprintf(stderr, "%s: system error [%s] occurred on input stream fin\n",
		progname, s1);
		fclose(fin);
		exit(14);
	}
	else /* it was EOF, not error */
	{
		fclose(fin);
		fflush(fout);
		if (sa.print_dest[0] != '\0')
		{
			pclose(fout);
		}
		fprintf(stderr, "%s: done\n", progname);
	}
```
改为
```go
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
```

以上即为代码改写的主要部分

## 实验测试

1. `$ selpg -s1 -e1 input_file`
![](https://github.com/HaiRG/SC/raw/master/selpg/image/1.PNG)
2. `$ selpg -s1 -e1 < input_file`
![](https://github.com/HaiRG/SC/raw/master/selpg/image/2.PNG)
3. `$ other_command | selpg -s10 -e20`
![](https://github.com/HaiRG/SC/raw/master/selpg/image/3.PNG)
4. `$ selpg -s10 -e20 input_file >output_file`
![](https://github.com/HaiRG/SC/raw/master/selpg/image/4.PNG)
5. `$ selpg -s10 -e20 input_file 2>error_file`
![](https://github.com/HaiRG/SC/raw/master/selpg/image/5.PNG)
6. `$ selpg -s10 -e20 input_file >output_file 2>error_file`
![](https://github.com/HaiRG/SC/raw/master/selpg/image/6.PNG)
7. `$ selpg -s10 -e20 input_file >output_file 2>/dev/null`
![](https://github.com/HaiRG/SC/raw/master/selpg/image/7.PNG)
8. `$ selpg -s10 -e20 input_file >/dev/null`
![](https://github.com/HaiRG/SC/raw/master/selpg/image/8.PNG)
9. `$ selpg -s10 -e20 input_file | other_command`
![](https://github.com/HaiRG/SC/raw/master/selpg/image/9.PNG)
10. `$ selpg -s10 -e20 input_file 2>error_file | other_command`
![](https://github.com/HaiRG/SC/raw/master/selpg/image/10.PNG)
11. `$ selpg -s10 -e20 -l66 input_file`
![](https://github.com/HaiRG/SC/raw/master/selpg/image/11.PNG)
12. `$ selpg -s10 -e20 -f input_file`
![](https://github.com/HaiRG/SC/raw/master/selpg/image/12.PNG)
13. `$ selpg -s10 -e20 -dlp1 input_file`
![](https://github.com/HaiRG/SC/raw/master/selpg/image/13.PNG)
14. `$ selpg -s10 -e20 input_file > output_file 2>error_file &`
![](https://github.com/HaiRG/SC/raw/master/selpg/image/14.PNG)

以上即为本次实验的全部内容

## 实验心得与体会

本次实验是对一份命令行程序的改编，将原用c语言编写的程序，改为利用go语言实现。改写过程中，我首先研读了源代码，了解了其中各个模块的作用，然后，通过实验要求中的提示的所需要的各个库，将其功能替换到源代码中，而源程序中的大体框架并不需要进行改变，只是利用pflag库进行简化，以及将输入输出改写为go语言的形式。经过本次实验，我对于go语言的一些库比如os，pflag，flag等有了初步的了解，也加深了我对go语言的了解。
