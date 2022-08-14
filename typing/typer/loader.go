package typer

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func LoadWordList(filePath string) ([]string, error) {
	var wordList []string

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordList, nil
}

func GenerateString(wordList []string) string {
	var phrase []string
	for i := 0; i < 5; i++ {
		randomIndex := rand.Intn(len(wordList))
		phrase = append(phrase, wordList[randomIndex])
	}
	phraseString := strings.Join(phrase, " ")
	return phraseString
}

func FlushScreen() {
	fmt.Print("\033[H\033[2J")
}
