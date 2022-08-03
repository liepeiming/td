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

// KeyboardButtonTypeText represents TL type `keyboardButtonTypeText#96519938`.
type KeyboardButtonTypeText struct {
}

// KeyboardButtonTypeTextTypeID is TL type id of KeyboardButtonTypeText.
const KeyboardButtonTypeTextTypeID = 0x96519938

// construct implements constructor of KeyboardButtonTypeClass.
func (k KeyboardButtonTypeText) construct() KeyboardButtonTypeClass { return &k }

// Ensuring interfaces in compile-time for KeyboardButtonTypeText.
var (
	_ bin.Encoder     = &KeyboardButtonTypeText{}
	_ bin.Decoder     = &KeyboardButtonTypeText{}
	_ bin.BareEncoder = &KeyboardButtonTypeText{}
	_ bin.BareDecoder = &KeyboardButtonTypeText{}

	_ KeyboardButtonTypeClass = &KeyboardButtonTypeText{}
)

func (k *KeyboardButtonTypeText) Zero() bool {
	if k == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (k *KeyboardButtonTypeText) String() string {
	if k == nil {
		return "KeyboardButtonTypeText(nil)"
	}
	type Alias KeyboardButtonTypeText
	return fmt.Sprintf("KeyboardButtonTypeText%+v", Alias(*k))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*KeyboardButtonTypeText) TypeID() uint32 {
	return KeyboardButtonTypeTextTypeID
}

// TypeName returns name of type in TL schema.
func (*KeyboardButtonTypeText) TypeName() string {
	return "keyboardButtonTypeText"
}

// TypeInfo returns info about TL type.
func (k *KeyboardButtonTypeText) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "keyboardButtonTypeText",
		ID:   KeyboardButtonTypeTextTypeID,
	}
	if k == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (k *KeyboardButtonTypeText) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonTypeText#96519938 as nil")
	}
	b.PutID(KeyboardButtonTypeTextTypeID)
	return k.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (k *KeyboardButtonTypeText) EncodeBare(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonTypeText#96519938 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (k *KeyboardButtonTypeText) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonTypeText#96519938 to nil")
	}
	if err := b.ConsumeID(KeyboardButtonTypeTextTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButtonTypeText#96519938: %w", err)
	}
	return k.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (k *KeyboardButtonTypeText) DecodeBare(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonTypeText#96519938 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (k *KeyboardButtonTypeText) EncodeTDLibJSON(b tdjson.Encoder) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonTypeText#96519938 as nil")
	}
	b.ObjStart()
	b.PutID("keyboardButtonTypeText")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (k *KeyboardButtonTypeText) DecodeTDLibJSON(b tdjson.Decoder) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonTypeText#96519938 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("keyboardButtonTypeText"); err != nil {
				return fmt.Errorf("unable to decode keyboardButtonTypeText#96519938: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// KeyboardButtonTypeRequestPhoneNumber represents TL type `keyboardButtonTypeRequestPhoneNumber#a4d9b7b9`.
type KeyboardButtonTypeRequestPhoneNumber struct {
}

// KeyboardButtonTypeRequestPhoneNumberTypeID is TL type id of KeyboardButtonTypeRequestPhoneNumber.
const KeyboardButtonTypeRequestPhoneNumberTypeID = 0xa4d9b7b9

// construct implements constructor of KeyboardButtonTypeClass.
func (k KeyboardButtonTypeRequestPhoneNumber) construct() KeyboardButtonTypeClass { return &k }

// Ensuring interfaces in compile-time for KeyboardButtonTypeRequestPhoneNumber.
var (
	_ bin.Encoder     = &KeyboardButtonTypeRequestPhoneNumber{}
	_ bin.Decoder     = &KeyboardButtonTypeRequestPhoneNumber{}
	_ bin.BareEncoder = &KeyboardButtonTypeRequestPhoneNumber{}
	_ bin.BareDecoder = &KeyboardButtonTypeRequestPhoneNumber{}

	_ KeyboardButtonTypeClass = &KeyboardButtonTypeRequestPhoneNumber{}
)

func (k *KeyboardButtonTypeRequestPhoneNumber) Zero() bool {
	if k == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (k *KeyboardButtonTypeRequestPhoneNumber) String() string {
	if k == nil {
		return "KeyboardButtonTypeRequestPhoneNumber(nil)"
	}
	type Alias KeyboardButtonTypeRequestPhoneNumber
	return fmt.Sprintf("KeyboardButtonTypeRequestPhoneNumber%+v", Alias(*k))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*KeyboardButtonTypeRequestPhoneNumber) TypeID() uint32 {
	return KeyboardButtonTypeRequestPhoneNumberTypeID
}

// TypeName returns name of type in TL schema.
func (*KeyboardButtonTypeRequestPhoneNumber) TypeName() string {
	return "keyboardButtonTypeRequestPhoneNumber"
}

// TypeInfo returns info about TL type.
func (k *KeyboardButtonTypeRequestPhoneNumber) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "keyboardButtonTypeRequestPhoneNumber",
		ID:   KeyboardButtonTypeRequestPhoneNumberTypeID,
	}
	if k == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (k *KeyboardButtonTypeRequestPhoneNumber) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonTypeRequestPhoneNumber#a4d9b7b9 as nil")
	}
	b.PutID(KeyboardButtonTypeRequestPhoneNumberTypeID)
	return k.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (k *KeyboardButtonTypeRequestPhoneNumber) EncodeBare(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonTypeRequestPhoneNumber#a4d9b7b9 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (k *KeyboardButtonTypeRequestPhoneNumber) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonTypeRequestPhoneNumber#a4d9b7b9 to nil")
	}
	if err := b.ConsumeID(KeyboardButtonTypeRequestPhoneNumberTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButtonTypeRequestPhoneNumber#a4d9b7b9: %w", err)
	}
	return k.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (k *KeyboardButtonTypeRequestPhoneNumber) DecodeBare(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonTypeRequestPhoneNumber#a4d9b7b9 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (k *KeyboardButtonTypeRequestPhoneNumber) EncodeTDLibJSON(b tdjson.Encoder) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonTypeRequestPhoneNumber#a4d9b7b9 as nil")
	}
	b.ObjStart()
	b.PutID("keyboardButtonTypeRequestPhoneNumber")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (k *KeyboardButtonTypeRequestPhoneNumber) DecodeTDLibJSON(b tdjson.Decoder) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonTypeRequestPhoneNumber#a4d9b7b9 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("keyboardButtonTypeRequestPhoneNumber"); err != nil {
				return fmt.Errorf("unable to decode keyboardButtonTypeRequestPhoneNumber#a4d9b7b9: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// KeyboardButtonTypeRequestLocation represents TL type `keyboardButtonTypeRequestLocation#f8828cfd`.
type KeyboardButtonTypeRequestLocation struct {
}

// KeyboardButtonTypeRequestLocationTypeID is TL type id of KeyboardButtonTypeRequestLocation.
const KeyboardButtonTypeRequestLocationTypeID = 0xf8828cfd

// construct implements constructor of KeyboardButtonTypeClass.
func (k KeyboardButtonTypeRequestLocation) construct() KeyboardButtonTypeClass { return &k }

// Ensuring interfaces in compile-time for KeyboardButtonTypeRequestLocation.
var (
	_ bin.Encoder     = &KeyboardButtonTypeRequestLocation{}
	_ bin.Decoder     = &KeyboardButtonTypeRequestLocation{}
	_ bin.BareEncoder = &KeyboardButtonTypeRequestLocation{}
	_ bin.BareDecoder = &KeyboardButtonTypeRequestLocation{}

	_ KeyboardButtonTypeClass = &KeyboardButtonTypeRequestLocation{}
)

func (k *KeyboardButtonTypeRequestLocation) Zero() bool {
	if k == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (k *KeyboardButtonTypeRequestLocation) String() string {
	if k == nil {
		return "KeyboardButtonTypeRequestLocation(nil)"
	}
	type Alias KeyboardButtonTypeRequestLocation
	return fmt.Sprintf("KeyboardButtonTypeRequestLocation%+v", Alias(*k))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*KeyboardButtonTypeRequestLocation) TypeID() uint32 {
	return KeyboardButtonTypeRequestLocationTypeID
}

// TypeName returns name of type in TL schema.
func (*KeyboardButtonTypeRequestLocation) TypeName() string {
	return "keyboardButtonTypeRequestLocation"
}

// TypeInfo returns info about TL type.
func (k *KeyboardButtonTypeRequestLocation) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "keyboardButtonTypeRequestLocation",
		ID:   KeyboardButtonTypeRequestLocationTypeID,
	}
	if k == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (k *KeyboardButtonTypeRequestLocation) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonTypeRequestLocation#f8828cfd as nil")
	}
	b.PutID(KeyboardButtonTypeRequestLocationTypeID)
	return k.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (k *KeyboardButtonTypeRequestLocation) EncodeBare(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonTypeRequestLocation#f8828cfd as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (k *KeyboardButtonTypeRequestLocation) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonTypeRequestLocation#f8828cfd to nil")
	}
	if err := b.ConsumeID(KeyboardButtonTypeRequestLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButtonTypeRequestLocation#f8828cfd: %w", err)
	}
	return k.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (k *KeyboardButtonTypeRequestLocation) DecodeBare(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonTypeRequestLocation#f8828cfd to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (k *KeyboardButtonTypeRequestLocation) EncodeTDLibJSON(b tdjson.Encoder) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonTypeRequestLocation#f8828cfd as nil")
	}
	b.ObjStart()
	b.PutID("keyboardButtonTypeRequestLocation")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (k *KeyboardButtonTypeRequestLocation) DecodeTDLibJSON(b tdjson.Decoder) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonTypeRequestLocation#f8828cfd to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("keyboardButtonTypeRequestLocation"); err != nil {
				return fmt.Errorf("unable to decode keyboardButtonTypeRequestLocation#f8828cfd: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// KeyboardButtonTypeRequestPoll represents TL type `keyboardButtonTypeRequestPoll#7164dcb8`.
type KeyboardButtonTypeRequestPoll struct {
	// If true, only regular polls must be allowed to create
	ForceRegular bool
	// If true, only polls in quiz mode must be allowed to create
	ForceQuiz bool
}

// KeyboardButtonTypeRequestPollTypeID is TL type id of KeyboardButtonTypeRequestPoll.
const KeyboardButtonTypeRequestPollTypeID = 0x7164dcb8

// construct implements constructor of KeyboardButtonTypeClass.
func (k KeyboardButtonTypeRequestPoll) construct() KeyboardButtonTypeClass { return &k }

// Ensuring interfaces in compile-time for KeyboardButtonTypeRequestPoll.
var (
	_ bin.Encoder     = &KeyboardButtonTypeRequestPoll{}
	_ bin.Decoder     = &KeyboardButtonTypeRequestPoll{}
	_ bin.BareEncoder = &KeyboardButtonTypeRequestPoll{}
	_ bin.BareDecoder = &KeyboardButtonTypeRequestPoll{}

	_ KeyboardButtonTypeClass = &KeyboardButtonTypeRequestPoll{}
)

func (k *KeyboardButtonTypeRequestPoll) Zero() bool {
	if k == nil {
		return true
	}
	if !(k.ForceRegular == false) {
		return false
	}
	if !(k.ForceQuiz == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (k *KeyboardButtonTypeRequestPoll) String() string {
	if k == nil {
		return "KeyboardButtonTypeRequestPoll(nil)"
	}
	type Alias KeyboardButtonTypeRequestPoll
	return fmt.Sprintf("KeyboardButtonTypeRequestPoll%+v", Alias(*k))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*KeyboardButtonTypeRequestPoll) TypeID() uint32 {
	return KeyboardButtonTypeRequestPollTypeID
}

// TypeName returns name of type in TL schema.
func (*KeyboardButtonTypeRequestPoll) TypeName() string {
	return "keyboardButtonTypeRequestPoll"
}

// TypeInfo returns info about TL type.
func (k *KeyboardButtonTypeRequestPoll) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "keyboardButtonTypeRequestPoll",
		ID:   KeyboardButtonTypeRequestPollTypeID,
	}
	if k == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ForceRegular",
			SchemaName: "force_regular",
		},
		{
			Name:       "ForceQuiz",
			SchemaName: "force_quiz",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (k *KeyboardButtonTypeRequestPoll) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonTypeRequestPoll#7164dcb8 as nil")
	}
	b.PutID(KeyboardButtonTypeRequestPollTypeID)
	return k.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (k *KeyboardButtonTypeRequestPoll) EncodeBare(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonTypeRequestPoll#7164dcb8 as nil")
	}
	b.PutBool(k.ForceRegular)
	b.PutBool(k.ForceQuiz)
	return nil
}

