package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	filename := "myfile.txt"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	homePath, err := os.UserHomeDir()
	fmt.Println(filepath.Ext(homePath))

	dest, err := os.OpenFile("/home/blaze/Documents/neworkspace/file.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	bytesWritten, err := io.Copy(dest, file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bytes written: %d\n", bytesWritten)

}
