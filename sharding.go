// @Author: Jiekun
// @Date: 2022/8/8

package sharding

import "fmt"

type Sharding interface {
	AddConfig(Config) error
	Rewrite(sql string) string
}

type sharding struct{}

func NewSharding() Sharding {
	return &sharding{}
}

func (s *sharding) AddConfig(config Config) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := configMap.Load(config.tableName); ok {
		return errDuplicateConfig
	}

	configMap.Store(config.tableName, config)
	return nil
}

func (s *sharding) Rewrite(sql string) string {
	stmtNode, err := s.parse(sql)
	if err != nil {
		return sql
	}
	fmt.Println((*stmtNode).Text())
	return sql
}
