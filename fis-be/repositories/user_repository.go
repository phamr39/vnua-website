package repositories

import (
	"fis/be/database"
	"fis/be/models"
	"log"
)

func CreateNewUser(user models.User) error {
	db := database.Connect()
	defer db.Close()
	_, err := db.Query("insert into users (name, phone_number, email, password, role_id) "+
		"values ( ?, ?, ?, ?, ?)", user.Name, user.PhoneNumber, user.Email, user.Password, 2)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func GetAllUsers(start int) []models.User {
	db := database.Connect()
	defer db.Close()
	rows, _ := db.Query("select u.id, u.name, u.phone_number, u.email, r.name " +
		"from users u join roles r on r.id = u.role_id limit 20 offset ?", start)
	var users []models.User
	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Email, &user.Role)
		users = append(users, user)
	}
	return users
}

func GetUserById(id int64) models.User {
	db := database.Connect()
	defer db.Close()
	var user models.User
	row, _ := db.Query("select u.id, u.name, u.phone_number, u.email, r.name "+
		"from users u join roles r on r.id = u.role_id where u.id = ? ", id)
	if row.Next() {
		row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Email, &user.Role)
	}
	return user
}

func GetUserByEmail(email string) (models.User, error) {
	db := database.Connect()
	defer db.Close()
	var user models.User
	row, err := db.Query("select u.id, u.name, u.phone_number, u.email, r.name "+
		"from users u join roles r on r.id = u.role_id where u.email = ? ", email)
	if err != nil {
		log.Println(err)
		return user, err
	}
	if row.Next() {
		row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Email, &user.Role)
	}
	return user, nil
}

func UpdateUserById(id int64, user models.User) {
	db := database.Connect()
	defer db.Close()
	var roleId int64
	if user.Role == "ADMIN" {
		roleId = 1
	} else {
		roleId = 2
	}
	_, err := db.Query("update users set name = ?, phone_number = ?, email = ?, role_id = ? "+
		"where id = ?", user.Name, user.PhoneNumber, user.Email, roleId, id)
	if err != nil {
		log.Println(err)
		return
	}
}

func DeleteUserById(id int64) {
	db := database.Connect()
	defer db.Close()
	_, err := db.Query("delete from users where id = ?", id)
	if err != nil {
		return
	}
}

func GetTotalUsers() (int, error) {
	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT COUNT(*) FROM users")
	if err != nil {
		return 0, err
	}
	var total int
	if rows.Next() {
		rows.Scan(&total)
	}
	return total, err
}
