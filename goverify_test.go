package goverify_test

import (
	"log"
	"os"
	"testing"

	"github.com/herla97/goverify"
	"github.com/joho/godotenv"
)

func init() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func TestVerify_NewVerify(t *testing.T) {
	client := goverify.NewVerify(
		os.Getenv("ACCOUNT_SID"),
		os.Getenv("AUTH_TOKEN"),
		os.Getenv("TWILIO_SERVICE"),
	)

	if client == nil {
		t.Error("error got: nil, want *goverify.TwilioClient")
	}
}

func TestVerify_Send(t *testing.T) {
	client := goverify.NewVerify(
		os.Getenv("ACCOUNT_SID"),
		os.Getenv("AUTH_TOKEN"),
		os.Getenv("TWILIO_SERVICE"),
	)

	input := &goverify.VerifyInput{
		To:      "+521234567890", // Random number, don't call :), use your own number
		Channel: "sms",
	}

	resp, err := client.VerifySend(input)
	if err != nil {
		t.Errorf("twilio error, got: %v", err)
	}

	if resp.Sid == "" {
		t.Error("error got: '', want: VEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	}
}

func TestVerify_Check(t *testing.T) {
	client := goverify.NewVerify(
		os.Getenv("ACCOUNT_SID"),
		os.Getenv("AUTH_TOKEN"),
		os.Getenv("TWILIO_SERVICE"),
	)

	input := &goverify.VerifyInput{
		To:   "+521234567890", // Random number, don't call :), use your own number
		Code: "308870",
	}

	resp, err := client.VerifyCheck(input)
	if err != nil {
		t.Errorf("twilio error, got: %v", err)
	}

	if !resp.Valid {
		t.Error("error invalid code")
	}
}
