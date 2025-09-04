package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var multipliers = map[string]float64{
	"ribu":    1e3,
	"k":       1e3,
	"juta":    1e6,
	"jt":      1e6,
	"m":       1e6,
	"miliar":  1e9,
	"b":       1e9,
	"triliun": 1e12,
	"t":       1e12,
}

func NormalizeEstimationRevenueToMonthly(input string) (float64, error) {
	text := strings.ToLower(input)
	text = strings.ReplaceAll(text, "rp", "")
	text = strings.ReplaceAll(text, "idr", "")
	text = strings.ReplaceAll(text, ",", "")
	text = strings.TrimSpace(text)

	// handle ranges like "50-100 juta"
	if strings.Contains(text, "-") {
		parts := strings.Split(text, "-")
		// keep the right-most number part
		text = strings.TrimSpace(parts[len(parts)-1])
	}

	// regex number + optional multiplier word
	re := regexp.MustCompile(`([\d\.]+)\s*([a-z]*)`)
	matches := re.FindStringSubmatch(text)
	if len(matches) < 2 {
		return 0, fmt.Errorf("could not parse: %s", input)
	}

	// parse number
	num, err := strconv.ParseFloat(strings.ReplaceAll(matches[1], ".", ""), 64)
	if err != nil {
		return 0, err
	}

	// apply multiplier
	if mult, ok := multipliers[matches[2]]; ok {
		num *= mult
	}

	// adjust if yearly
	if strings.Contains(text, "tahun") {
		num = num / 12
	}

	return num, nil
}
