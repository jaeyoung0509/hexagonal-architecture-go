package msgpack

import (
	"errors"

	"github.com/jaeyoung0509/hexagonal-go/shortner"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*shortner.Redirect, error) {
	redirect := &shortner.Redirect{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")

	}
	return nil
}

func (r *Redirect) Encode(input *shortner.Redirect) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
}
