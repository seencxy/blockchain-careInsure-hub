package pkg

import (
	"backed/app/core/consts"
	"gorm.io/gen"
)

// NewGenerator 生成mysql操作代码
func NewGenerator(moduleName string) *gen.Generator {
	generator := gen.NewGenerator(gen.Config{
		OutPath: consts.MysqlFilePath,
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	return generator
}
