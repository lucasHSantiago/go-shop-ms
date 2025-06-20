package main

type handler struct {
	repo *repo
}

func NewHandler(repo *repo) *handler {
	return &handler{
		repo: repo,
	}
}
