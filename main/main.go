package main

import (
	"fmt"
	"log"
	"net/http"

	"modul3/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/rooms", controllers.GetAllRooms).Methods("GET")
	router.HandleFunc("/roomsDetail", controllers.GetDetailRooms).Methods("GET")
	router.HandleFunc("/rooms", controllers.JoinRoom).Methods("POST")

	// router.HandleFunc("/products", controllers.UpdateProduct).Methods("PUT")
	// router.HandleFunc("/products", controllers.DeleteSingleProduct).Methods("DELETE")

	// router.HandleFunc("/transactions", controllers.GetAllTransactions).Methods("GET")
	// router.HandleFunc("/transactionsDetail", controllers.GetDetailUserTransactions).Methods("GET")
	// router.HandleFunc("/transactionsDetailbyID", controllers.GetDetailUserTransactionsByID).Methods("GET")
	// router.HandleFunc("/transactions", controllers.InsertNewTransaction).Methods("POST")
	// router.HandleFunc("/transactions", controllers.UpdateTransaction).Methods("PUT")
	// router.HandleFunc("/products", controllers.DeleteTransaction).Methods("DELETE")

	// router.HandleFunc("/login", controllers.Login).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
