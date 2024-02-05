package filemanagement

import (
	"fmt"
	"os"
)

// ListFiles lists files in a directory
func ListFiles(directory string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Files in", directory, ":")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
