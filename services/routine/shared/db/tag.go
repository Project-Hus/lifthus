package db

import (
	"context"
	"routine/ent"
	"routine/ent/tag"
)

// QueryAndCreateTags queries tags from db and create tags if not exists.
// and returns tag ids.
func QueryAndCreateTags(dbClient *ent.Client, c context.Context, tags []string) ([]uint64, error) {
	tagIDs := make([]uint64, 0, len(tags))
	for _, t := range tags {
		tagID, err := dbClient.Tag.Query().Where(tag.NameEQ(t)).OnlyID(c)
		if ent.IsNotFound(err) {
			newTag, err := dbClient.Tag.Create().SetName(t).Save(c)
			if err != nil {
				return nil, err
			}
			tagIDs = append(tagIDs, newTag.ID)
		} else if err != nil {
			return nil, err
		} else {
			tagIDs = append(tagIDs, tagID)
		}
	}
	return tagIDs, nil
}
