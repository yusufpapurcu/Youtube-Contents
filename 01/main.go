package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/tarm/serial"
)

var commands = make(chan string)

func main() {
	go func() {
		for {
			mySerial()
		}
	}()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), "[@]") {
				commands <- strings.SplitAfter(scanner.Text(), "[@] ")[1]
			}
		}

		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}
}

// [22:41:18] [Server thread/INFO]: [@]
func mySerial() {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Println(err)
		s.Close()
		return
	}
	for {
		command := <-commands
		_, err := s.Write([]byte(command))
		if err != nil {
			log.Println(err)
			s.Close()
			return
		}
	}
}
