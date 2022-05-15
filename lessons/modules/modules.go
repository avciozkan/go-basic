package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
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
	arr, _ := json.Marshal(k)
	fmt.Println(string(arr))
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello, World!")
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.SendString(string(arr))
	})

	app.Listen(":3000")
}
