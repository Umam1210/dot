package artikelsdto

import (
	"journey/models"
	"time"
)

type ArtikelResponse struct {
	ID     int                 `json:"id"`
	Title  string              `json:"title" form:"title" validate:"required"`
	Image  string              `json:"image" form:"image" validate:"required"`
	Desc   string              `json:"desc" form:"desc" validate:"required"`
	UserID int                 `json:"user_id"`
	User   models.UserResponse `json:"user"`
	Date   time.Time           `json:"date"`
}
