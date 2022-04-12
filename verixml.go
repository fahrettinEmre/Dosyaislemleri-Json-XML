//15 kişilk takımdan en çok gol atan 11 i bulup bir takım oluşturacağız.
//Çıkan takım başka bir xml dosyasına eklenecek ve ek olarak xml dosyası json dosyasına çevirilecek.

package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//xml verisi ile çalısırken bir tag ın altında başka
//bir tag var ise onun struct ını oluştur.

type Takım struct {
	Oyuncular []Oyuncu `xml:"oyuncu"`
}
type Oyuncu struct {
	İsim   string `xml:"isim"`
	Mevki  string `xml:"mevki"`
	Gol    int    `xml:"gol"`
	Sosyal struct {
		Facebook string `xml:"facebook"`
		Twitter  string `xml:"twitter"`
		Youtube  string `xml:"youtube"`
	} `xml:"sosyal"`
	//Sosyal Sosyal `xml:"sosyal"`
}

/* type Sosyal struct {						// ' türlü tanımlama yapabiliriz. 1. ayrı bir struct 2. Oyuncunun
	Facebook string `xml:"facebook"`		// içerisine bir struct.
	Twitter  string `xml:"twitter"`
	Youtube  string `xml:"youtube"`
} */
type jsonOyuncu struct { // xml i json a cevirmek için struct oluşturduk.
	İsim   string `json:"isim"`
	Mevki  string `json:"mevki"`
	Gol    int    `json:"gol"`
	Sosyal struct {
		Facebook string `json:"facebook"`
		Twitter  string `json:"twitter"`
		Youtube  string `json:"youtube"`
	} `json:"sosyal"`
}

func main() {
	var takim Takım
	file, err := ioutil.ReadFile("veri.xml") // xml dosyasını okuduk.
	check(err)
	err = xml.Unmarshal(file, &takim) // takım structına gönderdik.
	check(err)
	for i := 0; i < len(takim.Oyuncular); i++ {
		//fmt.Println(takim.Oyuncular[i])  // bütün takımdaki oyuncuların özelikleri yazdırıldı.
		for j := i + 1; j < len(takim.Oyuncular); j++ {
			if takim.Oyuncular[j].Gol > takim.Oyuncular[i].Gol { //takımın içindeki en fazla gol atanları sıraladık.
				gecici := takim.Oyuncular[i]
				takim.Oyuncular[i] = takim.Oyuncular[j]
				takim.Oyuncular[j] = gecici

			}
		}
	}
	for i := 0; i < len(takim.Oyuncular); i++ {
		fmt.Println(takim.Oyuncular[i])
	}

	takim.Oyuncular = takim.Oyuncular[0:11] // 11 tane oyuncuyu aldı.
	fmt.Println(len(takim.Oyuncular))

	newFile, err := os.Create("Atakımı.xml") // xml oluşturduk.
	check(err)
	xmlbyte, err := xml.Marshal(takim)
	check(err)
	_, err = newFile.Write(xmlbyte) //xml dosyasına yazdık.

	var jsonOyuncuları []jsonOyuncu // Json oyuncularının bir slıcesini oluşturduk
	var jsonOyuncusu jsonOyuncu     // Json oyuncusunun bir objesini oluşturduk.
	for i := 0; i < len(takim.Oyuncular); i++ {
		jsonOyuncusu.İsim = takim.Oyuncular[i].İsim
		jsonOyuncusu.Mevki = takim.Oyuncular[i].Mevki
		jsonOyuncusu.Gol = takim.Oyuncular[i].Gol
		jsonOyuncusu.Sosyal.Facebook = takim.Oyuncular[i].Sosyal.Facebook
		jsonOyuncusu.Sosyal.Twitter = takim.Oyuncular[i].Sosyal.Twitter
		jsonOyuncusu.Sosyal.Youtube = takim.Oyuncular[i].Sosyal.Youtube
		jsonOyuncuları = append(jsonOyuncuları, jsonOyuncusu)

	}
	jsonfile, err := os.Create("Atakımı.json") //json dosyası oluşturduk.
	check(err)
	jsonbyte, err := json.Marshal(jsonOyuncuları)
	check(err)
	_, err = jsonfile.Write(jsonbyte)
	check(err)

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
