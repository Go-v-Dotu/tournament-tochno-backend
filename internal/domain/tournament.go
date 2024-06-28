package domain

import (
	"slices"
	"time"
)

type Tournament struct {
	ID      string
	HostID  string
	Title   string
	Players []*EnrolledPlayer
	Date    time.Time
	State   TournamentState
}

func NewTournament(
	id string,
	hostID string,
	title string,
	players []*EnrolledPlayer,
	date time.Time,
	state TournamentState,
) *Tournament {
	t := &Tournament{
		ID:      id,
		HostID:  hostID,
		Title:   title,
		Players: players,
		Date:    date,
		State:   state,
	}
	return t
}

func (t *Tournament) IsParticipant(player *Player) bool {
	if player == nil {
		return false
	}
	return slices.ContainsFunc(t.Players, func(x *EnrolledPlayer) bool {
		return x.PlayerID == player.ID
	})
}
