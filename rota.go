package rota

type Rota struct {
	Pattern    string
	CaptureFns CaptureFns
}

func New(pattern string) *Rota {
	return &Rota{pattern, ParseRotaPattern(pattern)}
}

// Match checks if the given path matches with the rota.
func (r *Rota) Match(path string) bool {
	for _, cf := range r.CaptureFns {
		match, rest := cf(path)
		if !match {
			return false
		}
		path = rest
	}

	if path == "" {
		return true
	}

	return false
}
