package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	// Trace logger. Use: Trace.Println("I have something standard to say")
	Trace *log.Logger
	// Info logger. Use: Info.Println("Special Information")
	Info *log.Logger
	// Warning logger. Use: Warning.Println("There is something you need to know about")
	Warning *log.Logger
	// Error logger. Use: Error.Println("Something has failed")
	Error *log.Logger
)

var dirDefault = "static"
var dirEnv = "GOWS_DIR"

// Init ...
func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func createDir(dir string, mode os.FileMode) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		Warning.Printf("Content directory `%s` did not exist. It will be created, but this means there will be no content to serve. Specify a different directory with %s.", dir, dirEnv)
		os.Mkdir(dir, mode)
	}
	if empty, _ := IsEmpty(dir); empty == true {
		Warning.Printf("Content directory `%s` is empty. No content to serve.", dir)
	}
}

// IsEmpty checks whether a directory is Empty
func IsEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

func setup() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	createDir(getenv(dirEnv, dirDefault), 0700)
}

func main() {
	setup()

	Info.Printf("Serving content from %s", getenv(dirEnv, dirDefault))
	fs := http.FileServer(http.Dir(getenv(dirEnv, dirDefault)))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":8080", nil)
}
