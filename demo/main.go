package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

type Bot struct {
	db *sql.DB
}

type KnowledgeBaseEntry struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func NewBot() (*Bot, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Bot{db: db}, nil
}

func (b *Bot) Close() {
	b.db.Close()
}

func (b *Bot) GetAnswer(question string) (string, error) {
	query := `SELECT answer FROM knowledge_base WHERE question ILIKE $1 LIMIT 1`

	var answer string
	err := b.db.QueryRow(query, "%"+question+"%").Scan(&answer)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // Return nil if no answer found
		}
		return "", err
	}

	return answer, nil
}

func (b *Bot) AddEntry(question, answer string) error {
	query := `INSERT INTO knowledge_base (question, answer) VALUES ($1, $2)`

	_, err := b.db.Exec(query, question, answer)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	bot, err := NewBot()
	if err != nil {
		log.Fatal("Error creating bot:", err)
	}
	defer bot.Close()

	router := mux.NewRouter()
	router.HandleFunc("/knowledgebase", AddEntryHandler).Methods("POST")
	router.HandleFunc("/knowledgebase", GetAnswerHandler).Methods("GET")

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func AddEntryHandler(w http.ResponseWriter, r *http.Request) {
	var entry KnowledgeBaseEntry
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	bot := getBotInstance()
	err = bot.AddEntry(entry.Question, entry.Answer)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetAnswerHandler(w http.ResponseWriter, r *http.Request) {
	
	question := r.URL.Query().Get("question")
	if question == "" {
		http.Error(w, "Missing question parameter", http.StatusBadRequest)
		return
	}

	bot := getBotInstance()
	answer, err := bot.GetAnswer(question)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"question": question,
		"answer":   answer,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getBotInstance() *Bot {
	// You can implement a singleton pattern here if required
	bot, err := NewBot()
	if err != nil {
		log.Fatal("Error creating bot:", err)
	}
	return bot
}