// Decode implements bin.Decoder.
func (k *KeyboardButtonTypeRequestPoll) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonTypeRequestPoll#7164dcb8 to nil")
	}
	if err := b.ConsumeID(KeyboardButtonTypeRequestPollTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButtonTypeRequestPoll#7164dcb8: %w", err)
	}
	return k.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (k *KeyboardButtonTypeRequestPoll) DecodeBare(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonTypeRequestPoll#7164dcb8 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonTypeRequestPoll#7164dcb8: field force_regular: %w", err)
		}
		k.ForceRegular = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonTypeRequestPoll#7164dcb8: field force_quiz: %w", err)
		}
		k.ForceQuiz = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (k *KeyboardButtonTypeRequestPoll) EncodeTDLibJSON(b tdjson.Encoder) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonTypeRequestPoll#7164dcb8 as nil")
	}
	b.ObjStart()
	b.PutID("keyboardButtonTypeRequestPoll")
	b.Comma()
	b.FieldStart("force_regular")
	b.PutBool(k.ForceRegular)
	b.Comma()
	b.FieldStart("force_quiz")
	b.PutBool(k.ForceQuiz)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (k *KeyboardButtonTypeRequestPoll) DecodeTDLibJSON(b tdjson.Decoder) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonTypeRequestPoll#7164dcb8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("keyboardButtonTypeRequestPoll"); err != nil {
				return fmt.Errorf("unable to decode keyboardButtonTypeRequestPoll#7164dcb8: %w", err)
			}
		case "force_regular":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode keyboardButtonTypeRequestPoll#7164dcb8: field force_regular: %w", err)
			}
			k.ForceRegular = value
		case "force_quiz":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode keyboardButtonTypeRequestPoll#7164dcb8: field force_quiz: %w", err)
			}
			k.ForceQuiz = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetForceRegular returns value of ForceRegular field.
