package test

import (
	"context"
	"testing"

	"entexample/pkg/ent"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

func NewEntClient(t *testing.T) *ent.Client {
	t.Helper()

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		t.Error(err)
	}

	err = client.Schema.Create(context.Background())
	if err != nil {
		t.Error(err)
	}

	return client
}
