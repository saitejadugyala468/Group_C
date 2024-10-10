package main

//"encoding/json"
//"log"
//"net/http"
//"strconv"

//"github.com/gorilla/mux"

// Task represents a to-do item with an ID, Title, Description, and Status
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` // "pending" or "completed"
}

var tasks []Task // In-memory task storage
var nextID = 1   // Task ID counter

func main() {
}
