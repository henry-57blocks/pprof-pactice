package main

import (
	"log"
	"net/http"
	// Handler will be automatically registered to the HTTP server
	// for easy access to the program running sample report through the HTTP interface
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"

	"github.com/ohdata/pprof/human"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	// Limit CPU usage tp avoid overload, this line of code is only used for the test of this program,
	// and has nothing to do with pprof.
	runtime.GOMAXPROCS(1)

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	for {
		for _, v := range human.AllHumans {
			v.Skills()
		}
		time.Sleep(time.Second)
	}
}
