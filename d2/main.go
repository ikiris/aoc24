package main

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

func p1(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)

	var ttl int64
	for s.Scan() {
		l := strings.Split(s.Text(), " ")

		if testline(l) {
			ttl++
		}
	}

	return ttl, nil
}

func p2(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)

	var ttl int64
	for s.Scan() {
		l := strings.Split(s.Text(), " ")

		if testlineB(l) {
			ttl++
		}
	}

	return ttl, nil
}

func testlineB(l []string) bool {
	if testline(l) {
		return true
	}

	for i := range l {
		tl := []string{}
		tl = append(tl, l[:i]...)
		tl = append(tl, l[i+1:]...)
		if testline(tl) {
			return true
		}
	}

	return false
}

func testline(l []string) bool {
	var last int64
	n := []int64{}
	for i, v := range l {
		num, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}

		n = append(n, num)
		if i == 0 {
			continue
		}

		diff := n[i] - n[i-1]
		if adiff := math.Abs(float64(diff)); adiff < 1 || adiff > 3 {
			return false
		}

		if last == 0 {
			last = diff
			continue
		}

		if math.Signbit(float64(diff)) != math.Signbit(float64(last)) {
			return false
		}

		last = diff
	}

	return true
}
