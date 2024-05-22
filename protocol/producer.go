package protocol

type Producer struct {
	ID      string       `json:"id"`
	Name    string       `json:"name"`
	Group   string       `json:"group"`
	Modules []ModuleName `json:"modules"`
}
