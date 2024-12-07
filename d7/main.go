package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func p1(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)

	var ttl int64

	for s.Scan() {
		t1 := strings.Split(s.Text(), ":")

		desired, err := strconv.ParseInt(t1[0], 10, 64)
		if err != nil {
			return 0, err
		}

		textNums := strings.Split(t1[1], " ")

		nums := []int64{}

		for _, v := range textNums {
			if v == "" {
				continue
			}

			n, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("failed to parse: %w", err)
			}

			nums = append(nums, n)
		}

		if !tryget(desired, nums) {
			continue
		}

		ttl += desired
	}

	return ttl, nil
}

func tryget(desired int64, nums []int64) bool {
	if len(nums) == 1 {
		return nums[0] == desired
	}

	numsb := []int64{nums[0] + nums[1]}
	numsb = append(numsb, nums[2:]...)

	if tryget(desired, numsb) {
		return true
	}

	numsc := []int64{nums[0] * nums[1]}
	numsc = append(numsc, nums[2:]...)

	return tryget(desired, numsc)
}
