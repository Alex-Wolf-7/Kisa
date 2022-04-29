package keybindings

import (
	"fmt"
	"strings"
)

const KEY_FORMAT = "[%s]"

type Binding string

func NewBinding(keys []string) Binding {
	var builder strings.Builder
	for _, key := range keys {
		builder.WriteString(fmt.Sprintf(KEY_FORMAT, key))
	}

	return Binding(builder.String())
}
