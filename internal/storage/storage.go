package storage

// Storage .
type Storage struct {
	datasource I
}

// I .
type I interface {
	GetUsers()
	GetTweets()
	GetLastToken()

	InsertTweets()
	InsertTokens()
	InsertUsers()

	UpdateUserByUserIDTwitter()
}

// GetUsers .
func (s *Storage) GetUsers() {
	s.datasource.GetUsers()
}

// NewStorage .
func NewStorage(s I) *Storage {
	return &Storage{}
}
