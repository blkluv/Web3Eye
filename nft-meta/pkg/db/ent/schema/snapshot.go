package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
)

type Snapshot struct {
	ent.Schema
}

func (Snapshot) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (Snapshot) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique(),
		field.Uint64("index"),
		field.String("snapshot_comm_p"),
		field.String("snapshot_root"),
		field.String("snapshot_uri"),
		field.String("backup_state"),
	}
}

func (Snapshot) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("index", "backup_state").
			Unique(),
	}
}
