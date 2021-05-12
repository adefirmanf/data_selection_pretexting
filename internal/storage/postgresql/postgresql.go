package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/adefirmanf/data_selection_pretexting/internal/storage"
)

// Postgresql .
type Postgresql struct {
	conn *sql.DB
}

// Config .
type Config struct {
	driver           string
	connectionString string
}

// NewConfig .
func NewConfig(connectionString string) *Config {
	return &Config{
		driver:           "postgres",
		connectionString: connectionString,
	}
}

// OpenConnection .
func (c *Config) OpenConnection() *Postgresql {
	db, err := sql.Open(c.driver, c.connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return &Postgresql{
		conn: db,
	}
}

// GetUsers .
func (p *Postgresql) GetUsers() ([]*storage.User, error) {
	q := fmt.Sprintf(`select * from "users"`)
	rows, err := p.conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*storage.User

	for rows.Next() {
		var user storage.User
		var err = rows.Scan(&user.UserUsername)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

// GetUserByUserAuthorID .
func (p *Postgresql) GetUserByUserAuthorID(TweetAuthorUserID string) (*storage.User, error) {
	q := fmt.Sprintf(`select user_id, is_verified, total_following, total_followers, username, name, user_created_at from users where user_id = $1`)
	var user storage.User

	err := p.conn.QueryRow(q, TweetAuthorUserID).Scan(&user.UserID, &user.UserIsVerified, &user.UserTotalFollowing, &user.UserTotalFollowers, &user.UserUsername, &user.UserFullName, &user.UserCreatedAt)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetTweets .
func (p *Postgresql) GetTweets() ([]*storage.Tweet, error) {
	q := fmt.Sprintf("select tweet_text from tweets")
	rows, err := p.conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tweets []*storage.Tweet

	for rows.Next() {
		var tweet storage.Tweet
		var err = rows.Scan(&tweet.TweetText)
		if err != nil {
			return nil, err
		}
		tweets = append(tweets, &tweet)
	}
	return tweets, nil
}

// InsertTweets .
func (p *Postgresql) InsertTweets(TweetID, TweetAuthorID, TweetText, SuspiciousKeywords, TweetMentionedAccount,
	OptionalParameters string,
	TokenID int,
	TweetCreatedAt time.Time,
	TweetPossiblySensitive bool) error {

	now := time.Now()
	q := fmt.Sprintf(`insert into tweets(tweet_id, tweet_author_id, tweet_possibly_sensitive,
			tweet_created_at, tweet_text, created_at, updated_at, token_id, sus_keywords, 
			tweet_mentioned_account, optional_parameters) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`)

	_, err := p.conn.Exec(q, TweetID, TweetAuthorID, TweetPossiblySensitive, TweetCreatedAt, TweetText,
		now, now, TokenID, SuspiciousKeywords, TweetMentionedAccount, OptionalParameters)
	if err != nil {
		return err
	}
	return nil
}

// GetLastToken .
func (p *Postgresql) GetLastToken() (*storage.Token, error) {
	q := fmt.Sprintf(`select tweet_next_token from tokens where id=(select max(id) from tokens);`)
	var token storage.Token

	err := p.conn.QueryRow(q).Scan(&token.TweetNextToken)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

// InsertToken .
func (p *Postgresql) InsertToken(TweetNextToken, URL string) error {
	now := time.Now()
	q := fmt.Sprintf(`insert into tokens(tweet_next_token, created_at, url) values ($1, $2, $3);`)
	_, err := p.conn.Exec(q, TweetNextToken, now, URL)
	if err != nil {
		return err
	}
	return nil
}

// INSERT INTO tweets(tweet_id, tweet_author_id, tweet_possibly_sensitive," +
//             "tweet_created_at, tweet_text, created_at, updated_at, token_id," +
//             "sus_keywords, tweet_mentioned_account, optional_parameters) " +
//             "VALUES(:a, :b, :c, :d, :e, now(), now(), :f, :g, :h, :i)
