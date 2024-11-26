package services

import (
	"fmt"
	"net/http"
	"os"
)

func SendOTPWhatsApp(phoneNumber, otp string) error {
	apiKey := os.Getenv("WHATSAPP_API_KEY")
	url := fmt.Sprintf("https://api.whatsapp.com/send?phone=%s&text=Your OTP is: %s", phoneNumber, otp)

	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send OTP")
	}

	return nil
}
