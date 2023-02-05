package db

import (
	_ "embed" // support embedding files in variables
	"fmt"
	"log"
	"strings"

	"atlassian.carcgl.com/bitbucket/ls/lms/pkg/model"
	"github.com/jmoiron/sqlx"
)

//go:embed schema.sql
var schema string

// CreateSchema creates the database schema.
func CreateSchema(db *sqlx.DB) error {
	for n, statement := range strings.Split(schema, ";") {
		_, err := db.Exec(statement)
		if err != nil {
			return fmt.Errorf("statement %d failed: \"%s\" : %w", n+1, statement, err)
		}
	}
	return nil
}

func AddOperator(db *sqlx.DB, operator *model.Operator) error {
	_, err := db.NamedExec(
		`INSERT INTO operators (id, name, valid, created)
         VALUES (:id, :name, :valid, :created)`,
		operator,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func GetOperator(db *sqlx.DB, id string) (model.Operator, error) {
	var operator model.Operator
	return operator, db.Get(&operator, "SELECT * FROM operators WHERE id = $1", id)
}

func UpdateOperator(db *sqlx.DB, operator *model.Operator) error {
	_, err := db.NamedExec(
		`UPDATE operators SET name = :name, valid = :valid, created = :created WHERE id = :id`,
		operator,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func DeleteOperator(db *sqlx.DB, id string) error {
	_, err := db.Exec("DELETE FROM operators WHERE id = $1", id)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func ListOperators(db *sqlx.DB) ([]model.Operator, error) {
	var operators []model.Operator
	return operators, db.Select(&operators, "SELECT * FROM operators")
}
