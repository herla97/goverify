# goverify

## Golang Twilio Verify API Integration

### How to use?
It is necessary to have a Twilio account and an active Verify Service.
See more: https://www.twilio.com/docs/verify

#### VerifySend
```go
	client := goverify.NewVerify(
		"ACCOUNT_SID",
		"AUTH_TOKEN",
		"TWILIO_SERVICE",
	)

	input := &goverify.VerifyInput{
		To:      "+521234567890", // Random number, don't call :), use your own number
		Channel: "sms",
	}

	resp, err := client.VerifySend(input)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
```

#### VerifyCheck
```go
	client := goverify.NewVerify(
		"ACCOUNT_SID",
		"AUTH_TOKEN",
		"TWILIO_SERVICE",
	)

	input := &goverify.VerifyInput{
		To:      "+521234567890", // Random number, don't call :), use your own number
		Code: "308870",
	}

	resp, err := client.VerifyCheck(input)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
```

