package service

//go:generate mockgen -source=get_rows.go -destination=mocks/mock.go

type RowsGetter interface {
	GetRows() []string
}

// After we got table from repository we should transform it to slice of names
func (s *Service) GetRows() []string {
	rowsFromRepo := s.repo.GetRows()
	rows := make([]string, len(rowsFromRepo[0]))
	for i := 0; i < len(rowsFromRepo[0]); i++ {
		rows[i] = rowsFromRepo[0][i] + " " + rowsFromRepo[1][i]
	}
	return rows
}
