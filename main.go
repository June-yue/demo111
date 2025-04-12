package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

type Words struct {
	Word         string         `json:"word"`
	Phrases      []Phrase1      `json:"phrases"`
	Translations []Translation1 `json:"translations"`
}

type Phrase1 struct {
	Phrase      string `json:"phrase"`
	Translation string `json:"translation"`
}

type Translation1 struct {
	Translation string `json:"translation"`
	Type        string `json:"type"`
}

func readjson(db *sql.DB, files string) {
	file, err := os.Open(files)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	jsonData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var words []Words
	err = json.Unmarshal(jsonData, &words)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	wordtx, err := tx.Prepare("INSERT OR IGNORE INTO words(word) VALUES(?)")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer wordtx.Close()
	phrasetx, err := tx.Prepare("INSERT OR IGNORE INTO phrases(word_id,phrase,translation) VALUES(?,?,?)")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer phrasetx.Close()
	translationtx, err := tx.Prepare("INSERT OR IGNORE INTO translations(word_id,translation,type) VALUES(?,?,?)")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer translationtx.Close()

	for _, word := range words {
		res, err := wordtx.Exec(word.Word)
		if err != nil {
			tx.Rollback()
			fmt.Println("Error:", err)
			return
		}
		wordId, err := res.LastInsertId()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for _, phrase := range word.Phrases {
			_, err := phrasetx.Exec(wordId, phrase.Phrase, phrase.Translation)
			if err != nil {
				tx.Rollback()
				fmt.Println("Error:", err)
				return
			}
		}
		for _, translation := range word.Translations {
			_, err := translationtx.Exec(wordId, translation.Translation, translation.Type)
			if err != nil {
				tx.Rollback()
				fmt.Println("Error:", err)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func main() {
	db, err := sql.Open("sqlite", "data.db")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer db.Close()
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS words(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		word TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS translations(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		word_id INTEGER NOT NULL,
		translation TEXT NOT NULL,
		type TEXT,
		FOREIGN KEY(word_id) REFERENCES words(id)
	);
	CREATE TABLE IF NOT EXISTS phrases(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		word_id INTEGER NOT NULL,
		phrase TEXT NOT NULL,
		translation TEXT NOT NULL,
		FOREIGN KEY(word_id) REFERENCES words(id)
	);`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	start := time.Now()
	files := []string{"./3-CET4-顺序.json", "./4-CET6-顺序.json"}
	for _, file := range files {
		readjson(db, file)
	}
	//readjson(db, "./4-CET6-顺序.json")
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println(duration.Seconds())
	fmt.Println("Data import completed!")

}