func (k *KeyboardButtonTypeRequestPoll) GetForceRegular() (value bool) {
	if k == nil {
		return
	}
	return k.ForceRegular
}

// GetForceQuiz returns value of ForceQuiz field.
func (k *KeyboardButtonTypeRequestPoll) GetForceQuiz() (value bool) {
	if k == nil {
		return
	}
	return k.ForceQuiz
}

// KeyboardButtonTypeWebApp represents TL type `keyboardButtonTypeWebApp#70c8ff62`.
type KeyboardButtonTypeWebApp struct {
	// An HTTP URL to pass to getWebAppUrl
	URL string
}

// KeyboardButtonTypeWebAppTypeID is TL type id of KeyboardButtonTypeWebApp.
const KeyboardButtonTypeWebAppTypeID = 0x70c8ff62

// construct implements constructor of KeyboardButtonTypeClass.
func (k KeyboardButtonTypeWebApp) construct() KeyboardButtonTypeClass { return &k }

// Ensuring interfaces in compile-time for KeyboardButtonTypeWebApp.
var (
	_ bin.Encoder     = &KeyboardButtonTypeWebApp{}
	_ bin.Decoder     = &KeyboardButtonTypeWebApp{}
	_ bin.BareEncoder = &KeyboardButtonTypeWebApp{}
	_ bin.BareDecoder = &KeyboardButtonTypeWebApp{}

	_ KeyboardButtonTypeClass = &KeyboardButtonTypeWebApp{}
)

