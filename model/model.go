package model

type HistoryData struct {
	Date string `json:"date"`
	URL  string `json:"url"`
	Data Data   `json:"data"`
}

type Data struct {
	Events []Event `json:"Events"`
	Births []Birth `json:"Births"`
	Deaths []Death `json:"Deaths"`
}

type Event struct {
	Year  string `json:"year"`
	HTML  string `json:"html" gorm:"PRIMARY_KEY"`
	Text  string `json:"text"`
	Links []Link `json:"links" gorm:"polymorphic:Link"`
}

type Death struct {
	Year  string `json:"year"`
	Text  string `json:"text" gorm:"PRIMARY_KEY"`
	HTML  string `json:"html"`
	Links []Link `json:"links" gorm:"polymorphic:Link"`
}

type Birth struct {
	Year  string `json:"year"`
	HTML  string `json:"html"`
	Text  string `json:"text" gorm:"PRIMARY_KEY"`
	Links []Link `json:"links" gorm:"polymorphic:Link"`
}

type Link struct {
	Title string `json:"title"`
	Link  string `json:"link" gorm:"PRIMARY_KEY"`
}
