package router

import (
	"PBLOG/model"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"github.com/unknwon/com"
)

var PAGESIZE int = 20

type Page struct {
	Last    int
	Next    int
	Current int
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// 返回具体的博文
func GetBlog(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	fmt.Printf("id is %v\n", id)
	blog, _ := model.GetBlog(id)

	html := blackfriday.MarkdownBasic([]byte(blog.Content))

	c.HTML(http.StatusOK, "blog.html", gin.H{"id": blog.ID, "title": blog.Title, "content": template.HTML(html)})

}

// 返回博文的总数量和博文的全部标题
func GetBlogs(c *gin.Context) {

	count, _ := model.GetBlogsNum()
	maxPage := count / PAGESIZE
	pageNum := com.StrTo(c.Query("page")).MustInt()

	if pageNum > maxPage {
		fmt.Printf("访问过界")
		// TODO redirect
		location := url.URL{Path: "/blogs"}
		c.Redirect(http.StatusFound, location.RequestURI())
		return
	} else {
		blogs, _ := model.GetBlogs(pageNum, 50)
		page := Page{Last: max(0, pageNum-1), Next: min(maxPage, pageNum+1), Current: pageNum}

		c.HTML(http.StatusOK, "index.html", gin.H{"blogs": blogs, "time": time.Unix(int64(blogs[0].CreatedOn), 0).String(), "paginate": page})
	}

}

// 添加博文
func AddBlog(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "add.html", gin.H{})
	} else if c.Request.Method == "POST" {

		c.Request.ParseForm()

		data := make(map[string]interface{})
		data["title"] = c.Request.PostForm.Get("title")
		data["content"] = c.Request.PostForm.Get("content")
		id := com.StrTo(c.Param("id")).MustInt()
		data["id"] = id
		model.CreateBlog(data)
	}
	location := url.URL{Path: "/blogs"}
	c.Redirect(http.StatusFound, location.RequestURI())

}

// 更新博文
func UpdateBlog(c *gin.Context) {
	if c.Request.Method == "GET" {
		id := com.StrTo(c.Param("id")).MustInt()
		blog, err := model.GetBlog(id)
		if err != nil {
			fmt.Printf("models fun GetBlog Failed, id is %v!", id)
		}

		c.HTML(http.StatusOK, "edit.html", blog)
	}

}

func DeleteBlog(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	model.DeleteBlog(id)
	location := url.URL{Path: "/blogs"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func RouterInit() *gin.Engine {
	r := gin.New()

	r.GET("/blog/:id", GetBlog)
	r.GET("/blogs", GetBlogs)
	r.GET("/blogadd", AddBlog)
	r.POST("/save/:id", AddBlog)
	r.GET("/edit/:id", UpdateBlog)
	r.GET("/delete/:id", DeleteBlog)
	// v1.DELETE("/blog/:id")

	return r
}
