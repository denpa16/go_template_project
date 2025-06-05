package products

type RepoQueries struct {
	db DBTX
}

type Repository struct {
	queries RepoQueries
}

func NewProductsRepository(db DBTX) *Repository {
	return &Repository{
		queries: *New(db),
	}
}

func New(db DBTX) *RepoQueries {
	return &RepoQueries{db: db}
}
