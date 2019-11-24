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

	var err error

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var hostport string
	flag.StringVar(&hostport, "hostport", "localhost:8080", "A string of the form \"{host}:{port}\"")

	var rootd string
	flag.StringVar(&rootd, "rootdir", cwd, "The root directory of static files to be served.")

	rootd, err = filepath.Abs(rootd)
	if err != nil {
		log.Fatal(err)
	}
	flag.Parse()

	//Need to verify rootd is an existing directory

	fmt.Printf("hostport=%s\n", hostport)
	fmt.Printf("cwd=%s\n", cwd)
	fmt.Printf("rootd=%s\n", rootd)

	log.Fatal(http.ListenAndServe(hostport, http.FileServer(http.Dir(rootd))))
}
