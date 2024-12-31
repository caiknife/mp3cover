package main

import (
	"os"
	"path/filepath"

	"github.com/caiknife/mp3lister/lib/types"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/caiknife/ncmdl/v2"

	"github.com/caiknife/mp3cover"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			ncmdl.AppLogger.Fatalln("程序发生了异常", r)
		}
	}()

	if err := newApp().Run(os.Args); err != nil {
		ncmdl.AppLogger.Fatalln(err)
		return
	}
}

func newApp() *cli.App {
	app := &cli.App{
		Name:  "MP3封面设置工具",
		Usage: "设置MP3文件的封面",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "cover",
				Aliases:  []string{"c"},
				Usage:    "封面文件路径",
				Required: true,
				Value:    "",
			},
		},
		Action: action(),
	}

	return app
}

func action() cli.ActionFunc {
	return func(c *cli.Context) error {
		files := types.Slice[string](c.Args().Slice())
		if files.IsEmpty() {
			return mp3cover.ErrInputIsEmpty
		}

		cover := c.String("cover")
		if cover == "" {
			return mp3cover.ErrCoverIsEmpty
		}

		absCover, err := filepath.Abs(cover)
		if err != nil {
			err = errors.WithMessage(err, "get abs cover")
			return err
		}

		if !fileutil.IsExist(absCover) {
			return mp3cover.ErrCoverNotExist
		}

		files.ForEach(func(s string, i int) {
			err := mp3cover.SetCover(absCover, s)
			if err != nil {
				err = errors.WithMessage(err, "set cover")
				ncmdl.AppLogger.Errorln(err)
				return
			}
		})

		return nil
	}
}
