package fluent

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/ugorji/go/codec"
)

// https://github.com/fluent/fluentd/wiki/Forward-Protocol-Specification-v0

type Message struct {
	Tag    string
	Time   time.Time
	Record map[string]interface{}
	Option map[string]interface{}
}

func (m *Message) CodecEncodeSelf(*codec.Encoder) {
	panic("not implemented")
}

func (m *Message) CodecDecodeSelf(d *codec.Decoder) {
	var _m *struct {
		Tag    string
		Time   interface{}
		Record map[string]interface{}
		Option map[string]interface{}
	}

	d.MustDecode(&_m)

	*m = Message{
		Tag:    _m.Tag,
		Time:   decodeTime(_m.Time),
		Record: _m.Record,
		Option: _m.Option,
	}
}

type Forward struct {
	Tag     string
	Entries []*Entry
	Option  map[string]interface{}
}

type Entry struct {
	Time   time.Time
	Record map[string]interface{}
}

func (e *Entry) CodecEncodeSelf(*codec.Encoder) {
	panic("not implemented")
}

func (e *Entry) CodecDecodeSelf(d *codec.Decoder) {
	var _e *struct {
		Time   interface{}
		Record map[string]interface{}
	}

	d.MustDecode(&_e)

	*e = Entry{
		Time:   decodeTime(_e.Time),
		Record: _e.Record,
	}
}

type Ack struct {
	Ack string `codec:"ack,omitempty"`
}

func decodeTime(i interface{}) time.Time {
	switch i := i.(type) {
	case uint64:
		return time.Unix(int64(i), 0)

	case codec.RawExt:
		if i.Tag != 0 {
			panic(fmt.Sprintf("invalid tag %d", i.Tag))
		}
		if len(i.Data) != 8 {
			panic(fmt.Sprintf("invalid data length %d", len(i.Data)))
		}

		sec := binary.BigEndian.Uint32(i.Data[0:4])
		nsec := binary.BigEndian.Uint32(i.Data[4:8])

		return time.Unix(int64(sec), int64(nsec))
	}

	panic(fmt.Sprintf("invalid type %T", i))
}
