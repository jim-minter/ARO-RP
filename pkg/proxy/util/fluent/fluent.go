package fluent

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
	"github.com/ugorji/go/codec"

	"github.com/Azure/ARO-RP/pkg/util/recover"
)

type conn struct {
	log *logrus.Entry

	h *codec.MsgpackHandle
	d *codec.Decoder
	e *codec.Encoder

	ch chan *Message
}

func Messages(log *logrus.Entry, rw io.ReadWriter) <-chan *Message {
	h := &codec.MsgpackHandle{
		BasicHandle: codec.BasicHandle{
			DecodeOptions: codec.DecodeOptions{
				RawToString: true,
			},
		},
	}

	c := &conn{
		log: log,

		h: h,
		d: codec.NewDecoder(rw, h),
		e: codec.NewEncoder(rw, h),

		ch: make(chan *Message),
	}

	go c.decode()

	return c.ch
}

func (c *conn) decode() error {
	defer recover.Panic(c.log)
	defer close(c.ch)

	for {
		err := c.decodeOne()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			c.log.Error(err)
			return err
		}
	}
}

func (c *conn) decodeOne() error {
	var b codec.Raw

	err := c.d.Decode(&b)
	if err != nil {
		return err
	}

	{
		var f *Forward
		err := codec.NewDecoderBytes(b, c.h).Decode(&f)
		if err == nil {
			for _, e := range f.Entries {
				c.ch <- &Message{
					Tag:    f.Tag,
					Time:   e.Time,
					Record: e.Record,
					Option: f.Option,
				}
			}

			if chunk, ok := f.Option["chunk"].(string); ok {
				err = c.e.Encode(&Ack{Ack: chunk})
			}

			return err
		}
	}

	{
		var m *Message
		err := codec.NewDecoderBytes(b, c.h).Decode(&m)
		if err == nil {
			c.ch <- m
			return err
		}
	}

	return fmt.Errorf("unrecognised message")
}
