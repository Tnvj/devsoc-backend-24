package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/google/uuid"
)

func CheckTeamCode(code string) bool {
	query := `SELECT COUNT(*) FROM teams WHERE code = $1`
	var temp int
	err := database.DB.QueryRow(query, code).Scan(&temp)
	if err != nil || temp == 0 {
		return false
	}
	return true
}

func CheckTeamSize(team_id uuid.UUID) bool {
	query := `SELECT COUNT(*) WHERE team_id = $1`

	var no_of_member int

	err := database.DB.QueryRow(query, team_id).Scan(&no_of_member)
	if err != nil || no_of_member > 4 {
		return false
	}

	return true
}

func CheckUserInTeam(email string, teamid uuid.UUID) bool {
	query := `SELECT COUNT(*) FROM users WHERE team_id = $1 AND email = $2`
	var check int
	err := database.DB.QueryRow(query, teamid, email).Scan(&check)
	if err != nil || check == 0 {
		return false
	}
	return true
}
