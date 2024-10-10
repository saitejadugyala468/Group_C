# Creating a simple CRUD API in Go

This project is simple **CRUD**( Create, Read, Update and Delete) API built with Go version 1.23.1 using only the built-in Go package. This API demonstrates the basic principles of a CRUD application by handling HTTP requests for managing a collection of users stored in-memory. The API manages a collection of tasks, each with a title, description, and status.

# installation and setup
1. At first we have to install the GO on your local machine.
2. Clone the repository  git clone https://github.com/yourusername/your-repo-name.git 

## Folder structure

In this project the main.go file plays a key role because it contains the application logic. Where we perform all CRUD operations.

## Task Struct

Define a Task struct in Go with the following fields:


type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`  // "pending" or "completed"
}

## Implement CRUD operations

1. Create (POST /tasks):- Create a new task by accepting a JSON payload with a title and description.
2. Read:
a. Get all tasks (GET /tasks): Return a list of all tasks.
b. Get a task by ID (GET /tasks/{id}): Return a specific task by its ID.
3. Update (PUT /tasks/{id}): Update the title, description, or status of an existing task.
4. Delete (DELETE /tasks/{id}): Delete a task by its ID.

## Data Storage

For this project, store tasks in an in-memory slice of Task structs (no database required).

## start the server
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", createTask).Methods("POST")
	r.HandleFunc("/tasks", getTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", getTaskByID).Methods("GET")
	r.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	log.Println("Server starting on port 8089...")
	log.Fatal(http.ListenAndServe(":8089", r))
}

## Testing API
Using curl instructions or a program like Postman, you can test the API.
# curl first method
1. create a new task
curl -X POST http://localhost:8089/tasks \
-H "Content-Type: application/json" \
-d '{"title": "Test Task", "description": "This is a test task for group C", "status": "Pending"}'

2. Get ALL Tasks
curl http://localhost:8089/tasks

3. Get a task by ID
curl http://localhost:8089/tasks/1

4. update a task
curl -X PUT http://localhost:8089/tasks/1 \
-H "Content-Type: application/json" \
-d '{"title": "Updated Task for Clod 2003", "description": "Group C", "status": "completed"}'

5. Delete a Task

curl -X DELETE http://localhost:8089/tasks/1

## Explanation of CRUD Operations
CRUD is an acronym for the four basic operations that can be performed on data in persistent storage:

1. Create: Adding a new record (HTTP POST).
2. Read: Retrieving records (HTTP GET).
3. Update: Modifying an existing record (HTTP PUT).
4. Delete: Removing a record (HTTP DELETE).

## Running the application
 To run the API we use the command mentioned below
**go run main.go**

## Conclusion
This is a basic example of a CRUD API built with Go using only the built-in standard library.


