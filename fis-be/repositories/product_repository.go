package repositories

import (
	"database/sql"
	"fis/be/database"
	"fis/be/models"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"log"
	"net/url"
	"strconv"
	"strings"
	"unicode"
)

func GetAllProducts(category string, orders string, page int, name string) ([]models.Product, error) {
	db := database.Connect()
	defer db.Close()

	var products []models.Product

	query := "SELECT products.* FROM products "
	if category != ""{
		query += "JOIN product_category ON products.id = product_category.product_id " +
			" JOIN categories ON categories.id = product_category.category_id" +
			" WHERE (categories.id = " + category + " OR categories.root_id = " + category + ") AND products.name LIKE ?"
	}else{
		query += "WHERE products.name LIKE ?"
	}
	query += " ORDER BY products.name " + orders + " LIMIT 12 OFFSET " + strconv.Itoa(page)

	rows, err := db.Query(query, name)
	if err != nil {
		return products, err
	}
	for rows.Next() {
		var product models.Product
		rows.Scan(
			&product.ID,
			&product.Name,
			&product.Handle,
			&product.Description,
			&product.ProductInfo,
			&product.UserManual,
			&product.CreatedAt,
		)
		product.Images = getImages(product.ID)
		product.Categories = getCategories(product.ID)
		products = append(products, product)
	}
	return products, nil
}

func GetProductByID(id int64) models.Product {
	db := database.Connect()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM products WHERE id = ?", id)
	if err != nil {
		log.Println(err)
	}
	var product models.Product
	if rows.Next() {
		rows.Scan(
			&product.ID,
			&product.Name,
			&product.Handle,
			&product.Description,
			&product.ProductInfo,
			&product.UserManual,
			&product.CreatedAt,
		)
	}
	product.Images = getImages(product.ID)
	product.Categories = getCategories(product.ID)
	return product
}

func GetProductByName(name string) models.Product {
	db := database.Connect()
	defer db.Close()

	decode, _ := url.QueryUnescape(name)
	rows, err := db.Query("SELECT * FROM products WHERE name = ?", decode)
	if err != nil {
		log.Println(err)
	}
	var product models.Product
	if rows.Next() {
		rows.Scan(
			&product.ID,
			&product.Name,
			&product.Handle,
			&product.Description,
			&product.ProductInfo,
			&product.UserManual,
			&product.CreatedAt,
		)
	}
	product.Images = getImages(product.ID)
	product.Categories = getCategories(product.ID)
	return product
}

func GetProductByHandle(handle string) models.Product {
	db := database.Connect()
	defer db.Close()

	decode, _ := url.QueryUnescape(handle)
	rows, err := db.Query("SELECT * FROM products WHERE handle = ?", decode)
	if err != nil {
		log.Println(err)
	}
	var product models.Product
	if rows.Next() {
		rows.Scan(
			&product.ID,
			&product.Name,
			&product.Handle,
			&product.Description,
			&product.ProductInfo,
			&product.UserManual,
			&product.CreatedAt,
		)
	}
	product.Images = getImages(product.ID)
	product.Categories = getCategories(product.ID)
	return product
}

func GetProductsByCategory(category string) []models.Product {
	db := database.Connect()
	defer db.Close()

	decode, _ := url.QueryUnescape(category)
	rows, err := db.Query("select p.id, p.name, p.handle,p.description, p.product_info, p.user_manual, p.created_at " +
		"from products p " +
		"join product_category pc on pc.product_id = p.id " +
		"join categories c on c.id = pc.category_id " +
		"where c.name = ? limit 6", decode)
	if err != nil {
		log.Println(err)
	}

	var products []models.Product
	for rows.Next() {
		var product models.Product
		rows.Scan(
			&product.ID,
			&product.Name,
			&product.Handle,
			&product.Description,
			&product.ProductInfo,
			&product.UserManual,
			&product.CreatedAt,
		)
		product.Images = getImages(product.ID)
		product.Categories = getCategories(product.ID)
		products = append(products, product)
	}
	return products
}

