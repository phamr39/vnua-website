package utils

import (
	"encoding/csv"
	"fis/be/database"
	"fis/be/models"
	"log"
	"os"
	"strconv"
)

func CreateFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	db := database.Connect()
	defer db.Close()

	var customers []models.Customer
	rows, err := db.Query("SELECT * FROM customers")
	if err != nil {
		log.Println(err)
		return err
	}
	for rows.Next(){
		var customer models.Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.PhoneNumber, &customer.ProductInfo, &customer.Subject, &customer.Message)
		if err != nil {
			log.Println(err)
			return err
		}
		customers = append(customers, customer)
	}


	writer.Write([]string{
		"ID", "Name", "Email", "Phone Number", "Products", "Subject" , "Message",
	})

	for _, customer := range customers {
		data := []string{
			strconv.Itoa(int(customer.ID)),
			customer.Name,
			customer.Email,
			customer.PhoneNumber,
			customer.ProductInfo,
			customer.Subject,
			customer.Message,
		}

		if err := writer.Write(data); err != nil {
			return err
		}
	}

	return nil
}
