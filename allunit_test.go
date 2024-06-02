package unit

import (
	"testing"
)

func Test_CheckAllValid(t *testing.T) {
	if err := CheckAllValid(ALL); err != nil {
		t.Error(err)
	}
}
