package utils

type IntegerRange struct {
	From int
	To   int
}

func (r IntegerRange) Contains(i int) bool {
	return r.From <= i && i <= r.To
}

func (r IntegerRange) Includes(r2 IntegerRange) bool {
	return r.From <= r2.From && r2.To <= r.To
}

func (r IntegerRange) IntegersCount() int {
	return r.To - r.From + 1
}

func ReduceIntegerRanges(ranges []IntegerRange) []IntegerRange {
	if len(ranges) == 0 {
		return []IntegerRange{}
	}

	// Sort ranges by From value
	sortedRanges := make([]IntegerRange, len(ranges))
	copy(sortedRanges, ranges)

	for i := 0; i < len(sortedRanges)-1; i++ {
		for j := i + 1; j < len(sortedRanges); j++ {
			if sortedRanges[j].From < sortedRanges[i].From {
				sortedRanges[i], sortedRanges[j] = sortedRanges[j], sortedRanges[i]
			}
		}
	}

	reducedRanges := make([]IntegerRange, 0)
	current := sortedRanges[0]

	for i := 1; i < len(sortedRanges); i++ {
		if current.Includes(sortedRanges[i]) {
			continue
		} else if sortedRanges[i].From <= current.To+1 {
			// Ranges overlap or are adjacent, merge them
			if sortedRanges[i].To > current.To {
				current.To = sortedRanges[i].To
			}
		} else {
			// No overlap, add current range and start new one
			reducedRanges = append(reducedRanges, current)
			current = sortedRanges[i]
		}
	}

	// Add the last range
	reducedRanges = append(reducedRanges, current)

	return reducedRanges
}
