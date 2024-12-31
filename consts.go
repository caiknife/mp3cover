package mp3cover

import (
	"github.com/caiknife/mp3lister/lib/types"
)

const (
	ErrInputIsEmpty  types.Error = "请输入MP3文件路径"
	ErrCoverIsEmpty  types.Error = "请输入封面文件路径"
	ErrCoverNotExist types.Error = "封面文件不存在"
)
