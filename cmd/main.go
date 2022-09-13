package main

import (
	"flag"
	"gena/internal/generator"
	"log"
)

func main() {
	dest := flag.String("dest", "dest", "path to gen")

	log.Println("dest: ", *dest)

	gena := generator.New(*dest)

	if err := gena.Gen(); err != nil {
		log.Fatalln(err)
	}
}
