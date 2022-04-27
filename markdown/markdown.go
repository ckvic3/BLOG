package markdown

import (
	"fmt"
	"os"

	"github.com/russross/blackfriday"
)

func Parser() string {

	input, err := os.ReadFile("./doc/arch.md")
	if err != nil {
		fmt.Printf("failed open markdown file %v", err.Error())
	}

	html := blackfriday.MarkdownBasic(input)
	// file, err := os.OpenFile("./doc/blog.html", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println("file open success!")
	// }
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// file.Write(html)
	// file.Close()
	// fmt.Printf("file write ok!")
	return string(html)
}

// html := markdown.Parser()

// 	data := make(map[string]interface{})
// 	data["title"] = "first blog"
// 	data["content"] = html
// 	if model.CreateBlog(data) {
// 		fmt.Println("create a new blog success!")
// 	}

// 	// if model.DeleteBlog(3) {
// 	// 	fmt.Printf("delete blog which id is %v\n", 1)
// 	// }

// 	if _, err := model.GetBlog(1); err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	data["title"] = "first blog version 2"
// 	if model.UpdataBlog(2, data) {
// 		fmt.Printf("update blog which id is %v success", 2)
// 	}
