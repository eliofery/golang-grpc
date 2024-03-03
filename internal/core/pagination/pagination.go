package pagination

// Pagination ...
type Pagination struct {
	limit  uint64
	offset uint64
}

// New ...
func New(config *Config) *Pagination {
	return &Pagination{
		limit:  config.Limit,
		offset: 0,
	}
}

// SetLimit ...
func (p *Pagination) SetLimit(limit uint64) *Pagination {
	p.limit = limit

	return p
}

// SetOffset ...
func (p *Pagination) SetOffset(page uint64) *Pagination {
	if page == 0 {
		page = 1
	}

	p.offset = (page - 1) * p.limit

	return p
}

// Limit ...
func (p *Pagination) Limit() uint64 {
	return p.limit
}

// Offset ...
func (p *Pagination) Offset() uint64 {
	return p.offset
}
