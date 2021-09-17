package crd

type statusDesc struct {
	src  string
	curr string
	dest []string
}

func (s *statusDesc) NewStatusDesc(src string, curr string, dest []string) *statusDesc {
	return &statusDesc{
		src:  src,
		curr: curr,
		dest: dest,
	}
}

func (s *statusDesc) SetSrc(src string) {
	s.src = src
}

func (s *statusDesc) SetCurr(curr string) {
	s.curr = curr
}

func (s *statusDesc) SetDest(dest []string) {
	s.dest = dest
}

func (s *statusDesc) GetSrc() string {
	return s.src
}

func (s *statusDesc) GetCurr() string {
	return s.curr
}

func (s *statusDesc) GetDest() []string {
	return s.dest
}
