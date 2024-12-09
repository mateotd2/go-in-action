package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	// "os"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "example",
		Net:    "tcp",
		Addr:   "db:3306",
		DBName: "exampledb",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// Pruebo el buscar album por artista
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	// Pruebo el añadir el album
	var newAlbum = Album{
		Title:  "Titulo Test",
		Artist: "Artista Test",
		Price:  30.00,
	}

	id, _ := addAlbum(newAlbum)

	fmt.Printf("Nuevo album con titulo %s del artista %s, se añadio con valor %d \n", newAlbum.Title, newAlbum.Artist, id)
}

// Ejemplo donde busco todos los albumes por nombre del artista
func albumsByArtist(name string) ([]Album, error) {
	// Un slice de albums
	var albums []Album
	rows, err := db.Query("SELECT * FROM album WHERE artist =?", name)

	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	defer rows.Close()

	for rows.Next() {
		var alb Album
		err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
		if err != nil {
			return nil, fmt.Errorf("AlbumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("albumByArtist %q: %v", name, err)
	}
	return albums, nil

}

func addAlbum(alb Album) (int64, error) {
	rslt, err := db.Exec("INSERT INTO album(artist, title, price) VALUES (?, ?, ?)", alb.Artist, alb.Title, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum :%v", err)
	}

	id, err := rslt.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum :%v", err)
	}

	return id, nil
}
