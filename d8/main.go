package main

import (
	"bufio"
	"io"
	"math"
	"unicode"
)

type key struct {
	row, col int
}

type mapArr [][]rune

type antenna struct {
	pos  key
	tune rune
}

func readMap(r io.Reader) (mapArr, map[rune][]antenna) {
	s := bufio.NewScanner(r)

	row := 0

	m := [][]rune{}

	atMap := make(map[rune][]antenna)

	for s.Scan() {
		m = append(m, []rune{})

		for col, v := range s.Text() {
			m[row] = append(m[row], v)

			if !isAntenna(v) {
				continue
			}

			l, ok := atMap[v]
			if !ok {
				l = []antenna{}
			}

			l = append(l, antenna{key{row, col}, v})

			atMap[v] = l
		}

		row++
	}

	return m, atMap
}

func p1(r io.Reader) (int64, error) {
	m, atMap := readMap(r)

	seen := make(map[key]struct{})

	for _, freq := range atMap {
		for a1i, a1 := range freq[:len(freq)-1] {
			for _, a2 := range freq[a1i+1:] {
				for _, n := range aNodes(a1, a2) {
					if n.row < 0 || n.row >= len(m) {
						continue
					}

					if n.col < 0 || n.col >= len(m[0]) {
						continue
					}

					seen[n] = struct{}{}
				}
			}
		}
	}

	return int64(len(seen)), nil
}

func aNodes(a1, a2 antenna) []key {
	rd, cd := a2.pos.row-a1.pos.row, a2.pos.col-a1.pos.col

	rd, cd = int(math.Abs(float64(rd))), int(math.Abs(float64(cd)))

	if a1.pos.row > a2.pos.row {
		rd = -rd
	}

	if a1.pos.col > a2.pos.col {
		cd = -cd
	}

	return []key{
		{a1.pos.row - rd, a1.pos.col - cd},
		{a2.pos.row + rd, a2.pos.col + cd},
	}
}

func aNodesRepeat(a1, a2 antenna, maxRow, maxCol int) []key {
	rd, cd := a2.pos.row-a1.pos.row, a2.pos.col-a1.pos.col

	rd, cd = int(math.Abs(float64(rd))), int(math.Abs(float64(cd)))

	if a1.pos.row > a2.pos.row {
		rd = -rd
	}

	if a1.pos.col > a2.pos.col {
		cd = -cd
	}

	r := []key{}

	lfuse, hifuse := false, false
	for i := 0; !lfuse || !hifuse; i++ {
		r1, c1 := a1.pos.row-(rd*i), a1.pos.col-(cd*i)
		r2, c2 := a2.pos.row+(rd*i), a2.pos.col+(cd*i)

		if r1 >= 0 && r1 <= maxRow && c1 >= 0 && c1 <= maxCol {
			r = append(r, key{r1, c1})
		} else {
			lfuse = true
		}

		if r2 >= 0 && r2 <= maxRow && c2 >= 0 && c2 <= maxCol {
			r = append(r, key{r2, c2})
		} else {
			hifuse = true
		}
	}

	return r
}

func isAntenna(r rune) bool {
	switch {
	case r == '.':
		return false
	case unicode.IsLetter(r):
		return true
	case unicode.IsDigit(r):
		return true
	default:
		return false
	}
}

func p2(r io.Reader) (int64, error) {
	m, atMap := readMap(r)

	seen := make(map[key]struct{})

	maxRow, maxCol := len(m), len(m[0])

	for _, freq := range atMap {
		for a1i, a1 := range freq[:len(freq)-1] {
			for _, a2 := range freq[a1i+1:] {
				for _, n := range aNodesRepeat(a1, a2, maxRow, maxCol) {
					if n.row < 0 || n.row >= maxRow {
						continue
					}

					if n.col < 0 || n.col >= maxCol {
						continue
					}

					seen[n] = struct{}{}
				}
			}
		}
	}

	return int64(len(seen)), nil
}
