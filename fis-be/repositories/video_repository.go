package repositories

import (
	"fis/be/database"
	"fis/be/models"
	"log"
	"net/url"
	"strings"
)

var base_url = []string{"https://www.youtube.com/watch?v=", "https://youtu.be/"}
var base = "https://www.youtube.com/embed/"

func GetVideos(key string, start int) ([]models.Video, int, error){
	db := database.Connect()
	defer db.Close()

	var videos []models.Video
	rows, err := db.Query("SELECT * FROM videos WHERE title LIKE ? LIMIT 12 OFFSET ?", key, start)
	if err != nil {
		log.Println(err)
		return videos, 0, err
	}

	for rows.Next(){
		var video models.Video
		err := rows.Scan(&video.ID, &video.Title, &video.Description, &video.VideoUrl)
		if err != nil {
			log.Println(err)
			return videos, 0, err
		}
		videos = append(videos, video)
	}
	return videos, getTotal(key) ,err
}

func getTotal(key string) int {
	db := database.Connect()
	defer db.Close()

	row, err := db.Query("SELECT COUNT(*) FROM videos WHERE title LIKE ?", key)
	if err != nil{
		return 0
	}
	var total int
	if row.Next(){
		row.Scan(&total)
	}
	return total
}

func GetVideoByName(name string) (models.Video, error){
	db := database.Connect()
	defer db.Close()
	decode, _ := url.QueryUnescape(name)

	rows, err := db.Query("SELECT * FROM videos WHERE title = ?", decode)
	if err != nil {
		log.Println(err)
	}
	var video models.Video
	if rows.Next(){
		err = rows.Scan(&video.ID, &video.Title, &video.Description, &video.VideoUrl)
	}
	return video, err
}

func CreateVideo(video models.Video) (models.Video, error){
	db := database.Connect()
	defer db.Close()

	video.VideoUrl = getIdFromLinkYoutube(video.VideoUrl)
	query, err := db.Exec("INSERT INTO videos (title, description, video_url) value (?, ?, ?)", video.Title, video.Description, video.VideoUrl)
	if err != nil {
		log.Println(err)
		return video, err
	}
	video.ID, err = query.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return video, err
}

func getIdFromLinkYoutube(link string) string {
	for _, base := range base_url{
		if strings.HasPrefix(link, base) {
			link = link[len(base):]
		}
	}
	return base +  link + "?enablejsapi=1"
}


func UpdateVideoByID(video models.Video) error {
	db := database.Connect()
	defer db.Close()
	_, err := db.Query("UPDATE videos "+
		"SET title = ?, description = ?, video_url = ? "+
		"WHERE id = ?", video.Title, video.Description, video.VideoUrl, video.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteVideo(id int) error {
	db := database.Connect()
	defer db.Close()
	_, err := db.Query("DELETE FROM videos WHERE id = ?", id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetNewestVideo()(models.Video, error){
	db := database.Connect()
	defer db.Close()
	var video models.Video
	row, err := db.Query("SELECT * FROM videos ORDER BY id desc limit 1")
	if err != nil{
		return video, err
	}
	if row.Next(){
		row.Scan(&video.ID, &video.Title, &video.Description, &video.VideoUrl)
	}
	return video, err
}