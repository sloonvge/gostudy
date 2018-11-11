package cash

func CreateCash(t string) (c AcceptCash){
	switch t {
	case "8折":
		c = NewRebate(0.8)
	case "满10减2":
		c = NewReturn(10, 2)
	default:
		c = NewNormal()
	}
	return
}

