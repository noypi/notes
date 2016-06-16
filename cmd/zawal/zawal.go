package main

import(
	"fmt"
	"time"
)

func main(){
	// sample sunrise
	tbegin, _ := time.Parse("15:04", "05:23")
	// sample sunset
	tend, _ := time.Parse("15:04", "18:00")
	elapsed := tend.Sub(tbegin)
	zawal := (elapsed/2)
	zawal -= 15*time.Minute
	zawalBegin := tbegin.Add(zawal)
	zawalEnd := zawalBegin.Add(30*time.Minute)
	fmt.Println("from:", zawalBegin.Format("15:04"), " - to:", zawalEnd.Format("15:04")  )
}