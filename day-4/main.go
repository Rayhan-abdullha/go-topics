package day4

import (
	"fmt"
	"strings"
)

// Middleware টাইপ: ফাংশন যা একটা string function নেয় আর নতুন string function return করে
type Middleware func(func(string) string) func(string) string

// Chain ফাংশন: একাধিক middleware একসাথে apply করে
func Chain(f func(string) string, middlewares ...Middleware) func(string) string {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
func UppercaseMiddleware(next func(string) string) func(string) string {
	return func(s string) string {
		fmt.Println("UppercaseMiddleware before")
		result := next(s)
		result = strings.ToUpper(result)
		fmt.Println("UppercaseMiddleware after")
		return result
	}
}
func HelloMiddleware(next func(string) string) func(string) string {
	return func(s string) string {
		fmt.Println("HelloMiddleware before")
		result := next(s)
		result = "Hello " + result
		fmt.Println("HelloMiddleware after")
		return result
	}
}

// মূল ফাংশন
func FinalFunction(s string) string {
	return s + " World"
}

func Main() {
	chained := Chain(
		FinalFunction,
		UppercaseMiddleware,
		HelloMiddleware,
	)
	result := chained("golang")
	fmt.Println("Result:", result)
}
