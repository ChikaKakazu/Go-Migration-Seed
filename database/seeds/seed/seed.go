package seed

import (
	"database/sql"
)

type Seed struct {
	Name string
	Run  func(*sql.DB)
}

func All() []Seed {
	return []Seed{
		{
			Name: "Users",
			Run: func(db *sql.DB) {
				Delete(db)
				Create(db)
			},
		},
	}
}
