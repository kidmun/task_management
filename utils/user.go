package utils

import "task_management/models"

func IsValidRole(role models.Role) bool {
	switch role {
	case models.Admin, models.NormalUser:
		return true
	}
	return false
}
