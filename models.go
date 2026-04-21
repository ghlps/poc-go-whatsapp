package main

type Menu struct {
	Restaurant *Restaurant       `json:"restaurant" `
	Date       string            `json:"date"      `
	ImgMenu    *string           `json:"imgMenu"    `
	Meals      map[string][]Meal `json:"meals"      `
}

type Meal struct {
	Name    string   `json:"name"  `
	Icons   []string `json:"icons" `
	Changed bool     `json:"changed"`
}

type Restaurant struct {
	Name string `json:"name" `
	Code string `json:"code" `
	Url  string `json:"url" `
}

type RequestPayload struct {
	RuCode         string `json:"ruCode"`
	WhatsAppLink   string `json:"whatsAppLink"`
	WhatsAppNumber string `json:"whatsAppNumber"`
}

type EventLambda struct {
	RequestPayload  RequestPayload `json:"requestPayload"`
	ResponsePayload Menu           `json:"responsePayload"`
}
