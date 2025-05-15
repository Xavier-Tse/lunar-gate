package query

import (
	"fmt"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"gorm.io/gorm"
)

type Option struct {
	Page     model.Page
	Likes    []string
	Where    *gorm.DB
	Debug    bool
	Preloads []string
	Callback func(list any)
}

func List[T any](model T, option Option) (list []T, count int64, err error) {
	baseDB := global.DB.Model(model).Where(model)
	if option.Debug {
		baseDB = baseDB.Debug()
	}
	if option.Where != nil {
		baseDB = baseDB.Where(option.Where)
	}

	if option.Page.Key != "" && len(option.Likes) != 0 {
		query := global.DB.Where("")
		for _, column := range option.Likes {
			query.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Page.Key))
		}
		baseDB = baseDB.Where(query)
	}
	for _, preload := range option.Preloads {
		baseDB = baseDB.Preload(preload)
	}

	if option.Page.Limit <= 0 {
		option.Page.Limit = 10
	}
	if option.Page.Page <= 0 {
		option.Page.Page = 1
	}

	offset := (option.Page.Page - 1) * option.Page.Limit
	if option.Page.Sort == "" {
		option.Page.Sort = "created_at desc"
	}

	baseDB.Limit(option.Page.Limit).Offset(offset).Order(option.Page.Sort).Find(&list)
	if option.Callback != nil {
		option.Callback(list)
	}
	baseDB.Model(model).Count(&count)
	return
}
