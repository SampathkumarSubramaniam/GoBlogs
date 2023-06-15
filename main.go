package main

import (
	"GoBlogs/modernWeb"
)

func main() {
	modernWeb.StartServer()
	json := `{
  "status": "some context",
  "class": "first",
  "grade": "A++"
}`
}
