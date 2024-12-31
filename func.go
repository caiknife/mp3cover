package mp3cover

import (
	"io"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2/v2"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/caiknife/ncmdl/v2"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func SetCover(coverFile string, pathOrFile string) error {
	pathOrFile = ncmdl.Path(pathOrFile)
	if fileutil.IsDir(pathOrFile) {

	} else {

	}
	return nil
}

func ReadMP3FilesFromPath(path string) (files types.Slice[string]) {
	if !fileutil.IsDir(path) {
		return nil
	}

	names, err := fileutil.ListFileNames(path)
	if err != nil {
		return nil
	}
	files = lo.Filter[string](names, func(item string, index int) bool {
		ext := strings.ToLower(filepath.Ext(item))
		return ext == ".mp3"
	})
	lo.Map[string, string](files, func(item string, index int) string {
		return filepath.Join(path, item)
	})
	return files
}

func SetCoverForFile(coverFile string, fileName string) error {
	open, err := id3v2.Open(fileName, id3v2.Options{Parse: true})
	if err != nil {
		err = errors.WithMessage(err, "id3 open file")
		return err
	}
	defer open.Close()

	open.SetDefaultEncoding(id3v2.EncodingUTF8)
	file, f, err := fileutil.ReadFile(coverFile)
	if err != nil {
		err = errors.WithMessage(err, "get cover file")
		return err
	}
	defer f()

	all, err := io.ReadAll(file)
	if err != nil {
		err = errors.WithMessage(err, "read cover file")
		return err
	}
	cover := id3v2.PictureFrame{
		Encoding:    id3v2.EncodingUTF8,
		MimeType:    "image/jpeg",
		PictureType: id3v2.PTFrontCover,
		Picture:     all,
	}
	open.AddAttachedPicture(cover)

	err = open.Save()
	if err != nil {
		err = errors.WithMessage(err, "save id3")
		return err
	}
	return nil
}
