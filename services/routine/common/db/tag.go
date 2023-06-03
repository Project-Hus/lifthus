package db

import (
	"context"
	"routine/ent"
	"routine/ent/tag"
)

// QueryAndCreateTags queries tags from db and create tags if not exists.
// and returns tag ids.
func QueryAndCreateTags(dbClient *ent.Client, c context.Context, tags []string) ([]uint64, error) {
	var tagIDs []uint64
	for _, t := range tags {
		tagID, err := dbClient.Tag.Query().Where(tag.NameEQ(t)).OnlyID(c)
		if ent.IsNotFound(err) {
			newTag, err := dbClient.Tag.Create().SetName(t).Save(c)
			tagIDs = append(tagIDs, newTag.ID)
			if err != nil {
				return nil, err
			}
		} else if err != nil {
			return nil, err
		}
		tagIDs = append(tagIDs, tagID)
	}
	return tagIDs, nil
}
