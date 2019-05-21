package spaserver

// IndexContent is exported for use with the templating engine
type IndexContent struct {
	IndexTitle string
}

// ExampleContent is exported
type ExampleContent struct {
	ExampleTitle       string `json:"example_title"`
	ExampleDescription string `json:"example_description"`
}
