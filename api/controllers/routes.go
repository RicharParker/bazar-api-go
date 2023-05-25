package controllers

import "github.com/rchavez-dev/fullstack/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareAuthentication(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

	//Products routes
	s.Router.HandleFunc("/products", middlewares.SetMiddlewareJSON(s.CreateProduct)).Methods("POST")
	s.Router.HandleFunc("/products", middlewares.SetMiddlewareJSON(s.GetProducts)).Methods("GET")
	s.Router.HandleFunc("/products/{id}", middlewares.SetMiddlewareJSON(s.GetProduct)).Methods("GET")
	s.Router.HandleFunc("/products/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateProduct))).Methods("PUT")
	s.Router.HandleFunc("/products/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteProduct)).Methods("DELETE")

	//Orders routes
	s.Router.HandleFunc("/orders", middlewares.SetMiddlewareJSON(s.CreateOrder)).Methods("POST")
	s.Router.HandleFunc("/orders", middlewares.SetMiddlewareJSON(s.GetOrders)).Methods("GET")
	s.Router.HandleFunc("/orders/{id}", middlewares.SetMiddlewareJSON(s.GetOrder)).Methods("GET")
	s.Router.HandleFunc("/orders/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateOrder))).Methods("PUT")
	s.Router.HandleFunc("/orders/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteOrder)).Methods("DELETE")

	//Cart routes
	s.Router.HandleFunc("/cart", middlewares.SetMiddlewareJSON(s.CreateCart)).Methods("POST")
	s.Router.HandleFunc("/cart", middlewares.SetMiddlewareJSON(s.GetCarts)).Methods("GET")
	s.Router.HandleFunc("/cart/{id}", middlewares.SetMiddlewareJSON(s.GetCart)).Methods("GET")
	s.Router.HandleFunc("/cart/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateCart))).Methods("PUT")
	s.Router.HandleFunc("/cart/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteCart)).Methods("DELETE")

}
