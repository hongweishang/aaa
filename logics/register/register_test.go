package register

import (
	"bytes"
	"encoding/json"
	"errcode"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLocalRegister(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	var ctx struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	ctx.Username = "u1"
	ctx.Password = "password1"

	jsonData, _ := json.Marshal(&ctx)
	body := bytes.NewBuffer(jsonData)
	c.Request, _ = http.NewRequest("POST", "", body)

	status, msg, ec := localRegister(c)
	assert.Equal(t, status, http.StatusOK)
	assert.Equal(t, msg, errcode.ErrStrOk)
	assert.Equal(t, errcode.ErrCodeType(ec), errcode.ErrCodeType(errcode.ErrOk))

	c, _ = gin.CreateTestContext(httptest.NewRecorder())
	body = bytes.NewBuffer(jsonData)
	c.Request, _ = http.NewRequest("POST", "", body)
	status, msg, ec = localRegister(c)
	assert.Equal(t, status, http.StatusOK)
	assert.Equal(t, msg, errcode.ErrStrAe)
	assert.Equal(t, errcode.ErrCodeType(ec), errcode.ErrCodeType(errcode.ErrAe))

	ctx.Username = "u2"
	ctx.Password = "password1"

	jsonData, _ = json.Marshal(&ctx)
	body = bytes.NewBuffer(jsonData)
	c.Request, _ = http.NewRequest("POST", "", body)

	status, msg, ec = localRegister(c)
	assert.Equal(t, status, http.StatusOK)
	assert.Equal(t, msg, errcode.ErrStrOk)
	assert.Equal(t, errcode.ErrCodeType(ec), errcode.ErrCodeType(errcode.ErrOk))

}
