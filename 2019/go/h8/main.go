package main

import (
	"bufio"
	"fmt"
	"strings"

	"camille.codes/aoc/utils"
)

// It's not 1536 (too low)
// It's not 2060

func main() {
	file := utils.GetFile("h8/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := strings.Split(scanner.Text(), "")

	width := 25
	height := 6

	imagePixelSize := width * height

	// The number of layers are the total number of pixels
	// divided by the number of pixels per image
	numLayers := len(input) / imagePixelSize
	start := 0
	end := imagePixelSize
	// For each number in the layer, count the number of zeros
	counts := make(map[int]int, 0)
	ones := make(map[int]int, 0)
	twos := make(map[int]int, 0)

	layers := make([][]string, numLayers)
	for layer := 0; layer < numLayers; layer++ {
		layers[layer] = input[start:end]
		start += imagePixelSize
		end += imagePixelSize

		// [[1 2 3 4 5 6] [7 8 9 0 1 2]]

		// For each number in the layer, count the number of zeros
		for _, pixel := range layers[layer] {
			// For each pixel in this layer, check if it is a zero
			if pixel == "0" {
				counts[layer]++
			}

			if pixel == "1" {
				ones[layer]++
			}

			if pixel == "2" {
				twos[layer]++
			}
		}
	}

	// Which layer has the fewest 0 digits?
	lowestLayer := 0
	lowest := 9999999
	for layer, pixels := range layers {
		fmt.Println(layer, pixels)
		fmt.Println("layer", layer, "has", counts[layer], "zeros")
		if counts[layer] < lowest {
			lowestLayer = layer
			lowest = counts[layer]
		}
	}
	fmt.Println("lowest zeros are in layer", lowestLayer)

	fmt.Println("1s*2s:", ones[lowestLayer]*twos[lowestLayer])
}
