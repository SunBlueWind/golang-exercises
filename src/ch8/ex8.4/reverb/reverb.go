package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(conn net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToLower(shout))
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	var wg sync.WaitGroup
	input := bufio.NewScanner(conn)
	for input.Scan() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			echo(conn, input.Text(), 1*time.Second)
		}()
	}
	if input.Err() != nil {
		log.Printf("handleConn error: %s\n", input.Err())
	}
	wg.Wait()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}
