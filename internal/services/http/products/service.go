package products

type Handler struct {
	repository repository
}

func New(repo repository) Handler {
	return Handler{
		repository: repo,
	}
}
