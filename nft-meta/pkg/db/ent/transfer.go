// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/transfer"
)

// Transfer is the model entity for the Transfer schema.
type Transfer struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// ChainType holds the value of the "chain_type" field.
	ChainType string `json:"chain_type,omitempty"`
	// ChainID holds the value of the "chain_id" field.
	ChainID string `json:"chain_id,omitempty"`
	// Contract holds the value of the "contract" field.
	Contract string `json:"contract,omitempty"`
	// TokenType holds the value of the "token_type" field.
	TokenType string `json:"token_type,omitempty"`
	// TokenID holds the value of the "token_id" field.
	TokenID string `json:"token_id,omitempty"`
	// From holds the value of the "from" field.
	From string `json:"from,omitempty"`
	// To holds the value of the "to" field.
	To string `json:"to,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount uint64 `json:"amount,omitempty"`
	// BlockNumber holds the value of the "block_number" field.
	BlockNumber uint64 `json:"block_number,omitempty"`
	// TxHash holds the value of the "tx_hash" field.
	TxHash string `json:"tx_hash,omitempty"`
	// BlockHash holds the value of the "block_hash" field.
	BlockHash string `json:"block_hash,omitempty"`
	// TxTime holds the value of the "tx_time" field.
	TxTime uint32 `json:"tx_time,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transfer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case transfer.FieldCreatedAt, transfer.FieldUpdatedAt, transfer.FieldDeletedAt, transfer.FieldAmount, transfer.FieldBlockNumber, transfer.FieldTxTime:
			values[i] = new(sql.NullInt64)
		case transfer.FieldChainType, transfer.FieldChainID, transfer.FieldContract, transfer.FieldTokenType, transfer.FieldTokenID, transfer.FieldFrom, transfer.FieldTo, transfer.FieldTxHash, transfer.FieldBlockHash, transfer.FieldRemark:
			values[i] = new(sql.NullString)
		case transfer.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Transfer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transfer fields.
func (t *Transfer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transfer.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case transfer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = uint32(value.Int64)
			}
		case transfer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = uint32(value.Int64)
			}
		case transfer.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				t.DeletedAt = uint32(value.Int64)
			}
		case transfer.FieldChainType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chain_type", values[i])
			} else if value.Valid {
				t.ChainType = value.String
			}
		case transfer.FieldChainID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chain_id", values[i])
			} else if value.Valid {
				t.ChainID = value.String
			}
		case transfer.FieldContract:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contract", values[i])
			} else if value.Valid {
				t.Contract = value.String
			}
		case transfer.FieldTokenType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_type", values[i])
			} else if value.Valid {
				t.TokenType = value.String
			}
		case transfer.FieldTokenID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_id", values[i])
			} else if value.Valid {
				t.TokenID = value.String
			}
		case transfer.FieldFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				t.From = value.String
			}
		case transfer.FieldTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field to", values[i])
			} else if value.Valid {
				t.To = value.String
			}
		case transfer.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				t.Amount = uint64(value.Int64)
			}
		case transfer.FieldBlockNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field block_number", values[i])
			} else if value.Valid {
				t.BlockNumber = uint64(value.Int64)
			}
		case transfer.FieldTxHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tx_hash", values[i])
			} else if value.Valid {
				t.TxHash = value.String
			}
		case transfer.FieldBlockHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field block_hash", values[i])
			} else if value.Valid {
				t.BlockHash = value.String
			}
		case transfer.FieldTxTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tx_time", values[i])
			} else if value.Valid {
				t.TxTime = uint32(value.Int64)
			}
		case transfer.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				t.Remark = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Transfer.
// Note that you need to call Transfer.Unwrap() before calling this method if this Transfer
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transfer) Update() *TransferUpdateOne {
	return (&TransferClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Transfer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transfer) Unwrap() *Transfer {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transfer is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transfer) String() string {
	var builder strings.Builder
	builder.WriteString("Transfer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", t.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", t.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", t.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("chain_type=")
	builder.WriteString(t.ChainType)
	builder.WriteString(", ")
	builder.WriteString("chain_id=")
	builder.WriteString(t.ChainID)
	builder.WriteString(", ")
	builder.WriteString("contract=")
	builder.WriteString(t.Contract)
	builder.WriteString(", ")
	builder.WriteString("token_type=")
	builder.WriteString(t.TokenType)
	builder.WriteString(", ")
	builder.WriteString("token_id=")
	builder.WriteString(t.TokenID)
	builder.WriteString(", ")
	builder.WriteString("from=")
	builder.WriteString(t.From)
	builder.WriteString(", ")
	builder.WriteString("to=")
	builder.WriteString(t.To)
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", t.Amount))
	builder.WriteString(", ")
	builder.WriteString("block_number=")
	builder.WriteString(fmt.Sprintf("%v", t.BlockNumber))
	builder.WriteString(", ")
	builder.WriteString("tx_hash=")
	builder.WriteString(t.TxHash)
	builder.WriteString(", ")
	builder.WriteString("block_hash=")
	builder.WriteString(t.BlockHash)
	builder.WriteString(", ")
	builder.WriteString("tx_time=")
	builder.WriteString(fmt.Sprintf("%v", t.TxTime))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(t.Remark)
	builder.WriteByte(')')
	return builder.String()
}

// Transfers is a parsable slice of Transfer.
type Transfers []*Transfer

func (t Transfers) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
