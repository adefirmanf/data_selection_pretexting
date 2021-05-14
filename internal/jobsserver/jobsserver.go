package jobsserver

import (
	"github.com/adefirmanf/data_selection_pretexting/internal/queue"
	"github.com/adefirmanf/data_selection_pretexting/internal/storage"
)

// Config .
type Config struct {
	IntervalTime int
	Range        string
}

// NewConfig .
func NewConfig(IntervalTime int, Range string) *Config {
	return &Config{
		IntervalTime: IntervalTime,
		Range:        Range,
	}
}

// JobServer .
type JobServer struct {
	*Config
	storage *storage.Storage
	queue   *queue.Queue
}

// NewJobServer .
func NewJobServer(cfg *Config, storage *storage.Storage, q *queue.Queue) *JobServer {
	return &JobServer{
		Config:  cfg,
		storage: storage,
		queue:   q,
	}
}
