package main

import "fmt"

var map_1 map[int]string

var map_4 map[string]string = make(map[string]string)

func main() {
	map_1 = make(map[int]string)

	map_2 := make(map[int]int)
	map_3 := make(map[int]int, 3)
	//sira garantisi yoktur
	map_4["asds1"] = "asd"
	map_4["asds2"] = "asd"
	map_4["asds3"] = "asd"
	map_4["asds4"] = "asd"
	map_4["asds5"] = "asd"
	map_4["asds6"] = "asd"
	map_1[0] = "asdsa"
	map_2[0] = 1
	map_2[1] = 2
	fmt.Println(map_1, "\n", map_2, "\n", map_3, "\n", map_4)
}
