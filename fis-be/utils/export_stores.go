package utils

import (
	"encoding/csv"
	"fis/be/database"
	"fis/be/models"
	"log"
	"os"
	"strconv"
)

func CreateCSVStores(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	db := database.Connect()
	defer db.Close()

	var stores []models.Store
	rows, err := db.Query("SELECT * FROM stores")
	if err != nil {
		log.Println(err)
		return err
	}
	for rows.Next(){
		var store models.Store
		err := rows.Scan(&store.ID, &store.Name, &store.Email, &store.PhoneNumber, &store.Address, &store.BusinessModel, &store.Message)
		if err != nil {
			log.Println(err)
			return err
		}
		stores = append(stores, store)
	}


	writer.Write([]string{
		"ID", "Name", "Email", "Phone Number", "Address", "BusinessModel" , "Message",
	})

	for _, store := range stores {
		data := []string{
			strconv.Itoa(int(store.ID)),
			store.Name,
			store.Email,
			store.PhoneNumber,
			store.Address,
			store.BusinessModel,
			store.Message,
		}

		if err := writer.Write(data); err != nil {
			return err
		}
	}

	return nil
}
