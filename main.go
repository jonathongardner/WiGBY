package main

import (
	"fmt"
  "embed"
	"io/fs"

	"github.com/jonathongardner/wegyb/wegyb"
)

//go:embed public/*
var uiFolder embed.FS

func main() {
	fmt.Println("Starting...")
	ui, _ := fs.Sub(uiFolder, "public")
	// Block forever
	wegyb.ListenAndServe(":3000", ui)
	fmt.Println("...Closing")

}
