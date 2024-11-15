package resources

import (
	_ "embed"
)

//go:embed home.html
var HomePage string

//go:embed new-joke.html
var NewJoke string
