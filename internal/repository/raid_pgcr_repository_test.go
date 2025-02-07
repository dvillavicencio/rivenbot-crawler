package repository

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"rivenbot/internal/model"
	"testing"
)

func TestSavePgcr(t *testing.T) {
	// given: a raid pgcr to save
	pgcr := model.RaidPgcr{
		InstanceId: 12377100310231,
		Blob:       []byte{},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error stablishing stub connection to database")
	}

	defer db.Close()

	pgcrRepository := PgcrRepository{
		Conn: db,
	}

	// when: save is called
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO raid_pgcr").WithArgs(pgcr.InstanceId, pgcr.Blob).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result, err := pgcrRepository.Save(pgcr)

	// then: the result didn't return an error
	assert := assert.New(t)

	assert.Nil(err)
	assert.Equal(pgcr.InstanceId, result.InstanceId, "Raid PGCR instanceIds should match")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSavePgcrShouldRollback(t *testing.T) {
	// given: a raid pgcr to save
	pgcr := model.RaidPgcr{
		InstanceId: 123389859102,
		Blob:       []byte{},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error stablishing stub connection to database")
	}

	defer db.Close()

	pgcrRepository := PgcrRepository{
		Conn: db,
	}

	// when: save is called
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO raid_pgcr").WithArgs(pgcr.InstanceId, pgcr.Blob).WillReturnError(fmt.Errorf("Some error when inserting into database"))
	mock.ExpectRollback()

	// then: we expect an error to be raised when saving
	if _, err = pgcrRepository.Save(pgcr); err == nil {
		t.Errorf("Was expecting error, got none")
	}

	// and: we expect a rollback to be done
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
