package g

import (
	"errors"
	"testing"
)

func TestFoo(t *testing.T) {
    t.Errorf("Got bar error: %g", errors.New("bar!"))
}
