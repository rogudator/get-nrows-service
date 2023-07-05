package repository

type Repository struct {
	RowsGetter
	table [][]string
}

// Since we store information in memory. Repository will just return table containing two columns: names and surnames
func NewRepository() *Repository {
	return &Repository{
		table: [][]string{{"John", "Sarah", "Michael", "Emily", "David", "Olivia", "Daniel", "Sophia", "Matthew", "Isabella", "Andrew", "Ava", "Christopher", "Mia", "Joseph"},
			{"Johnson", "Williams", "Jones", "Brown", "Davis", "Miller", "Wilson", "Moore", "Taylor", "Anderson", "Thomas", "Jackson", "White", "Harris", "Clark"}},
	}
}
