package utils

import "github.com/gin-gonic/gin"

func Https(code int, version string, endpoint string, data string) gin.H {
	mydata := make(gin.H)

	mydata["status_code"] = code
	mydata["version"] = version
	mydata["endpoint"] = endpoint
	mydata["data"] = data

	return mydata
}
