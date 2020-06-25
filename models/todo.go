package models

type Todo struct {
	ID    int    `json:"id" gorm:"AUTO_INCREMENT"`
	Title string `json:"title" gorm:"type:VARCHAR(225)"`
	Done  bool   `json:"done" gorm:"boolean"`
}
