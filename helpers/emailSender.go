package helpers

import (
	"TEST_SERVER/utils"
	"log"
	"os"
)

func SendEmailVerification(data utils.VerificationEmailDataTemplate, emailAdress string) (error, bool) {
	content, err := utils.ReadFileToString("./assets/verifyEmail.html")
	if err != nil {
		log.Fatal(" Failed to get html file from assets folder to string")
		return err, false
	}

	verificationEmailHtml, err := utils.ParseHtmlVariables(content, data)
	if err != nil {
		log.Fatal(" Failed to get the verification emal")
		return err, false
	}

	sender := utils.GmailSender{Address: os.Getenv("SENDER_EMAIL"), Password: os.Getenv("APP_PASSWORD"), AppName: os.Getenv("SENDER_APP")}
	sender.SendEmail("Verify your email.", string(verificationEmailHtml), []string{emailAdress}, nil, nil, nil)
	if err != nil {
		log.Fatal("Error:  ", err)
		return err, false
	}

	return nil, true
}

func SendEmailVerificationClient(data utils.VerificationEmailDataTemplate, emailAdress string) (error, bool) {
	content, err := utils.ReadFileToString("./assets/customerVerification.html")
	if err != nil {
		log.Fatal(" Failed to get html file from assets folder to string")
		return err, false
	}

	verificationEmailHtml, err := utils.ParseHtmlVariables(content, data)
	if err != nil {
		log.Fatal(" Failed to get the verification emal")
		return err, false
	}

	sender := utils.GmailSender{Address: os.Getenv("SENDER_EMAIL"), Password: os.Getenv("APP_PASSWORD"), AppName: os.Getenv("SENDER_APP")}
	sender.SendEmail("Verify your email.", string(verificationEmailHtml), []string{emailAdress}, nil, nil, nil)
	if err != nil {
		log.Fatal("Error:  ", err)
		return err, false
	}

	return nil, true
}


func SendRecoverEmail(data utils.VerificationEmailDataTemplate, emailAdress string) (error, bool) {
	content, err := utils.ReadFileToString("./assets/recoveryEmail.html")
	if err != nil {
		log.Fatal(" Failed to get html file from assets folder to string")
		return err, false
	}

	verificationEmailHtml, err := utils.ParseHtmlVariables(content, data)
	if err != nil {
		log.Fatal(" Failed to get the Recovery Email")
		return err, false
	}

	sender := utils.GmailSender{Address: os.Getenv("SENDER_EMAIL"), Password: os.Getenv("APP_PASSWORD"), AppName: os.Getenv("SENDER_APP")}
	sender.SendEmail("Reset your Account.", string(verificationEmailHtml), []string{emailAdress}, nil, nil, nil)
	if err != nil {
		log.Fatal("Error:  ", err)
		return err, false
	}

	return nil, true
}

func SendRecoverEmailClient(data utils.VerificationEmailDataTemplate, emailAdress string) (error, bool) {
	content, err := utils.ReadFileToString("./assets/recoverClient.html")
	if err != nil {
		log.Fatal(" Failed to get html file from assets folder to string")
		return err, false
	}

	verificationEmailHtml, err := utils.ParseHtmlVariables(content, data)
	if err != nil {
		log.Fatal(" Failed to get the Recovery Email")
		return err, false
	}

	sender := utils.GmailSender{Address: os.Getenv("SENDER_EMAIL"), Password: os.Getenv("APP_PASSWORD"), AppName: os.Getenv("SENDER_APP")}
	sender.SendEmail("Reset your Account.", string(verificationEmailHtml), []string{emailAdress}, nil, nil, nil)
	if err != nil {
		log.Fatal("Error:  ", err)
		return err, false
	}

	return nil, true
}