func CreateProduct(product models.Product) models.Product {
	db := database.Connect()
	defer db.Close()

	product.Handle = generateHandle(product.Name)

	row, err := db.Exec("INSERT INTO products "+
		"(name, handle, description, product_info, user_manual, created_at) "+
		"VALUES (?, ?, ?, ?, ?, NOW())", product.Name, product.Handle, product.Description, product.ProductInfo, product.UserManual)

	if err != nil {
		log.Println(err)
	}

	id, err := row.LastInsertId()

	if err != nil {
		log.Println(err)
	}

	for _, image := range product.Images {
		_, err = db.Query("INSERT INTO images VALUES (?, ?)", id, image.ImageUrl)
	}
	for _, category := range product.Categories {
		_, err = db.Query("INSERT INTO product_category VALUES (?, ?)", id, category.ID)
	}
	if err != nil {
		log.Println(err)
	}
	return product
}

func UpdateProductByID(product models.Product) error {
	db := database.Connect()
	defer db.Close()
	_, err := db.Query("UPDATE products "+
		"SET name = ?, description = ?, product_info = ?, user_manual= ? "+
		"WHERE id = ?", product.Name, product.Description, product.ProductInfo, product.UserManual, product.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteProduct(id int64) error {
	db := database.Connect()
	defer db.Close()

	_, err := db.Query("DELETE FROM images WHERE product_id = ?", id)
	if err != nil {
		return err
	}

	_, err = db.Query("DELETE FROM product_category WHERE product_id = ?", id)
	if err != nil {
		return err
	}

	_, err = db.Query("DELETE FROM products where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func GetNewProducts() []models.Product {
	db := database.Connect()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM products ORDER BY created_at DESC LIMIT 10")
	var products []models.Product
	for rows.Next() {
		var product models.Product
		rows.Scan(
			&product.ID,
			&product.Name,
			&product.Handle,
			&product.Description,
			&product.ProductInfo,
			&product.UserManual,
			&product.CreatedAt,
		)
		product.Images = getImages(product.ID)
		product.Categories = getCategories(product.ID)
		products = append(products, product)
	}
	return products
}

func GetRandomProducts() []models.Product {
	db := database.Connect()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM products ORDER BY RAND() LIMIT 6")
	var products []models.Product
	for rows.Next() {
		var product models.Product
		rows.Scan(
			&product.ID,
			&product.Name,
			&product.Handle,
			&product.Description,
			&product.ProductInfo,
			&product.UserManual,
			&product.CreatedAt,
		)
		product.Images = getImages(product.ID)
		product.Categories = getCategories(product.ID)
		products = append(products, product)
	}
	return products
}

func GetTotalProductByRequest(category string, name string) (int, error) {
	db := database.Connect()
	defer db.Close()

	var rows *sql.Rows
	var err error

	if category == "" {
		rows, err = db.Query("SELECT COUNT(*) FROM products "+
			"WHERE name LIKE ? ", name)
	} else {
		rows, err = db.Query("SELECT COUNT(*) FROM products p "+
			"JOIN product_category c ON p.id = c.product_id "+
			"WHERE c.category_id = ? AND p.name LIKE ? ", category, name)
	}

	if err != nil {
		return 0, err
	}
	var total int
	if rows.Next() {
		rows.Scan(&total)
	}
	return total, err
}

func getImages(id int64) []models.Image {
	db := database.Connect()
	defer db.Close()
	rows, _ := db.Query("SELECT image_url FROM images "+
		"WHERE product_id = ?", id)
	var images []models.Image
	for rows.Next() {
		var image models.Image
		rows.Scan(&image.ImageUrl)
		images = append(images, image)
	}
	return images
}

func getCategories(productId int64) []models.Category {
	db := database.Connect()
	defer db.Close()
	rows, _ := db.Query("select c.id ,c.name from categories c " +
		"join product_category pc on pc.category_id = c.id " +
		"join products p on p.id = pc.product_id " +
		"where p.id = ?;", productId)
	var categories []models.Category
	for rows.Next() {
		var category models.Category
		rows.Scan(&category.ID ,&category.Name)
		categories = append(categories, category)
	}
	return categories
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

func generateHandle(productName string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, productName)
	result = strings.Trim(result, " ")
	result = strings.ReplaceAll(result, " ", "-")
	result = strings.ToLower(result)

	return result
}
