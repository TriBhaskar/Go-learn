package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Api call to /hello")
    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func handleGetRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Api call to /greet")
	fmt.Fprintf(w, "GET request received\n")
}

func handlePostRequest(w http.ResponseWriter, req *http.Request) {
    // Check content type
    contentType := req.Header.Get("Content-Type")
    if contentType != "application/json" {
        http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
        return
    }

    // Read the body
    body, err := io.ReadAll(req.Body)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error reading body: %v", err), http.StatusBadRequest)
        return
    }
    defer req.Body.Close()

    // Handle empty body
    if len(body) == 0 {
        http.Error(w, "Request body is empty", http.StatusBadRequest)
        return
    }

    // Parse JSON
    var data map[string]interface{}
    if err := json.Unmarshal(body, &data); err != nil {
        http.Error(w, fmt.Sprintf("Error parsing JSON: %v", err), http.StatusBadRequest)
        return
    }

    // Log the parsed data
    fmt.Printf("Api call to /data with body: %v\n", data)

    // Send JSON response
    w.Header().Set("Content-Type", "application/json")
    response := map[string]interface{}{
        "message": "POST request received",
        "status": "success",
        "received_data": data,
    }
    
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
        return
    }
}

func handlePutRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Api call to /update")
	fmt.Fprintf(w, "PUT request received\n")
}

func handleDeleteRequest(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Api call to /delete")
	fmt.Fprintf(w, "DELETE request received\n")
}

func handlePatchRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Api call to /modify")
	fmt.Fprintf(w, "PATCH request received\n")
}

func handleOptionsRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Api call to /options")
	fmt.Fprintf(w, "OPTIONS request received\n")
}

func handleHeadRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Api call to /head")
	fmt.Fprintf(w, "HEAD request received\n")
}

func handleTraceRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Api call to /trace")
	fmt.Fprintf(w, "TRACE request received\n")
}

func main() {
	mux := http.NewServeMux()
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

	mux.HandleFunc("GET /greet", handleGetRequest)

	mux.HandleFunc("POST /data", handlePostRequest) 

	mux.HandleFunc("PUT /update", handlePutRequest)

	mux.HandleFunc("DELETE /delete", handleDeleteRequest)

	mux.HandleFunc("PATCH /modify", handlePatchRequest)

	mux.HandleFunc("OPTIONS /options", handleOptionsRequest)

	mux.HandleFunc("HEAD /head", handleHeadRequest)

	mux.HandleFunc("TRACE /trace", handleTraceRequest)

    http.ListenAndServe(":8090", mux)
}