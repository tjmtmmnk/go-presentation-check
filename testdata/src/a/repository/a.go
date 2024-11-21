package repository

type X struct{}

func Find() (*X, error) { // want `report function Find`
	return &X{}, nil
}
