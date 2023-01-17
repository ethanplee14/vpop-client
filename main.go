package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/pkg/browser"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1337")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	defer conn.Close()
	hostname, _ := os.Hostname()
	conn.Write([]byte(hostname + "\n"))
	for {
		log.Print("Waiting for url...")
		scanner := bufio.NewScanner(conn)
		scanner.Scan()
		if scanner.Err() != nil {
			log.Fatalln("Error reading from server!")
			return
		}
		message := scanner.Text()

		if message == "ping" {
			conn.Write([]byte("pong"))
		}
		log.Print("Got url: " + message)
		browser.OpenURL(message)
	}
}
