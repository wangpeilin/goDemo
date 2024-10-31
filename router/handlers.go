package router

import (
	"fmt"
	"goDemo/logic"
	"html/template"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html", "templates/multiply.html", "templates/result.html"))

// HomeHandler 处理主页请求
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HomeHandler here")
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

// MultiplyHandler 处理乘法计算请求
func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		num1Str := r.FormValue("num1")
		num2Str := r.FormValue("num2")

		// 验证输入
		if !logic.IsValidInput(num1Str, num2Str) {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// 计算结果
		fmt.Println("multiply input", num1Str, num2Str)
		num1, _ := strconv.Atoi(num1Str)
		num2, _ := strconv.Atoi(num2Str)
		result := logic.Multiply(num1, num2)
		fmt.Println("multiply result", result)

		// 重定向到结果页面
		http.Redirect(w, r, fmt.Sprintf("/result?result=%d", result), http.StatusSeeOther)
		return
	}
	tmpl.ExecuteTemplate(w, "multiply.html", nil)
}

// AddHandler 处理加法计算请求
func AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		num1Str := r.FormValue("num1")
		num2Str := r.FormValue("num2")

		// 验证输入
		if !logic.IsValidInput(num1Str, num2Str) {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// 计算结果
		fmt.Println("add input", num1Str, num2Str)
		num1, _ := strconv.Atoi(num1Str)
		num2, _ := strconv.Atoi(num2Str)
		result := logic.Add(num1, num2)
		fmt.Println("add result", result)

		// 重定向到结果页面
		http.Redirect(w, r, fmt.Sprintf("/result?result=%d", result), http.StatusSeeOther)
		return
	}
	tmpl.ExecuteTemplate(w, "multiply.html", nil)
}

// ResultHandler 处理结果页面请求
func ResultHandler(w http.ResponseWriter, r *http.Request) {
	result := r.URL.Query().Get("result")
	tmpl.ExecuteTemplate(w, "result.html", result)
}
