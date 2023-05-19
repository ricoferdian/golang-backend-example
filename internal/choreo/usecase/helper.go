package usecase

import "kora-backend/internal/model"

func (c ChoreoUseCaseImpl) checkContentUnlockStatus(choreoData model.ChoreographyModel) bool {
	// TODO remove after dynamic pricing implemented
	if choreoData.TempPrice.Int64 == 0 {
		return true
	}
	return false
}
