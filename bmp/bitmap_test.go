package bmp

import (
	"image"
	"image/png"
	"os"
	"strings"
	"testing"
)

var (
	testdata = "./testdata"
)

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestBitmapFromImage(t *testing.T) {
	gopher := []string{
		"1111111111111000000000000001111111111111",
		"1111111111100000011111111000011111011111",
		"1111001100011111111111110001100000000111",
		"1100000001000001111111001110011001110011",
		"1011110010111110011110111111101101111001",
		"0011000100111111001110111111100110011101",
		"0010001101011111101101001111110110011101",
		"0011001100001111100100000111110111011011",
		"1011011100001111100110000111110111010011",
		"1000011100001111100110000111110111000111",
		"1110111101111111101110011111100111101111",
		"1110111110111111000000011111101111101111",
		"1110111111011110100000100000011111101111",
		"1110111111100000000000001111111111100111",
		"1100111111111110111111101111111111100111",
		"1100111111111110110001101111111111100111",
		"1100111111111111001000011111111111110111",
		"1100111111111111101011111111111111110111",
		"1100111111111111101010011111111111110111",
		"1100111111111111110000111111111111110111",
		"1100111111111111111111111111111111110111",
		"1110111111111111111111111111111111110111",
		"1110111111111111111111111111111111110111",
		"1110111111111111111111111111111111110111",
		"1110111111111111111111111111111111110111",
		"1110011111111111111111111111111111110111",
		"1110011111111111111111111111111111110001",
		"1010011111111111111111111111111111110110",
		"0110011111111111111111111111111111110110",
		"0000111111111111111111111111111111110000",
		"1010111111111111111111111111111111110111",
		"1110111111111111111111111111111111110111",
		"1110111111111111111111111111111111110111",
		"1110111111111111111111111111111111110011",
		"1110111111111111111111111111111111110011",
		"1110111111111111111111111111111111110011",
		"1110111111111111111111111111111111110011",
		"1110111111111111111111111111111111110011",
		"1110111111111111111111111111111111110011",
		"1110111111111111111111111111111111110111",
		"1110111111111111111111111111111111110111",
		"1110111111111111111111111111111111110111",
		"1110111111111111111111111111111111100111",
		"1110111111111111111111111111111111100111",
		"1110011111111111111111111111111111101111",
		"1111011111111111111111111111111111001111",
		"1111001111111111111111111111111111011111",
		"1111101111111111111111111111111110111111",
		"1111100111111111111111111111111100111111",
		"1111110000111111111111111111110010111111",
		"1111101110011111111111111111100111011111",
		"1111101110001111111111111111000011011111",
		"1111100100100000000000000000111101011111",
		"1111101011111111000000011111111100011111"}

	f, err := os.Open(testdata + "/gopher.png")
	check(t, err)
	defer f.Close()

	var (
		img image.Image
		bmp *Bitmap
		exp string
	)
	img, err = png.Decode(f)
	check(t, err)

	bmp = NewFromImage(img)
	exp = strings.Join(gopher, "\n") + "\n"
	if bmp.String() != exp {
		t.Errorf("NewFromImage() expected gopher, didn't have one")
	}
}

func TestBlackImage(t *testing.T) {
	var testTbl = []struct {
		w, h                   int
		minx, miny, maxx, maxy int
	}{
		{2, 2, 0, 0, 1, 1},
		{1, 2, 0, 0, 0, 1},
		{2, 1, 0, 0, 1, 0},
		{4, 2, 0, 0, 3, 1},
		{4, 2, 2, 2, 3, 1},
		{2, 4, 0, 0, 1, 3},
		{2, 4, 2, 2, 1, 3},
	}

	for _, tt := range testTbl {
		scanner := bruteForceScanner{b: New(tt.w, tt.h)}

		if !scanner.IsBlack(image.Point{tt.minx, tt.miny}, image.Point{tt.maxx, tt.maxy}) {
			t.Errorf("TestBlackImage (dim:%dx%d)(%d,%d|%d,%d): expected true, got false", tt.w, tt.h, tt.minx, tt.miny, tt.maxx, tt.maxy)
		}
	}
}
