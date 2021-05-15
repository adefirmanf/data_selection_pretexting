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
	GetLastToken() (*Token, error)
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

	InsertUser(UserID, Username, Name, Location string,
		TotalFollowing, TotalFollowers int,
		IsVerified bool,
		UserCreatedAt time.Time) error

	UpdateUserByUserIDTwitter(UserID string, TotalFollowing, TotalFollowers int) error
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
	return s.datasource.InsertTweets(TweetID, TweetAuthorID, TweetText, SuspiciousKeywords, TweetMentionedAccount,
		OptionalParameters,
		TokenID,
		TweetCreatedAt,
		TweetPossiblySensitive)
}

// GetLastToken .
func (s *Storage) GetLastToken() (*Token, error) {
	return s.datasource.GetLastToken()
}

// InsertToken .
func (s *Storage) InsertToken(TweetNextToken, URL string) error {
	return s.datasource.InsertToken(TweetNextToken, URL)
}

// InsertUser .
func (s *Storage) InsertUser(UserID, Username, Name, Location string,
	TotalFollowing, TotalFollowers int,
	IsVerified bool,
	UserCreatedAt time.Time,
) error {
	return s.datasource.InsertUser(UserID, Username, Name, Location,
		TotalFollowing, TotalFollowers,
		IsVerified,
		UserCreatedAt)
}

// UpdateUserByUserIDTwitter .
func (s *Storage) UpdateUserByUserIDTwitter(UserID string, TotalFollowing, TotalFollowers int) error {
	return s.datasource.UpdateUserByUserIDTwitter(UserID, TotalFollowing, TotalFollowers)
}

// NewStorage .
func NewStorage(s I) *Storage {
	return &Storage{
		datasource: s,
	}
}
