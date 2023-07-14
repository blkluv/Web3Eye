// Code generated by ent, DO NOT EDIT.

package block

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the block type in the database.
	Label = "block"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldChainType holds the string denoting the chain_type field in the database.
	FieldChainType = "chain_type"
	// FieldChainID holds the string denoting the chain_id field in the database.
	FieldChainID = "chain_id"
	// FieldBlockNumber holds the string denoting the block_number field in the database.
	FieldBlockNumber = "block_number"
	// FieldBlockHash holds the string denoting the block_hash field in the database.
	FieldBlockHash = "block_hash"
	// FieldBlockTime holds the string denoting the block_time field in the database.
	FieldBlockTime = "block_time"
	// Table holds the table name of the block in the database.
	Table = "blocks"
)

// Columns holds all SQL columns for block fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldChainType,
	FieldChainID,
	FieldBlockNumber,
	FieldBlockHash,
	FieldBlockTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/runtime"
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)