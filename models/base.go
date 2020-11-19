package models

import (
	"car-tms/libs"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"gorm.io/gorm"
	"strconv"
)

// Filed 查询字段结构体
type Filed struct {
	Condition string      `json:"condition"`
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
}

// Search 查询参数结构体
type Search struct {
	Fields  []*Filed `json:"fields"`
	OrderBy string   `json:"order_by"`
	Sort    string   `json:"sort"`
	Sorter  string   `json:"sorter"`
	Limit   int      `json:"limit"`
	Offset  int      `json:"offset"`
}

// Found 查询条件
func Found(s *Search) *gorm.DB {
	color.Yellow(fmt.Sprintf("Searach :%+v", s))
	return libs.Db.Scopes(FoundByWhere(s.Fields), FoundOrder(s.OrderBy, s.Sort))
}

func IsNotFound(err error) bool {
	if ok := errors.Is(err, gorm.ErrRecordNotFound); ok {
		color.Yellow("查询数据不存在")
		return true
	}
	return false
}

// GetAll 批量查询
func GetAll(model interface{}, s *Search) *gorm.DB {
	db := libs.Db.Model(model)
	sort := "desc"
	orderBy := "created_at"
	if len(s.Sort) > 0 {
		sort = s.Sort
	}
	if len(s.OrderBy) > 0 {
		orderBy = s.OrderBy
	}

	if len(s.Sorter) > 0 {
		var sorterMapResult map[string]interface{}
		err := json.Unmarshal([]byte(s.Sorter), &sorterMapResult)
		if err == nil {
			for index, value := range sorterMapResult {
				orderBy = index
				switch value {
				case "ascend":
					sort = "asc"
					break
				case "descend":
					sort = "desc"
					break

				}

			}
		}

	}

	db = db.Order(fmt.Sprintf("%s %s", orderBy, sort))

	db.Scopes(FoundByWhere(s.Fields))

	return db
}

// FoundByWhere 查询条件
func FoundByWhere(fields []*Filed) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(fields) > 0 {
			for _, field := range fields {
				if field != nil {
					if field.Condition == "" {
						field.Condition = "="
					}
					if value, ok := field.Value.(int); ok {
						if value > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.(uint); ok {
						if value > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.(string); ok {
						if len(value) > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.([]int); ok {
						if len(value) > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.([]string); ok {
						if len(value) > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else {
						color.Red(fmt.Sprintf("未知数据类型：%+v", field.Value))
					}
				}
			}
		}
		return db
	}
}

// FoundOrder 查询条件
func FoundOrder(order string, sort string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tempSort := "asc"
		if len(order) > 0 {
			if len(sort) > 0 {
				tempSort = sort
			}
			db = db.Order(fmt.Sprintf("%s %s", order, tempSort))
		}
		return db
	}
}

// GetSearche 转换前端查询关系为 *Filed
func GetSearche(key, search string, condition string) *Filed {
	if len(search) > 0 {
		if len(condition) > 0 {
			switch condition {
			case "like":
				{
					value := fmt.Sprintf("%%%s%%", search)
					return &Filed{
						Key:       key,
						Condition: condition,
						Value:     value,
					}
					break
				}
			default:
				return &Filed{
					Key:       key,
					Condition: "=",

					Value: search,
				}
				break
			}
		} else {
			return &Filed{
				Condition: "=",
				Key:       key,
				Value:     search,
			}
		}
	}
	return nil
}

// Paginate 分页
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize < 0:
			pageSize = -1
		case pageSize == 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		if page < 0 {
			offset = -1
		}
		return db.Offset(offset).Limit(pageSize)
	}
}

func Update(v, d interface{}, id uint) error {
	if err := libs.Db.Model(v).Where("id = ?", id).Save(d).Error; err != nil {
		color.Red(fmt.Sprintf("Update %+v to %+v\n", v, d))
		return err
	}
	return nil
}

func GetRolesForUser(uid uint) []string {
	uids, err := libs.Enforcer.GetRolesForUser(strconv.FormatUint(uint64(uid), 10))
	if err != nil {
		color.Red(fmt.Sprintf("GetRolesForUser 错误: %v", err))
		return []string{}
	}

	return uids
}

func GetPermissionsForUser(uid uint) [][]string {
	return libs.Enforcer.GetPermissionsForUser(strconv.FormatUint(uint64(uid), 10))
}

func DropTables() {
	libs.Db.Migrator().DropTable(libs.Config.DB.Prefix+"users", libs.Config.DB.Prefix+"roles", libs.Config.DB.Prefix+"permissions", libs.Config.DB.Prefix+"oauth_tokens", "casbin_rule")
}
