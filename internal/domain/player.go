package domain

type Player struct {
	ID       string
	UserID   string
	Username string
}

func NewPlayer(id string, userID string, username string) *Player {
	return &Player{
		ID:       id,
		UserID:   userID,
		Username: username,
	}
}
