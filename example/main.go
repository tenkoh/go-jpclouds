package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"image/color"
	"image/png"
	"os"

	"github.com/tenkoh/go-jpclouds"
	"github.com/tenkoh/go-wordclouds"
)

//go:embed BIZUDPGothic-Bold.ttf
var fontBytes []byte

//go:embed mask.png
var maskImage []byte

//go:embed stopwords.txt
var stopWordBytes []byte

var defaultColors = []color.RGBA{
	{0x1b, 0x1b, 0x1b, 0xff},
	{0x48, 0x48, 0x4B, 0xff},
	{0x59, 0x3a, 0xee, 0xff},
	{0x65, 0xCD, 0xFA, 0xff},
	{0x70, 0xD6, 0xBF, 0xff},
}

type config struct {
	width           int
	height          int
	maskColor       color.RGBA
	colors          []color.RGBA
	backgroundColor color.RGBA
	fontMaxSize     int
	fontMinSize     int
	randomPlacement bool
	sizeFunction    string
}

var conf = config{
	width:           2048,
	height:          2048,
	maskColor:       color.RGBA{0, 0, 0, 0},
	colors:          defaultColors,
	backgroundColor: color.RGBA{255, 255, 255, 255},
	fontMaxSize:     300,
	fontMinSize:     30,
	randomPlacement: false,
	sizeFunction:    "linear",
}

func main() {
	// Load words
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Load stopwords
	var stopWords []string
	sc := bufio.NewScanner(bytes.NewReader(stopWordBytes))
	for sc.Scan() {
		stopWords = append(stopWords, sc.Text())
	}
	fmt.Println(len(stopWords))

	r, err := jpclouds.Preprocess(f)
	if err != nil {
		panic(err)
	}
	words, err := jpclouds.CollectWords(r, jpclouds.Noun, jpclouds.Verb, jpclouds.Adjective)
	wc, err := jpclouds.WordCount(words)
	if err != nil {
		panic(err)
	}
	jpclouds.RemoveStopWords(wc, stopWords)

	boxes := wordclouds.Mask(
		maskImage,
		conf.width,
		conf.height,
		conf.maskColor,
	)

	colors := make([]color.Color, 0)
	for _, c := range conf.colors {
		colors = append(colors, c)
	}

	oarr := []wordclouds.Option{
		wordclouds.FontFile(fontBytes),
		wordclouds.FontMaxSize(conf.fontMaxSize),
		wordclouds.FontMinSize(conf.fontMinSize),
		wordclouds.Colors(colors),
		wordclouds.MaskBoxes(boxes),
		wordclouds.Height(conf.height),
		wordclouds.Width(conf.width),
		wordclouds.RandomPlacement(conf.randomPlacement),
		wordclouds.BackgroundColor(conf.backgroundColor),
		wordclouds.WordSizeFunction(conf.sizeFunction),
	}
	w := wordclouds.NewWordcloud(wc, oarr...)

	img := w.Draw()
	outputFile, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	png.Encode(outputFile, img)
}
