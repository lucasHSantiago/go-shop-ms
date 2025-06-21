module github.com/lucasHSantiago/go-shop-ms/product

go 1.24.4

require (
	github.com/ardanlabs/conf/v3 v3.8.0
	github.com/go-chi/chi/v5 v5.2.2
	github.com/lucasHSantiago/go-shop-ms/logger v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.34.0
	golang.org/x/sync v0.15.0
)

replace github.com/lucasHSantiago/go-shop-ms/logger => ../logger

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	golang.org/x/sys v0.12.0 // indirect
)
