package utils

import (
	"encoding/json"
	"io"

	"gorm.io/gorm"
)

func MarshalAndInsert(data any, box any) {
	marshaledData, _ := json.Marshal(data)
	json.Unmarshal(marshaledData, box)
}

func UnmarshalFromReader(body io.Reader, box any) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(box)
	if err != nil {
		return err
	}
	return nil
}

func Paginator(db *gorm.DB, page, perPage int32) *gorm.DB {
	if page <= 0 {
		page = 1
	}
	switch {
	case perPage > 100:
		perPage = 100
	case perPage <= 0:
		perPage = 10
	}

	return db.Offset(int((page - 1) * perPage)).Limit(int(perPage))
}
