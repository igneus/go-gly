package gly

type Score struct {
	lyrics Lyrics
	music  Music
}

type Lyrics struct{}

func (l Lyrics) isEmpty() bool { return true }

type Music struct{}

func (m Music) isEmpty() bool { return true }
