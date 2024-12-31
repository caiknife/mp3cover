package mp3cover

import (
	"testing"

	"github.com/caiknife/ncmdl/v2"
)

func TestReadMP3FilesFromPath(t *testing.T) {
	files := ReadMP3FilesFromPath(ncmdl.Path("."))
	files.ForEach(func(s string, i int) {
		t.Log(s)
	})
}
