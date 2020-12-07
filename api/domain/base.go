package domain

import (
	"net/http"
	"reflect"
	"time"

	"github.com/midnight-trigger/todo/api/error_handling"
)

type Base struct {
}

type Result struct {
	// field
	Data       interface{}
	Pagination *Pagination

	// error
	error_handling.ErrorHandling
}

type Pagination struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func (r *Result) New() {
	r.Code = http.StatusOK
}

func (b *Base) SetStructOnSameField(fromStruct interface{}, toStruct interface{}) {
	from := reflect.Indirect(reflect.ValueOf(fromStruct))
	to := reflect.Indirect(reflect.ValueOf(toStruct))

	for i := 0; i < from.NumField(); i++ {
		fromFieldValue := from.Field(i).Interface()
		if t, ok := to.Type().FieldByName(from.Type().Field(i).Name); ok {
			switch v := fromFieldValue.(type) {
			case int64:
				to.Field(t.Index[0]).SetInt(v)
			case int:
				to.Field(t.Index[0]).SetInt(int64(v))
			case string:
				to.Field(t.Index[0]).SetString(v)
			case time.Time:
				to.Field(t.Index[0]).SetString(v.Format("2006-01-02 15:04:05"))
			case *time.Time:
				if v != nil {
					to.Field(t.Index[0]).SetString(v.Format("2006-01-02 15:04:05"))
				}
			}
		}
	}
	return
}

func (b *Base) SliceFind(slice interface{}, val interface{}) bool {
	switch val.(type) {
	case string:
		for _, item := range slice.([]string) {
			if item == val {
				return true
			}
		}
	case int64:
		for _, item := range slice.([]int64) {
			if item == val {
				return true
			}
		}
	}
	return false
}
