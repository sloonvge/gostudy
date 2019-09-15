package base

type T struct {

}

type F interface {
	E(t *T) error
}

type G struct {

}

func (g G) E(t *T) error {
	return nil
}

var f F = &G{}

