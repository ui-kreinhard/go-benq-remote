package main

import (
	"github.com/ui-kreinhard/go-benq-remote/beamer"
	"github.com/ui-kreinhard/go-benq-remote/web"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalln("Usage: /dev/serialPortname baudrate listenAddr")
	}
	uartDevicePath := os.Args[1]
	baudRateStr := os.Args[2]
	listenAddr := os.Args[3]
	
	baudRate, err := strconv.Atoi(baudRateStr)
	if err != nil {
		log.Fatalln("Invalid baudrate")
	}
	
	log.Println("Using", uartDevicePath, "with", baudRate, "baud")
	beamerInstance := beamer.NewBeamer(uartDevicePath, uint(baudRate))
	webInstance := web.NewWeb(beamerInstance)
	err = webInstance.Start(listenAddr)
	if err != nil {
		log.Fatalln(err)
	}
}