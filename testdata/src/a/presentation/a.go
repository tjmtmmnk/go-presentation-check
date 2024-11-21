package presentation

import (
	"a/repository"
	"a/service"
)

func valid() { // OK
	_, err := repository.Find()
	if err != nil {
		return
	}
}

func invalidTooComplex() {
	got, err := repository.Find()
	if err != nil {
		return
	}
	if got == nil { // want `ロジックを書いてはいけません`
		return
	}
}

func invalidTooComplexInVar() {
	var invalid = func() {
		got, err := repository.Find()
		if err != nil {
			return
		}
		if got == nil { // want `ロジックを書いてはいけません`
			return
		}
	}
	invalid()
}

func invalidTooManyCallRepository() {
	_, err := repository.Find()
	_, err = repository.Find() // want `repositoryの呼び出しは1回までです。usecaseを作ってください。`

	if err != nil {
		return
	}
}

func invalidCallService() {
	_, err := service.Find() // want `serviceの呼び出しは禁止です。usecaseを作ってください。`
	if err != nil {
		return
	}
}

type T struct{}

func (t *T) valid() { // OK
	_, err := repository.Find()
	if err != nil {
		return
	}
}

func (t *T) invalidTooComplex() {
	got, err := repository.Find()
	if err != nil {
		return
	}
	if got == nil { // want `ロジックを書いてはいけません`
		return
	}
}
