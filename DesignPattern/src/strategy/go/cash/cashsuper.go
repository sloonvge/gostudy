package cash

type Super struct {}

type AcceptCash interface {
	Cash(m float64) float64
}
