package core

import (
	"github.com/redis/go-redis/v9"
	"healthmonitor/ent"
)

// Data holds the database and Redis clients.
type Data struct {
	DB  *ent.Client
	RDB *redis.Client
}
