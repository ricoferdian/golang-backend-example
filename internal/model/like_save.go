package model

type BaseLikeSave struct {
	ChoreoID int64 `db:"choreo_id"`
	UserID   int64 `db:"user_id"`
}

type ChoreoLikeModel struct {
	BaseLikeSave
}

type ChoreoSaveModel struct {
	BaseLikeSave
}
