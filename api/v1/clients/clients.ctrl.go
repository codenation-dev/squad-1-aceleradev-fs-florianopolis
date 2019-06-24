package clients

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"gitlab.com/codenation-squad-1/backend/database"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Cliente struct {
	Nome string
}

const clientsCollection = "clientes-uati"

func addCSVToDatabase(c *gin.Context) {
	file, _ := c.FormFile("clientes.csv")
	_ = c.SaveUploadedFile(file, "clientes.csv")

	defer removeFile()

	csvFile, err := os.Open("clientes.csv")
	if err != nil {
		log.Println(err)
	}
	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	var clientes []interface{}
	for {
		item, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		clientes = append(clientes, Cliente{Nome: item[0]})
	}

	var collection = database.GetCollection(clientsCollection)
	insertManyResult, err := collection.InsertMany(context.TODO(), clientes)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "F")
		return
	}
	log.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func removeFile() {
	err := os.Remove("clientes.csv")
	if err != nil {
		log.Println(err)
	}
}
