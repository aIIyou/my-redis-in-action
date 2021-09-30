package chapter1

import "context"

func (rdb *RClient) AddAndRemoveGroups(articleID string, toAdd []string, toRemove []string) error {
	for _, g := range toAdd {
		key := "group:" + g
		intCmd := rdb.SAdd(context.Background(), key, articleID)
		err := intCmd.Err()
		if err != nil {
			return err
		}
	}

	for _, g := range toRemove {
		key := "group" + g
		intCmd := rdb.SRem(context.Background(), key, articleID)
		err := intCmd.Err()
		if err != nil {
			return err
		}
	}
	return nil
}
