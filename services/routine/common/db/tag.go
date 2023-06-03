package db

import (
	"context"
	"routine/ent"
)

// QueryAndCreateTags queries tags from db and create tags if not exists.
// and returns tag ids.
func QueryAndCreateTags(dbClient *ent.Client, c context.Context, tags []string) ([]uint64, error) {
	return nil, nil
}
