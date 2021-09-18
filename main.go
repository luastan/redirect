package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var (
	InfoLogger   *log.Logger
	ErrorLogger  *log.Logger
	DumpRequests *bool
	BaseUrl      *url.URL
	StatusCode   *int
	FixedPath    bool
)

func init() {
	// Global loggers set-up
	InfoLogger = log.New(os.Stderr, "INFO: ", log.Ltime)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ltime)
}

// Simple handler to perform redirects and optionally dump requests
func requestHandler(writer http.ResponseWriter, req *http.Request) {
	InfoLogger.Printf("request from %s", req.RemoteAddr)
	if *DumpRequests {
		requestDump, err := httputil.DumpRequest(req, true)
		if err != nil {
			ErrorLogger.Println(err.Error())
		} else {
			println(string(requestDump))
		}
	}

	if !FixedPath {
		BaseUrl.Path = req.RequestURI
	}

	http.Redirect(writer, req, BaseUrl.String(), *StatusCode)
}

func main() {
	//	Flag parsing
	listenAddr := flag.String("addr", ":8888", "Address used to set up the listener")
	DumpRequests = flag.Bool("dump", false, "Dump requests reaching the server to stdout")
	StatusCode = flag.Int("status", http.StatusMovedPermanently, "Status code to send with the response")
	targetPath := flag.String("path", "", "Path to redirect to")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		ErrorLogger.Fatalln("Missing host to redirect to")
	} else if flag.NArg() > 1 {
		flag.Usage()
		ErrorLogger.Fatalln("Unexpected arguments were provided")
	}

	var err error

	// Parse the URL to redirect to
	BaseUrl, err = url.Parse(flag.Arg(0))
	if err != nil {
		ErrorLogger.Fatalln(err.Error())
	}

	// Check if the path is fixed and set it
	FixedPath = len(*targetPath) > 0
	if FixedPath {
		BaseUrl.Path = *targetPath
	}

	// Setup the HTTP handlers
	http.HandleFunc("/", requestHandler)
	InfoLogger.Printf("listening on %s\n", *listenAddr)
	err = http.ListenAndServe(*listenAddr, nil)
	if err != nil {
		ErrorLogger.Fatalln(err.Error())
	}

}
