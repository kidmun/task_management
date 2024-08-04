package utils

import "task_management/internal/core/models"

func IsValidRole(role models.Role) bool {
	switch role {
	case models.Admin, models.NormalUser:
		return true
	}
	return false
}
