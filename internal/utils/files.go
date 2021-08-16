package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetFileContent(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(fmt.Printf("Can't open file %s", filePath))
		return "", err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(fmt.Printf("Error during file closure. File: %s", filePath))
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Println(fmt.Printf("Content reading error. File: %s", filePath))
		return "", err
	}

	return content, nil
}
