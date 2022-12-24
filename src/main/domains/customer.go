package domains

type Customer struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Active  bool   `json:"active,omitempty"`
}
