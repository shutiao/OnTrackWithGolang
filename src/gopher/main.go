package main

import (
	"fmt"
	"gopher/model"
)


//func validateAge(g *gopher) {
//	g.isAdult = g.age > 21
//}

func main() {
	gopherList := model.GetLists()
	for _, gopher := range gopherList {
		fmt.Println(gopher.Jump())
	}
}

