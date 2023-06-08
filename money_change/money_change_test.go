package money_change

import (
	"testing"
)

func TestChange(t *testing.T) {
	if Change(28) != 6 {
		t.Errorf("Expected 6, got %d", Change(28))
	}
}
