package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/Cainaz/test-go/typing/typer"
)

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	wordsFilePath := fmt.Sprintf("%s/data/words", filepath.Dir(filename))

	typer, err := typer.New(wordsFilePath)
	if err != nil {
		log.Fatal(err)
	}

	typer.Run()
}
