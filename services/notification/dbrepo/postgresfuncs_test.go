package dbrepo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	db, _, _ := sqlmock.New()
	repo := New(db)
	_ , ok :=  repo.(*postgresRepo)
	assert.True(t, ok, "Expected a *PostgresRepository")
}