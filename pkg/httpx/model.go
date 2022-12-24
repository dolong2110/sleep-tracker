package httpx

type ApiJson struct {
	Error   []error           `json:"error,omitempty"`
	Message string            `json:"message"`
	Data    []interface{}     `json:"data,omitempty"`
	Links   map[string]string `json:"links,omitempty"`
}
