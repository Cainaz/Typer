package typer

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Cainaz/Go/utils"
)

type Typer struct {
	wordsPerLine int
	start        time.Time
	wordList     []string
	duration     float64
	wordCount    int
	inputString  string
	line         string
}

func New(filePath string) (*Typer, error) {
	t := &Typer{wordsPerLine: 5}

	loadedWordList, err := t.loadWordList(filePath)
	if err != nil {
		return nil, err
	}

	t.wordList = loadedWordList
	return t, nil
}

func (t *Typer) init() {
	t.flushScreen()
	t.start = time.Now()
	t.wordCount = 0
	t.duration = 0
}

func (t *Typer) showSessionReport() {
	duration := time.Since(t.start).Minutes()
	fmt.Println("---------------Session-Result--------------")
	fmt.Printf("Duration: %.2f m\n", duration)
	fmt.Printf("Correctly typed words: %d\n", t.wordCount)
	fmt.Printf("Typing speed: %.2f W/m\n", float64(t.wordCount)/duration)
	fmt.Println("-------------------------------------------")
}

func (t *Typer) scanInput() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		t.inputString = scanner.Text()
	}
}

func (t *Typer) Run() {
	t.init()
	for {
		t.nextLine(t.wordList)
		t.showLine()
		t.scanInput()
		t.countCorrectWords()
		t.flushScreen()

		if !t.IsCorrectlyTyped() {
			t.showMisstyping()
			t.showSessionReport()
			t.close()
		}
	}
}

// Load a word list from file
func (t Typer) loadWordList(filePath string) ([]string, error) {
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

// Gets next string from wordlist
func (t *Typer) nextLine(wordList []string) {
	var lineWords []string
	for i := 0; i < t.wordsPerLine; i++ {
		randomIndex := rand.Intn(len(wordList))
		lineWords = append(lineWords, wordList[randomIndex])
	}
	t.line = strings.Join(lineWords, " ")
}

// Flush terminal
func (t *Typer) flushScreen() {
	fmt.Print("\033[H\033[2J")
}

func (t *Typer) close() {
	os.Exit(0)
}

func (t *Typer) countCorrectWords() {
	typedWords := strings.Split(t.inputString, " ")
	lineWords := strings.Split(t.line, " ")

	// Count correct words
	for _, word := range typedWords {
		if utils.StringInList(word, lineWords) {
			t.wordCount += 1
		}
	}
}

func (t *Typer) showMisstyping() {
	fmt.Println("############### Misstyping ################")
	fmt.Printf("Phrase: %s\n", t.line)
	fmt.Printf("Typed : %s\n", t.inputString)
	fmt.Printf("###########################################\n\n")
}

func (t *Typer) IsCorrectlyTyped() bool {
	return t.inputString == t.line
}

func (t *Typer) showLine() {
	fmt.Printf("%s\n", t.line)
}
