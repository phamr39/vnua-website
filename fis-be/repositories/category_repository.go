package repositories

import (
	"fis/be/database"
	"fis/be/models"
	"log"
)

func GetAllCategories(root string) ([]models.Category, error) {
	db := database.Connect()
	defer db.Close()

	query := "SELECT id, name FROM categories "
	if root != "" {
		query += "WHERE root_id = " + root
	} else {
		query += "WHERE root_id IS NULL"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var categories []models.Category
	for rows.Next() {
		var cate models.Category
		err = rows.Scan(&cate.ID, &cate.Name)
		categories = append(categories, cate)
		if err != nil {
			log.Println(err)
		}
	}
	return categories, nil
}
