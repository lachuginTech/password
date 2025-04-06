package files

import (
	"fmt"
	"os"
	"password/output"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	date, err := os.ReadFile(db.filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return date, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		output.PrintError(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		output.PrintError(err)
		return
	}
	fmt.Println("Запись успешна")
}
