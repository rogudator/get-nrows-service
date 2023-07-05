package repository

type RowsGetter interface {
	getRows() [][]string
}

func (r *Repository) GetRows() [][]string {
	return r.table
}
