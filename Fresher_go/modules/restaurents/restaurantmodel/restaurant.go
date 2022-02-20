package restaurantmodel

import (
	"Fresher_go/common"
	"strings"
)

const EntityName = "Restaurent"

// create restaurants
type Restaurants struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name"`
	Addr            string         `json:"addr" gorm:"column:addr"`
	CityId          int            `json:"city_id" gorm:"column:city_id"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
	LikedCount      int            `json:"liked_count" gorm:"-"`
}

func (Restaurants) TableName() string {
	return "restaurants"
}

// check input
func (res *Restaurants) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	if len(res.Name) == 0 {
		return ErrNameCannotBeEmpty
	}
	return nil
}

//func (res *Restaurants) Mask(isAdminOwner bool) {
//	res.GenUID(common.DbTypeRetaurent)
//}

var (
	ErrNameCannotBeEmpty = common.NewCustomError(nil, "restuarant name can't be black", "ErrNameCannotBeEmpty")
)

func (data *Restaurants) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRetaurent)
}
