package main

import (
	"log"
	"net/http"

	"transaction_processor/facade"
	"transaction_processor/utils"

	"github.com/gin-gonic/gin"
)

var config utils.FileConfig

func init() {
	var err error
	config, err = utils.LoadConfigFromEnv()
	if err != nil {
		log.Fatalf("Error al cargar la configuración: %v", err)
	}

	log.Printf("Configuración cargada: %+v", config)
}

func main() {
	router := gin.Default()

	router.POST("/process", func(c *gin.Context) {
		processor := facade.NewProcessorFacade(config)
		if err := processor.Run(); err != nil {
			log.Printf("Error during processing: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		log.Println("Processing completed successfully.")
		c.JSON(http.StatusOK, gin.H{"message": "Transactions processed successfully"})
	})

	port := config.Port
	if port == "" {
		port = "8080" // default port
	}

	router.Run(":" + port)
}
