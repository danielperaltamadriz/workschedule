package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSchedules(t *testing.T) {
	tests := []struct {
		workHours int
		dayHours  int
		pattern   string
		expected  []string
	}{
		{
			workHours: 24,
			dayHours:  4,
			pattern:   "08??840",
			expected: []string{
				"0804840",
				"0813840",
				"0822840",
				"0831840",
				"0840840",
			},
		},
		{
			workHours: 24,
			dayHours:  4,
			pattern:   "28??8?2",
			expected: []string{
				"2800842",
				"2801832",
				"2802822",
				"2803812",
				"2804802",
				"2810832",
				"2811822",
				"2812812",
				"2813802",
				"2820822",
				"2821812",
				"2822802",
				"2830812",
				"2831802",
				"2840802",
			},
		},
		{
			workHours: 56,
			dayHours:  8,
			pattern:   "?8?????",
			expected: []string{
				"8888888",
			},
		},
		{
			workHours: 55,
			dayHours:  8,
			pattern:   "?8?????",
			expected: []string{
				"7888888",
				"8878888",
				"8887888",
				"8888788",
				"8888878",
				"8888887",
			},
		},
		{
			workHours: 56,
			dayHours:  8,
			pattern:   "8888888",
			expected: []string{
				"8888888",
			},
		},
	}

	for _, test := range tests {
		result := findSchedules(test.workHours, test.dayHours, test.pattern)
		assert.Equal(t, result, test.expected)
	}
}
