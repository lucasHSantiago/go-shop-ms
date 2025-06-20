package main

type storer interface {
}

type handler struct {
	storer storer
}

func NewHandler(storer storer) *handler {
	return &handler{
		storer: storer,
	}
}
