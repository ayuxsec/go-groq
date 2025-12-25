package groq

import "errors"

var ErrEmptyApiKey = errors.New("empty API key")
var ErrNilHttpClient = errors.New("http.Client pointer is nil")
