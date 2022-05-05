package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(IsEmpty("deneme"))
}

func IsEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
