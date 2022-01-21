package repositories

import (
	"fis/be/database"
	"fis/be/models"
	"log"
)

func GetAllCustomers(start int) ([]models.Customer, error){
	db := database.Connect()
	defer db.Close()

	var customers []models.Customer
	rows, err := db.Query("SELECT * FROM customers LIMIT 20 OFFSET ?", start)
	if err != nil {
		log.Println(err)
		return customers, err
	}

	for rows.Next(){
		var customer models.Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.PhoneNumber, &customer.ProductInfo, &customer.Subject, &customer.Message)
		if err != nil {
			log.Println(err)
			return customers, err
		}
		customers = append(customers, customer)
	}
	return customers, err
}

func CreateCustomerInfo(customer models.Customer) (models.Customer, error) {
	db := database.Connect()
	defer db.Close()

	query, err := db.Exec("INSERT INTO customers (name, email, phone_number, product_info, subject, message) value (?, ?, ?, ?, ?, ?)",
		customer.Name, customer.Email, customer.PhoneNumber, customer.ProductInfo, customer.Subject, customer.Message)
	if err != nil {
		log.Println(err)
		return customer, err
	}
	customer.ID, err = query.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return customer, err
}

func GetTotalCustomers() (int, error) {
	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT COUNT(*) FROM customers")
	if err != nil {
		return 0, err
	}
	var total int
	if rows.Next() {
		rows.Scan(&total)
	}
	return total, err
}
