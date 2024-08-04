package utils

import "task_management/internal/core/models"

func IsValidStatus(status models.Status) bool {
	switch status {
	case models.Pending, models.InProgress, models.Done:
		return true
	}
	return false
}
