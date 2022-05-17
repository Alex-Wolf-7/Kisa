package keybindings

import (
	"fmt"
	"strings"
)

const KEY_FORMAT = "[%s]"

type Binding struct {
	keys []string
}

func NewBinding(keys ...string) Binding {
	return Binding{
		keys: keys,
	}
}

func (b Binding) MarshalJSON() ([]byte, error) {
	var builder strings.Builder
	builder.WriteRune('"')
	for _, key := range b.keys {
		builder.WriteString(fmt.Sprintf(KEY_FORMAT, key))
	}
	builder.WriteRune('"')

	strBinding := builder.String()
	return []byte(strBinding), nil
}

func (b *Binding) UnmarshalJSON(data []byte) error {
	jsonString := string(data)
	if jsonString == "\"\"" {
		return nil
	}

	if len(data) < 4 || jsonString[0:2] != "\"[" || jsonString[len(jsonString)-2:] != "]\"" {
		return fmt.Errorf("Binding not in expected format: %s", jsonString)
	}

	jsonString = jsonString[2 : len(jsonString)-2]

	b.keys = strings.Split(jsonString, "][")
	return nil
}
