package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Stap 1: Open en lees het invoerbestand
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

	// Stap 2: Sorteer de regels met sort.SliceStable en aangepaste vergelijking
	sort.SliceStable(lines, func(i, j int) bool {
		fieldsI := strings.Fields(lines[i])
		fieldsJ := strings.Fields(lines[j])

		// Extract rating (laatste veld)
		ratingI, _ := strconv.Atoi(fieldsI[len(fieldsI)-1])
		ratingJ, _ := strconv.Atoi(fieldsJ[len(fieldsJ)-1])

		// Extract level (voorlaatste veld)
		levelI, _ := strconv.Atoi(fieldsI[len(fieldsI)-2])
		levelJ, _ := strconv.Atoi(fieldsJ[len(fieldsJ)-2])

		// Primaire sleutel: rating aflopend
		if ratingI != ratingJ {
			return ratingI > ratingJ
		}

		// Secundaire sleutel: level aflopend
		if levelI != levelJ {
			return levelI > levelJ
		}

		// Tertiaire sleutel: non-"---" vóór "---"
		isDashI := fieldsI[0] == "---"
		isDashJ := fieldsJ[0] == "---"
		if !isDashI && isDashJ {
			return true // i vóór j
		}
		if isDashI && !isDashJ {
			return false // j vóór i
		}
		return false // Behoud originele volgorde voor gelijke elementen
	})

	// Stap 3: Schrijf de gesorteerde regels naar het uitvoerbestand
	outFile, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	for _, line := range lines {
		_, err := outFile.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}
}
