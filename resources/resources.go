package resources

import (
	_ "embed"
)

//go:embed home.html
var HomePage string

//go:embed new-joke.html
var NewJokePage string

//go:embed terms.html
var TermsPage string

//go:embed thanks.html
var ThanksPage string
