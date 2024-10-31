package router

import (
	"net/http"
)

// SetupRoutes 设置路由
func SetupRoutes() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/multiply", MultiplyHandler)
	http.HandleFunc("/add", AddHandler)
	http.HandleFunc("/result", ResultHandler)
}
