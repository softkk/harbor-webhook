package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/softkk/harbor-webhook/database"
	_ "github.com/softkk/harbor-webhook/database"
	"github.com/softkk/harbor-webhook/models"
	"gopkg.in/yaml.v2"

	"github.com/gin-gonic/gin"
)

var (
	//ConfigFilePath -
	ConfigFilePath string
)

func init() {
	ConfigFilePath = getenv("CONFIG_FILEPATH", "rule.yaml")
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.Default()
	router.Use(gin.Recovery())

	//config
	var rule models.Rule
	fileBytes, err := ioutil.ReadFile(ConfigFilePath)
	if err != nil {
		log.Panic(err.Error())
		return
	}
	err2 := yaml.Unmarshal(fileBytes, &rule)
	if err2 != nil {
		log.Panic(err2.Error())
		return
	}

	//DB
	defer database.MysqlDB.Close()

	router.POST("/", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println(string(body))

		hookMessage := models.HookMessage{}
		jsonerr := json.Unmarshal(body, &hookMessage)
		if jsonerr != nil {
			log.Println("JSON ERR", jsonerr.Error())
			return
		}

		hookName := hookMessage.EventData.Repository.Name
		hookNamespace := hookMessage.EventData.Repository.Namespace
		switch {
		case hookMessage.Type == "pushImage":
			for _, value := range rule.Pushimage {
				if value.Action == "updateDB" {
					for _, repo := range value.Repository { // docker.io/namespace/name:tag
						if repo.Namespace == hookNamespace && repo.Name == hookName {
							updateDB(hookMessage)
							break
						}
					}
				}
			}
		case hookMessage.Type == "pullImage":
			pullImage(hookMessage)
		default:
			log.Printf("No Definitation, event is %s .", hookMessage.Type)
		}

	})

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	router.Run(":20001")

}
