package main

import (
	"bufio"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	inputPath, outputPath := ReadInputData()
	fmt.Println(inputPath)
	for true {
		convertImagesFromDirectory(inputPath, outputPath)
		time.Sleep(30 * time.Second)
	}
}

func convertImagesFromDirectory(inputPath string, outputPath string) {
	files, err := ioutil.ReadDir(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	db := getDBConn()
	fileChan := make(chan [2]string)
	dbChan := make(chan [4]string)
	outDBChan := make(chan int)
	outConverterChan := make(chan int)
	go writerToDB(dbChan, db, outDBChan)
	for i := 0; i < workersCount; i++ {
		go converterImageFileToBlackWhite(fileChan, dbChan, outConverterChan)
	}
	for _, fileInfo := range files {
		isImage := false
		for _, suffix := range imgSuffixes {
			isImage = isImage || strings.HasSuffix(fileInfo.Name(), suffix)
		}
		if isImage {
			newImagePath := outputPath + fileInfo.Name()
			currentPath := inputPath + fileInfo.Name()
			if _, err := os.Stat(newImagePath); err == nil {
				continue
			} else if !errors.Is(err, os.ErrNotExist) {
				log.Fatal(err)
			}
			fileChan <- [2]string{currentPath, newImagePath}
		}
	}
	close(fileChan)
	for i := 0; i < workersCount; i++ {
		<-outConverterChan
	}
	close(dbChan)
	<-outDBChan
	defer db.Close()
}

func ReadInputData() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input directory: ")
	inputPath, _ := reader.ReadString('\n')
	fmt.Print("Enter output directory: ")
	outputPath, _ := reader.ReadString('\n')

	inputPath = strings.TrimSuffix(inputPath, "\n")
	if inputPath == "" {
		inputPath, _ = filepath.Abs("./exampleInputDir/")
	}
	outputPath = strings.TrimSuffix(outputPath, "\n")
	if outputPath == "" {
		outputPath, _ = filepath.Abs("./exampleOutputDir/")
	}
	if !strings.HasSuffix(inputPath, "/") {
		inputPath += "/"
	}
	if !strings.HasSuffix(outputPath, "/") {
		outputPath += "/"
	}
	return inputPath, outputPath
}
