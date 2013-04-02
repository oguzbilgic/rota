package rota

type Rota struct {
	Pattern    string
	CaptureFns CaptureFns
}

func New(pattern string) *Rota {
	r := &Rota{
		Pattern:    pattern,
		CaptureFns: ParseRotaPattern(pattern),
	}
	return r
}

// Match checks if the given path matches with the rota.
func (r *Rota) Match(path string) bool {
	var rest string
	var match bool

	for _, cf := range r.CaptureFns {
		match, rest = cf(path)
		if !match {
			return false
		}
		path = rest
	}

	if rest == "" {
		return true
	}

	return false
}
