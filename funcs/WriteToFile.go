package funcs

import (
	"fmt"
	"os"
)

func WriteToFile(content string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	content= content +"\n\n"
	_, err = fmt.Fprint(file, content)
	if err != nil {
		return err
	}

	return nil
}
