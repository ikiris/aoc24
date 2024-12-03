package main

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

func p1(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)

	r1 := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

	var ttl int64

	for s.Scan() {
		r := r1.FindAllStringSubmatch(s.Text(), -1) // Lazy way. lets hope.
		for _, v := range r {
			m1, err := strconv.ParseInt(v[1], 10, 64)
			if err != nil {
				return 0, err
			}

			m2, err := strconv.ParseInt(v[2], 10, 64)
			if err != nil {
				return 0, err
			}

			ttl += m1 * m2
		}
	}

	return ttl, nil
}

func p2(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)

	r1 := regexp.MustCompile(`(mul\(([0-9]+),([0-9]+)\)|do\(\)|don't\(\))`) // yes i should have used a lexer parser, but i wanted to see if i could be lazy

	var ttl int64

	enabled := true

	for s.Scan() {
		r := r1.FindAllStringSubmatch(s.Text(), -1) // Lazy way. lets hope.
		for _, v := range r {
			if v[0] == "do()" {
				enabled = true

				continue
			}

			if v[0] == "don't()" {
				enabled = false

				continue
			}

			if !enabled {
				continue
			}

			m1, err := strconv.ParseInt(v[2], 10, 64)
			if err != nil {
				return 0, err
			}

			m2, err := strconv.ParseInt(v[3], 10, 64)
			if err != nil {
				return 0, err
			}

			ttl += m1 * m2
		}
	}

	return ttl, nil
}
