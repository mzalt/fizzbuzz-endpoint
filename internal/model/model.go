package model

// FizzbuzzResponse struct of response.
type FizzbuzzResponse struct {
	Output string `json:"output,omitempty"`
}

// FizzbuzzInpute struct for input.
type FizzbuzzInpute struct {
	Number1 int
	Number2 int
	Limit   int
	Str1    string
	Str2    string
}

// CheckInput checks if numbers are zero
func (r *FizzbuzzInpute) CheckInput() bool {
	if r.Number1 == 0 || r.Number2 == 0 {
		return false
	}

	return true
}
