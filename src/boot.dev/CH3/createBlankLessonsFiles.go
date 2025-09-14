package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	n := 21
	baseName := "L"
	customContent := `
	package main

	import "fmt"

	func main() {
		// go nuts
	}
`

	for i := 1; i <= n; i++ {
		fileName := fmt.Sprintf("%s%d.go", baseName, i)

		file, err := os.Create(fileName)
		if err != nil {
			log.Fatalf("Failed to create file %s: %v", fileName, err)
		}

		_, err = file.WriteString(customContent)
		if err != nil {
			log.Fatalf("Failed to write to file %s: %v", fileName, err)
		}

		file.Close()
		log.Printf("Created file %s with go boilerplate content", fileName)
	}
}
