package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"unicode/utf8"
)

const (
	background = 0xe9
)

type Symbol []string

func init() {
	termbox.SetOutputMode(termbox.Output256)
}

func (s Symbol) width() int {
	return utf8.RuneCountInString(s[0])
}

func (s Symbol) height() int {
	return len(s)
}

type Text []Symbol

func (t Text) width() int {
	w := 0
	for _, s := range t {
		w += s.width()
	}
	return w
}

func (t Text) height() int {
	return len(t[0])
}

func toText(str string) Text {
	symbols := make(Text, 0)
	for _, r := range str {
		if s, ok := defaultFont[r]; ok {
			symbols = append(symbols, s)
		}
	}
	return symbols
}

type Font map[rune]Symbol

func echo(s Symbol, startX, startY int, fg termbox.Attribute) {
	x, y := startX, startY
	for _, line := range s {
		for _, r := range line {
			termbox.SetCell(x, y, r, fg, background)
			x++
		}
		x = startX
		y++
	}
}

func clear() {
	err := termbox.Clear(background, background)
	if err != nil {
		panic(err)
	}
}

func flush() {
	err := termbox.Flush()
	if err != nil {
		panic(err)
	}
}

func stderr(s string, a ...interface{}) {
	_, err := fmt.Fprintf(os.Stderr, s, a...)
	if err != nil {
		panic(err)
	}
}
