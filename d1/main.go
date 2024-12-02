package main

import (
	"bufio"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"
)

func p1(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)

	var a1, a2 []int64

	for s.Scan() {
		t := strings.Split(s.Text(), "   ")

		n1, err := strconv.ParseInt(t[0], 10, 64)
		if err != nil {
			return 0, err
		}

		n2, err := strconv.ParseInt(t[1], 10, 64)
		if err != nil {
			return 0, err
		}

		a1 = append(a1, n1)
		a2 = append(a2, n2)
	}

	slices.Sort(a1)
	slices.Sort(a2)

	var ttl int64
	for i := range a1 {
		ttl += int64(math.Abs(float64(a1[i] - a2[i])))
	}

	return ttl, nil
}

func p2(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)

	var a1 []int64

	nh := make(map[int64]int)

	for s.Scan() {
		t := strings.Split(s.Text(), "   ")

		n1, err := strconv.ParseInt(t[0], 10, 64)
		if err != nil {
			return 0, err
		}

		n2, err := strconv.ParseInt(t[1], 10, 64)
		if err != nil {
			return 0, err
		}

		a1 = append(a1, n1)
		nh[n2]++
	}

	var ttl int64
	for _, v := range a1 {
		ttl += v * int64(nh[v])
	}

	return ttl, nil
}
