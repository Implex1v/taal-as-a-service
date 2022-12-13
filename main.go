package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"math/rand"
	"os"
	"strings"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.GET("/words", getWords)

	err := e.Start(":8080")
	if err != nil {
		return
	}
}

func getWords(c echo.Context) error {
	words, err := readWords()

	if err != nil {
		return err
	}

	return c.JSON(200, words)
}

func readWords() (*[]string, error) {
	line, err := readFile("taal.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(*line, "\n")
	length := len(lines)

	println(length)

	var words []string

	for i := 0; i < 10; i += 1 {
		words = append(words, lines[rand.Intn(length)])
	}

	return &words, nil
}

func readFile(fileName string) (*string, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	line := string(data)
	return &line, nil
}
