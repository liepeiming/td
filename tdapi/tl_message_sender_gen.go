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

// MessageSenderUser represents TL type `messageSenderUser#ebf760e3`.
type MessageSenderUser struct {
	// Identifier of the user that sent the message
	UserID int64
}

// MessageSenderUserTypeID is TL type id of MessageSenderUser.
const MessageSenderUserTypeID = 0xebf760e3

// construct implements constructor of MessageSenderClass.
func (m MessageSenderUser) construct() MessageSenderClass { return &m }

// Ensuring interfaces in compile-time for MessageSenderUser.
var (
	_ bin.Encoder     = &MessageSenderUser{}
	_ bin.Decoder     = &MessageSenderUser{}
	_ bin.BareEncoder = &MessageSenderUser{}
	_ bin.BareDecoder = &MessageSenderUser{}

	_ MessageSenderClass = &MessageSenderUser{}
)

func (m *MessageSenderUser) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageSenderUser) String() string {
	if m == nil {
		return "MessageSenderUser(nil)"
	}
	type Alias MessageSenderUser
	return fmt.Sprintf("MessageSenderUser%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageSenderUser) TypeID() uint32 {
	return MessageSenderUserTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageSenderUser) TypeName() string {
	return "messageSenderUser"
}

// TypeInfo returns info about TL type.
func (m *MessageSenderUser) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageSenderUser",
		ID:   MessageSenderUserTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageSenderUser) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSenderUser#ebf760e3 as nil")
	}
	b.PutID(MessageSenderUserTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageSenderUser) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSenderUser#ebf760e3 as nil")
	}
	b.PutInt53(m.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageSenderUser) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSenderUser#ebf760e3 to nil")
	}
	if err := b.ConsumeID(MessageSenderUserTypeID); err != nil {
		return fmt.Errorf("unable to decode messageSenderUser#ebf760e3: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageSenderUser) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSenderUser#ebf760e3 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode messageSenderUser#ebf760e3: field user_id: %w", err)
		}
		m.UserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageSenderUser) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSenderUser#ebf760e3 as nil")
	}
	b.ObjStart()
	b.PutID("messageSenderUser")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(m.UserID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageSenderUser) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSenderUser#ebf760e3 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageSenderUser"); err != nil {
				return fmt.Errorf("unable to decode messageSenderUser#ebf760e3: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode messageSenderUser#ebf760e3: field user_id: %w", err)
			}
			m.UserID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (m *MessageSenderUser) GetUserID() (value int64) {
	if m == nil {
		return
	}
	return m.UserID
}

// MessageSenderChat represents TL type `messageSenderChat#f1b71131`.
type MessageSenderChat struct {
	// Identifier of the chat that sent the message
	ChatID int64
}

// MessageSenderChatTypeID is TL type id of MessageSenderChat.
const MessageSenderChatTypeID = 0xf1b71131

// construct implements constructor of MessageSenderClass.
func (m MessageSenderChat) construct() MessageSenderClass { return &m }

// Ensuring interfaces in compile-time for MessageSenderChat.
var (
	_ bin.Encoder     = &MessageSenderChat{}
	_ bin.Decoder     = &MessageSenderChat{}
	_ bin.BareEncoder = &MessageSenderChat{}
	_ bin.BareDecoder = &MessageSenderChat{}

	_ MessageSenderClass = &MessageSenderChat{}
)

func (m *MessageSenderChat) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageSenderChat) String() string {
	if m == nil {
		return "MessageSenderChat(nil)"
	}
	type Alias MessageSenderChat
	return fmt.Sprintf("MessageSenderChat%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageSenderChat) TypeID() uint32 {
	return MessageSenderChatTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageSenderChat) TypeName() string {
	return "messageSenderChat"
}

// TypeInfo returns info about TL type.
func (m *MessageSenderChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageSenderChat",
		ID:   MessageSenderChatTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageSenderChat) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSenderChat#f1b71131 as nil")
	}
	b.PutID(MessageSenderChatTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageSenderChat) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSenderChat#f1b71131 as nil")
	}
	b.PutInt53(m.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageSenderChat) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSenderChat#f1b71131 to nil")
	}
	if err := b.ConsumeID(MessageSenderChatTypeID); err != nil {
		return fmt.Errorf("unable to decode messageSenderChat#f1b71131: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageSenderChat) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSenderChat#f1b71131 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode messageSenderChat#f1b71131: field chat_id: %w", err)
		}
		m.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageSenderChat) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSenderChat#f1b71131 as nil")
	}
	b.ObjStart()
	b.PutID("messageSenderChat")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(m.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageSenderChat) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSenderChat#f1b71131 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageSenderChat"); err != nil {
				return fmt.Errorf("unable to decode messageSenderChat#f1b71131: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode messageSenderChat#f1b71131: field chat_id: %w", err)
			}
			m.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (m *MessageSenderChat) GetChatID() (value int64) {
	if m == nil {
		return
	}
	return m.ChatID
}

// MessageSenderClassName is schema name of MessageSenderClass.
const MessageSenderClassName = "MessageSender"

// MessageSenderClass represents MessageSender generic type.
//
// Example:
//
//	g, err := tdapi.DecodeMessageSender(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.MessageSenderUser: // messageSenderUser#ebf760e3
//	case *tdapi.MessageSenderChat: // messageSenderChat#f1b71131
//	default: panic(v)
//	}
type MessageSenderClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessageSenderClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeMessageSender implements binary de-serialization for MessageSenderClass.
func DecodeMessageSender(buf *bin.Buffer) (MessageSenderClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessageSenderUserTypeID:
		// Decoding messageSenderUser#ebf760e3.
		v := MessageSenderUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageSenderClass: %w", err)
		}
		return &v, nil
	case MessageSenderChatTypeID:
		// Decoding messageSenderChat#f1b71131.
		v := MessageSenderChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageSenderClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageSenderClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONMessageSender implements binary de-serialization for MessageSenderClass.
func DecodeTDLibJSONMessageSender(buf tdjson.Decoder) (MessageSenderClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "messageSenderUser":
		// Decoding messageSenderUser#ebf760e3.
		v := MessageSenderUser{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageSenderClass: %w", err)
		}
		return &v, nil
	case "messageSenderChat":
		// Decoding messageSenderChat#f1b71131.
		v := MessageSenderChat{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageSenderClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageSenderClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// MessageSender boxes the MessageSenderClass providing a helper.
type MessageSenderBox struct {
	MessageSender MessageSenderClass
}

// Decode implements bin.Decoder for MessageSenderBox.
func (b *MessageSenderBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageSenderBox to nil")
	}
	v, err := DecodeMessageSender(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessageSender = v
	return nil
}

// Encode implements bin.Encode for MessageSenderBox.
func (b *MessageSenderBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.MessageSender == nil {
		return fmt.Errorf("unable to encode MessageSenderClass as nil")
	}
	return b.MessageSender.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for MessageSenderBox.
func (b *MessageSenderBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageSenderBox to nil")
	}
	v, err := DecodeTDLibJSONMessageSender(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessageSender = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for MessageSenderBox.
func (b *MessageSenderBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.MessageSender == nil {
		return fmt.Errorf("unable to encode MessageSenderClass as nil")
	}
	return b.MessageSender.EncodeTDLibJSON(buf)
}
