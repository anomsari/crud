package main

import (
	"fmt"
	"anom/model"
	// "time"
)
func main(){

	var segitiga = model.Segitiga{Alas:12, Sisi:10}
	fmt.Printf("Keliling= %v", segitiga.Luass())

}
