package keybindings_test

import (
	"testing"

	"github.com/Alex-Wolf-7/Kisa/keybindings"
	"github.com/stretchr/testify/assert"
)

func TestBindingMarshaling(t *testing.T) {
	key1 := []byte("")
	key2 := []byte("\"[a]\"")
	key3 := []byte("\"asdf\"")
	key4 := []byte("\"[[][]][]]\"")
	key5 := []byte("\"\"")

	var b1 keybindings.Binding
	var b2 keybindings.Binding
	var b3 keybindings.Binding
	var b4 keybindings.Binding
	var b5 keybindings.Binding

	err1 := b1.UnmarshalJSON(key1)
	err2 := b2.UnmarshalJSON(key2)
	err3 := b3.UnmarshalJSON(key3)
	err4 := b4.UnmarshalJSON(key4)
	err5 := b5.UnmarshalJSON(key5)

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Error(t, err3)
	assert.NoError(t, err4)
	assert.Error(t, err5)

	assert.Equal(t, keybindings.NewBinding(), b1)
	assert.Equal(t, keybindings.NewBinding("a"), b2)
	assert.Equal(t, keybindings.NewBinding(), b3)
	assert.Equal(t, keybindings.NewBinding("[", "]", "]"), b4)
	assert.Equal(t, keybindings.NewBinding(), b5)
}
