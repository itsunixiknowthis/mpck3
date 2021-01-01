package abc

import "testing"
import "github.com/itsunixiknowthis/mpck3/b"

func TestA(t *testing.T) {
    if 788 != A { t.Errorf("!") }
    if false != B { t.Errorf("?") }
    if notb.X != 1243 { t.Errorf("@") }
}

