package config

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func GetDatabaseURL() string {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
}

func SendVerificationEmail(email, token string) error {
	from := "hanggoroseto8@gmail.com"
	password := "pcxf rviq wvfz nfyy"
	to := []string{email}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	verificationLink := fmt.Sprintf("https://cleanconnect-app-336q.vercel.app//verify/%s", token)
	message := []byte("Subject: Email Verification\n\n" +
		"Please click the following link to verify your email address:\n" +
		verificationLink)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}

	fmt.Println("Verification email sent to:", email)
	return nil
}

func SendNotification(email string, billDescription string, dueDate time.Time) error {
	from := "hanggoroseto8@gmail.com"
	password := "pcxf rviq wvfz nfyy"
	to := []string{email}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte("Subject: Tagihan Reminder\n\n"+billDescription+" is due on "+dueDate.Format("2006-01-02")))
	if err != nil {
		return err
	}

	fmt.Println("Notification email sent to:", email)
	return nil
}

func GenerateVerificationToken() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func UploadToCloudinary(file *multipart.FileHeader) (string, error) {
	cld, _ := cloudinary.NewFromURL("cloudinary://633714464826515:u1W6hqq-Gb8y-SMpXe7tzs4mH44@dvrhf8d9t")

	f, _ := file.Open()
	ctx := context.Background()
	uploadResult, _ := cld.Upload.Upload(ctx, f, uploader.UploadParams{})

	return uploadResult.SecureURL, nil
}

func GenerateMayarQRCode(amount int) (string, error) {
	client := &http.Client{}
	url := "https://api.mayar.id/hl/v1/qrcode/create"

	payload := map[string]interface{}{
		"amount": amount,
	}
	payloadBytes, _ := json.Marshal(payload)

	// Buat request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	// Set header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIyMjhlODczYi0wYjM1LTQwZGUtOWEwNC05NTM3ZjY1YzcxMTUiLCJhY2NvdW50SWQiOiJlNDdiMzJmOS0wM2FiLTQwNGMtOGI0NC03OWIyNmRmNzRmMDciLCJjcmVhdGVkQXQiOiIxNzMzOTM4MDQ1OTIwIiwicm9sZSI6ImRldmVsb3BlciIsInN1YiI6Imhhbmdnb3Jvc2V0bzhAZ21haWwuY29tIiwibmFtZSI6IkNsZWFuIENvbm5lY3QiLCJsaW5rIjoiaGFuZ2dvcm8tc2V0byIsImlzU2VsZkRvbWFpbiI6bnVsbCwiaWF0IjoxNzMzOTM4MDQ1fQ.Nwx65ZKh9P7wllxKAu1NetDLDC5-9jXIe5reuCMmSW6QfJ9EnVWRwLM3dN1pJEAZTFWDpgRUu-D0h_4tNZT7caHLnlVgeqRCpmDfGMjRr-yEponNGwXZZBtGPF4rt3uvcl9u9gYJ5nkakUO4PMNhyc8SA_6koaebapaZNeZRJGH8jALryTrckyezaiibf68ZvOev0VsmrjMxG-K0Ro4YBFMTsw1pCbni1XxVxI0473FxELG54qpiNSOi8vLCowQe3hEYyX3k_XQlzj83Xta2mKj8jGfFU-GylV7RM_GLzqBx6SFiP5-IShUg_aKEa8dOi_XVuT5yue6XXMMh9QHHZA")

	// Kirim request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Parsing response
	var response struct {
		StatusCode int    `json:"statusCode"`
		Messages   string `json:"messages"`
		Data       struct {
			URL string `json:"url"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	if response.StatusCode != 200 {
		return "", errors.New(response.Messages)
	}

	return response.Data.URL, nil
}
