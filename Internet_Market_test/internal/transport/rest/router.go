package rest

import (
	handler2 "github.com/Dimoonevs/Online_store/internal/transport/rest/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"net/http"
)

// Router: struct defines a router that handles HTTP requests.
type Router struct {
	handler *handler2.Handler
}

// NewRouter: creates a new Router instance with the provided handler.
func NewRouter(handler *handler2.Handler) *Router {
	return &Router{
		handler: handler,
	}
}

// Router: returns an HTTP handler that routes requests to the appropriate handler functions.
func (r *Router) Router() http.Handler {
	mux := chi.NewRouter()

	// Set up CORS middleware.
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Define routes.
	mux.Post("/auth", r.handler.Authentication)
	mux.Group(func(mux chi.Router) {
		mux.Use(r.handler.ValidateMiddleware)
		mux.Get("/", r.handler.Test)
		mux.Get("/seller/all", r.handler.GetSellers)
		mux.Get("/seller", r.handler.GetSellerByID)
		mux.Post("/seller", r.handler.CreateSeller)
		mux.Put("/seller", r.handler.UpdateSeller)
		mux.Delete("/seller", r.handler.DeleteSeller)
		mux.Get("/product/seller", r.handler.GetProductsSeller)
		mux.Get("/product", r.handler.GetProductByID)
		mux.Post("/product", r.handler.CreateProduct)
		mux.Put("/product", r.handler.UpdateProduct)
		mux.Delete("/product", r.handler.DeleteProduct)
		mux.Get("/product/all", r.handler.GetProducts)
	})
	return mux
}
