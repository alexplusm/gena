package main

import (
	"flag"
	"fmt"
	"log"

	"gena/internal/generator"
	"github.com/manifoldco/promptui"
)

func main() {
	dest := flag.String("dest", "dest", "path to gen")

	p := promptui.Select{
		Label: "choose go version",
		Items: []string{
			"1.8",
			"1.9",
		},
	}

	i, val, err := p.Run()

	fmt.Println("val, err", i, val, err)

	log.Println("dest: ", *dest)

	gena := generator.New(*dest)

	if err := gena.Gen(); err != nil {
		log.Fatalln(err)
	}
}
