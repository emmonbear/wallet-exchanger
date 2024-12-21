package postgres

import (
	"context"
	"strings"
)

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (c *Connector) Create(ctx context.Context, report interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c *Connector) DeleteAll(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
