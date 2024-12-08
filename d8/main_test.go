package main

import (
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ikiris/aoc24/generic/testgeneric"
)

func TestD8P1(t *testing.T) {
	tests := []struct {
		name    string
		data    io.Reader
		want    int64
		wantErr bool
	}{
		{
			"basic",
			testgeneric.GetHandle(t, "testdata/basic.txt"),
			14,
			false,
		},
		{
			"d8p1",
			testgeneric.GetHandle(t, "testdata/d8p1.txt"),
			0,
			false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := p1(tc.data)
			if (err != nil) != tc.wantErr {
				t.Fatalf("func %s goterr: %v wanted: %v", tc.name, err, tc.wantErr)
			}

			if got != tc.want {
				t.Fatalf("func got: %d, want: %d", got, tc.want)
			}
		})
	}
}

func TestANodes(t *testing.T) {
	tests := []struct {
		name   string
		a1, a2 antenna
		want   []key
	}{
		{
			"Basic",
			antenna{
				key{2, 2},
				'a',
			},
			antenna{
				key{4, 3},
				'a',
			},
			[]key{
				{0, 1},
				{6, 4},
			},
		},
		{
			"updiag",
			antenna{
				key{1, 8},
				'a',
			},
			antenna{
				key{2, 5},
				'a',
			},
			[]key{
				{0, 11},
				{3, 2},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := aNodes(tc.a1, tc.a2)

			if diff := cmp.Diff(tc.want, got, cmp.Comparer(func(x, y key) bool {
				return x.row == y.row && x.col == y.col
			})); diff != "" {
				t.Errorf("%s mismatch (-want +got):\n%s", tc.name, diff)
			}
		})
	}
}

func TestD8P2(t *testing.T) {
	tests := []struct {
		name    string
		data    io.Reader
		want    int64
		wantErr bool
	}{
		{
			"basic",
			testgeneric.GetHandle(t, "testdata/basic.txt"),
			34,
			false,
		},
		{
			"d8p2",
			testgeneric.GetHandle(t, "testdata/d8p1.txt"),
			0,
			false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := p2(tc.data)
			if (err != nil) != tc.wantErr {
				t.Fatalf("func %s goterr: %v wanted: %v", tc.name, err, tc.wantErr)
			}

			if got != tc.want {
				t.Fatalf("func got: %d, want: %d", got, tc.want)
			}
		})
	}
}
