//go:generate easytags $GOFILE json
package entity

type ChoreoPurchaseEntity struct {
	ChoreoPurchaseID int64  `json:"choreo_purchase_id"`
	UserID           int64  `json:"user_id"`
	ChoreoID         int64  `json:"choreo_id"`
	Receipt          string `json:"receipt"`
	Status           int16  `json:"status"`
	// Relation
	Choreography *ChoreographyEntity `json:"choreography_data,omitempty"`
}
