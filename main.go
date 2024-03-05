package main

import (
	"files/fileutils"
	"fmt"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/text.txt"
	c, err := fileutils.ReadTextFile(filePath)

	if err == nil {
		fmt.Println(c)
		WriteContet := fmt.Sprintf("Origianl: %v/n Double Original: %v", c, c)
		newContent := WriteContet
		fileutils.WriteToFile(filePath+".output.txt", newContent)
	} else {
		fmt.Printf("Error Panic %v", err)
	}
}
