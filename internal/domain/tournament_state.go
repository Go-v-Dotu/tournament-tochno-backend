package domain

type TournamentState uint8

const (
	TournamentStateUndefined TournamentState = iota
	TournamentStateCreated
	TournamentStateStarted
	TournamentStateFinished
)
