package main

import (
	"encoding/json"
	"net/http"
	//"html"
	"fmt"
	"log"
	"os/exec"
	"os/signal"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"sync"
	//"os/exec"
	//"code.google.com/p/go.crypto/openpgp"
)

type Message struct {
	Body string
	Signature string
	From string
}

type ApiHandler struct {

}

type ExternalApiHandler struct {

}

func main() {

	startTor()
	onionHostname, _ := readOnionHostname()
	log.Print(onionHostname);

	// Register and update namecoin address

	startHttpServers()

	// Wait forever while the http servers run

	var wg sync.WaitGroup
	wg.Add(1);
	wg.Wait();
}

// Start tor hidden service

func startTor() {
	torCmd := exec.Command("bin/tor", "-f", "./torrc")
	torCmd.Stdout = os.Stdout
    torCmd.Stderr = os.Stderr

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		for sig := range c {
			log.Printf("got %v, killing tor", sig)
			torCmd.Process.Kill()
			os.Exit(1)
		}
	}()

	err := torCmd.Start()

	if err != nil {
		log.Fatal(err)
	}
}

func startHttpServers() {

	// Start external http server

	externalApiHandler := new(ExternalApiHandler)

	externalApiServer := &http.Server {
		Addr:           ":7777",
		Handler:        externalApiHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go externalApiServer.ListenAndServe()

	// Start local API server

	apiHandler := new(ApiHandler)

	localApiServer := &http.Server {
		Addr:           ":8080",
		Handler:        apiHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go localApiServer.ListenAndServe()

}

func (h *ExternalApiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
		var msg Message
		body := make([]byte, req.ContentLength)
		req.Body.Read(body)
		json.Unmarshal(body, &msg)
		fmt.Fprintf(w, "%s", msg.Signature)
}

func (h *ApiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s", "foo")

	// publish leet
	// follow litter name
	// get leets of followed names
}


func readOnionHostname() (string, error){
	// Waiting for the hostname file to be
	// written if it doesn't yet exist

	hiddenServiceHostname := "hidden_service/hostname"
	onionHostnameFileExists := false
	for i := 0; i < 5 ; i++ {
		var err error
		onionHostnameFileExists, err = exists(hiddenServiceHostname)
		if err != nil {
			log.Print("Couldn't read .onion hostname file. Exiting.");
			os.Exit(1)
		}
		if onionHostnameFileExists {
			break	
		}
		log.Print("Waiting for hostname file to be written...");
		time.Sleep(1000 * time.Millisecond)
	}

	// Read hostname file

	content, err := ioutil.ReadFile(hiddenServiceHostname)
	if err != nil {
		log.Print("Couldn't find .onion hostname file. Exiting.");
		os.Exit(1)
	}
	lines := strings.Split(string(content), "\n")
	onionHostname := lines[0]

	return onionHostname, err
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}
