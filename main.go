package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

const VERSION = "v1.2"
const PORT = "8080"

type Endpoint struct {
	Path   string
	Type   string
	Active bool
}

type TemplateData struct {
	ServerIP  string
	Version   string
	Deployed  string
	Port      string
	Endpoints []Endpoint
}

const FORMAT = "2006-01-02 15:04:05"

func getServerIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr).IP.String()
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = PORT
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return port
}

func main() {

	port := getPort()
	log.Println("[-] Listening on...", port)

	tmpl := template.Must(template.ParseFiles("temp.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TemplateData{
			ServerIP: getServerIP(),
			Version:  VERSION,
			Deployed: time.Now().Local().Format(FORMAT),
			Port:     PORT,
			Endpoints: []Endpoint{
				{Path: "/flip", Type: "GET", Active: true},
				{Path: "/time", Type: "GET", Active: true},
				{Path: "/version", Type: "GET", Active: false},
			},
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		current := time.Now().Local()
		fmt.Fprintf(w, current.Format(FORMAT))
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, VERSION)
	})

	http.HandleFunc("/flip", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		textToFlip := r.FormValue("text")
		fmt.Fprintf(w, Flip(textToFlip))
	})

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}

}
