package main

import (
	"fmt"
	"github.com/ssalvatori/fondosMutuosGo/crawler"
	"time"
)

func main() {
	fmt.Printf("Fondos Mutuos")
	fmt.Println("Starting ** ", time.Now().String())
	c := crawler.NewCrawler()
	c.Run()
}
