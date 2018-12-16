package proxy

type Girl struct {
	Name string
}

func NewGirl(name string) (g *Girl) {
	g = &Girl{Name:name}
	return
}