func (k *KeyboardButtonTypeWebApp) Zero() bool {
	if k == nil {
		return true
	}
	if !(k.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (k *KeyboardButtonTypeWebApp) String() string {
	if k == nil {
		return "KeyboardButtonTypeWebApp(nil)"
	}
	type Alias KeyboardButtonTypeWebApp
	return fmt.Sprintf("KeyboardButtonTypeWebApp%+v", Alias(*k))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*KeyboardButtonTypeWebApp) TypeID() uint32 {
	return KeyboardButtonTypeWebAppTypeID
}

// TypeName returns name of type in TL schema.
func (*KeyboardButtonTypeWebApp) TypeName() string {
	return "keyboardButtonTypeWebApp"
}

// TypeInfo returns info about TL type.
func (k *KeyboardButtonTypeWebApp) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "keyboardButtonTypeWebApp",
		ID:   KeyboardButtonTypeWebAppTypeID,
	}
	if k == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (k *KeyboardButtonTypeWebApp) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonTypeWebApp#70c8ff62 as nil")
	}
	b.PutID(KeyboardButtonTypeWebAppTypeID)
	return k.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (k *KeyboardButtonTypeWebApp) EncodeBare(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonTypeWebApp#70c8ff62 as nil")
	}
	b.PutString(k.URL)
	return nil
}

