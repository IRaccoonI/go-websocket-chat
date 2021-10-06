package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var db *sql.DB

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// remove the previous fmt statement
	// fmt.Fprintf(w, "Hello World")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// helpful log statement to show connections
	log.Println("Client Connected")

	reader(ws)

}

func migrateDB() {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Doesn't get db driver: %v\n", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://C:/Users/Depo/Documents/go-websocket-chat/back/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Doesn't get migrate instance: %v\n", err)
	}
	m.Steps(999)
}

func connect2db() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Error connect to db: %v\n", err)
	}
	return db
}

func loadDotEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func setupRoutes() {
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Run server!")

	fmt.Println("Load .env...")
	loadDotEnv()

	fmt.Println("Connect to db...")
	db = connect2db()

	fmt.Println("Try to migrate...")
	migrateDB()

	fmt.Println("Setup routes...")
	setupRoutes()

	fmt.Println("Listen!")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
