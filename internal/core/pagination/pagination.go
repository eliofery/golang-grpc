package pagination

// Pagination ...
type Pagination interface {
	SetLimit(limit uint64) Pagination
	SetOffset(page uint64) Pagination
	Limit() uint64
	Offset() uint64
}

// pagination ...
type pagination struct {
	limit  uint64
	offset uint64
}

// New ...
func New(config *Config) Pagination {
	return &pagination{
		limit:  config.Limit,
		offset: 0,
	}
}

// SetLimit ...
func (p *pagination) SetLimit(limit uint64) Pagination {
	p.limit = limit

	return p
}

// SetOffset ...
func (p *pagination) SetOffset(page uint64) Pagination {
	if page < 1 {
		page = 1
	}

	p.offset = (page - 1) * p.limit

	return p
}

// Limit ...
func (p *pagination) Limit() uint64 {
	return p.limit
}

// Offset ...
func (p *pagination) Offset() uint64 {
	return p.offset
}
