package gly

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Parser struct {
	// state switch influencing how parser interprets the input
	currentBlock int8

	// the score currently being built
	currentScore *Score
}

const bSCORE, bHEADER, bMUSIC, bLYRICS, bMARKUP int8 = 0, 1, 2, 3, 4

func (p Parser) Parse(file *os.File) {
	p.currentBlock = bSCORE
	p.currentScore = new(Score)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p.parseLine(stripComment(scanner.Text()))
	}
}

func (p Parser) parseLine(line string) {
	if isEmpty(line) && p.currentBlock != bMARKUP {
		// do nothing
	} else if isNewScore(line) {
		p.currentBlock = bSCORE
	} else if isHeaderStart(line) {
		p.currentBlock = bHEADER
	} else if isLyricBlockStart(line) {
		p.currentBlock = bLYRICS
	} else if isMusicBlockStart(line) {
		p.currentBlock = bMUSIC
	} else if isExplicitLyrics(line) {
		p.parseLyrics(line)
	} else if isExplicitMusic(line) {
		p.parseMusic(line)
	} else if p.inSpecialBlock() {
		p.parseInBlock(line)
	} else if isMarkup(line) {
		p.parseMarkup(line)
	} else if p.isHeader(line) {
		p.parseHeader(line)
	} else if isLyrics(line) {
		p.parseLyrics(line)
	} else {
		p.parseDefault(line)
	}
}

func (p Parser) parseInBlock(line string) {
	switch p.currentBlock {
	case bHEADER:
		p.parseHeader(line)
	case bMUSIC:
		p.parseMusic(line)
	case bLYRICS:
		p.parseLyrics(line)
	case bMARKUP:
		p.parseMarkup(line)
	}
}

// current parser state
func (p Parser) inSpecialBlock() bool {
	return p.currentBlock != bSCORE
}

func (p Parser) isHeader(line string) bool {
	return p.currentBlock == bSCORE && p.currentScore.lyrics.isEmpty() && p.currentScore.music.isEmpty() && isHeaderField(line)
}

// line processing according to the detected meaning

func (p Parser) parseMusic(line string)   {}
func (p Parser) parseLyrics(line string)  { fmt.Println(line) }
func (p Parser) parseHeader(line string)  {}
func (p Parser) parseMarkup(line string)  {}
func (p Parser) parseDefault(line string) {}

// common regexps
var reEmpty = regexp.MustCompile("^\\s*$")
var reSyllSep = regexp.MustCompile("\\s*--\\s*")
var reHeaderField = regexp.MustCompile("^[\\w_-]+:")
var reComment = regexp.MustCompile("%.*$")

func stripComment(line string) string {
	return reComment.ReplaceAllString(line, "")
}

// detection of line meaning

func isEmpty(line string) bool {
	return reEmpty.MatchString(line)
}

func isNewScore(line string) bool        { return false }
func isHeaderStart(line string) bool     { return false }
func isLyricBlockStart(line string) bool { return false }
func isMusicBlockStart(line string) bool { return false }
func isExplicitLyrics(line string) bool  { return false }
func isExplicitMusic(line string) bool   { return false }
func isMarkup(line string) bool          { return false }

func isLyrics(line string) bool {
	return (!strings.ContainsAny(line, "[]")) &&
		(reSyllSep.MatchString(line) ||
			containsUnmusicalChars(line))
}

func isHeaderField(line string) bool {
	return reHeaderField.MatchString(line)
}

func containsUnmusicalChars(line string) bool {
	musicalCharacters := "abcdefghijklmorsvwxz"
	reUninteresting := regexp.MustCompile("[\\W\\d_]+")

	// delete numbers, interpunction etc.
	interesting := reUninteresting.ReplaceAllString(line, "")
	// delete all word characters used to encode music
	return len(strings.Trim(strings.ToLower(interesting), musicalCharacters)) > 0
}