// Decode implements bin.Decoder.
func (k *KeyboardButtonTypeWebApp) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonTypeWebApp#70c8ff62 to nil")
	}
	if err := b.ConsumeID(KeyboardButtonTypeWebAppTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButtonTypeWebApp#70c8ff62: %w", err)
	}
	return k.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (k *KeyboardButtonTypeWebApp) DecodeBare(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonTypeWebApp#70c8ff62 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonTypeWebApp#70c8ff62: field url: %w", err)
		}
		k.URL = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (k *KeyboardButtonTypeWebApp) EncodeTDLibJSON(b tdjson.Encoder) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonTypeWebApp#70c8ff62 as nil")
	}
	b.ObjStart()
	b.PutID("keyboardButtonTypeWebApp")
	b.Comma()
	b.FieldStart("url")
	b.PutString(k.URL)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (k *KeyboardButtonTypeWebApp) DecodeTDLibJSON(b tdjson.Decoder) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonTypeWebApp#70c8ff62 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("keyboardButtonTypeWebApp"); err != nil {
				return fmt.Errorf("unable to decode keyboardButtonTypeWebApp#70c8ff62: %w", err)
			}
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode keyboardButtonTypeWebApp#70c8ff62: field url: %w", err)
			}
			k.URL = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetURL returns value of URL field.
func (k *KeyboardButtonTypeWebApp) GetURL() (value string) {
	if k == nil {
		return
	}
	return k.URL
}

// KeyboardButtonTypeClassName is schema name of KeyboardButtonTypeClass.
const KeyboardButtonTypeClassName = "KeyboardButtonType"

// KeyboardButtonTypeClass represents KeyboardButtonType generic type.
//
// Example:
//
//	g, err := tdapi.DecodeKeyboardButtonType(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.KeyboardButtonTypeText: // keyboardButtonTypeText#96519938
//	case *tdapi.KeyboardButtonTypeRequestPhoneNumber: // keyboardButtonTypeRequestPhoneNumber#a4d9b7b9
//	case *tdapi.KeyboardButtonTypeRequestLocation: // keyboardButtonTypeRequestLocation#f8828cfd
//	case *tdapi.KeyboardButtonTypeRequestPoll: // keyboardButtonTypeRequestPoll#7164dcb8
//	case *tdapi.KeyboardButtonTypeWebApp: // keyboardButtonTypeWebApp#70c8ff62
//	default: panic(v)
//	}
type KeyboardButtonTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() KeyboardButtonTypeClass

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

