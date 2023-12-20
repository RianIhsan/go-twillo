package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error load env")
	}
	from := "+12059315655"
	to := os.Getenv("MY_PHONE_NUMBER")
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	message := "Ini message test : Anjay twillo"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{
		To:   &to,
		From: &from,
		Body: &message,
	}

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
