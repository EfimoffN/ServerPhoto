package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/ConvertJPG", convertJPG)
	router.POST("/ConvertPNG", convertPNG)

	router.Run(":8080")
}

func convertJPG(con *gin.Context) {

	type postTest struct {
		User string `json:"User" binding:"User"` // `form:"User" json:"User"`
		Key  string `json:"Key" binding:"Key"`   // `form:"Key" json:"Key"`
	}

	var data postTest

	// con.Bind(data)

	body, err := ioutil.ReadAll(con.Request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	b := string(body)
	fmt.Println(b)
	err = json.Unmarshal(body, &data)

	user := data.User
	userD := con.DefaultQuery("User", "")
	userP := con.PostForm("User")
	fmt.Println(user)
	fmt.Println(userD)
	fmt.Println(userP)

	con.JSON(http.StatusOK, "ConvertJPG")
}

func convertPNG(con *gin.Context) {

	photoPath := "./photos/"

	file, _ := con.FormFile("file")
	log.Println(file.Filename)

	con.SaveUploadedFile(file, photoPath+file.Filename)

	con.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
