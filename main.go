package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	openai "github.com/sashabaranov/go-openai"
)

type Event struct {
	Title       string `json:"title"`
	Location    string `json:"location"`
	DateTime    string `json:"dateTime"`
	EndTime     string `json:"endTime"`
	Description string `json:"description"`
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/generate", handleGenerate)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	input := r.FormValue("input")

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		r.Context(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `Extract event details from the input and provide a JSON output with the following structure:
					{
					  "title": "Event title",
					  "location": "Event location (if available, otherwise empty string)",
					  "dateTime": "Event start date and time in ISO 8601 format (YYYY-MM-DDTHH:MM:SS)",
					  "endTime": "Event end date and time in ISO 8601 format (YYYY-MM-DDTHH:MM:SS), if available",
					  "description": "Event description"
					}`,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: input,
				},
			},
		},
	)

	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	var event Event
	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &event)
	if err != nil {
		http.Error(w, "Error parsing response", http.StatusInternalServerError)
		return
	}

	// If endTime is empty, set it to 1 hour after dateTime
	if event.EndTime == "" {
		event.EndTime = strings.Replace(event.DateTime, "T", "T", 1)
	}

	jsonResponse, err := json.Marshal(event)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
