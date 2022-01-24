// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package rasa

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"io/ioutil"
	"net/http"
	"os"
)

var log = logging.GetLogger("rasa")

type Rasa struct {
	ModelPath string
}

func NewRasa(modelPath string) *Rasa {
	return &Rasa{
		ModelPath:    modelPath,
	}
}

func (rs *Rasa) Run() {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:4200"}
	corsConfig.AllowAllOrigins = false
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))

	//Added to server readinessProbe and livenessProbe
	router.GET("/models", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.GET("/models/:model", rs.handleModel)

	err := router.Run("0.0.0.0:" + "8080")
	if err != nil {
		log.Fatal(err)
	}
}

func (rs *Rasa) handleModel(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		model, okay := c.Params.Get("model")
		if !okay {
			c.String(http.StatusNotFound, "error")
		}
		ifNonMatchValue := c.Request.Header.Get("If-None-Match")

		modelPath := rs.ModelPath + model + ".tar.gz"

		//Check for the file exist or not
		_, err := os.Stat(modelPath)
		if os.IsNotExist(err) {
			c.Status(http.StatusNotFound)
			return
		}

		log.Infof("Serving file %s", modelPath)

		////Check hash value for the model file
		hashValue, err := getHashValue(modelPath)
		if err != nil {
				log.Error(err.Error())
				c.Status(http.StatusInternalServerError)
				return
		}

		if ifNonMatchValue == hashValue {
			c.Status(http.StatusNotModified)
		} else {
			//c.Header("Content-Type", "application/tar+gzip") // Found on web for tar.gz but seems doesnot working
			//c.Header("Content-Type", "application/x-gzip") // Its also not working
			c.Header("Content-Type", "application/gzip") //It is used for .gz
			c.Header("ETag", hashValue)
			c.File(modelPath)
			c.Status(http.StatusOK)
		}

	default:
		c.Status(http.StatusNotImplemented)
	}
}

func getHashValue(modelPath string) (string, error){
	newHash := sha256.New()
	byteData, err := ioutil.ReadFile(modelPath)
	newHash.Write(byteData)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(newHash.Sum(nil)), nil
}
