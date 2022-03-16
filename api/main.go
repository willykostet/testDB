package main

import (
	// "config"
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

type IncomingPostback struct {
	Id        int32     `db:"id"`
	TrackerId int32     `db:"trackerid"`
	CnvStatus string    `db:"cnv_status"`
	Payout    float64   `db:"payout"`
	Currency  string    `db:"currency"`
	UrlQuery  string    `db:"url_query"`
	RequestIp string    `db:"request_ip"`
	CreatedAt time.Time `db:"created_at"`
}
type Tracker struct {
	Trackername sql.NullString `db:"tracker_name"`
}
type SendPostback struct {
	Trackername sql.NullString `db:"tracker_name"`
}
type SendPostbackFailed struct {
	Trackername sql.NullString `db:"tracker_name"`
}
type Storage struct {
	db *pgx.Conn
}

func envOr(envor, or string) string {
	if os.Getenv(envor) != "" {
		return os.Getenv(envor)
	}
	return or
}
func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, envOr(os.Getenv("DB_CONN"), "postgres://postgres:@localhost:5432/postgres"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)
	str := &Storage{db: conn}
	r := mux.NewRouter()
	// GET
	r.HandleFunc("/incoming", str.getIncomingPostback)
	// r.HandleFunc("/tracker", str.getTracker)
	// r.HandleFunc("/sendpostback", str.getSendPostback)
	// r.HandleFunc("/sendpostbackfailed", str.getTracker)
	log.Fatal(http.ListenAndServe(":8000", r))
}

// func GeneratePostgresDatabaseURL(conf config.DBConfig) string {
// 	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
// 		conf.User,
// 		conf.Password,
// 		conf.Host,
// 		conf.Port,
// 		conf.Name,
// 		conf.SSLMode,
// 	)
// }

func (s *Storage) getIncomingPostback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	incoming, err := s.currentIncoming()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(incoming)
}

// func (s *Storage) getTracker(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Origin", "*")
// 	w.Header().Set("Content-Type", "application/json")
// 	tracker, err := s.currentTracker()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(tracker)
// }

// func (s *Storage) getSendPostback(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Origin", "*")
// 	w.Header().Set("Content-Type", "application/json")
// 	spb, err := s.currentSendPostback()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(spb)
// }

// func (s *Storage) getSendPostbackFailed(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Origin", "*")
// 	w.Header().Set("Content-Type", "application/json")
// 	spbfailed, err := s.currentSendPostbackFailed()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(spbfailed)
// }
func (s *Storage) currentIncoming() ([]IncomingPostback, error) {
	var Incoming []IncomingPostback

	err := s.db.QueryRow(context.Background(), "SELECT * FROM incoming_postback").Scan(&Incoming)
	if err != nil {
		log.Printf(``)
		return nil, err
	}

	return Incoming, nil
}

// func (s *Storage) currentTracker() ([]Tracker, error) {
// 	var currentTracker []Tracker

// 	db, err := sql.Open("mysql", dsn("pb_reciever_db"))
// 	if err != nil {
// 		log.Fatal(err)
// 		return currentTracker, err
// 	}

// 	rows, err := db.Query("SELECT * FROM tracker")
// 	if err != nil {

// 		return currentTracker, err
// 	}

// 	for rows.Next() {
// 		t := Tracker{}
// 		err = rows.Scan(&t.Trackername)
// 		if err != nil {
// 			log.Fatal(err)
// 			return currentTracker, err
// 		}
// 		currentTracker = append(currentTracker, t)
// 	}
// 	return currentTracker, nil
// }

// func (s *Storage) currentSendPostback() ([]Tracker, error) {
// 	var currentTracker []Tracker

// 	// db, err := sql.Open("mysql", dsn("pb_reciever_db"))
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// 	return currentTracker, err
// 	// }

// 	rows, err := db.Query("SELECT * FROM tracker")
// 	if err != nil {

// 		return currentTracker, err
// 	}

// 	for rows.Next() {
// 		t := Tracker{}
// 		err = rows.Scan(&t.Trackername)
// 		if err != nil {
// 			log.Fatal(err)
// 			return currentTracker, err
// 		}
// 		currentTracker = append(currentTracker, t)
// 	}
// 	return currentTracker, nil
// }
