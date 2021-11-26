// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// SetSupergroupUsernameRequest represents TL type `setSupergroupUsername#503f4b04`.
type SetSupergroupUsernameRequest struct {
	// Identifier of the supergroup or channel
	SupergroupID int64
	// New value of the username. Use an empty string to remove the username
	Username string
}

// SetSupergroupUsernameRequestTypeID is TL type id of SetSupergroupUsernameRequest.
const SetSupergroupUsernameRequestTypeID = 0x503f4b04

// Ensuring interfaces in compile-time for SetSupergroupUsernameRequest.
var (
	_ bin.Encoder     = &SetSupergroupUsernameRequest{}
	_ bin.Decoder     = &SetSupergroupUsernameRequest{}
	_ bin.BareEncoder = &SetSupergroupUsernameRequest{}
	_ bin.BareDecoder = &SetSupergroupUsernameRequest{}
)

func (s *SetSupergroupUsernameRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.SupergroupID == 0) {
		return false
	}
	if !(s.Username == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetSupergroupUsernameRequest) String() string {
	if s == nil {
		return "SetSupergroupUsernameRequest(nil)"
	}
	type Alias SetSupergroupUsernameRequest
	return fmt.Sprintf("SetSupergroupUsernameRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetSupergroupUsernameRequest) TypeID() uint32 {
	return SetSupergroupUsernameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetSupergroupUsernameRequest) TypeName() string {
	return "setSupergroupUsername"
}

// TypeInfo returns info about TL type.
func (s *SetSupergroupUsernameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setSupergroupUsername",
		ID:   SetSupergroupUsernameRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
		{
			Name:       "Username",
			SchemaName: "username",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetSupergroupUsernameRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setSupergroupUsername#503f4b04 as nil")
	}
	b.PutID(SetSupergroupUsernameRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetSupergroupUsernameRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setSupergroupUsername#503f4b04 as nil")
	}
	b.PutLong(s.SupergroupID)
	b.PutString(s.Username)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetSupergroupUsernameRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setSupergroupUsername#503f4b04 to nil")
	}
	if err := b.ConsumeID(SetSupergroupUsernameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setSupergroupUsername#503f4b04: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetSupergroupUsernameRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setSupergroupUsername#503f4b04 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode setSupergroupUsername#503f4b04: field supergroup_id: %w", err)
		}
		s.SupergroupID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setSupergroupUsername#503f4b04: field username: %w", err)
		}
		s.Username = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetSupergroupUsernameRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setSupergroupUsername#503f4b04 as nil")
	}
	b.ObjStart()
	b.PutID("setSupergroupUsername")
	b.FieldStart("supergroup_id")
	b.PutLong(s.SupergroupID)
	b.FieldStart("username")
	b.PutString(s.Username)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetSupergroupUsernameRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setSupergroupUsername#503f4b04 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setSupergroupUsername"); err != nil {
				return fmt.Errorf("unable to decode setSupergroupUsername#503f4b04: %w", err)
			}
		case "supergroup_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode setSupergroupUsername#503f4b04: field supergroup_id: %w", err)
			}
			s.SupergroupID = value
		case "username":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setSupergroupUsername#503f4b04: field username: %w", err)
			}
			s.Username = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (s *SetSupergroupUsernameRequest) GetSupergroupID() (value int64) {
	return s.SupergroupID
}

// GetUsername returns value of Username field.
func (s *SetSupergroupUsernameRequest) GetUsername() (value string) {
	return s.Username
}

// SetSupergroupUsername invokes method setSupergroupUsername#503f4b04 returning error if any.
func (c *Client) SetSupergroupUsername(ctx context.Context, request *SetSupergroupUsernameRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}