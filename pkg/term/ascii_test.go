package term // import "github.com/dims/mobyutils/pkg/term"

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestToBytes(t *testing.T) {
	codes, err := ToBytes("ctrl-a,a")
	assert.NilError(t, err)
	assert.Check(t, is.DeepEqual([]byte{1, 97}, codes))

	_, err = ToBytes("shift-z")
	assert.Check(t, is.ErrorContains(err, ""))

	codes, err = ToBytes("ctrl-@,ctrl-[,~,ctrl-o")
	assert.NilError(t, err)
	assert.Check(t, is.DeepEqual([]byte{0, 27, 126, 15}, codes))

	codes, err = ToBytes("DEL,+")
	assert.NilError(t, err)
	assert.Check(t, is.DeepEqual([]byte{127, 43}, codes))
}
