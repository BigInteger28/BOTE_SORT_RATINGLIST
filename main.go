package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Open en lees het invoerbestand
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Sorteer de regels
	sort.SliceStable(lines, func(i, j int) bool {
		fieldsI := strings.Split(lines[i], "   ")
		fieldsJ := strings.Split(lines[j], "   ")

		// Extract name (eerste veld)
		nameI := fieldsI[0]
		nameJ := fieldsJ[0]

		// Extract level (tweede veld)
		levelI, _ := strconv.Atoi(fieldsI[1])
		levelJ, _ := strconv.Atoi(fieldsJ[1])

		// Extract rating (derde veld)
		ratingI, _ := strconv.Atoi(fieldsI[2])
		ratingJ, _ := strconv.Atoi(fieldsJ[2])

		// Sorteer op rating (aflopend)
		if ratingI != ratingJ {
			return ratingI > ratingJ
		}

		// Bij gelijke rating, sorteer op level (aflopend)
		if levelI != levelJ {
			return levelI > levelJ
		}

		// Bij gelijke rating en level, non-"---" vóór "---"
		isDashI := strings.HasPrefix(nameI, "---")
		isDashJ := strings.HasPrefix(nameJ, "---")
		if isDashI != isDashJ {
			return !isDashI // non-"---" komt vóór "---"
		}

		return false // Behoud originele volgorde
	})

	// Schrijf naar uitvoerbestand
	outFile, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
}
