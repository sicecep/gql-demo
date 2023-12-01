package models

type Book struct {
	ID        int      `gorm:"primary key;autoIncrement" json:"id"`
	Title     string   `json:"title"`
	Author    string   `json:"author"`
	Publisher string   `json:"publisher"`
	CountryID int      `json:"country_id"`
	Country   *Country `json:"country"`
}

type Country struct {
	ID   int    `gorm:"primary key;autoIncrement" json:"id"`
	Name string `json:"name"`
}
