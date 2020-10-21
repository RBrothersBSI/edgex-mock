package handlers

import (
	"server/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

var devices domain.Devices

type PongRes struct {
	Message string `json:"message"`
}

func Pong(c *gin.Context) {
	pong := PongRes{"pong"}
	c.JSON(200, pong)
}

func Device(c *gin.Context) {
	count := c.Params.ByName("count")
	val, _ := strconv.Atoi(count)
	devices = make(domain.Devices, val)
	generateDevices(val)
}

func GetAll(c *gin.Context) {
	c.JSON(200, devices)
}

func generateDevices(count int) domain.Devices {
	for i := 0; i < count; i++ {
		deviceNum := strconv.Itoa(i)
		devices[i] = domain.Device{
			Id:             deviceNum,
			Name:           "TestDevice" + deviceNum,
			AdminState:     "ENABLED",
			OperatingState: "Unlocked",
			Labels:         []string{"test", deviceNum},
			Commands: []domain.Command{
				{
					Created:  123123223,
					Modified: 323423133,
					Id:       deviceNum,
					Name:     "Name",
					Get: domain.CommandResponse{
						Path: "/api/vi/device/" + deviceNum + "/Name",
						Url:  "http://core-command:48082/api/v1/device/" + deviceNum + "/comand/" + deviceNum,
						Responses: []domain.Response{
							{
								Code:           "200",
								Description:    "Success",
								ExpectedValues: []string{"Name"},
							},
						},
					},
				},
			},
		}
	}
	return devices
}
