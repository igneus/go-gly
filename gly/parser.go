package gly

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Parse(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parseLine(scanner.Text())
	}
}

func parseLine(line string) {
	if isEmpty(line) {
		// do nothing
	} else if isNewScore(line) {
		//
	} else if isHeaderStart(line) {
		//
	} else if isBlockStart(line) {
		//
	} else if isExplicitLyrics(line) {
		parseLyrics(line)
	} else if isExplicitMusic(line) {
		parseMusic(line)
	} else if inSpecialBlock(line) {
		//
	} else if isMarkup(line) {
		parseMarkup(line)
	} else if isHeader(line) {
		parseHeader(line)
	} else if isLyrics(line) {
		parseLyrics(line)
	} else {
		parseDefault(line)
	}
}

// detection of line meaning

func isEmpty(line string) bool          { return false }
func isNewScore(line string) bool       { return false }
func isHeaderStart(line string) bool    { return false }
func isBlockStart(line string) bool     { return false }
func isExplicitLyrics(line string) bool { return false }
func isExplicitMusic(line string) bool  { return false }
func isMarkup(line string) bool         { return false }
func isHeader(line string) bool         { return false }

func isLyrics(line string) bool {
	return (!strings.ContainsAny(line, "[]")) &&
		(strings.Contains(line, " -- ") ||
			containsUnmusicalChars(line))
}

func containsUnmusicalChars(line string) bool {
	musicalCharacters := " abcdefghijklmorsvwxz"
	reUninteresting := regexp.MustCompile("[\\W\\d_]+")

	// delete numbers, interpunction etc.
	interesting := reUninteresting.ReplaceAllString(line, "")
	// delete all word characters used to encode music
	return len(strings.Trim(strings.ToLower(interesting), musicalCharacters)) > 0
}

// current parser state

func inSpecialBlock(line string) bool { return false }

// line processing according to the detected meaning

func parseMusic(line string)   {}
func parseLyrics(line string)  { fmt.Println(line) }
func parseHeader(line string)  {}
func parseMarkup(line string)  {}
func parseDefault(line string) {}
