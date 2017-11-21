package register

import (
	"encoding/json"
	"errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"model"
	"net/http"
)

var json_temp = gin.H{
	"errcode": 0,
	"msg":     "",
}

func localRegister(c *gin.Context) (int, string, errcode.ErrCodeType) {

	var bodyMap map[string]interface{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &bodyMap)
	username, exist := bodyMap["username"]
	if !exist {
		return http.StatusOK, errcode.ErrStrPm + ":username", errcode.ErrPm
	}
	password, paexist := bodyMap["password"]
	if !paexist {
		return http.StatusOK, errcode.ErrStrPm + ":password", errcode.ErrPm
	}
	uuidV4 := uuid.NewV4()
	uuidV4Str := uuidV4.String()
	usernameUserId := model.UsernameUseridIndex{username.(string), uuidV4Str}
	fmt.Println(usernameUserId)
	var user model.Users
	user.Id = uuidV4Str
	user.Name = username.(string)
	user.Password = password.(string)

	engine1 := model.GetEngine(model.UsernameUserIdIndexConstModel)
	session1 := engine1.NewSession()
	engine2 := model.ShardUuid(model.UsersConstModel, uuidV4)
	session2 := engine2.NewSession()
	defer session1.Close()
	defer session2.Close()
	err := session1.Begin()
	if err != nil {
		panic(err)
	}
	err = session2.Begin()
	if err != nil {
		panic(err)
	}
	_, err = session1.Insert(&usernameUserId)
	if err != nil {
		fmt.Println(err)
		session1.Rollback()
		return http.StatusOK, errcode.ErrStrAe, errcode.ErrAe
	}
	_, err = session2.Insert(&user)
	if err != nil {
		fmt.Println(err)
		session1.Rollback()
		session2.Rollback()
		return http.StatusOK, errcode.ErrStrAe, errcode.ErrAe
	}
	err = session1.Commit()
	if err != nil {
		panic(err)
	}
	err = session2.Commit()
	if err != nil {
		panic(err)
	}
	return http.StatusOK, errcode.ErrStrOk, errcode.ErrOk
}

func Register(c *gin.Context) {
	way := c.Param("way")
	switch way {
	case "phone":
		c.JSON(http.StatusOK, gin.H{
			"errcode": errcode.ErrNotImpl,
			"msg":     errcode.ErrStrNotImpl,
		})
	case "local":
		status, msg, err := localRegister(c)
		c.JSON(status, gin.H{
			"errcode": err,
			"msg":     msg})
	case "wechat":
	case "qq":
	}
}
