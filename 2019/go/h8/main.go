package main

import (
	"bufio"
	"fmt"
	"strings"

	"camille.codes/aoc/utils"
)

func main() {
	file := utils.GetFile("h8/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := strings.Split(scanner.Text(), "")

	width := 25
	height := 6

	imagePixelSize := width * height

	// We only care about 0s, 1s, and 2s
	// e.g. map[layer]count
	zeros := make(map[int]int, 0)
	ones := make(map[int]int, 0)
	twos := make(map[int]int, 0)

	// The number of layers are the total number of pixels
	// divided by the number of pixels per image
	numLayers := len(input) / imagePixelSize
	start := 0
	end := imagePixelSize

	layers := make([][]string, numLayers)
	for layer := 0; layer < numLayers; layer++ {
		layers[layer] = input[start:end]
		start += imagePixelSize
		end += imagePixelSize

		for _, pixel := range layers[layer] {
			switch pixel {
			case "0":
				zeros[layer]++
			case "1":
				ones[layer]++
			case "2":
				twos[layer]++
			}
		}
	}

	// Which layer has the fewest 0 digits?
	lowestLayer := 0
	lowest := zeros[0]
	for layer := range layers {
		if zeros[layer] < lowest {
			lowestLayer = layer
			lowest = zeros[layer]
		}
	}

	fmt.Println(
		"layer", lowestLayer,
		"has", lowest, "zeros,",
		ones[lowestLayer], "ones,",
		"and", twos[lowestLayer], "twos",
	)
	fmt.Println("Part 1 (1s*2s):", ones[lowestLayer]*twos[lowestLayer])

	decodedImage := make([]string, imagePixelSize)
	for i := 0; i < imagePixelSize; i++ {
		for _, layer := range layers {
			if layer[i] != "2" {
				decodedImage[i] = layer[i]
				break
			}
		}
	}

	for pixel := range decodedImage {
		if decodedImage[pixel] == "0" {
			decodedImage[pixel] = " "
		}
	}

	fmt.Println("Part 2 (decoded image):")
	for i := 0; i < imagePixelSize; i += width {
		fmt.Println(strings.Join(decodedImage[i:i+width], ""))
	}
}

// Part 2
// 0222112222120000

// Layers with their pixels
// Layer 1: 02
//          22

// Layer 2: 11
//          22

// Layer 3: 22
//          12

// Layer 4: 00
//          00

// 0 - black
// 1 - white
// 2 - transparent

// Check layer one, pixel by pixel
// 1. If pixel is 0 or 1, then we're done (capture it in the final image)
// 2. If pixel is 2, check the next layer
// 	2a. In the next layer, return to step 1
// 3. If there are no more layers to check, then the pixel is transparent

// Decoded image
// 01
// 10
