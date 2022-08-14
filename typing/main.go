package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/Cainaz/test-go/typing/typer"
)

func main() {
	// Guarantee of a random number
	rand.Seed(time.Now().UnixNano())

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	wordsFilePath := fmt.Sprintf("%s/data/words", filepath.Dir(filename))

	wordList, err := typer.LoadWordList(wordsFilePath)
	if err != nil {
		panic("Error loading word list")
	}

	typer.FlushScreen()

	for {
		phraseString := typer.GenerateString(wordList)
		fmt.Printf("%s \n", phraseString)

		var inputString string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			inputString = scanner.Text()
		}

		if inputString != phraseString {
			typer.FlushScreen()
			fmt.Printf("###### Misstyping #######\n'%s' is different from\n'%s'\n", inputString, phraseString)

			// Print report and ask if you wanna keep playing
			os.Exit(0)
		} else {
			typer.FlushScreen()
		}
	}

}
