package bot

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"
)

const (
	driver = "sqlite3"
	dataSource = "vmd.db"
	dateFormat = "2006-01-02T15:04:05-0700"
)

type Statistic struct {
	Username string
	DeletedCount int
}

func getQuery(name string) string {
	path := filepath.Join("queries", name + ".sql")
	c, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(c)
}

func getDB() *sql.DB {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		panic(err)
	}
	return db
}

func getDate() string {
	return time.Now().Format(dateFormat)
}

func InitDB() {
	db := getDB()
	defer db.Close()
	_, err := db.Exec(getQuery("init"))
	if err != nil {
		panic(err)
	}
}


func getLastSeen(chatID int64, userID int) {
	db := getDB()
	defer db.Close()
	row := db.QueryRow(getQuery("getLastSeen"), chatID, userID)
	log.Println(row.Scan())
}

func insertUser(chatID int64, userID int) {
	db := getDB()
	defer db.Close()
	log.Println(
		db.Exec(
			getQuery("insertUser"),
			chatID,
			userID,
			getDate()))
}

func updateUser(chatID int64, userID int) {
	db := getDB()
	defer db.Close()
	_, err := db.Exec(
		getQuery("updateUser"),
		getDate(),
		chatID,
		userID)

	if err != nil {
		panic(err)
	}
}

func createStatistic(chatID int64, userID int, username string) error {
	db := getDB()
	defer db.Close()
	_, err := db.Exec(
		getQuery("createStatistic"),
		chatID,
		userID,
		username,
		getDate())

	return err
}

func updateStatistic(chatID int64, userID int) {
	db := getDB()
	defer db.Close()
	_, err := db.Exec(
		getQuery("updateStatistic"),
		getDate(),
		chatID,
		userID)

	if err != nil {
		panic(err)
	}
}

func getStatistic(chatID int64) []Statistic {
	db := getDB()
	defer db.Close()
	rows, err := db.Query(
		getQuery("getStatistic"), chatID)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	var statistics []Statistic

	for rows.Next() {
		s := Statistic{}
		err := rows.Scan(&s.Username, &s.DeletedCount)
		if err != nil {
			log.Println(err)
			continue
		}
		statistics = append(statistics, s)
	}
	return statistics
}
