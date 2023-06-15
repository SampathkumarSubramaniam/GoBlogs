package modernWebSubPackage

import (
	"GoBlogs/modernWeb"
	"log"
)

func tryfromSubPackage() {
	tryMe := modernWeb.CheckMe{
		Name: "Jonasa",
		Age:  36,
	}
	log.Println(tryMe)
}
