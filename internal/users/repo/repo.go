package repo

import "bitmex-subscriber/pkg/db"

type Repo struct {
	*db.DB
}

func New() (*Repo, error) {
	return nil, nil
}
