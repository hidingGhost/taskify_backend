package main

import (
	"io"
	"log"
	"os"
)

func initLogger() {
	file, err := os.OpenFile("taskify.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Logger konnte nicht initialisiert werden:", err)
	}

	// Schreibt gleichzeitig in Datei und Konsole
	multi := io.MultiWriter(file, os.Stdout)
	log.SetOutput(multi)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
