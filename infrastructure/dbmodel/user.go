package dbmodel

import "time"

type (
	User struct {
		Id        string    `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
		Name      string    `gorm:"column:name;type:varchar(255);not null"`
		Email     string    `gorm:"column:email;type:varchar(255);not null;unique"`
		Birthday  time.Time `gorm:"column:birthday;type:date;not null"`
		CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:now()"`
		UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:now()"`
	}
)
