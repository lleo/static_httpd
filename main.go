package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	//"path"
	"path/filepath"
)

func main() {

	fmt.Println(os.Args)

	var lgr = log.New(os.Stdout, "Logger:", log.Lshortfile|log.Ldate|log.Ltime)

	var err error

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var hostport string
	flag.StringVar(&hostport, "hostport", "localhost:8081", "A string of the form \"{host}:{port}\"")

	flag.Parse()

	var twistedDocs string = "/Users/seanegan/source/Twisted-19.10.0/docs/_build/html"

	twistedRootStr, err := filepath.Abs(twistedDocs)
	if err != nil {
		log.Fatal(err)
	}

	var twistedRoot = http.Dir(twistedRootStr)

	//Need to verify rootd is an existing directory

	fmt.Printf("hostport=%s\n", hostport)
	fmt.Printf("cwd=%s\n", cwd)
	fmt.Printf("twistedRoot=%s\n", twistedRoot)

	// create mux
	var mux = http.NewServeMux()

	// register /twisted/
	var twistedHandler = http.StripPrefix("/twisted/", http.FileServer(twistedRoot))
	mux.Handle("/twisted/", twistedHandler)

	// build Server
	var svr = &http.Server{
		Addr:    hostport,
		Handler: mux,
		//TLSConfig: *tls.Config,
		//ReadTimeout: time.Duration,
		//ReadHeaderTimeout: time.Duration,
		//WriteTimeout: time.Duration,
		//IdleTimeout: time.Duration,
		//MaxHeaderBytes: int,
		//TLSNextProto: map[string]func(net.Conn, ConnState),
		//ConnState: func(net.Conn, ConnState),
		ErrorLog: lgr,
		//BaseContext: func(net.Listener) context.Context,
		//ConnContext: func(context.Context, net.Conn) context.Context,
	}

	lgr.Println("starting server...")
	lgr.Fatal(svr.ListenAndServe())
}
