package repository

type X struct{}

func Find() (*X, error) {
	return &X{}, nil
}
