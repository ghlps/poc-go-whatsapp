package main

type Menu struct {
	Restaurant *Restaurant       `json:"restaurant" `
	Date       string            `json:"date"      `
	ImgMenu    *string           `json:"imgMenu"    `
	Served     []string          `json:"served"    `
	Meals      map[string][]Meal `json:"meals"      `
}

type Meal struct {
	Name  string   `json:"name"  `
	Icons []string `json:"icons" `
}

type Restaurant struct {
	Name string `json:"name" `
	Code string `json:"code" `
	Url  string `json:"url" `
}

type EventLambda struct {
	RuCode          string `json:"ruCode"`
	RunType         string `json:"runType"`
	ResponsePayload Menu   `json:"responsePayload"`
}
