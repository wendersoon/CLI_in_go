package main

import (
	"command-line/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Generate()
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
