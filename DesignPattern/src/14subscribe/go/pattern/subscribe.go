package pattern

/**
* Created by wanjx in 2018/12/24 22:35
**/
 
type Subscribe struct {
	Name string
	Browses []Browser
	Action string
}

func (s *Subscribe) Add(sb Browser) {
	s.Browses = append(s.Browses, sb)
}

func (s *Subscribe) Inform() {
	for _, b := range s.Browses {
		b.Update()
	}
}

func (s *Subscribe) ActionT() string{
	return s.Action
}