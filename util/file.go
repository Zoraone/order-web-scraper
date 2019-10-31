package util

import (
	"log"
	"os"
)

func WriteToFile(content []byte) {
	f, err := os.Create("datadump.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		log.Println(err)
	}
}