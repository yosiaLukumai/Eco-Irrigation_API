package handlers

import (
	"TEST_SERVER/database"
	"TEST_SERVER/helpers"
	"TEST_SERVER/model"
	"TEST_SERVER/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterClient(w http.ResponseWriter, r *http.Request) {
	var ClientData struct {
		FirstName string `json:"firstname" validate:"required"`
		LastName  string `json:"lastname" validate:"required"`
		Email     string `json:"email" validate:"required"`
		Phone     string `json:"phone" validate:"required"`
		PumpID    string `json:"pumpid" validate:"required"`
		Packae    string `json:"package" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&ClientData); err != nil {
		utils.CreateOutput(w, fmt.Errorf("JSON decoding error"), false, nil)
		return
	}

	client, err := database.FindEmailClient(ClientData.Email)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("failed to check the email"), false, nil)
		return
	}
	if client.Email == ClientData.Email {
		utils.CreateOutput(w, fmt.Errorf("sorry email taken"), false, nil)
		return
	}

	newClient := model.CreateNewFarmer(ClientData.Email, ClientData.Packae, ClientData.FirstName, ClientData.LastName, ClientData.Phone)
	randomData, err := utils.GenerateRandomStr32(32)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("FAILED to send you verification Email visit recovery center"), false, nil)
		return
	}
	verificationDetails := model.NewVerificationObjectClient(newClient.ID, ClientData.Email, randomData)
	_, err = database.SaveVerification(verificationDetails)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("FAILED to save you verification Email visit clients recovery center"), false, nil)
		return
	}
	// save the client
	_, err = database.InsertOne(database.Farmers, newClient)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("failed to insert the client"), false, nil)
		return
	}

	//update the field of the meter being assigned....
	filter := bson.M{"_id": utils.IDHex(ClientData.PumpID)}
	update := bson.M{"$set": bson.M{"assigned": true, "farmer": newClient.ID}}
	succes, err := database.UpdateOne(database.Pumps, filter, update)
	if err != nil || !succes {
		utils.CreateOutput(w, fmt.Errorf("FAILED to update the Meter assignement Go recovery page"), false, nil)
		return
	}
	// sending the client verification Email
	err, success := helpers.SendEmailVerificationClient(utils.VerificationEmailDataTemplate{
		AppName:    os.Getenv("SENDER_APP"),
		VerifyLink: fmt.Sprintf("%s/verify/client/%s", os.Getenv("UI_URL"), randomData),
		Name:       ClientData.FirstName + "-" + ClientData.LastName,
		Year:       utils.Year(),
	}, ClientData.Email)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("failed to save you verification Email visit recovery center"), false, nil)
		return
	}
	if !success {
		utils.CreateOutput(w, fmt.Errorf("failed to save you verification Email visit recovery center"), false, nil)
		return
	}

	// Meter model has to be modified as assigned...

	utils.CreateOutput(w, fmt.Errorf(""), true, nil)
}

// {

// 	// save that the user is now not verified
// 	succes, err := database.UpdateOne(database.User, user.ID)
// 	if err != nil || !succes {
// 		utils.CreateOutput(w, fmt.Errorf("FAILED to save you verification Email visit recovery center"), false, nil)
// 		return
// 	}
// 	err, success := helpers.SendRecoverEmail(utils.VerificationEmailDataTemplate{
// 		AppName:    "AFM Technologies",
// 		VerifyLink: fmt.Sprintf("%s/recover/account/%s", os.Getenv("UI_URL"), randomData),
// 		Name:       user.FirstName + "-" + user.LastName,
// 		Year:       time.Now().Year(),
// 	}, user.Email)
// 	if err != nil {
// 		utils.CreateOutput(w, fmt.Errorf("FAILED to save you verification Email visit recovery center"), false, nil)
// 		return
// 	}
// 	if !success {
// 		utils.CreateOutput(w, fmt.Errorf("FAILED to save you verification Email visit recovery center"), false, nil)
// 		return
// 	}
// 	utils.CreateOutput(w, fmt.Errorf(""), true, nil)
// }

func FindFarmers(w http.ResponseWriter, r *http.Request) {
	var tableOptions struct {
		RowsPerPage int16 `json:"rowperpage"   validate:"required"`
		CurrentPage int16 `json:"currentpage"   validate:"required"`
		Initial     bool  `json:"initial"`
	}

	if err := json.NewDecoder(r.Body).Decode(&tableOptions); err != nil {
		utils.CreateOutput(w, fmt.Errorf("JSON decoding error"), false, nil)
		return
	}

	msg, err := utils.ValidateIncoming(tableOptions)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf(msg), false, nil)
		return
	}

	skipValue := utils.SkipValue(tableOptions.CurrentPage, tableOptions.RowsPerPage)

	lookup := utils.MD("$lookup", utils.MDs(
		utils.ME("from", "pumps"),
		utils.ME("localField", "_id"),
		utils.ME("foreignField", "farmer"),
		utils.ME("as", "pumpD"),
	))

	lookUp2 := utils.MD("$lookup", utils.MDs(
		utils.ME("from", "package"),
		utils.ME("localField", "package"),
		utils.ME("foreignField", "_id"),
		utils.ME("as", "packageInfo"),
	))

	unwind := utils.MD("$unwind", utils.MDs(utils.ME("path", "$pumpD")))
	unwind2 := utils.MD("$unwind", utils.MDs(utils.ME("path", "$packageInfo")))
	projection := utils.MD("$project", utils.MDs(
		utils.ME("pumpID", "$pumpD._id"),
		utils.ME("balance", "$pumpD.balance"),
		utils.ME("email", 1),
		utils.ME("phone", 1),
		utils.ME("package", "$packageInfo.name"),
	))

	pipelineClient := mongo.Pipeline{
		utils.FacetCreatorMain(lookup, unwind, lookUp2, unwind2, projection, utils.MD("$limit", tableOptions.RowsPerPage), utils.MD("$skip", skipValue)),
	}
	data, err := database.FindCollArrayTableMain(database.Farmers, pipelineClient, tableOptions.Initial)
	if err != nil {
		fmt.Println(err)
		utils.CreateOutput(w, fmt.Errorf(" can't find companie's meter"), false, nil)
		return
	}
	utils.CreateOutput(w, fmt.Errorf(""), true, data)
}
