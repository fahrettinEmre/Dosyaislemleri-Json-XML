package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

//puanı en yüksek olan kişinin maaşına veri.jsonda zam yapılacak
func main() {
	file, err := ioutil.ReadFile("veri.json")
	check(err)
	var elemanlar []Eleman
	err = json.Unmarshal(file, &elemanlar) //json dosyasındanki verileri elemanlar slicesına attık.
	check(err)
	fmt.Println(elemanlar)

	//içerideki en iyi puanlı elemanı bulma
	var eniyi Eleman = elemanlar[0]
	fmt.Println(eniyi) //Elemanların ilkini yazdırdık.
	for i := 0; i < len(elemanlar); i++ {
		if elemanlar[i].Puan > eniyi.Puan {
			eniyi = elemanlar[i]
		}
	}
	fmt.Println(eniyi) // slicesin içindeki en iyi puanlı eleman bulundu.

	for i := 0; i < len(elemanlar); i++ {
		if eniyi == elemanlar[i] {
			elemanlar[i].Maas += elemanlar[i].Maas * 10 / 100
		}
	}
	fmt.Println(eniyi)

	/* fmt.Println(elemanlar)
	for i := 0; i < len(elemanlar); i++ {     // burada en iyi eleman hangisi diye tekrardan görmek istedik.
		if elemanlar[i].Puan > eniyi.Puan {
			eniyi = elemanlar[i]
		}
	}
	fmt.Println(eniyi) */

	fmt.Println(elemanlar) // zam yapılan elemanı görmek için baktık.

	newFile, err := json.Marshal(elemanlar) //json dosyasına yazdırmak için kullanılır.
	check(err)
	err = ioutil.WriteFile("veri.json", newFile, 0666) // 0666 dosya izni okuma ve yazma
	check(err)                                         // ve json dosyasına yazdırma işlemi

}

type Eleman struct {
	İsim string `json:"isim"`
	Maas int    `json:"maas"`
	Puan int    `json:"puan"`
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
