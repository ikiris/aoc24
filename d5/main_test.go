package main

import (
	"io"
	"testing"

	"github.com/ikiris/aoc24/generic/testgeneric"
)

func TestD5P1(t *testing.T) {
	tests := []struct {
		name    string
		data    io.Reader
		want    int64
		wantErr bool
	}{
		{
			"basic",
			testgeneric.GetHandle(t, "testdata/basic.txt"),
			143,
			false,
		},
		{
			"d5p1",
			testgeneric.GetHandle(t, "testdata/d5p1.txt"),
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

func TestD5P2(t *testing.T) {
	tests := []struct {
		name    string
		data    io.Reader
		want    int64
		wantErr bool
	}{
		{
			"basic",
			testgeneric.GetHandle(t, "testdata/basic.txt"),
			123,
			false,
		},
		{
			"d5p2",
			testgeneric.GetHandle(t, "testdata/d5p1.txt"),
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
