package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type location struct {
	name string
	conn net.Conn
}

func main() {
	var locations []location
	for _, clock := range os.Args[1:] {
		input := strings.Split(clock, "=")
		if len(input) != 2 {
			log.Printf("invalid option: %s\n", clock)
			continue
		}
		name, address := input[0], input[1]
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Printf("connection error to %s: %v\n", address, err)
			continue
		}
		defer conn.Close()
		locations = append(locations, location{name, conn})
	}
	displayClock(locations)
	time.Sleep(10 * time.Second)
}

func displayClock(locations []location) {
	for _, loc := range locations {
		go func(loc location) {
			s := bufio.NewScanner(loc.conn)
			for s.Scan() {
				fmt.Printf("%s: %s\n", loc.name, s.Text())
			}
			if s.Err() != nil {
				log.Printf("scan error: %v\n", s.Err())
			}
		}(loc)
	}
}