// DecodeKeyboardButtonType implements binary de-serialization for KeyboardButtonTypeClass.
func DecodeKeyboardButtonType(buf *bin.Buffer) (KeyboardButtonTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case KeyboardButtonTypeTextTypeID:
		// Decoding keyboardButtonTypeText#96519938.
		v := KeyboardButtonTypeText{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	case KeyboardButtonTypeRequestPhoneNumberTypeID:
		// Decoding keyboardButtonTypeRequestPhoneNumber#a4d9b7b9.
		v := KeyboardButtonTypeRequestPhoneNumber{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	case KeyboardButtonTypeRequestLocationTypeID:
		// Decoding keyboardButtonTypeRequestLocation#f8828cfd.
		v := KeyboardButtonTypeRequestLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	case KeyboardButtonTypeRequestPollTypeID:
		// Decoding keyboardButtonTypeRequestPoll#7164dcb8.
		v := KeyboardButtonTypeRequestPoll{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	case KeyboardButtonTypeWebAppTypeID:
		// Decoding keyboardButtonTypeWebApp#70c8ff62.
		v := KeyboardButtonTypeWebApp{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode KeyboardButtonTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONKeyboardButtonType implements binary de-serialization for KeyboardButtonTypeClass.
func DecodeTDLibJSONKeyboardButtonType(buf tdjson.Decoder) (KeyboardButtonTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "keyboardButtonTypeText":
		// Decoding keyboardButtonTypeText#96519938.
		v := KeyboardButtonTypeText{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	case "keyboardButtonTypeRequestPhoneNumber":
		// Decoding keyboardButtonTypeRequestPhoneNumber#a4d9b7b9.
		v := KeyboardButtonTypeRequestPhoneNumber{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	case "keyboardButtonTypeRequestLocation":
		// Decoding keyboardButtonTypeRequestLocation#f8828cfd.
		v := KeyboardButtonTypeRequestLocation{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	case "keyboardButtonTypeRequestPoll":
		// Decoding keyboardButtonTypeRequestPoll#7164dcb8.
		v := KeyboardButtonTypeRequestPoll{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	case "keyboardButtonTypeWebApp":
		// Decoding keyboardButtonTypeWebApp#70c8ff62.
		v := KeyboardButtonTypeWebApp{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode KeyboardButtonTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode KeyboardButtonTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// KeyboardButtonType boxes the KeyboardButtonTypeClass providing a helper.
type KeyboardButtonTypeBox struct {
	KeyboardButtonType KeyboardButtonTypeClass
}

// Decode implements bin.Decoder for KeyboardButtonTypeBox.
func (b *KeyboardButtonTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode KeyboardButtonTypeBox to nil")
	}
	v, err := DecodeKeyboardButtonType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.KeyboardButtonType = v
	return nil
}

// Encode implements bin.Encode for KeyboardButtonTypeBox.
func (b *KeyboardButtonTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.KeyboardButtonType == nil {
		return fmt.Errorf("unable to encode KeyboardButtonTypeClass as nil")
	}
	return b.KeyboardButtonType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for KeyboardButtonTypeBox.
func (b *KeyboardButtonTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode KeyboardButtonTypeBox to nil")
	}
	v, err := DecodeTDLibJSONKeyboardButtonType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.KeyboardButtonType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for KeyboardButtonTypeBox.
func (b *KeyboardButtonTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.KeyboardButtonType == nil {
		return fmt.Errorf("unable to encode KeyboardButtonTypeClass as nil")
	}
	return b.KeyboardButtonType.EncodeTDLibJSON(buf)
}
