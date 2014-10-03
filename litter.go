package main

import (
	"net/http"
	"log"
	"fmt"
	"os/exec"
	"os/signal"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"sync"
	"github.com/codegangsta/cli"
	"github.com/awt/litter/public"
	"github.com/awt/litter/private"
	"github.com/awt/litter/store"
	"github.com/awt/litter/config"
	"github.com/awt/litter/nmc"
	//"code.google.com/p/go.crypto/openpgp"
)

var Config struct {
}

func main() {

	var conf  = &config.Config{}
	conf.SetEnvironment("local")
	conf.Set("dbpath", "./litter.db")
	store.Config = conf
	nmc.Config = conf

	app := cli.NewApp()
	app.Name = "litter"
	app.Usage = "Spreading litter across the dark web"

	app.Commands = []cli.Command{
		{
			Name:      "daemon",
			ShortName: "d",
			Usage:     "start the daemon process",
			Action: func(c *cli.Context) {
				startTor()
				onionHostname, _ := readOnionHostname()
				log.Print(onionHostname);

				// Register and update namecoin address

				startHttpServers()
					
				nmc.FetchLeets()
						

				// Set up persistent connections with all friends

				// Wait forever while the http servers run

				var wg sync.WaitGroup
				wg.Add(1);
				wg.Wait();
			},
		},
		{
			Name:      "register",
			ShortName: "d",
			Usage:     "register <name>",
			Action: func(c *cli.Context) {
				name := c.Args().First()

				// check if name is taken
				if !nmc.IsRegistered(name) {
					log.Println(name)

					nmc.ReserveName(name)
				} else {
					fmt.Printf("%s is already registered in the Namecoin network.\n", name)	
				}
			},
		},
	}
	app.Run(os.Args)

}

func initializeDatabase() {

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
			log.Printf("Got %v, shutting down.", sig)
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

	externalApiHandler := new(public.ApiHandler)

	externalApiServer := &http.Server {
		Addr:           ":7777",
		Handler:        externalApiHandler,
	}

	go externalApiServer.ListenAndServe()

	// Start local API server

	apiHandler := new(private.ApiHandler)

	localApiServer := &http.Server {
		Addr:           ":8080",
		Handler:        apiHandler,
	}

	go localApiServer.ListenAndServe()

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
