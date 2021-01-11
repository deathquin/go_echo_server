package routes

import (
	"fmt"
	"github.com/labstack/echo"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// Handler
func sayHello(c echo.Context) error {
	fmt.Println("Say Hello")
	return c.String(http.StatusOK, "Hello, World!")
}

func sayHello2(c echo.Context) error {

	var content struct {
		Response  string    `json:"response"`
		Timestamp time.Time `json:"timestamp"`
		Random    int       `json:"random"`
	}

	content.Response = "Sent via JSONP"
	content.Timestamp = time.Now().UTC()
	content.Random = rand.Intn(1000)

	return c.JSON(http.StatusOK, &content)
}

func IndexRoutes(group *echo.Group) {
	/*group.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})*/
	group.GET("", sayHello)
	group.GET("/hello", sayHello2)
	group.POST("/single/upload", upload)
	group.POST("/multiple/upload", MultipleUpload)
}

func MultipleUpload(c echo.Context) error {

	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	//------------
	// Read files
	//------------

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	defer form.RemoveAll()

	files := form.File["files"]

	for _, file := range files {

		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create("public/" + file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>Uploaded successfully %d files with fields name=%s and email=%s.</p>", len(files), name, email))
}

// single file upload
func upload(c echo.Context) error {

	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	//-----------
	// Read file
	//-----------

	// https://prometheo.tistory.com/68
	// 용량이 큰 파일을 업로드하는 경우에 디스크 사용량이 파일 사이즈의 2배가 늘어났다.
	// 원인을 찾아보니 echo에서 설정한 메모리(32 MB)보다 큰 경우 /tmp/에 파일을 저장하고 있었다.
	//다음과 같이 하면 파일을 원하는 곳에 복사한 다음 임시 파일을 모두 제거할 수 있다.
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	defer form.RemoveAll()

	// Source
	// file html tag name file
	file, err := c.FormFile("file")

	if err != nil {
		return err
	}

	src, err := file.Open()

	if err != nil {
		return err
	}

	defer src.Close()

	// Destination
	dst, err := os.Create("public/" + file.Filename)

	if err != nil {
		return err
	}

	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))

}
