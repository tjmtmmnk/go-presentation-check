package presentation

import "github.com/tjmtmmnk/go-presentation-check/repository"

func valid() {
	_, err := repository.Find()
	if err != nil {
		return
	}
}

func invalid() {
	got, err := repository.Find()
	if err != nil {
		return
	}
	if got == nil {
		return
	}
}
