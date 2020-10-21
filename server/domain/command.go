package domain

type Command struct {
	Created  float32         `json:"created"`
	Modified float32         `json:"modified"`
	Id       string          `json:"id"`
	Name     string          `json:"name"`
	Get      CommandResponse `json:"get"`
}

type CommandResponse struct {
	Path      string     `json:"path"`
	Url       string     `json:"url"`
	Responses []Response `json:"responses"`
}

type Response struct {
	Code           string   `json:"code"`
	Description    string   `json:"description"`
	ExpectedValues []string `json:"expectedValues"`
}
