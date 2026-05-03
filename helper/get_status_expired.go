package helper

import (
	"inventory-indra/model"
	"time"
)


func GetExpiredStatus(expiredDate time.Time) model.ExpiredStatus {
	now := time.Now()
	daysLeft := time.Until(expiredDate).Hours() / 24

	if expiredDate.Before(now) {
		return model.StatusExpired
	} else if daysLeft <= 30 {
		return model.StatusWarning
	}

	return model.StatusSafe
}