package views

import (
	"time"
)

type ErrorResponse struct{}

type ParticipatedTournamentsResponse struct {
	Tournaments []*TournamentPreview `json:"tournaments"`
}

type GetTournamentsResponse struct {
	Tournaments []*TournamentPreview `json:"tournaments"`
}

type GetTournamentResponse struct {
	Tournament *Tournament `json:"tournament"`
}

type GetPlayersResponse struct {
	Players []*Player `json:"players"`
}

type Tournament struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Host           *Host     `json:"host"`
	Date           time.Time `json:"date"`
	TotalPlayers   int       `json:"total_players"`
	PlayerEnrolled bool      `json:"player_enrolled"`
}

type TournamentPreview struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Host         *Host     `json:"host"`
	Date         time.Time `json:"date"`
	TotalPlayers int       `json:"total_players"`
}

type Host struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type Player struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Dropped  bool   `json:"dropped"`
}
