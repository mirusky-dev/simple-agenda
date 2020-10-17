package dtos

const (
	pass = 10
)

// Pager is a simple pagination helper
type Pager struct {
	Limit  int
	Offset int
	Page   int
}

// Check make limit an offset auto
func (p *Pager) Check() {
	if p.Limit <= 0 {
		p.Limit = pass
	}
	if p.Page >= 0 {
		p.Offset = pass * p.Page
	}
}

// Next give the Next page
func (p *Pager) Next() int {
	return p.Page + 1
}

// Previous give the Previous page
func (p *Pager) Previous() int {
	if p.Offset != pass*p.Page {
		return -1
	}
	if p.Page == 0 {
		return 0
	}
	return p.Page - 1
}
