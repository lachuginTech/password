package files

import (
	"fmt"
	"os"
)

func ReadFile(path string) {}

func WriteFile(content string, name string) {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(content)
	if err != nil {
		file.Close()
		panic(err)
		return
	}
	fmt.Println("Запись успешна")
	file.Close()
}
