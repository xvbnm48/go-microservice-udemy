# go-microservice-udemy
go microservice chapter 15 10.29
link to course [udemy](https://www.udemy.com/course/building-modern-web-applications-with-go/learn/lecture/22875035?start=0#overview)

today is task for boilerplate!!!!

SERVER_ADDRESS=localhost SERVER_PORT=8080 DB_USER=root DB_PASSWD= DB_ADDR=localhost DB_PORT=3306 DB_NAME=banking go run main.go


// router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

    // func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, vars["customer_id"])
// }

// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "post request received")
// }