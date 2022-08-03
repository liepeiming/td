// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagesStartBotRequest represents TL type `messages.startBot#e6df7378`.
// Start a conversation with a bot using a deep linking parameter¹
//
// Links:
//  1. https://core.telegram.org/bots#deep-linking
//
// See https://core.telegram.org/method/messages.startBot for reference.
type MessagesStartBotRequest struct {
	// The bot
	Bot InputUserClass
	// The chat where to start the bot, can be the bot's private chat or a group
	Peer InputPeerClass
	// Random ID to avoid resending the same message
	RandomID int64
	// Deep linking parameter¹
	//
	// Links:
	//  1) https://core.telegram.org/bots#deep-linking
	StartParam string
}

// MessagesStartBotRequestTypeID is TL type id of MessagesStartBotRequest.
const MessagesStartBotRequestTypeID = 0xe6df7378

// Ensuring interfaces in compile-time for MessagesStartBotRequest.
var (
	_ bin.Encoder     = &MessagesStartBotRequest{}
	_ bin.Decoder     = &MessagesStartBotRequest{}
	_ bin.BareEncoder = &MessagesStartBotRequest{}
	_ bin.BareDecoder = &MessagesStartBotRequest{}
)

func (s *MessagesStartBotRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Bot == nil) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.RandomID == 0) {
		return false
	}
	if !(s.StartParam == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesStartBotRequest) String() string {
	if s == nil {
		return "MessagesStartBotRequest(nil)"
	}
	type Alias MessagesStartBotRequest
	return fmt.Sprintf("MessagesStartBotRequest%+v", Alias(*s))
}

// FillFrom fills MessagesStartBotRequest from given interface.
func (s *MessagesStartBotRequest) FillFrom(from interface {
	GetBot() (value InputUserClass)
	GetPeer() (value InputPeerClass)
	GetRandomID() (value int64)
	GetStartParam() (value string)
}) {
	s.Bot = from.GetBot()
	s.Peer = from.GetPeer()
	s.RandomID = from.GetRandomID()
	s.StartParam = from.GetStartParam()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesStartBotRequest) TypeID() uint32 {
	return MessagesStartBotRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesStartBotRequest) TypeName() string {
	return "messages.startBot"
}

// TypeInfo returns info about TL type.
func (s *MessagesStartBotRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.startBot",
		ID:   MessagesStartBotRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "StartParam",
			SchemaName: "start_param",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesStartBotRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.startBot#e6df7378 as nil")
	}
	b.PutID(MessagesStartBotRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesStartBotRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.startBot#e6df7378 as nil")
	}
	if s.Bot == nil {
		return fmt.Errorf("unable to encode messages.startBot#e6df7378: field bot is nil")
	}
	if err := s.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.startBot#e6df7378: field bot: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.startBot#e6df7378: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.startBot#e6df7378: field peer: %w", err)
	}
	b.PutLong(s.RandomID)
	b.PutString(s.StartParam)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesStartBotRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.startBot#e6df7378 to nil")
	}
	if err := b.ConsumeID(MessagesStartBotRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.startBot#e6df7378: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesStartBotRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.startBot#e6df7378 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.startBot#e6df7378: field bot: %w", err)
		}
		s.Bot = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.startBot#e6df7378: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.startBot#e6df7378: field random_id: %w", err)
		}
		s.RandomID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.startBot#e6df7378: field start_param: %w", err)
		}
		s.StartParam = value
	}
	return nil
}

// GetBot returns value of Bot field.
func (s *MessagesStartBotRequest) GetBot() (value InputUserClass) {
	if s == nil {
		return
	}
	return s.Bot
}

// GetPeer returns value of Peer field.
func (s *MessagesStartBotRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetRandomID returns value of RandomID field.
func (s *MessagesStartBotRequest) GetRandomID() (value int64) {
	if s == nil {
		return
	}
	return s.RandomID
}

// GetStartParam returns value of StartParam field.
func (s *MessagesStartBotRequest) GetStartParam() (value string) {
	if s == nil {
		return
	}
	return s.StartParam
}

// MessagesStartBot invokes method messages.startBot#e6df7378 returning error if any.
// Start a conversation with a bot using a deep linking parameter¹
//
// Links:
//  1. https://core.telegram.org/bots#deep-linking
//
// Possible errors:
//
//	400 BOT_INVALID: This is not a valid bot.
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 INPUT_USER_DEACTIVATED: The specified user was deleted.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	500 RANDOM_ID_DUPLICATE: You provided a random ID that was already used.
//	400 START_PARAM_EMPTY: The start parameter is empty.
//	400 START_PARAM_INVALID: Start parameter invalid.
//	400 START_PARAM_TOO_LONG: Start parameter is too long.
//
// See https://core.telegram.org/method/messages.startBot for reference.
func (c *Client) MessagesStartBot(ctx context.Context, request *MessagesStartBotRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
