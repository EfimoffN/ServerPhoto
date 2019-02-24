package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	convert "./convert"
)

func main() {

	router := gin.Default()

	router.MaxMultipartMemory = 16 << 32 // 16 MiB разобраться с MaxMultipartMemory и <<
	router.POST("/convertPhoto", convertPhoto)

	router.Run(":8080")
}

func convertPhoto(con *gin.Context) {

	photoPath := "./photos/bigPhotos/"

	file, _ := con.FormFile("file")
	log.Println(file.Filename)

	con.SaveUploadedFile(file, photoPath+file.Filename)

	con.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	convert.Convert(photoPath + file.Filename)
}
