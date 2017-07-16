package helper

import (
	"bytes"
	"fmt"
	"strings"
)

func Strcat(a string, b string) string {
	var buf bytes.Buffer
	buf.WriteString(a)
	buf.WriteString(b)
	return buf.String()
}

func Ucfirst(str string) string {
	fw := string(str[0])
	_str := fmt.Sprintf("%s%s", strings.ToUpper(fw), str[1:])
	return _str
}
