package repositories

import (
	"database/sql"
	"fis/be/database"
	"fis/be/models"
	"log"
)

func GetAllPosts(topic string, orders string, start int, title string, postType string) ([]models.Post, error) {
	db := database.Connect()
	defer db.Close()

	var posts []models.Post

	query := "SELECT posts.* FROM posts "
	if topic != ""{
		query += "JOIN post_topic ON posts.id = post_topic.post_id " +
			"WHERE post_topic.topic_id = " + topic + " AND posts.title LIKE ? 	"
	} else {
		query += "WHERE posts.title LIKE ? "
	}
	query += "AND posts.post_type = ? ORDER BY posts.created_at " + orders + " LIMIT 4 OFFSET ?"

	rows, err := db.Query(query, title, postType ,start)

	if err != nil {
		return posts, err
	}
	for rows.Next() {
		var post models.Post
		rows.Scan(
			&post.ID,
			&post.Title,
			&post.Description,
			&post.Content,
			&post.Author,
			&post.CreatedAt,
			&post.PostType,
			&post.Thumbnail,
		)
		post.Topics = getTopics(int(post.ID))
		posts = append(posts, post)
	}
	return posts, nil
}

func DeletePostById(id int64) error {
	db := database.Connect()
	defer db.Close()

	_, err := db.Query("DELETE FROM post_topic WHERE post_id = ?", id)
	if err != nil {
		return err
	}

	_, err = db.Query("DELETE FROM posts where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func GetPostById(id int64) models.Post {
	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM posts WHERE id = ?", id)
	if err != nil {
		log.Println(err)
	}

	var post models.Post
	if rows.Next() {
		rows.Scan(
			&post.ID,
			&post.Title,
			&post.Description,
			&post.Content,
			&post.Author,
			&post.CreatedAt,
			&post.PostType,
			&post.Thumbnail,
		)
	}
	post.Topics = getTopics(int(post.ID))
	return post
}

func GetNewestPosts(postType string) []models.Post {
	db := database.Connect()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM posts WHERE post_type = ? ORDER BY created_at DESC LIMIT 6", postType)
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		rows.Scan(
			&post.ID,
			&post.Title,
			&post.Description,
			&post.Content,
			&post.Author,
			&post.CreatedAt,
			&post.PostType,
			&post.Thumbnail,
		)
		topics := getTopics(int(post.ID))
		if len(topics) > 0{
			post.Topics = topics
		}
		post.Topics = getTopics(int(post.ID))
		posts = append(posts, post)
	}
	return posts
}

func GetRandomPosts(postType string) []models.Post {
	db := database.Connect()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM posts WHERE post_type = ? ORDER BY RAND() LIMIT 4", postType)
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		rows.Scan(
			&post.ID,
			&post.Title,
			&post.Description,
			&post.Content,
			&post.Author,
			&post.CreatedAt,
			&post.PostType,
			&post.Thumbnail,
		)
		topics := getTopics(int(post.ID))
		if len(topics) > 0{
			post.Topics = topics
		}
		post.Topics = getTopics(int(post.ID))
		posts = append(posts, post)
	}
	return posts
}

func CreatePost(post models.Post) (models.Post, error) {
	db := database.Connect()
	defer db.Close()

	query, err := db.Exec("insert into posts (title, thumbnail, description, content, post_type, author, created_at) values (?, ?, ?, ?, ?, ?, now())",
		post.Title,
		post.Thumbnail,
		post.Description,
		post.Content,
		post.PostType,
		post.Author,
	)

	post.ID, err = query.LastInsertId()

	if err != nil{
		return post, err
	}

	for _, topic := range post.Topics {
		_, err = db.Query("INSERT INTO post_topic VALUES (?, ?)", topic.ID, post.ID)
	}
	if err != nil {
		log.Println(err)
	}
	return post, err
}

func GetTotalPostRequest(topic string, title string, postType string) (int, error) {
	db := database.Connect()
	defer db.Close()

	var rows *sql.Rows
	var err error

	if topic == "" {
		rows, err = db.Query("SELECT COUNT(*) FROM posts "+
			"WHERE title LIKE ? AND post_type =  ?", title, postType)
	} else {
		rows, err = db.Query("SELECT COUNT(*) FROM posts p "+
			"JOIN post_topic c ON p.id = c.post_id "+
			"WHERE c.topic_id = ? AND p.title LIKE ? AND p.post_type = ?", topic, title, postType)
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

func GetTopics()[]models.Topic{
	db := database.Connect()
	defer db.Close()
	var topics []models.Topic
	rows, _ := db.Query("SELECT * FROM topics")
	for rows.Next(){
		var topic models.Topic
		rows.Scan(&topic.ID, &topic.Name)
		topics = append(topics, topic)
	}
	return topics
}

func getTopics(id int)[]models.Topic{
	db := database.Connect()
	defer db.Close()

	var topics []models.Topic
	rows, err := db.Query("SELECT topics.* FROM topics JOIN post_topic ON topics.id = post_topic.topic_id" +
		" WHERE post_topic.post_id = ?", id)
	if err!= nil{
		log.Println(err)
		return topics
	}
	for rows.Next(){
		var topic models.Topic
		rows.Scan(&topic.ID, &topic.Name)
		topics = append(topics, topic)
	}
	return topics
}