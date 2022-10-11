package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	ch := make(chan FreqMap)
	for i := range l {
		go func(i int) {
			ch <- Frequency(l[i])
		}(i)
	}

	freqMap := make(FreqMap)
	for range l {
		for letter, count := range <-ch {
			freqMap[letter] += count
		}
	}
	return freqMap
}
