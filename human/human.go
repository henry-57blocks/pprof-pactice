package human

import (
	"log"
	"time"
)

var (
	AllHumans = []Human{
		&Superman{},
		&IronMan{},
		&Batman{},
		&SpiderMan{},
	}
)

const (
	Ki = 1024
	Mi = Ki * Ki
	Gi = Ki * Mi
	Ti = Ki * Gi
	Pi = Ki * Ti
)

type Human interface {
	Skills()
}

type Superman struct {
}

func (s Superman) Skills() {
	log.Println("superman")
	_ = make([]byte, 16*Mi)
}

type IronMan struct {
}

func (i IronMan) Skills() {
	log.Println("IronMan")
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(30 * time.Second)
		}()
	}
}

type Batman struct {
}

func (b Batman) Skills() {
	log.Println("Batman")
	loop := 10000000000
	for i := 0; i < loop; i++ {

	}
}

type SpiderMan struct {
	buffer [][Mi]byte
}

func (s SpiderMan) Skills() {
	log.Println("spiderMan")
	max := Gi
	for len(s.buffer)*Mi < max {
		s.buffer = append(s.buffer, [Mi]byte{})
	}
}
