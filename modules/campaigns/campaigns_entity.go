package campaigns

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Campaing struct {
	gorm.Model
	Name string
	Rows pq.StringArray `gorm:"type:text[]"`
}
