package postgres

const (
	columnSelectAllPurchasedChoreo       = "choreo_purchase_id,choreo_id,user_id,receipt,status"
	columnSelectPurchasedChoreoNoReceipt = "choreo_purchase_id,choreo_id,user_id,status"
	columnInsertPurchasedChoreo          = "choreo_id,user_id,receipt,status"
	tablePurchasedChoreo                 = "t_choreo_purchase"
)
