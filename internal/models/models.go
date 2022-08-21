package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Url struct {
	Id        uint16
	Name      string
	CreatedAt string
}

func SaveUrl(db *sql.DB, url Url) bool {
	_, err := db.Exec("Insert into domain (name, createdAt) values (?, ?)", url.Name, time.Now())
	if err != nil {
		fmt.Println("error while inserting into database", err)

		return false
	}

	return true
}

func ListUrls(db *sql.DB) []Url {
	urls := []Url{}

	res, err := db.Query("SELECT * FROM domain")

	if err != nil {
		fmt.Println("cannot select from database", err)
		return []Url{}
	}

	for res.Next() {

		var url Url
		err := res.Scan(&url.Id, &url.Name, &url.CreatedAt)

		if err != nil {
			fmt.Println(err)
		}
		urls = append(urls, url)
	}

	return urls
}
