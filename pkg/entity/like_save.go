package entity

type BaseLikeSaveEntity struct {
	ChoreoID int64 `json:"choreo_id"`
	UserID   int64 `json:"user_id"`
}

type ChoreoLikeEntity struct {
	BaseLikeSaveEntity
}

type ChoreoSaveEntity struct {
	BaseLikeSaveEntity
}
