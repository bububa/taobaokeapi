package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"

	"github.com/bububa/taobaokeapi/util/token"
)

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func Nonce() string {
	t := token.New()
	return t.Encode()
}
