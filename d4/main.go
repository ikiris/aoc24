package main

import (
	"bufio"
	"io"

	"github.com/ikiris/aoc24/generic"
)

type key struct {
	row, col int
}

type track struct {
	row, col int
	trie     *generic.Trie
}

type tracker struct {
	tracks map[key][]track
	trie   *generic.Trie
}

func NewTracker(t *generic.Trie) *tracker {
	return &tracker{
		tracks: make(map[key][]track),
		trie:   t,
	}
}

func InitTrie(list []string) *generic.Trie {
	t := generic.NewTrie()
	for _, v := range list {
		t.Add(v)
	}

	return t
}

func p1(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)

	wordlist := []string{
		"XMAS",
		"SAMX",
	}

	row := 0

	var ttl int64

	t := NewTracker(InitTrie(wordlist))

	for s.Scan() {
		l := s.Text()
		for col, v := range l {
			res := t.Get(v, row, col)
			if res == 0 {
				continue
			}

			ttl += res
		}

		row++
	}

	return ttl, nil
}

func (t *tracker) Get(r rune, row, col int) int64 {
	k := key{row, col}

	tracks, ok := t.tracks[k]
	if !ok {
		tracks = []track{}
	}

	tracks = append(tracks, track{trie: t.trie})

	var ret int64

	for _, v := range tracks {
		tr, ok := v.trie.Kids[r]
		if !ok {
			continue
		}

		if len(tr.Kids) > 0 {
			t.appendTracks(tr, row, col, v)

			continue
		}

		ret++
	}

	delete(t.tracks, k)

	return ret
}

func (t *tracker) appendTracks(tr *generic.Trie, row, col int, v track) {
	if v.row == 0 && v.col == 0 {
		t.appendTrack(tr, row, col, track{row: 0, col: 1})  // Horizontal
		t.appendTrack(tr, row, col, track{row: 1, col: 0})  // Vertical
		t.appendTrack(tr, row, col, track{row: 1, col: 1})  // \
		t.appendTrack(tr, row, col, track{row: 1, col: -1}) // /

		return
	}

	t.appendTrack(tr, row, col, v)
}

func (t *tracker) appendTrack(tr *generic.Trie, row, col int, v track) {
	k := key{row + v.row, col + v.col}

	tracks, ok := t.tracks[k]
	if !ok {
		tracks = []track{}
	}

	t.tracks[k] = append(tracks, track{
		row:  v.row,
		col:  v.col,
		trie: tr,
	})
}

type xMas struct {
	pos key
	v   rune
}

type xPattern []xMas

func patternXGen(words []string) []xPattern {
	base := []key{
		{0, 0},
		{-1, -1},
		{-2, -2},
		{0, -2},
		{-2, 0},
	}

	r := []xPattern{}

	for w1, w := range words {
		fac := w1 * 2

		r = append(r, make([]xMas, len(base)), make([]xMas, len(base)))
		for i, l := range w {
			r[0+fac][i].pos = base[i]
			r[1+fac][i].pos = base[i]

			r[0+fac][i].v = l
			r[1+fac][i].v = l
		}

		for i2, w2 := range words {
			w2l := []rune(w2)

			r[(w1*2)+i2][3].pos = base[3]
			r[(w1*2)+i2][3].v = w2l[0]

			r[(w1*2)+i2][4].pos = base[4]
			r[(w1*2)+i2][4].v = w2l[2]
		}
	}

	return r
}

func p2(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)

	words := []string{"MAS", "SAM"}

	arr := [4][]rune{}
	modVal := len(arr)

	patterns := patternXGen(words)

	checkX := func(row, col int) bool {
	continueB:
		for _, xp := range patterns {
			for _, xM := range xp {
				pos := xM.pos
				pos.row += row
				pos.col += col

				if pos.row < 0 || pos.col < 0 {
					continue continueB
				}

				pos.row = pos.row % modVal

				if arr[pos.row][pos.col] != xM.v {
					continue continueB
				}
			}

			return true
		}

		return false
	}

	var ttl int64

	row := 0

	for s.Scan() {
		rm := row % modVal
		arr[rm] = []rune{}

		l := s.Text()
		for col, v := range l {
			arr[rm] = append(arr[rm], v)

			if checkX(row, col) {
				ttl++
			}
		}

		row++
	}

	return ttl, nil
}
