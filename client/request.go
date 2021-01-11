package client

import (
	"net/url"
)

type Request interface {
	Method() string
	Params() url.Values
}
