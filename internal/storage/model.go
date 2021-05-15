package storage

import "time"

// Tweet .
type Tweet struct {
	TweetID                string
	TweetAuthorID          string
	TweetCreatedAt         time.Time
	TweetText              string
	SuspiciousKeywords     string
	TweetMentionedAccount  string
	TweetPossiblySensitive bool
	TokenID                int
	OptionalParameters     string
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

// User .
type User struct {
	UserID             string
	UserIsVerified     bool
	UserTotalFollowing int
	UserTotalFollowers int
	UserUsername       string
	UserFullName       string
	Location           string
	UserCreatedAt      time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// Token .
type Token struct {
	TweetNextToken string
	URL            string
	CreatedAt      time.Time
}
