package repository

import "database/sql"

type profilePGSQL struct {
	db *sql.DB
}

func NewProfilePGSQL(db *sql.DB) *profilePGSQL {
	return &profilePGSQL{
		db: db,
	}
}

func (p *profilePGSQL) Create() error {
	return nil
}