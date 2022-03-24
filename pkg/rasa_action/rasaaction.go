// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rasa_action

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"net/http"
)

var log = logging.GetLogger("rasa_action")

type RasaAction struct {
	SmtpServer   string
	SmtpServPort int
	Port         string
	SmtpUser     string
	SmtpUserPass string
}

func NewRasaAction(smtpServ string, smtpSerPort int, port string, user string, pass string) *RasaAction {
	return &RasaAction{
		SmtpServer:   smtpServ,
		SmtpServPort: smtpSerPort,
		Port:         port,
		SmtpUser:     user,
		SmtpUserPass: pass,
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

	err := router.Run("0.0.0.0:" + ra.Port)
	if err != nil {
		log.Fatal(err)
	}
}

func (ra *RasaAction) handleAction(c *gin.Context) {
	switch c.Request.Method {
	case "POST":
		var data payload
		err := c.BindJSON(&data)
		if err != nil {
			log.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		if data.NextAction == "email" {
			err := ra.sendMail()
			if err != nil {
				log.Infof(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"action_name": "email",
					"error":       err.Error(),
				})
				c.Status(http.StatusInternalServerError)
				return
			}
			c.JSON(http.StatusOK, &response{
				Events:    nil,
				Responses: nil,
			})
		}
		c.Status(http.StatusBadRequest)
	default:
		c.Status(http.StatusNotImplemented)
	}
}

func (ra *RasaAction) sendMail() error {
	//TODO:Need to get the details dynamically
	m := gomail.NewMessage()
	m.SetHeader("From", "aliase@gmail.com")
	m.SetHeader("To", "b@gmail.com", "c@gmail.com")
	m.SetAddressHeader("Cc", "d@example.com", "D")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text", "Hello B and C")

	//TODO:Need to check for email userid and password and how we can use here
	d := gomail.NewDialer(ra.SmtpServer, ra.SmtpServPort, ra.SmtpUser, ra.SmtpUserPass)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
