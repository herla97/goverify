# goverify

## Golang Twilio Verify API Integration

### How to use?

#### VerifySendSMS
```go
	goverifyClient := goverify.NewVerify("ACCOUNT_SID", "AUTH_TOKEN", "TWILIO_SERVICE")

	resp, excep, err := goverifyClient.VerifySendSMS("+15017122661")
	if err != nil {
		log.Fatal(err)
	}
	if excep != nil {
		log.Println(excep)
	}

	log.Println(resp)
```

#### VerifyCheckSMS
```go
	goverifyClient := goverify.NewVerify("ACCOUNT_SID", "AUTH_TOKEN", "TWILIO_SERVICE")

	resp, excep, err := goverifyClient.VerifyCheckSMS("+15017122661", "6082")
	if err != nil {
		log.Fatal(err)
	}
	if excep != nil {
		log.Println(excep)
	}

	log.Println(resp)
```

