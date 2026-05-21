package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// Datenstrukturen für den Transfer
type TaskRequest struct {
	UserPrompt string `json:"user_prompt"`
}

type AIResponse struct {
	Response string `json:"response"`
}

func main() {
	// Logger initialisieren
	initLogger()

	http.HandleFunc("/api/chat", handleChat)
	http.HandleFunc("/api/tasks", handleTasks) // Platzhalter für CRUD

	port := ":8080"
	log.Printf("Taskify Backend läuft auf Port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// Kommunikation: Frontend -> Backend -> KI -> Frontend
func handleChat(w http.ResponseWriter, r *http.Request) {
	var req TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("ERROR: Ungültiger Request - %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("INFO: Anfrage erhalten: %s", req.UserPrompt)

	// KI-Anfrage (Llama.cpp läuft meist auf Port 8081 im Docker-Netz)
	aiResp, err := callAI(req.UserPrompt)
	if err != nil {
		log.Printf("ERROR: KI-Modell nicht erreichbar - %v", err)
		http.Error(w, "AI Service Unavailable", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aiResp)
}

func callAI(prompt string) (AIResponse, error) {
	// Hier wird der Traffic zur KI geroutet
	aiURL := os.Getenv("AI_ENDPOINT") // z.B. http://llama-cpp:8081/completion

	// Beispielhafte Implementation eines Proxy-Calls zur KI
	resp, err := http.Post(aiURL, "application/json", bytes.NewBuffer([]byte(prompt)))
	if err != nil {
		return AIResponse{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return AIResponse{Response: string(body)}, nil
}
