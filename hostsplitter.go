package main

import (
	"fmt"
	"github.com/VividCortex/godaemon"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

var (
	routedHostnames map[string]int
	Sites           []Site
	HTTPClient      http.Client
)

var (
	logFileLoc = kingpin.Flag("log", "Location of the log file").Default("stdout").String()
	daemonize  = kingpin.Flag("daemon", "If daemonized, the program will run in the background.").Default("true").Bool()
	sitesLoc   = kingpin.Flag("sites_dir", "Location of site files").Short('h').Default("/etc/hostsplitter/").String()
	bindAddr   = kingpin.Flag("bind", "Bind address").Short('b').Default(":80").String()
)

func main() {
	HTTPClient = http.Client{}

	kingpin.Parse()

	log.Print("Starting hostsplitter")

	if *logFileLoc != "stdout" {
		logFile, err := os.OpenFile(*logFileLoc, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0640)
		defer logFile.Close()
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		log.Print("Using ", *logFileLoc, " for logging")
		log.SetOutput(logFile)
	}

	if *daemonize {
		log.Print("Daemonizing... Bye Bye")
		godaemon.MakeDaemon(&godaemon.DaemonAttr{})
	}

	LoadConfig()

	go SignalHandler()

	log.Fatal(http.ListenAndServe(*bindAddr, &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			HTTPLogger(r)
			if i, ok := routedHostnames[string(r.Host)]; ok {
				r.Header.Set("X-Hostsplitter-Secret", Sites[i].Secret)
				r.Header.Set("Host", r.Host)
				r.URL.Scheme = "http"
				r.URL.Host = Sites[i].GetBackend()
				r.RequestURI = ""
			} else {
				log.Print("%q is not routed", r.Host)
			}
		},
	}))
}

func HTTPLogger(r *http.Request) {
	log.Print(fmt.Sprintf("httplog> %v %v (%v) (conlen %v)", r.Host, r.Method, r.RequestURI, r.RemoteAddr))
}
