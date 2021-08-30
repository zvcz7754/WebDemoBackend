package entity

type BookDO struct {
	ID        string //`json:"id"`
	Name      string
	Price     string
	Author    string
	Publisher string
}

type PublisherDO struct {
	ID      string
	Name    string
	Phone   string
	Country string
}

type BookVO struct {
	ID        string
	Name      string
	Price     string
	Author    string
	Publisher string
}
