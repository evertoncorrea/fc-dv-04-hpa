package message_test

import (
	"message"
	"testing"
)

func TestMessage(t *testing.T) {
	result := message.Message()
	if result != "Code.education Rocks! " {
		t.Errorf("Não teve resultado esperado")
	}
}
