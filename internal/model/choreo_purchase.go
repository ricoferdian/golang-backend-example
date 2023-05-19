//go:generate easytags $GOFILE db
package model

import "database/sql"

type ChoreoPurchaseModel struct {
	ChoreoPurchaseID int64          `db:"choreo_purchase_id"`
	UserID           int64          `db:"user_id"`
	ChoreoID         int64          `db:"choreo_id"`
	Receipt          sql.NullString `db:"receipt"`
	Status           int16          `db:"status"`
}
