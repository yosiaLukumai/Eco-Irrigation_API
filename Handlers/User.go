package handlers

import (
	"TEST_SERVER/database"
	"TEST_SERVER/helpers"
	"os"

	// "TEST_SERVER/helpers"
	"TEST_SERVER/model"
	"TEST_SERVER/utils"
	"encoding/json"
	"fmt"
	"net/http"

	// "os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	// "github.com/go-playground/validator/v10"
)

type LoginStruct struct {
	User    model.User
	Company model.Company
}

func SignIn(w http.ResponseWriter, r *http.Request) {

	var UserCredentials struct {
		Email    string `json:"email" bson:"email"  validate:"required,email"`
		Password string `json:"password" bson:"password"  validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&UserCredentials); err != nil {
		utils.CreateOutput(w, fmt.Errorf("JSON decoding error"), false, nil)
		return
	}
	msg, err := utils.ValidateIncoming(UserCredentials)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf(msg), false, nil)
		return
	}

	user, err := database.FindEmail(UserCredentials.Email)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("no such user the email"), false, nil)
		return
	}

	if !(user.Verified) {
		utils.CreateOutput(w, fmt.Errorf("email not verified"), false, nil)
		return
	}
	if !(user.Active) {
		utils.CreateOutput(w, fmt.Errorf("account disabled"), false, nil)
		return
	}

	if !(utils.ComparePassword(user.Password, UserCredentials.Password)) {
		utils.CreateOutput(w, fmt.Errorf("in-correct Password"), false, nil)
		return
	}

	utils.CreateOutput(w, fmt.Errorf(""), true, user)

}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var signUp struct {
		FirstName string `json:"firstname" validate:"required"`
		LastName  string `json:"lastname" validate:"required"`
		Email     string `json:"email" validate:"required"`
		Password  string `json:"password"  validate:"required"`
		Phone     string `json:"phone" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&signUp); err != nil {
		utils.CreateOutput(w, fmt.Errorf("json decoding error"), false, nil)
		return
	}
	msg, err := utils.ValidateIncoming(signUp)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf(msg), false, nil)
		return
	}

	// checking the email exists in the users's model...
	user, err := database.Find(database.User, "email", signUp.Email)
	if (err == nil) && (user != nil) {
		// how to pass the data of BSON.D to object...
		utils.CreateOutput(w, fmt.Errorf("email  taken"), false, nil)
		return
	}

	// let hash the password
	hashedPassword, err := utils.HashPassword(signUp.Password)
	if err != nil {
		response := Response{Error: err.Error(), Success: false, Data: "failed to hash the password"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&response)
		return
	}
	signUpD := model.CreateAdmin(hashedPassword, signUp.Email, signUp.FirstName, signUp.LastName, signUp.Phone)
	fmt.Println(signUpD)
	_, err = database.InsertUser(signUpD)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("failed to insert user"), false, nil)
		return
	}

	utils.CreateOutput(w, fmt.Errorf(""), true, nil)
}

// Password forgetting or having trouble logging in
// func RecoverAcc(w http.ResponseWriter, r *http.Request) {
// 	var UserCredentials struct {
// 		Email string `json:"email" validate:"required,email"`
// 		Type  string `json:"type" validate:"required"`
// 	}
// 	if err := json.NewDecoder(r.Body).Decode(&UserCredentials); err != nil {
// 		utils.CreateOutput(w, fmt.Errorf("JSON decoding error"), false, nil)
// 		return
// 	}
// 	msg, err := utils.ValidateIncoming(UserCredentials)
// 	if err != nil {
// 		utils.CreateOutput(w, fmt.Errorf(msg), false, nil)
// 		return
// 	}
// 	// var newData model.User
// 	if UserCredentials.Type == "user" {
// 		user, err := database.FindEmail(UserCredentials.Email)
// 		if err != nil {
// 			utils.CreateOutput(w, fmt.Errorf("sorry you don't have an account"), false, nil)
// 			return
// 		}

// 		randomData, err := utils.GenerateRandomStr32(32)
// 		if err != nil {
// 			utils.CreateOutput(w, fmt.Errorf("FAILED to send you verification Email visit recovery center"), false, nil)
// 			return
// 		}

