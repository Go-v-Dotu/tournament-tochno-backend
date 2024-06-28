package controllers

type GetTournamentRequest struct {
	ID string `param:"id"`
}

type GetPlayersRequest struct {
	TournamentID string `param:"id"`
}
