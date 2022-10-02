package main

import (
    "fmt"
    "net/http"
    "task.com/task/interfaces/controllers"
)

func main() {
    fmt.Print("Webサーバーを起動します")
    handleRequests()
   }
    
   func handleRequests() {
    mux := http.NewServeMux()
    mux.Handle("/task", http.HandlerFunc(TaskListView))
    http.ListenAndServe(":3000", mux)
   }
    
   func task (w http.ResponseWriter, r *http.Request){
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "こんにちは" + r.FormValue("name") + "さん！")
   }