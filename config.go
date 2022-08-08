// @Author: Jiekun
// @Date: 2022/8/8

package sharding

import "sync"

var (
	configMap sync.Map
	m         sync.Mutex
)

type Config struct {
	tableName string
	key       string
	algorithm func(value interface{}) (tableName string)
}
