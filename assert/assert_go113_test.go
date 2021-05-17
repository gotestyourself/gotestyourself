// +build go1.13

package assert

import (
	"fmt"
	"os"
	"testing"
)

func TestErrorIs(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		fakeT := &fakeTestingT{}

		var err error
		ErrorIs(fakeT, err, os.ErrNotExist)
		expectFailNowed(t, fakeT, "assertion failed: error is nil, not os.ErrNotExist")
	})
	t.Run("different error", func(t *testing.T) {
		fakeT := &fakeTestingT{}

		err := fmt.Errorf("the actual error")
		ErrorIs(fakeT, err, os.ErrNotExist)
		expected := `assertion failed: error is the actual error (*errors.errorString), not os.IsNotExist`
		expectFailNowed(t, fakeT, expected)
	})
}
