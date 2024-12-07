package main

import (
	"bufio"
	"io"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type rules struct {
	m  map[int][]int
	mm map[int]map[int]struct{}
}

func (r *rules) addRule(first, last int) {
	_, ok := r.m[first]
	if !ok {
		r.m[first] = []int{}
	}

	r.m[first] = append(r.m[first], last)

	_, ok = r.mm[first]
	if !ok {
		r.mm[first] = make(map[int]struct{})
	}

	r.mm[first][last] = struct{}{}
}

func (r *rules) before(n int) []int {
	ret, ok := r.m[n]
	if !ok {
		return nil
	}

	return ret
}

func (r *rules) isbefore(n int, n2 int) bool {
	ret, ok := r.mm[n]
	if !ok {
		return false
	}

	_, ok = ret[n2]

	return ok
}

func newRules() *rules {
	return &rules{
		make(map[int][]int),
		make(map[int]map[int]struct{}),
	}
}

func p1(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)

	var ttl int64

	rh := newRules()

	pOrder := regexp.MustCompile(`^([0-9]+)\|([0-9]+)$`)

continueB:
	for s.Scan() {
		l := s.Text()

		rOrder := pOrder.FindStringSubmatch(l)
		if len(rOrder) > 0 {
			first, err := strconv.ParseInt(rOrder[1], 10, 64)
			if err != nil {
				return 0, err
			}

			last, err := strconv.ParseInt(rOrder[2], 10, 64)
			if err != nil {
				return 0, err
			}

			rh.addRule(int(first), int(last))

			continue
		}

		n := strings.Split(l, ",")
		if len(n) == 0 || n[0] == "" {
			continue
		}

		seen := make(map[int]struct{})
		list := []int{}

		for _, num := range n {
			ni, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				return 0, err
			}

			seen[int(ni)] = struct{}{}
			b4 := rh.before(int(ni))
			for _, chk := range b4 {
				if _, exists := seen[chk]; exists {
					continue continueB
				}
			}

			list = append(list, int(ni))
		}

		m := list[len(list)/2]

		ttl += int64(m)
	}

	return ttl, nil
}

func p2(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)

	var ttl int64

	rh := newRules()

	pOrder := regexp.MustCompile(`^([0-9]+)\|([0-9]+)$`)

	for s.Scan() {
		l := s.Text()

		rOrder := pOrder.FindStringSubmatch(l)
		if len(rOrder) > 0 {
			first, err := strconv.ParseInt(rOrder[1], 10, 64)
			if err != nil {
				return 0, err
			}

			last, err := strconv.ParseInt(rOrder[2], 10, 64)
			if err != nil {
				return 0, err
			}

			rh.addRule(int(first), int(last))

			continue
		}

		n := strings.Split(l, ",")
		if len(n) == 0 || n[0] == "" {
			continue
		}

		list := []int{}

		for _, num := range n {
			ni, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				return 0, err
			}

			list = append(list, int(ni))
		}

		fuse := false

		slices.SortFunc(list, func(a, b int) int {
			if rh.isbefore(a, b) {
				fuse = true

				return -1
			}

			if rh.isbefore(b, a) {
				return 1
			}

			return 0
		})

		m := list[len(list)/2]

		if !fuse {
			continue
		}

		ttl += int64(m)
	}

	return ttl, nil
}
