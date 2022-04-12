package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	//klasör olşturma
	//err := os.Mkdir("Dosya", 0777)
	//check(err)

	// dosya oluşturma
	file, err := os.Create("dosya.txt")
	check(err)

	//dosyaya yazma
	_, err = file.WriteString("Dosya işlemleri")
	check(err)

	//dosya okuma
	//content, err := os.ReadFile("dosya.txt")
	content, err := ioutil.ReadFile("dosya.txt")

	check(err)
	fmt.Println(string(content)) // string koymazsam ascıı tablosunun değeri gelir.

	//dosyanın adını ve yolunu değiştirme
	err = os.Rename("dosya.txt", "test.txt")
	check(err)

	//dosya taşıma
	os.Rename("test.txt", "Dosya/test.txt")
	check(err)

	//dosyanın özelliklerini okumak
	info, err := os.Stat("Dosya/test.txt")
	check(err)
	fmt.Println(info.Size())

	//dosya silme
	//err = os.Remove("test.txt")
	//check(err)
}
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
