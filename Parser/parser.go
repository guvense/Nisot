package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Parser struct {
	parser   rune
	filename string
}

var stringArray []string

func Conf(parser rune, filename string) *Parser {

	p := Parser{
		parser:   parser,
		filename: filename,
	}

	stringArray = getStringArray(p.filename)
	return &p
}

func (p *Parser) ParseToMap() map[string]string {

	maper := make(map[string]string)

	for i, _ := range stringArray {

		s := strings.Split(stringArray[i], string(p.parser))
		maper[s[0]] = s[1]

	}

	return maper

}

func (p *Parser) Replace(filename string) {

	file := p.filename

	if filename != "" {
		file = filename
	}

	myFİle, err := os.Create(file)

	printAndExit(err)

	defer myFİle.Close()
	for i, _ := range stringArray {

		s := strings.Split(stringArray[i], string(p.parser))
		text := s[1] + string(p.parser) + s[0] + "\n"
		fmt.Fprint(myFİle, text)

	}

}

func getStringArray(filename string) []string {

	var values []string

	file, err := os.Open(filename)

	printAndExit(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		values = append(values, scanner.Text())

	}
	err = file.Close()

	printAndExit(err)
	printAndExit(scanner.Err())
	return values

}

func printAndExit(err error) {

	if err != nil {
		log.Fatal(err)
	}
}
