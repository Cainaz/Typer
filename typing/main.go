package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/Cainaz/test-go/typing/loader"
)

func main() {
	// Guarantee of a random number
	rand.Seed(time.Now().UnixNano())

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	wordsFilePath := fmt.Sprintf("%s/data/words", filepath.Dir(filename))

	wordList, err := loader.LoadWordList(wordsFilePath)
	if err != nil {
		panic("Error loading word list")
	}

	for {
		phraseString := loader.GenerateString(wordList)
		fmt.Printf("%s \n", phraseString)

		var inputString string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			inputString = scanner.Text()
		}

		if inputString != phraseString {
			fmt.Printf("'%s' is different from '%s'\n", inputString, phraseString)
			os.Exit(0)
		} else {
			fmt.Printf("You did it!\n")
		}
	}

}
