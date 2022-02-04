package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

type logWriter struct{}

func main() {

	fileToOpen := os.Args[1]
	fmt.Println(fileToOpen)

	file, err := os.OpenFile(fileToOpen, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file)

	fileType := reflect.TypeOf(file).Kind()
	fmt.Println(fileType)

	lw := logWriter{}

	io.Copy(lw, file)

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
