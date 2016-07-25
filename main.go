package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"./api"
	srv "./http"
)

type CliOption struct {
	port int
}

/*
  ↓みたいな使い方で動くはず

  $curl -w '\n' 'http://localhost:8080/' --data 'name=test,size=1'
  {"Status":0,"Message":"","Body":"name=test,size=1"}
  $ curl http://localhost:8080
  {"Status":1,"Message":"Method is Get","Body":"Nothing"}

 */
func main() {

	// 参考は↓あたりです
	// http://qiita.com/rerofumi/items/66be3c55405e03dbdcf0

	// shutdown defer
	defer srv.Stop()

	// parse cli option
	var option CliOption
	flag.IntVar(&option.port, "p", 8080, "int flag")
	flag.Parse()

	// routing
	mux := http.NewServeMux()
	mux.HandleFunc("/", api.GetHandler())

	// server start
	err := srv.Start(option.port, mux)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
