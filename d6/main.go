package main

import (
	"bufio"
	"io"
)

type guard struct {
	row, col, dir int
	nl            *key
}

var rl = map[rune]int{
	'^': 0,
	'>': 1,
	'v': 2,
	'<': 3,
}

type key struct {
	row, col int
}

var da = []key{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

type mapArr [][]rune

func p1(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)

	row := 0

	var g *guard

	m := [][]rune{}

	for s.Scan() {
		m = append(m, []rune{})

		for i, v := range s.Text() {
			m[row] = append(m[row], v)

			if d, ok := rl[v]; ok {
				g = &guard{
					row: row,
					col: i,
					dir: d,
				}

				m[row][i] = 'X'
			}
		}

		row++
	}

	var ttl int64

	for g.move(m) {
		v := m[g.row][g.col]
		if v == 'X' {
			continue
		}

		ttl++
		m[g.row][g.col] = 'X'
	}

	return ttl + 1, nil
}

func (g *guard) move(m mapArr) bool {
	g.forward()

	if g.row >= len(m) || g.row < 0 {
		return false
	}

	if g.col >= len(m[g.row]) || g.col < 0 {
		return false
	}

	v := m[g.row][g.col]

	if g.nl != nil && g.nl.row == g.row && g.nl.col == g.col {
		g.back()
		g.turn()

		return true
	}

	switch v {
	case '#':
		g.back()
		g.turn()

		return true
	default:
		return true
	}
}

func (g *guard) forward() {
	k := da[g.dir]

	g.row += k.row
	g.col += k.col
}

func (g *guard) back() {
	k := da[g.dir]

	g.row -= k.row
	g.col -= k.col
}

func (g *guard) turn() {
	g.dir++
	g.dir = g.dir % len(da)
}

func p2(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)

	row := 0

	var g *guard

	m := [][]rune{}

	for s.Scan() {
		m = append(m, []rune{})

		for i, v := range s.Text() {
			m[row] = append(m[row], v)

			if d, ok := rl[v]; ok {
				g = &guard{
					row: row,
					col: i,
					dir: d,
				}

				m[row][i] = 'X'
			}
		}

		row++
	}

	g1 := &guard{row: g.row, col: g.col, dir: g.dir}

	seenC := make(map[key]struct{})
	// Initial pass
	for g1.move(m) {
		seenC[key{g1.row, g1.col}] = struct{}{}
	}

	var ttl int64

contB:
	for k := range seenC {
		seen1 := make(map[dkey]struct{})

		g2 := &guard{row: g.row, col: g.col, dir: g.dir}

		g2.null(k)

		for g2.move(m) {
			if _, ok := seen1[dkey{g2.row, g2.col, g2.dir}]; ok {
				// loop detected!
				ttl++

				continue contB
			}

			seen1[dkey{g2.row, g2.col, g2.dir}] = struct{}{}
		}
	}

	return ttl, nil
}

type dkey struct {
	row, col, dir int
}

func (g *guard) null(k key) {
	g.nl = &k
}
