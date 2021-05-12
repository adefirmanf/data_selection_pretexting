package storage

import "time"

// Storage .
type Storage struct {
	datasource I
}

// I .
type I interface {
	GetUsers() ([]*User, error)
	GetTweets() ([]*Tweet, error)
	GetUserByUserAuthorID(TweetAuthorUserID string) (*User, error)
	// GetLastToken() (*Token, error)
	// GetTweets()
	// GetLastToken()

	InsertTweets(
		TweetID, TweetAuthorID, TweetText, SuspiciousKeywords, TweetMentionedAccount,
		OptionalParameters string,
		TokenID int,
		TweetCreatedAt time.Time,
		TweetPossiblySensitive bool,
	) error

	InsertToken(TweetNextToken, URL string) error
	// InsertUsers()

	// UpdateUserByUserIDTwitter()
}

// GetUsers .
func (s *Storage) GetUsers() ([]*User, error) {
	return s.datasource.GetUsers()
}

// GetUserByUserAuthorID .
func (s *Storage) GetUserByUserAuthorID(TweetAuthorUserID string) (*User, error) {
	return s.datasource.GetUserByUserAuthorID(TweetAuthorUserID)
}

// GetTweets .
func (s *Storage) GetTweets() ([]*Tweet, error) {
	return s.datasource.GetTweets()
}

// InsertTweets .
func (s *Storage) InsertTweets(
	TweetID, TweetAuthorID, TweetText, SuspiciousKeywords, TweetMentionedAccount,
	OptionalParameters string,
	TokenID int,
	TweetCreatedAt time.Time,
	TweetPossiblySensitive bool,
) error {
	return s.datasource.InsertTweets(TweetID, TweetAuthorID, TweetText, SuspiciousKeywords, TweetMentionedAccount, OptionalParameters, TokenID, TweetCreatedAt, TweetPossiblySensitive)
}

// InsertToken .
func (s *Storage) InsertToken(TweetNextToken, URL string) error {
	return s.datasource.InsertToken(TweetNextToken, URL)
}

// NewStorage .
func NewStorage(s I) *Storage {
	return &Storage{
		datasource: s,
	}
}
