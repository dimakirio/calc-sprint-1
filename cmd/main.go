package main

import (
	"fmt"

	"github.com/dimakirio/calc-sprint-1/internal"
)

func main() {
	app := application.New()
	fmt.Println("RunServer")
	// app.Run()
	app.RunServer()
}
