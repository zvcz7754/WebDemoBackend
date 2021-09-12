package entity

type BookDO struct {
	ID string `gorm:"ID INT auto_increment NOT NULL primary_key:id123"`
	// Name      string `gorm:"type:varchar NOT NULL;"`
	// Price     string `gorm:"type:varchar;"`
	// Author    string `gorm:"type:varchar;"`
	// Publisher string `gorm:"type:varchar;"`
}

func (bookDO BookDO) TableName() string {
	return "book"
}

type PublisherDO struct {
	ID      string
	Name    string
	Phone   string
	Country string
}

type BookVO struct {
	ID        string `json:"id" form:"id"` //binding:"required"
	Name      string `json:"name" form:"name"`
	Price     string `json:"price" form:"name"`
	Author    string `json:"author" form:"name"`
	Publisher string `json:"publisher" form:"publisher"`
}
