package main

import (
	"embed"
	"io/fs"

	"github.com/jonathongardner/wegyb/cli"

	log "github.com/sirupsen/logrus"
)

//go:embed public/*
var uiFolder embed.FS

func main() {
	ui, _ := fs.Sub(uiFolder, "public")
	err := cli.Run(ui)
	if err != nil {
		log.Fatal(err)
	}
}
