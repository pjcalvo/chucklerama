package resources

import (
	_ "embed"
)

//go:embed home.html
var HomePage string

//go:embed new-joke.html
var NewJokePage string

//go:embed new-joke.html
var TermsPage string
