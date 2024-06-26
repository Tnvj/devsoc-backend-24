package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
)

func InsertUser(user models.User) error {
	_, err := database.DB.Exec(
		"INSERT INTO users (id, first_name, last_name, reg_no, email, password, gender, phone, role, college, is_added, is_banned, is_vitian, is_verified, is_profile_complete, city, state) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)",
		user.ID,
		user.FirstName,
		user.LastName,
		user.RegNo,
		user.Email,
		user.Password,
		user.Gender,
		user.Phone,
		user.Role,
		user.College,
		user.IsAdded,
		user.IsBanned,
		user.IsVitian,
		user.IsVerified,
		user.IsProfileComplete,
		user.City,
		user.State,
	)
	return err
}
