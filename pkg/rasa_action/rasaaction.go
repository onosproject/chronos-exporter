// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package rasa_action

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"net/http"
)

var log = logging.GetLogger("rasa_action")

type RasaAction struct {
	SmtpServer string
}

func NewRasaAction() *RasaAction {
	return &RasaAction{
		SmtpServer: "example.com",
	}
}

func (ra *RasaAction) Run() {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:4200"}
	corsConfig.AllowAllOrigins = false
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))

	//Added to server readinessProbe and livenessProbe
	router.GET("/webhook", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.POST("/webhook", ra.handleAction)

	err := router.Run("0.0.0.0:" + "8081")
	if err != nil {
		log.Fatal(err)
	}
}

func (rs *RasaAction) handleAction(c *gin.Context) {
	switch c.Request.Method {
	case "POST":
		var data payload
		err := c.BindJSON(&data)
		if err != nil {
			fmt.Println("Error in json:",err.Error() )
		}
		if data.NextAction == "email" {
			err := sendMail()
			if err != nil {
				log.Infof(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"action_name": "email",
						"error": err.Error(),
				})
				c.Status(http.StatusInternalServerError)
				return
			}
			c.JSON(http.StatusOK, &response{
				Events:    nil,
				Responses: nil,
			})

		}

	default:
		c.Status(http.StatusNotImplemented)
	}
}

func sendMail() error {
	//TODO:Need to get the details dynamically
	m := gomail.NewMessage()
	m.SetHeader("From", "aliase@gmail.com")
	m.SetHeader("To", "b@gmail.com", "c@gmail.com")
	m.SetAddressHeader("Cc", "d@example.com", "D")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text", "Hello B and C")

	//TODO:Need to check for email userid and password and how we can use here
	d := gomail.NewDialer("smtp.gmail.com", 587, "user", "123456")

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}