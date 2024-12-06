// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"github.com/go-sql-driver/mysql"
// 	"log"
// 	"os"
// )

// var db *sql.DB

// func main() {
// 	// Capture connection properties.
// 	cfg := mysql.Config{
// 		User:   os.Getenv("DBUSER"),
// 		Passwd: os.Getenv("DBPASS"),
// 		Net:    "tcp",
// 		Addr:   "127.0.0.1:3306",
// 		DBName: "recordings",
// 	}
// 	// Get a database handle.
// 	var err error
// 	db, err = sql.Open("mysql", cfg.FormatDSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	fmt.Println("Connected!")
// }
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Define la cadena de conexión
    dsn := "user:password@tcp(db:3306)/exampledb" // Cambia "user", "password" y "exampledb" según tu configuración

    // Abre la conexión a la base de datos
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error al conectar con la base de datos: %v", err)
    }
    defer db.Close()

    // Prueba la conexión
    err = db.Ping()
    if err != nil {
        log.Fatalf("No se pudo conectar a la base de datos: %v", err)
    }

    fmt.Println("Conexión a la base de datos exitosa!")
}
