package cfl

/*ElementsPerSplit Returns the number of elements per split and leftover

 */
func ElementsPerSplit(totalSize, numberOfSplits int) (int, int) {
	return totalSize / numberOfSplits, totalSize % numberOfSplits
}

type SplitRange struct {
	start int
	end   int
}

func SplitRanges(totalSize, numberOfSplits int) []SplitRange {
	elementsPerSplit, leftover := ElementsPerSplit(totalSize, numberOfSplits)

	result := make([]SplitRange, numberOfSplits)

	for i := 0; i < numberOfSplits; i++ {
		if i == numberOfSplits-1 {
			result[i] = SplitRange{
				start: i * elementsPerSplit,
				end:   (i+1)*elementsPerSplit + leftover,
			}

		} else {
			result[i] = SplitRange{
				start: i * elementsPerSplit,
				end:   (i + 1) * elementsPerSplit,
			}
		}

	}

	return result
}
