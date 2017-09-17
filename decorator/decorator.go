package main

import "fmt"

type Request func(string) string

func RequireAuth(req Request) Request {
	return func(data string) string {
		fmt.Println("ensure header contains auth token")
		res := req(data)
		return res
	}
}

func EnsureParam(param string, req Request) Request {
	return func(data string) string {
		fmt.Println("ensure " + param + " is passed in body")
		res := req(data)
		return res
	}
}

func LogJson(req Request) Request {
	return func(data string) string {
		fmt.Println("enable log request response in json format")
		res := req(data)
		return res
	}
}

func FacebookAPI(data string) string {
	return "calling facebook api with -- " + data
}

func main() {
	client := LogJson(RequireAuth(EnsureParam("profile", FacebookAPI)))
	res := client("auth: key, profile: trump")
	fmt.Println(res)
}
