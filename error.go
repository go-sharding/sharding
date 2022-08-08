// @Author: Jiekun
// @Date: 2022/8/8

package sharding

import "errors"

var (
	errDuplicateConfig = errors.New("duplicate config")
	errParseTooMany    = errors.New("too many sql to parse")
)