// 		verificationDetails := model.NewVerificationObject(user, randomData)
// 		_, err = database.SaveVerification(verificationDetails)
// 		if err != nil {
// 			utils.CreateOutput(w, fmt.Errorf("FAILED to save you verification Email visit recovery center"), false, nil)
// 			return
// 		}

// 		// save that the user is now not verified
// 		filter := bson.M{"_id": user.ID}
// 		update := bson.M{"$set": bson.M{"verified": false}}
// 		succes, err := database.UpdateOne(database.User, filter, update)
// 		if err != nil || !succes {
// 			utils.CreateOutput(w, fmt.Errorf("FAILED to save you verification Email visit recovery center"), false, nil)
// 			return
// 		}
// 		err, success := helpers.SendRecoverEmail(utils.VerificationEmailDataTemplate{
// 			AppName:    "AFM Technologies",
// 			VerifyLink: fmt.Sprintf("%s/recover/account/%s", os.Getenv("UI_URL"), randomData),
// 			Name:       user.FirstName + "-" + user.LastName,
// 			Year:       time.Now().Year(),
// 		}, user.Email)
// 		if err != nil {
// 			utils.CreateOutput(w, fmt.Errorf("FAILED to save you verification Email visit recovery center"), false, nil)
// 			return
// 		}
// 		if !success {
// 			utils.CreateOutput(w, fmt.Errorf("FAILED to save you verification Email visit recovery center"), false, nil)
// 			return
// 		}
// 		utils.CreateOutput(w, fmt.Errorf(""), true, nil)
// 	} else {
// 		client, err := database.FindEmailClient(UserCredentials.Email)
// 		if err != nil {
// 			utils.CreateOutput(w, fmt.Errorf("sorry you don't have an account"), false, nil)
// 			return
// 		}

// 		randomData, err := utils.GenerateRandomStr32(32)
// 		if err != nil {
// 			utils.CreateOutput(w, fmt.Errorf("FAILED to send you verification Email visit recovery center"), false, nil)
// 			return
// 		}

// 		verificationDetails := model.NewVerificationObjectClient(client, randomData)
// 		_, err = database.SaveVerification(verificationDetails)
// 		if err != nil {
// 			utils.CreateOutput(w, fmt.Errorf("FAILED to save you verification Email visit recovery center"), false, nil)
// 			return
// 		}

// 		// save that the user is now not verified

// 		filter := bson.M{"_id": client.ID}
// 		update := bson.M{"$set": bson.M{"verified": false}}
// 		succes, err := database.UpdateOne(database.Client, filter, update)
// 		if err != nil || !succes {
// 			utils.CreateOutput(w, fmt.Errorf("FAILED to save you verification Email visit recovery center"), false, nil)
// 			return
// 		}
// 		err, success := helpers.SendRecoverEmailClient(utils.VerificationEmailDataTemplate{
// 			AppName:    "AFM Technologies",
// 			VerifyLink: fmt.Sprintf("%s/recover/account/client/%s", os.Getenv("UI_URL"), randomData),
// 			Name:       client.Name,
// 			Year:       time.Now().Year(),
// 		}, client.Email)
// 		if err != nil {
// 			utils.CreateOutput(w, fmt.Errorf("FAILED to save you verification Email visit recovery center"), false, nil)
// 			return
// 		}
// 		if !success {
// 			utils.CreateOutput(w, fmt.Errorf("FAILED to save you verification Email visit recovery center"), false, nil)
// 			return
// 		}
// 		utils.CreateOutput(w, fmt.Errorf(""), true, nil)
// 	}
// }

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var UserCredentials struct {
		FirstName string `json:"firstname" validate:"required,firstname"`
		LastName  string `json:"lastname" validate:"required,lastname"`
		Email     string `json:"email" validate:"required,email"`
		Station   string `json:"station" validate:"required,station"`
		RoleName  string `json:"rolename" validate:"required, rolename"`
		Phone     string `json:"phone" validate:"required, phone"`
	}
	if err := json.NewDecoder(r.Body).Decode(&UserCredentials); err != nil {
		utils.CreateOutput(w, fmt.Errorf("JSON decoding error"), false, nil)
		return
	}
	msg, err := utils.ValidateIncoming(UserCredentials)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf(msg), false, nil)
		return
	}
	// var newData model.User
	user, err := database.FindEmail(UserCredentials.Email)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("sorry you don't have an account"), false, nil)
		return
	}
	if user.Email == UserCredentials.Email {
		utils.CreateOutput(w, fmt.Errorf("sorry email taken"), false, nil)
		return
	}

	//
	randomData, err := utils.GenerateRandomStr32(32)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("FAILED to send you verification Email visit recovery center"), false, nil)
		return
	}

	verificationDetails := model.NewVerificationObject(user, randomData)

	_, err = database.SaveVerification(verificationDetails)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("FAILED to save you verification Email visit recovery center"), false, nil)
		return
	}

	// save that the user is now not verified
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{"verified": false}}
	succes, err := database.UpdateOne(database.User, filter, update)
	if err != nil || !succes {
		utils.CreateOutput(w, fmt.Errorf("FAILED to save you verification Email visit recovery center"), false, nil)
		return
	}
	err, success := helpers.SendRecoverEmail(utils.VerificationEmailDataTemplate{
		AppName:    "AFM Technologies",
		VerifyLink: fmt.Sprintf("%s/recover/account/%s", os.Getenv("UI_URL"), randomData),
		Name:       user.FirstName + "-" + user.LastName,
		Year:       time.Now().Year(),
	}, user.Email)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("FAILED to save you verification Email visit recovery center"), false, nil)
		return
	}
	if !success {
		utils.CreateOutput(w, fmt.Errorf("FAILED to save you verification Email visit recovery center"), false, nil)
		return
	}
	utils.CreateOutput(w, fmt.Errorf(""), true, nil)
}

