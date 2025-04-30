// controllers/twofa.go
package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	twilio "github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
	"github.com/waltertaya/SMS-Verification/db"
	"github.com/waltertaya/SMS-Verification/models"
)

func CodeRequest(ctx *gin.Context) {
	var body struct {
		Phone string `json:"phone"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	client := twilio.NewRestClient()
	params := &verify.CreateVerificationParams{}
	params.SetTo(body.Phone)
	params.SetChannel("sms")

	serviceSid := os.Getenv("TWILIO_VERIFY_SERVICE_SID")
	resp, err := client.VerifyV2.CreateVerification(serviceSid, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Verification code sent successfully",
		"sid": *resp.Sid,
	})
}

func VerifyCode(ctx *gin.Context) {
	var body struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
		ID    string `json:"id"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	client := twilio.NewRestClient()
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(body.Phone)
	params.SetCode(body.Code)

	serviceSid := os.Getenv("TWILIO_VERIFY_SERVICE_SID")
	resp, err := client.VerifyV2.CreateVerificationCheck(serviceSid, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// update the verified user in the database
	db.DB.Model(&models.User{}).Where("id = ?", body.ID).Updates(models.User{Verified: true})

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Verification code verified successfully",
		"sid": *resp.Sid,
	})
}
