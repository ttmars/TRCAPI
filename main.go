package main

import (
	"TRCAPI/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var Token = "hiseven"

func main()  {
	r := gin.Default()
	r.GET("/getTransInfo", GetTransInfoAPI)
	r.GET("/getNowBlockNum", GetNowBlockNumAPI)
	r.Run(":12345") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GetNowBlockNumAPI(c *gin.Context)  {
	r,err := api.DTRCNodeAPI.GetNowBlockNum()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "数据获取失败！",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": r,
	})
}

func GetTransInfoAPI(c *gin.Context)  {
	token,bool := c.GetQuery("token")
	if bool != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误！",
		})
		return
	}
	start,bool := c.GetQuery("start")
	if bool != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误！",
		})
		return
	}
	end,bool := c.GetQuery("end")
	if bool != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误！",
		})
		return
	}

	startNum,err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误！",
		})
		return
	}
	endNum,err := strconv.ParseInt(end, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误！",
		})
		return
	}

	if token != Token {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token认证失败！",
		})
		return
	}

	data,err := api.DTRCNodeAPI.GetTransInfo(startNum, endNum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "数据获取失败！",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
