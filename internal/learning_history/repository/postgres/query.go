package postgres

const (
	tableLearningHistory = "m_learning_history"

	columnSelectAllLearningHistory = "learning_history_id,choreo_detail,device,downloaded,expired,choreo_detail_id,progress,record_url,thumbnail_url,user_id"
	columnInsertLearningHistory    = "choreo_detail_id,user_id,choreo_detail,device,progress,record_url,thumbnail_url"
)
