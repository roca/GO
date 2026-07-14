package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"shopping-app/pkg/handlers"
	"shopping-app/pkg/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var db *sql.DB
var tmpl *template.Template
var Store = sessions.NewCookieStore([]byte("shoppingsecret"))

func init() {
	var err error

	db, err = initDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v\n", err)
	}
	log.Println("Connected to the database")

	err = initSchema(db)
	if err != nil {
		log.Fatalf("Failed to initialize schema: %v\n", err)
	}
	log.Println("Initialized the schema")

	// tmpl, err = template.ParseGlob("./templates/*.html")
	// if err != nil {
	// 	log.Fatalf("Failed to parse templates: %v\n", err)
	// }
	// log.Println("Parsed HTML templates")

	// // Set up Sessions
	// Store.Options = &sessions.Options{
	// 	Path:     "/",
	// 	MaxAge:   3600 * 3,
	// 	HttpOnly: true,
	// }
}

func initDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1)/shopping?parseTime=true")
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %v\n", err)
	}

	// Check the connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Database not responding to Pings: %v\n", err)
	}
	return db, nil
}

func initSchema(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS products(
		product_id 	VARCHAR(50) 	PRIMARY KEY NOT NULL, 
		product_name 	VARCHAR(100),
		price 		FLOAT,
		description 	MEDIUMTEXT,
		product_image 	VARCHAR(50),
		date_created 	DATE,
		date_modified 	DATE
	);
	`)
	if err != nil {
		return fmt.Errorf("Failed to create table products: %v", err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS orders(
		order_id 	VARCHAR(50) 	 PRIMARY KEY NOT NULL,
		user_id 	VARCHAR(50),
		order_status 	VARCHAR(15),
		order_date 	DATE
	)
	`)
	if err != nil {
		return fmt.Errorf("Failed to create table orders: %v", err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS order_items(
		order_id 	VARCHAR(50) 	NOT NULL,
		product_id 	VARCHAR(50),
		quantity 	INT DEFAULT 1,
		cost 		FLOAT
	);
	`)
	if err != nil {
		return fmt.Errorf("Failed to create table order_items: %v", err)
	}

	return nil
}

func main() {
	defer db.Close()

	gRouter := mux.NewRouter()

	// Setup Static file handling for images

	fs := http.FileServer(http.Dir("./static"))
	gRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))


	repo := repository.NewRepository(db)
	handler := handlers.NewHandler(repo)

	gRouter.HandleFunc("/seed-products", handler.SeedProducts).Methods("POST")


	//All dynamic routes

	// User routes

	gRouter.HandleFunc("/", handler.ShoppingHomepage).Methods("GET")
	gRouter.HandleFunc("/shoppingitems", handler.ShoppingItemsView).Methods("GET")
	gRouter.HandleFunc("/cartitems", handler.CartView).Methods("GET")
	gRouter.HandleFunc("/addtocart/{product_id}", handler.AddToCart).Methods("POST")
	gRouter.HandleFunc("/gotocart", handler.ShoppingCartView).Methods("GET")
	gRouter.HandleFunc("/updateorderitem", handler.UpdateOrderItemQuantity).Methods("PUT")
	gRouter.HandleFunc("/ordercomplete", handler.PlaceOrder).Methods("GET")

	// Admin routes

	gRouter.HandleFunc("/manageproducts", handler.ProductsPage).Methods("GET")
	gRouter.HandleFunc("/allproducts", handler.AllProductsView).Methods("GET")
	gRouter.HandleFunc("/products",handler.ListProducts).Methods("GET")
	gRouter.HandleFunc("/products/{id}",handler.GetProduct).Methods("GET")
	gRouter.HandleFunc("/createproduct",handler.CreateProductView).Methods("GET")
	gRouter.HandleFunc("/products",handler.CreateProduct).Methods("POST")
	gRouter.HandleFunc("/editproduct/{id}",handler.EditProductView).Methods("GET")
	gRouter.HandleFunc("/products/{id}",handler.UpdateProduct).Methods("PUT")
	gRouter.HandleFunc("/products/{id}",handler.DeleteProduct).Methods("DELETE")

	gRouter.HandleFunc("/manageorders", handler.OrdersPage).Methods("GET")
	gRouter.HandleFunc("/allorders", handler.AllOrdersView).Methods("GET")
	gRouter.HandleFunc("/orders",handler.ListOrders).Methods("GET")
	gRouter.HandleFunc("/orders/{id}",handler.GetOrder).Methods("GET")
  

	log.Println("Server started on http://localhost:5001")
	err := http.ListenAndServe(":5001", gRouter)
	if err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
