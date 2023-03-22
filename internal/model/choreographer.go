//go:generate easytags $GOFILE db
package model

type ChoreographerModel struct {
	ChoreographerID   int64  `db:"choreographer_id"`
	ChoreographerName string `db:"choreographer_name"`
}
