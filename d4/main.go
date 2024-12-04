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
