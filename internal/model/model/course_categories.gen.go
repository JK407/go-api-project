// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCourseCategory = "course_categories"

// CourseCategory 课程类别表
type CourseCategory struct {
	CategoryID   int64  `gorm:"column:category_id;type:int(11);primaryKey;autoIncrement:true;comment:课程类别ID" json:"category_id"` // 课程类别ID
	CategoryName string `gorm:"column:category_name;type:varchar(100);not null;comment:课程类别名称" json:"category_name"`             // 课程类别名称
	CreatedAt    int64  `gorm:"column:created_at;type:int(11);not null;comment:创建时间" json:"created_at"`                          // 创建时间
	UpdatedAt    int64  `gorm:"column:updated_at;type:int(11);not null;comment:更新时间" json:"updated_at"`                          // 更新时间
	Status       int64  `gorm:"column:status;type:int(11);not null;comment:状态:0正常、1禁用" json:"status"`                            // 状态:0正常、1禁用
}

// TableName CourseCategory's table name
func (*CourseCategory) TableName() string {
	return TableNameCourseCategory
}
