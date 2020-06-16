package autocomplete

import "math"

func Closest(c string, companies map[string]struct{}) (float64, string) {
	closestMatch := 0.0
	closestCompany := ""
	for company := range companies {
		longer, shorter := max(len(c), len(company)), min(len(c), len(company))
		numMatches, numMismatches := 0, 0
		cDict, companyDict := make(map[string]int), make(map[string]int)
		for i := 0; i < shorter; i++ {
			if c[i] == company[i] {
				numMatches++
			} else {
				numMismatches++
				cDict[string(c[i])]++
				companyDict[string(company[i])]++
			}
		}
		for i := shorter; i < len(c); i++ {
			cDict[string(c[i])]++
		}
		for i := shorter; i < len(company); i++ {
			companyDict[string(company[i])]++
		}

		actualErrors := 0.0
		typos := 0.0
		for k, v := range cDict {
			if _, ok := companyDict[k]; ok {
				if companyDict[k] > 0 && math.Abs(float64(companyDict[k] - v)) < 2 {
					companyDict[k] = 0
					typos += math.Abs(float64(companyDict[k] - v))
				} else {
					companyDict[k] -= v
					typos += math.Abs(float64(companyDict[k]))
				}
			}
		}
		for _, v := range companyDict {
			actualErrors += math.Abs(float64(v))
		}
		currMatch := float64(numMatches) / float64(longer) + ((float64(numMismatches) + typos / 2) / (float64(numMismatches) + float64(longer) + actualErrors))
		if currMatch > closestMatch {
			closestMatch = currMatch
			closestCompany = company
		}
	}
	return closestMatch, closestCompany
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
// two sigam ||||||| two sigma
