package shortener

type Redirect struct {
	Original      string `json:"original"`
	Alias         string `json:"short"`
	RedirectCount int    `json:"redirect_count"`
}
