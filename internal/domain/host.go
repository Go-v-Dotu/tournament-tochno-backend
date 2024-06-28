package domain

type Host struct {
	ID       string
	UserID   string
	Username string
}

func NewHost(id string, userID string, username string) *Host {
	return &Host{
		ID:       id,
		UserID:   userID,
		Username: username,
	}
}
