package repositories

import (
	"fis/be/database"
	"fis/be/models"
)

func GetStores(start int, name string, order string) ([]models.Store, int, error){
	db := database.Connect()
	defer db.Close()
	 var stores []models.Store
	query := "SELECT * FROM stores WHERE name LIKE ?" + " ORDER BY id " + order + " LIMIT 10 OFFSET ?"
	rows , err := db.Query(query, name, start)
	if err != nil{
		return stores, 0, err
	}
	for rows.Next() {
		var store models.Store
		err = rows.Scan(
			&store.ID,
			&store.Name,
			&store.Email,
			&store.PhoneNumber,
			&store.Address,
			&store.BusinessModel,
			&store.Message,
		)
		stores = append(stores, store)
	}
	row, _ := db.Query("SELECT COUNT(*) FROM stores WHERE name LIKE ?", name)
	var total int
	if row.Next(){
		row.Scan(&total)
	}
	return stores, total, err
}

func GetStoreByID(id int) (models.Store, error){
	db := database.Connect()
	defer db.Close()

	var store models.Store
	row, err := db.Query("SELECT * FROM stores WHERE id = ?", id)

	if err != nil{
		return store, err
	}
	if row.Next(){
		err = row.Scan(
			&store.ID,
			&store.Name,
			&store.Email,
			&store.PhoneNumber,
			&store.Address,
			&store.BusinessModel,
			&store.Message,
		)
	}
	return store, err
}

func CreateStore(store models.Store) (models.Store, error){
	db := database.Connect()
	defer db.Close()

	query, err := db.Exec("INSERT INTO stores (name, email, phone_number, address, business_model, message) " +
		"value (?, ?, ?, ?, ?, ?)", store.Name, store.Email, store.PhoneNumber, store.Address, store.BusinessModel, store.Message)
	if err != nil{
		return store, err
	}
	store.ID, err = query.LastInsertId()
	return store, err
}
func UpdateStore(store models.Store) (models.Store, error){
	db := database.Connect()
	defer db.Close()

	_, err := db.Query("UPDATE stores SET " +
		"name = ?, email = ?, phone_number = ?, address = ?, business_model = ?, message = ?" +
		" WHERE id = ?", store.Name, store.Email, store.PhoneNumber, store.Address, store.BusinessModel, store.Message, store.ID)
	return store, err
}

func DeleteStore(id int) error{
	db := database.Connect()
	defer db.Close()
	_, err := db.Query("DELETE FROM stores WHERE id = ?", id)
	return err
}