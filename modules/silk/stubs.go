// Package silk Silk编码核心模块
package silk

import (
	"github.com/rhwong/go-cqhttp-dev/internal/base"
)

func init() {
	base.EncodeSilk = encode
	base.ResampleSilk = resample
}
