package main

import (
	"fmt"
	"time"
)

func main() {
	timer(time.After(0 * time.Second))
	timer2(time.After(0*time.Second), "Selam dostum", "Selamlar dostlar")
	timer3(time.After(0 * time.Second))
	timer4(time.After(0 * time.Second))
}

func timer(c <-chan time.Time) {
	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}
	}
}

func timer2(c <-chan time.Time, message ...string) {
	b := false
	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
			if !b {
				for _, i := range message {
					fmt.Println(i)
				}
			}
			b = true
		}
	}
}

func timer3(c <-chan time.Time, message ...string) {
	defer fmt.Println("Bir sonraki aşamaya geçiyor")
	defer fmt.Println("Timer bitti")
	//defer work after function timer3 complete. From last to begin. Example, 2nd line defer works first, 1st line defer works second :)
	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}
	}

}

func timer4(c <-chan time.Time, message ...string) {
	defer fmt.Println("Bir sonraki aşamaya geçiyor")
	defer fmt.Println("Timer bitti")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//defer work after function timer3 complete. From last to begin. Example, 2nd line defer works first, 1st line defer works second :)
	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
			if time.Now().Day() == 12 {
				panic("Something went wrong") // java side throw Exception...
			}
		}
	}

}
