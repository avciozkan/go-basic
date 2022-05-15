package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Kullanici struct {
	Ad      string      `json:"adi"` //tag definition
	Soyad   string      `json:"soyadi"`
	Takipci []Kullanici `json:"takipci,omitempty"`
	Begeni  []Begen     `json:"begeni,omitempty"`
}

type Begen struct {
	Tarih time.Time
}

func main() {
	k := Kullanici{
		Ad:    "Özkan",
		Soyad: "Avcı",
		Takipci: []Kullanici{
			{
				Ad:    "LeBron",
				Soyad: "James",
				Takipci: []Kullanici{
					{
						Ad:    "Cristiano",
						Soyad: "Ronaldo",
						Begeni: []Begen{
							{
								Tarih: time.Now(),
							},
						},
					},
				},
			},
			{
				Ad:    "Lionel",
				Soyad: "Messi",
				Takipci: []Kullanici{
					{
						Ad:    "Erling",
						Soyad: "Haaland",
					},
				},
			},
		},
	}
	arr, err := json.Marshal(k)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arr))
	//mt.Println(k)

}
