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

func scan(conn net.Conn, line chan string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		line <- input.Text()
	}
}

func handleConn(conn net.Conn) {
	var wg sync.WaitGroup
	defer func() {
		wg.Wait()
		conn.Close()
	}()

	line := make(chan string)
	go scan(conn, line)
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("timeout...")
			return
		case text := <-line:
			wg.Add(1)
			go func() {
				defer wg.Done()
				echo(conn, text, 1*time.Second)
			}()
		}
	}
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
