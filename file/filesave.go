package file

import (
	"fmt"
	"os"
)

func SaveFile(filename string, content string) any {
	file, err := os.Create("file.go") // For read access.
	if err != nil {
		fmt.Println(err)
	}
	return file
}
