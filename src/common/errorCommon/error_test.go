package errorCommon

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewError(t *testing.T) {
	t.Run("new error msg equal test", func(t *testing.T) {
		got := NewError(ErrBadParameter, "ryan/src/test.go", "error test", ErrFromClient)
		want := fmt.Errorf("PARAM_BAD|ryan/src/test.go|error test|client")
		assert.Equal(t, got, want)
	})
}

func TestErrorParsing(t *testing.T) {
	t.Run("error parsing msg equal test", func(t *testing.T) {

		got := ErrorParsing("PARAM_BAD|ryan/src/test.go|error test|client")
		want := Err{400, "PARAM_BAD", "error test", "ryan/src/test.go", "client"}
		assert.Equal(t, got, want)
	})
}
