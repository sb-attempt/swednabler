package simplex

// The dataset struct
var SimpleData SimpleDataStruct

type SimpleDataStruct struct {
	Glossary []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		FullForm    string `json:"full_form"`
		Description string `json:"description"`
	} `json:"glossary_simple"`
}

type SimpleTerms struct {
	Name        string `json:"name"`
	FullForm    string `json:"full_form"`
	Description string `json:"description"`
}
