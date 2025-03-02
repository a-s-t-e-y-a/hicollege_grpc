package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	Id              int       `bun:"id,pk,autoincrement"`
	Name            string    `bun:"name,notnull"`
	College         string    `bun:"college"`
	Course_Name     string    `bun:"course_name"`
	Phone_Number    string    `bun:"phone_number"`
	Graduation_Year string    `bun:"graduation_year"`
	Dob             string    `bun:"dob"`
	Location        string    `bun:"location"`
	CreatedAt       time.Time `bun:"created_at,nullzero,default:current_timestamp"`
	UpdatedAt       time.Time `bun:"updated_at,nullzero,default:current_timestamp"`
}
