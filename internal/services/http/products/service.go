package products

type Handler struct {
	repository
}

func New(repo repository) Handler {
	return Handler{
		repository: repo,
	}
}
