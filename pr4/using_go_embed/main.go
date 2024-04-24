package main

import (
	_ "embed"
)

//go:embed text_file.txt
var readme string

func main() {
	print(readme)
}
