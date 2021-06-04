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

func ConcurrentFrequency(words []string) (mp FreqMap) {
	ch := make(chan FreqMap)

	for _, l := range words {
		go func(l string) {
			ch <- Frequency(l)
		}(l)
	}

	mp = FreqMap{}
	for range words {
		for k, v := range <-ch {
			mp[k] += v
		}
	}

	return
}
