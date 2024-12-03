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
