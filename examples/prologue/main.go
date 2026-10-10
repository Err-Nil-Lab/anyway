package main

import (
	_ "embed"
	"log"

	"github.com/Err-Nil-Lab/anyway"
)

//go:embed prologue.yaml
var prologue []byte

func main() {
	scene, err := anyway.Parse(prologue)
	if err != nil {
		log.Fatal(err)
	}
	if err := anyway.Run(scene, anyway.Config{Title: "prologue"}); err != nil {
		log.Fatal(err)
	}
}
