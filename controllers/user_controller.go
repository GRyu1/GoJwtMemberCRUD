package Controllers

import (
	"github.com/gin-gonic/gin"
	"goLangJwtPrac/models"
	"goLangJwtPrac/structures"
	"goLangJwtPrac/utils"
	"net/http"
	"strings"
)

func HandlePostInsertUser(c *gin.Context) {
	var user structures.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 형식입니다."})
		return
	}
	err := models.HandlePostInsertUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "내부 에러"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "회원가입 성공"})
}
func HandleGetAllUser(c *gin.Context) {
	users, err := models.HandleGetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "내부 에러"})
		return
	}
	c.JSON(http.StatusOK, users)
}
func HandleGetUser(c *gin.Context) {
	idStr := c.Param("id")
	user, err := models.HandleGetUser(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "내부 에러"})
		return
	}
	c.JSON(http.StatusOK, user)
}
func HandlePatchUser(c *gin.Context) {
	var inputUser structures.User

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "형식에 맞지 않는 요청"})
		return
	}
	idStr := c.Param("id")
	result, err := models.HandlePatchUser(idStr, &inputUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "내부에러"})
		return
	}
	c.JSON(http.StatusOK, result)

}
func HandleDeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	err := models.HandleDeleteUser(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "내부 에러"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "삭제 완료"})
}

func HandleAuthentication(c *gin.Context) {
	var loginReq structures.LoginForm
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 형식입니다."})
		return
	}
	accessToken, err := models.HandleAuthentication(&loginReq)
	if err != nil {
		//사실 경우에 따라 http status code 를 나눠주면 좋겠지만 귀찮다. 프젝 적용시에는 예외처리 해야징
		c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 형식입니다."})
		return
	}
	c.Header("Authorization", "Bearer "+accessToken)
	c.JSON(http.StatusOK, gin.H{"message": "로그인 성공"})
}
func HandleAuthorization(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if len(token) > 7 && strings.HasPrefix(token, "Bearer ") {
		token = token[7:]
	}
	_, err := utils.VerifyAccessToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "유효하지 않은 토큰"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "안녕하세요^^"})
}
