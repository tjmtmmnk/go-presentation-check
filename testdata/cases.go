package testdata

import (
	"errors"
)

func a() {
	_, err := b(0)
	if err != nil {
		return
	}
	return
}

func b(v int) (int, error) {
	return v, errors.New("error")
}
