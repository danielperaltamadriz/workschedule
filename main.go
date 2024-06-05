package main

import (
	"slices"
	"strconv"
)

func main() {}

func findSchedules(workHours int, dayHours int, pattern string) []string {
	missingWorkingHours := workHours
	var missingWorkingDays int

	for _, p := range pattern {
		if p == '?' {
			missingWorkingDays++
			continue
		}
		workingHours, _ := strconv.Atoi(string(p))
		missingWorkingHours -= workingHours
	}
	if missingWorkingDays == 0 {
		return []string{pattern}
	}
	if missingWorkingHours/(missingWorkingDays*dayHours) == 1 {
		var out string
		for i := 0; i < len(pattern); i++ {
			if pattern[i] != '?' {
				out += string(pattern[i])
				continue
			}
			out += strconv.Itoa(dayHours)
		}
		return []string{out}
	}
	combinations := fillMissingWorkingDays(missingWorkingDays, dayHours, missingWorkingHours)
	var schedules []string
	for i := 0; i < len(combinations); i++ {
		var o string
		var outIndex int
		for j := 0; j < len(pattern); j++ {
			if pattern[j] != '?' {
				o += string(pattern[j])
				continue
			}
			o += string(combinations[i][outIndex])
			outIndex++
		}
		schedules = append(schedules, o)
	}

	return schedules
}

func fillMissingWorkingDays(missingDays int, dayHours int, missingWorkingHours int) []string {
	maxHoursPerDay := min(dayHours, missingWorkingHours)
	var output []string
	for i := maxHoursPerDay; i >= 0; i-- {
		out := strconv.Itoa(i)
		diff := missingWorkingHours - i
		if missingDays == 2 {
			if diff > dayHours {
				break
			}
			output = append(output, out+strconv.Itoa(diff))
			continue
		}
		comb := fillMissingWorkingDays((missingDays - 1), dayHours, diff)
		if len(comb) == 0 {
			break
		}
		for j := 0; j < len(comb); j++ {
			output = append(output, out+comb[j])
		}

	}
	slices.Sort(output)
	return output
}
