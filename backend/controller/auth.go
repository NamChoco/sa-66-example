package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "github.com/go-delve/delve/service"
	"github.com/NamChoco/sa-66-example/entity"
	"github.com/NamChoco/sa-66-example/service"
	"golang.org/x/crypto/bcrypt"
)


//Login payload body
type LoginPayload struct {
	Email string    `json:"email"`
	Password string `json:"password"`
}

//Login Response body
type LoginResponse struct {
	Token string `json:"token"`
	ID 	  uint   `json:"id"`
}


// ทำการรับข้อมูลจากผู้ใช้ email and password
// ตรวจสอบข้อมูลว่าถูกต้องหรือไม่
// response token ให้ผู้ใช้
func Login(c *gin.Context){
	var payload  LoginPayload
	var user     entity.User
	
	if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	if err := entity.DB().Raw("SELECT * FROM users WHERE email = ?", payload.Email).Scan(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	//ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	jwtWrapper := service.JwtWrapper {
		SecretKey: "ABC",
		Issuer:    "AuthService",
        ExpirationHours: 24,	
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error generating token"})
        return
	}

	tokenResponse := LoginResponse{
		Token: signedToken,
        ID: user.ID,
	}
	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}