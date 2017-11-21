package session

import (
	"config"
	"encoding/json"
	"errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"model"
	"net/http"
	"time"
)

func createLocalSession(c *gin.Context) (int, string, errcode.ErrCodeType, *model.Users) {
	var bodyMap map[string]interface{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &bodyMap)
	username, exist := bodyMap["username"]
	if !exist {
		return http.StatusOK, errcode.ErrStrPm + ":username", errcode.ErrPm, nil
	}
	password, paexist := bodyMap["password"]
	if !paexist {
		return http.StatusOK, errcode.ErrStrPm + ":password", errcode.ErrPm, nil
	}
	//get uid from model UsernameUserIdIndex
	unmUidEngine := model.GetEngine(model.UsernameUserIdIndexConstModel)
	unmUidIndex := &model.UsernameUseridIndex{Username: username.(string)}
	has, err := unmUidEngine.Get(unmUidIndex)
	if err != nil {
		fmt.Println(err)
	}
	if !has {
		return http.StatusOK, errcode.ErrStrUnPw, errcode.ErrUnPw, nil
	}
	// get user by userid
	uuidv4, _ := uuid.FromString(unmUidIndex.Userid)
	userEngine := model.ShardUuid(model.UsersConstModel, uuidv4)
	user := &model.Users{Id: unmUidIndex.Userid, Password: password.(string)}
	has, err = userEngine.Get(user)
	if err != nil {
		fmt.Println(err)
	}
	if !has {
		return http.StatusOK, errcode.ErrStrUnPw, errcode.ErrUnPw, nil
	}
	// get sid from model UseridSessionidIndex
	uidSidEngine := model.GetEngine(model.UseridSessionidIndexConstModel)
	uidSidIndex := &model.UseridSessionidIndex{Userid: user.Id}
	has, err = uidSidEngine.Get(uidSidIndex)
	if err != nil {
		fmt.Println(err)
	}
	thisSession := &model.Sessions{}
	if has { // session exist
		thisSession = &model.Sessions{Id: uidSidIndex.Sessionid}
		uuidv4, _ := uuid.FromString(thisSession.Id)
		sessionEngine := model.ShardUuid(model.SessionsConstModel, uuidv4)
		has, err = sessionEngine.Get(thisSession)
		if err != nil {
			fmt.Println(err)
		}
		if !has {
			fmt.Println("Integret Error:sessionid userid")
			return http.StatusOK, errcode.ErrStrUnPw, errcode.ErrUnPw, nil
		}
		if thisSession.ExpiredAt.Before(time.Now()) {
			return http.StatusOK, errcode.ErrStrExp, errcode.ErrExp, nil
		}

	} else { //sesion not exist
		//todo: create_or_update duplicate
		uuidV4 := uuid.NewV4()
		uuidV4Str := uuidV4.String()
		thisSession = &model.Sessions{Id: uuidV4Str, Userid: user.Id, ExpiredAt: time.Now().Add(time.Duration(config.SessionAgeHour) * time.Hour)}
		sessionEngine := model.ShardUuid(model.SessionsConstModel, uuidV4)
		affected, inerr := sessionEngine.Insert(thisSession)
		if affected != 1 || inerr != nil {
			panic(inerr)
		}
		uidSidIndex.Userid = user.Id
		uidSidIndex.Sessionid = uuidV4Str
		affected, inerr = uidSidEngine.Insert(uidSidIndex)
		if affected != 1 || inerr != nil {
			panic(inerr)
		}
	}
	cookie := &http.Cookie{
		Name:     config.SessionName,
		Value:    thisSession.Id,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
	fmt.Println(user)
	return http.StatusOK, errcode.ErrStrOk, errcode.ErrOk, user
}

func Create(c *gin.Context) {
	way := c.Param("way")
	switch way {
	case "phone":
		c.JSON(http.StatusOK, gin.H{
			"errcode": errcode.ErrNotImpl,
			"msg":     errcode.ErrStrNotImpl,
		})
	case "local":
		status, msg, err, user := createLocalSession(c)
		c.JSON(status, gin.H{
			"errcode": err,
			"msg":     msg,
			//"data": {
			//	"userinfo": interface{}(user.JsonData),
			//},
			"data": map[string]interface{}{"userinfo": user.JsonData},
		})
	case "wechat":
	case "qq":
	}
}
