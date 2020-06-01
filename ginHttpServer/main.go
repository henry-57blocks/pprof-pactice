package main

import (
	"log"
	"os"
	"runtime"
	"time"

	// The gin framework use this library to register the pprof routing.
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/ohdata/pprof/human"
)

func main() {

	// Limit CPU usage tp avoid overload, this line of code is only used for the test of this program,
	// and has nothing to do with pprof.
	runtime.GOMAXPROCS(1)

	r := gin.New()

	// This line of code is used to register the pprof route.
	pprof.Register(r)

	go func() {
		if err := r.Run(":6060"); err != nil {
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
