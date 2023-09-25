package main

import (
	"bufio"
	"fmt"
	"strings"

	"camille.codes/aoc/utils"
)

// The universal center of mass is indicated as COM
// Except COM, every object in space is in orbit around
// exactly one other object

func main() {
	file := utils.GetFile("f6/input.txt")
	defer file.Close()

	// A map of objects and the objects they orbit
	// By querying the map for an object, we can find which object
	// it orbits... until we reach the universal center of mass
	orbitsMap := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		orbit := getLocalOrbit(scanner.Text())
		orbitsMap[orbit[1]] = orbit[0]
	}

	fmt.Println("Part 1:", getOrbitCounts(orbitsMap))
	fmt.Println("Part 2:", getOrbitalTransfers(orbitsMap))
}

func getLocalOrbit(input string) []string {
	return strings.Split(input, ")")
}

// getOrbitalTransfers returns the minumim number of transfers required to
// move from the object YOU are orbiting to the object SAN is orbiting
func getOrbitalTransfers(orbitsMap map[string]string) int {
	youPathSlice, youPathMap := getPath(orbitsMap, "YOU")
	sanPathSlice, sanPathMap := getPath(orbitsMap, "SAN")

	youPathLength := len(youPathSlice)
	sanPathLength := len(sanPathSlice)

	fmt.Println("youPathLength:", youPathLength)
	fmt.Println("sanPathLength:", sanPathLength)

	fmt.Println("youPathSlice:", youPathSlice)
	fmt.Println("sanPathSlice:", sanPathSlice)

	fmt.Println("youPathMap:", youPathMap)
	fmt.Println("sanPathMap:", sanPathMap)

	for uIndex, uObj := range youPathSlice {
		fmt.Println("uIndex:", uIndex, "uObj:", uObj)
		for sIndex, sObj := range sanPathSlice {
			fmt.Println("sIndex:", sIndex, "sObj:", sObj)
			if uObj == sObj {
				return uIndex + sIndex
			}
		}
	}

	return 0
}

// getPath (part 2)
//
// Returns a slice of the objects in the path from start to COM and a map of the
// objects in the path to the number of transfers away from the start object
func getPath(orbitsMap map[string]string, start string) ([]string, map[string]int) {
	pathSlice := make([]string, 0)
	pathMap := make(map[string]int)
	count := 0

	for {
		if start == "COM" {
			break
		}

		start = orbitsMap[start]
		pathSlice = append(pathSlice, start)

		count++
		pathMap[start] = count
	}

	return pathSlice, pathMap
}

// getOrbitCounts returns the total number of direct and indirect orbits in the map (for part 1)
func getOrbitCounts(orbitsMap map[string]string) int {
	allOrbitCounts := 0
	for obj := range orbitsMap {
		orbitCount := 0
		count := calculateOrbitCounts(orbitsMap, obj, orbitCount)

		allOrbitCounts += count
	}

	return allOrbitCounts
}

// calculateOrbitCounts returns the number of objects that the given object orbits
func calculateOrbitCounts(orbitsMap map[string]string, object string, count int) int {
	if orbitsMap[object] != "" {
		return calculateOrbitCounts(orbitsMap, orbitsMap[object], count+1)
	}

	return count
}

// Part 1
// Orbit Count Checksums
// The total number of direct and indirect orbits in the map
// Direct orbits are the number of keys in orbitsMap
// and indirect orbits must be calculated for each key/object in the orbitsMap

// For each object in the map, we need to find the number of objects in total
// that it orbits
// For example, if the map is:
// COM)B
// B)C
// C)D
// D)E

// COM orbits 0 objects
// B orbits 1 object (COM)
// C orbits 2 objects (B and COM)
// D orbits 3 objects (C, B, and COM)
// E orbits 4 objects (D, C, B, and COM)

// Part 2
// COM)B
// B)C
// C)D
// D)E
// E)F
// B)G
// G)H
// D)I
// E)J
// J)K
// K)L
// K)YOU
// I)SAN

// pathMap = map[string]int{"K": 0, "J": 1, "E": 2, "D": 3, "C": 4, "B": 5, "COM": 6}
// pathMap = map[string]int{"I": 0, "D": 1, "C": 2, "B": 3, "COM": 4}

// pathSlice = ["I", "D", "C", "B", "COM]
// pathSlice = ["K", "J", "E", "D", "C", "B", "COM"]
