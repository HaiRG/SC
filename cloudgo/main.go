package main

import( 
	"cloudgo/server"
	"os"
	"github.com/spf13/pflag"
)
func main(){
	port := os.Getenv("PORT");
	if len(port) == 0{
		port = "5556"
	} 
	newport := pflag.StringP("port", "p", "5555", "the port to listen")
	pflag.Parse()
	if len(*newport) != 0{
		port = *newport
	}
	server.Start(port)
}
