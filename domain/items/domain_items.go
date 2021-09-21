package items

type Item struct {
	Id          string    `json:"id"`
	Seller      int64     `json:"seller"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Pictures    []Picture `json:"pictures"`
	Video       string    `json:"video"`
	Price       float32   `json:"price"`
	Available   int       `json:"available"`
	Sold        int       `json:"sold"`
	Status      string    `json:"status"`
}

type Description struct {
	PlainText string `json:"plain_text"`
	Html      string `json:"html"`
}

type Picture struct {
	Id  int64  `json:"id"`
	Url string `json:"utl"`
}
