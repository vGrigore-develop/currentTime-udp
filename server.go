package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	dateFormat = "2006-January-02"
	timeFormat = "2006-01-02 15:04:05"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		currentTime := time.Now()
		fmt.Print("\n-> ", string(buffer[0:n]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exitin UDP server!")
			return
		}
		option := string(buffer[0:n])

		data := []byte("No option")

		switch option {
		case "-s":
			firstPartDate := strings.Split(currentTime.Format(dateFormat), "-")
			secondPartDate := strings.Split(currentTime.Format(timeFormat), " ")
			data = []byte(firstPartDate[1] + " " + firstPartDate[2] + " " + secondPartDate[1])
			break
		case "-u":
			secs := currentTime.Unix()
			data = []byte(strconv.Itoa(int(secs)))
		}

		fmt.Printf("\ndata: %s", string(data))
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
