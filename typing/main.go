package main

import (
	"fmt"
	"log"
	"math/rand"
	"path/filepath"
	"runtime"
	"time"

	"github.com/Cainaz/test-go/typing/typer"
)

func main() {
	// Guarantee of a "real" random number
	rand.Seed(time.Now().UnixNano())

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
