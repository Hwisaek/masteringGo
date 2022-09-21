package main

import (
	"io"
	"log"
	"os"
)

const (
	logFilePath = "/var/log/masteringGo/ch1/4/"
	logFile1    = logFilePath + "customLog1.log"
	logFile2    = logFilePath + "customLog2.log"
)

func main() {
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		_ = os.Mkdir(logFilePath, 0644)
	}

	f1, err := os.OpenFile(logFile1, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f1.Close()

	f2, err := os.OpenFile(logFile2, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f1.Close()

	log.SetOutput(io.MultiWriter(log.Writer(), f1, f2))
	log.SetPrefix("customLog ")
	log.SetFlags(log.LUTC | log.Lshortfile | log.LstdFlags)

	log.Println("test")

}
