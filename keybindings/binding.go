package keybindings

import (
	"fmt"
	"strings"
)

const KEY_FORMAT = "[%s]"

type Binding []string

func (b Binding) MarshalJSON() ([]byte, error) {
	var builder strings.Builder
	builder.WriteRune('"')
	for _, key := range b {
		builder.WriteString(fmt.Sprintf(KEY_FORMAT, key))
	}
	builder.WriteRune('"')

	strBinding := builder.String()
	return []byte(strBinding), nil
}
