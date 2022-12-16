package main

import (
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetWords(c echo.Context) error {
	var count = min(parseQueryAsInt(c, "count", 10), 1000)

	words, err := readWords(count)
	if err != nil {
		return err
	}

	return c.JSON(200, words)
}

func parseQueryAsInt(c echo.Context, name string, defaultValue int) int {
	var str = c.QueryParam(name)
	if str == "" {
		return defaultValue
	}

	num, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}

	return num
}

func GetHealth(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func readWords(count int) (*[]string, error) {
	line, err := readFile("taal.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(*line, "\n")
	length := len(lines)

	var words []string
	for i := 0; i < count; i += 1 {
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

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
