package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIHandler struct {}

func NewAPIHandler() *APIHandler {
	return &APIHandler{}
}

func (h *APIHandler) Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Api call to /hello")
    fmt.Fprintf(w, "hello\n")
}

func (h *APIHandler) Headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func (h *APIHandler) HandleGet(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Api call to /greet")
	fmt.Fprintf(w, "GET request received\n")
}

func (h *APIHandler) HandlePost(w http.ResponseWriter, req *http.Request) {
    contentType := req.Header.Get("Content-Type")
    if contentType != "application/json" {
        http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
        return
    }

    body, err := io.ReadAll(req.Body)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error reading body: %v", err), http.StatusBadRequest)
        return
    }
    defer req.Body.Close()

    if len(body) == 0 {
        http.Error(w, "Request body is empty", http.StatusBadRequest)
        return
    }

    var data map[string]interface{}
    if err := json.Unmarshal(body, &data); err != nil {
        http.Error(w, fmt.Sprintf("Error parsing JSON: %v", err), http.StatusBadRequest)
        return
    }

    fmt.Printf("Api call to /data with body: %v\n", data)

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

func (h *APIHandler) HandlePut(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Api call to /update")
	fmt.Fprintf(w, "PUT request received\n")
}

func (h *APIHandler) HandleDelete(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Api call to /delete")
	fmt.Fprintf(w, "DELETE request received\n")
}

func (h *APIHandler) HandlePatch(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Api call to /modify")
	fmt.Fprintf(w, "PATCH request received\n")
}

func (h *APIHandler) HandleOptions(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Api call to /options")
	fmt.Fprintf(w, "OPTIONS request received\n")
}

func (h *APIHandler) HandleHead(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Api call to /head")
	fmt.Fprintf(w, "HEAD request received\n")
}

func (h *APIHandler) HandleTrace(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Api call to /trace")
	fmt.Fprintf(w, "TRACE request received\n")
}