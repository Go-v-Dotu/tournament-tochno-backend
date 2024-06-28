package domain

type EnrolledPlayer struct {
	PlayerID      string
	Dropped       bool
	HasCommanders bool
}

func NewEnrolledPlayer(playerID string, dropped bool, hasCommander bool) *EnrolledPlayer {
	return &EnrolledPlayer{
		PlayerID:      playerID,
		Dropped:       dropped,
		HasCommanders: hasCommander,
	}
}

func (ep *EnrolledPlayer) Drop() {
	ep.Dropped = true
}

func (ep *EnrolledPlayer) Recover() {
	ep.Dropped = false
}
