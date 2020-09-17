package types

import (
	"errors"
)

type Data = int

type Size = int

type Less func(Data, Data) bool

var ErrUnimplemented = errors.New("unimplemented")
