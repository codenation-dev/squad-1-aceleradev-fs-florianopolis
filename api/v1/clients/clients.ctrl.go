package clients

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"gitlab.com/codenation-squad-1/backend/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
)

type Cliente struct {
	Nome string `json:"nome"`
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

func getClients(c *gin.Context) {
	number, err := strconv.ParseInt(c.Param("number"), 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, "O número de elementos por página deve ser um número válido")
		return
	}
	page, err := strconv.ParseInt(c.Param("page"), 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, "O número da página deve ser um número válido")
		return
	}
	skip := number * page
	findOptions := options.FindOptions{
		Limit: &number,
		Skip:  &skip,
		Sort:  bson.D{{"nome", 1}},
	}
	var collection = database.GetCollection(clientsCollection)
	results := make([]*Cliente, number)
	query, err := collection.Find(database.Context, bson.D{{}}, &findOptions)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao tentar consultar a base de dados")
		log.Println(err)
		return
	}
	err = query.All(database.Context, &results)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, map[string]interface{}{
		"clientes": results,
	})
	err = query.Close(database.Context)
	if err != nil {
		log.Println(err)
	}
}
