package common

import (
	"sync"
	"github.com/unrolled/render"
)

var (
	formatter *render.Render
	formatterOnce sync.Once
)

func GetFormatter() *render.Render {

	// 创建格式化模板
	formatterOnce.Do(func(){
		formatter = render.New(render.Options{IndentJSON: true})
	})
	return formatter
}


