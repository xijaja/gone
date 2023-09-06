package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type SomeTimesAt struct {
	CreatedAt *LocalTime `gorm:"comment:创建时间; type:timestamp(0);autoCreateTime;" json:"created_at"`
	UpdatedAt *LocalTime `gorm:"comment:更新时间; type:timestamp(0);autoUpdateTime;" json:"updated_at"`
	DeletedAt *LocalTime `gorm:"comment:删除时间; type:timestamp(0);" json:"deleted_at,omitempty"`
}

// LocalTime 自定义通用时间类型
type LocalTime time.Time

// MarshalJSON 在从数据库读取时调用
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// Value 在写入数据库时调用
func (t *LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(*t)
	// 判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

// Scan 在从数据库扫描时调用
func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// ToLocalTime 将 time.Time 转为 *LocalTime
func ToLocalTime(time time.Time) *LocalTime {
	return (*LocalTime)(&time)
}
