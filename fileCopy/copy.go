package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var sourceFile = "/Documents/vscode/go_learning/file.txt"
var destinationFile = "/Documents/vscode/go_learning/myfile.txt"

func main() {
	copy()
}

func copy()  {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	} 
	inputPath := filepath.Join(homeDir + sourceFile)
	input, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	
	destPath := filepath.Join(homeDir + destinationFile)
	err = ioutil.WriteFile(destPath, input, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Copied successfully!")
}
