package api

type Handler struct {
	name string
}

func NewHandler(name string) *Handler {
	return &Handler{
		name: name,
	}
}