package probe

// HTTPProbeCmd provides the HTTP probe parameters
type HTTPProbeCmd struct {
	Method   string   `json:"method"`
	Resource string   `json:"resource"`
	Port     int      `json:"port"`
	Protocol string   `json:"protocol"`
	Headers  []string `json:"headers"`
	Body     string   `json:"body"`
	BodyFile string   `json:"body_file"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Crawl    bool     `json:"crawl"`
}
