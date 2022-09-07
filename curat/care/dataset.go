package curat

// The dataset structs
var Data DataStruct

type DataStruct struct {
	Glossary []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"glossary"`
}

type Terms []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SimpleTerms struct {
	Name        string `json:"name"`
	FullForm    string `json:"full_form"`
	Description string `json:"description"`
}
