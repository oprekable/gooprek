package main

import (
	"embed"
	"os"
)

//go:embed all:embeds
var embedFS embed.FS
var exitFunc = os.Exit

func main() {

}
