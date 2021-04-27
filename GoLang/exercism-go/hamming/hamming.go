package hamming

import "errors"

func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("Different length on DNA's")
	}

	var dist int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist, nil
}
