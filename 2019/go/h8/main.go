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
		// fmt.Println("layer", layer, "has", zeros[layer], "zeros")
		if zeros[layer] < lowest {
			lowestLayer = layer
			lowest = zeros[layer]
		}
	}

	fmt.Println("lowest zeros are in layer", lowestLayer)
	fmt.Println("1s*2s:", ones[lowestLayer]*twos[lowestLayer])
}