func VerifyEmail(w http.ResponseWriter, r *http.Request) {

	var UserCredentials struct {
		EmailKey string `json:"emailKey"  validate:"required"`
		Password string `json:"password"   validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&UserCredentials); err != nil {
		utils.CreateOutput(w, fmt.Errorf("JSON decoding error"), false, nil)
		return
	}
	msg, err := utils.ValidateIncoming(UserCredentials)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf(msg), false, nil)
		return
	}
	user, err := database.FindKey(UserCredentials.EmailKey)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("no such user"), false, nil)
		return
	}
	// check if the Token has expired...

	expiredTime := time.Time(user.ExpiresAt.Time())
	if time.Now().After(expiredTime) {
		utils.CreateOutput(w, fmt.Errorf("token expired recover your account"), false, nil)
		return
	}

	if user.Used {
		utils.CreateOutput(w, fmt.Errorf("token already used"), false, nil)
		return
	}

	// update the password first lets hash password
	hashedPassword, err := utils.HashPassword(UserCredentials.Password)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("failed to hash password"), false, nil)
		return
	}

	if user.Type == "user" {
		_, err = database.UpdatePassword(user.UserID, hashedPassword)
		if err != nil {
			utils.CreateOutput(w, fmt.Errorf("failed to hash password"), false, nil)
			return
		}
	} else {
	_, err = database.UpdatePasswordClients(user.UserID, hashedPassword)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("failed to hash password"), false, nil)
		return
	}
	}
	// update the key object
	_, err = database.UpdateVerification(user.ID)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("failed to update verification"), false, nil)
		return
	}
	utils.CreateOutput(w, fmt.Errorf(""), true, nil)
}

func VerifyAdmin(w http.ResponseWriter, r *http.Request) {

	// check the token from the url param
	token := utils.RouteParam(r, "token")
	if token == "" {
		utils.CreateOutput(w, fmt.Errorf(" Can't verify empty credentials"), false, nil)
		return
	}

	emV, err := database.FindKey(token)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("no such token issued"), false, nil)
		return
	}

	if emV.Used {
		utils.CreateOutput(w, fmt.Errorf("token already used"), false, nil)
		return
	}

	// checking if the expired time has passed...
	expiredTime := time.Time(emV.ExpiresAt.Time())
	if time.Now().After(expiredTime) {
		utils.CreateOutput(w, fmt.Errorf("token expired recover your account"), false, nil)
		return
	}

	// update the flag
	success, err := database.Verify(emV.Email)
	if err != nil || !success {
		utils.CreateOutput(w, fmt.Errorf("failed to verify"), false, nil)
		return
	}

	_, err = database.UpdateVerification(emV.ID)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("failed to update verification"), false, nil)
		return
	}

	utils.CreateOutput(w, fmt.Errorf(""), true, nil)
}
