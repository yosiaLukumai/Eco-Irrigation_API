package handlers

import (
	"TEST_SERVER/database"
	"TEST_SERVER/model"
	"TEST_SERVER/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Response struct {
	Error   string
	Success bool
	Data    interface{}
}

func AddRoleCompany(w http.ResponseWriter, r *http.Request) {
	var CompanyCredentials struct {
		CompanyId string   `json:"companyId"  validate:"required"`
		RoleName  string   `json:"name"   validate:"required"`
		Roles     []string `json:"roles"   validate:"required"`
		RoleDesc  string   `json:"desc"   validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&CompanyCredentials); err != nil {
		utils.CreateOutput(w, fmt.Errorf("JSON decoding error"), false, nil)
		return
	}
	msg, err := utils.ValidateIncoming(CompanyCredentials)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf(msg), false, nil)
		return
	}

	comp, err := database.FindCompany(utils.IDHex(CompanyCredentials.CompanyId))
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf(" There is no such company"), false, nil)
		return
	}
	// create a the role
	newRole := model.Roles{
		ID:          primitive.NewObjectID(),
		Company:     comp.ID,
		Name:        CompanyCredentials.RoleName,
		Description: CompanyCredentials.RoleDesc,
		Access:      CompanyCredentials.Roles,
		CreatedAt:   utils.TimeLocal(),
		UpdatedAt:   utils.TimeLocal(),
	}

	_, err = database.InsertNewRole(newRole)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf(" Failed to insert the role"), false, nil)
		return
	}

	utils.CreateOutput(w, fmt.Errorf(""), true, nil)
}

func AddPump(w http.ResponseWriter, r *http.Request) {
	var PumpDetails struct {
		// FarmerID  string  `json:"farmerID"  validate:"required"`
		Head      float64 `json:"head" validate:"required"`
		Discharge float64 `json:"discharge" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&PumpDetails); err != nil {
		utils.CreateOutput(w, fmt.Errorf("JSON decoding error"), false, nil)
		return
	}

	msg, err := utils.ValidateIncoming(PumpDetails)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf(msg), false, nil)
		return
	}

	data := model.CreateNewPump(PumpDetails.Discharge, PumpDetails.Head)
	_, err = database.InsertOne(database.Pumps, data)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf("failed to insert the system"), false, nil)
		return
	}

	utils.CreateOutput(w, fmt.Errorf(""), true, nil)
}

func FindSystems(w http.ResponseWriter, r *http.Request) {
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
		utils.ME("from", "farmers"),
		utils.ME("localField", "farmer"),
		utils.ME("foreignField", "_id"),
		utils.ME("as", "farmerInfo"),
	))
	projection := utils.MD("$project", utils.MDs(
		utils.ME("farmer", "$farmerInfo.email"),
		utils.ME("head", 1),
		utils.ME("discharge", 1),
		utils.ME("status", utils.MDs(
			utils.ME("$cond", utils.MA("$status", "active", "not-active")),
		)),
		utils.ME("assigned", utils.MDs(
			utils.ME("$cond", utils.MA("$assigned", "assigned", "not-assigned")),
		)),
	))

	pipelineSystems := mongo.Pipeline{
		utils.FacetCreatorMain(lookup, projection, utils.MD("$limit", tableOptions.RowsPerPage), utils.MD("$skip", skipValue)),
	}
	data, err := database.FindCollArrayTableMain(database.Pumps, pipelineSystems, tableOptions.Initial)
	if err != nil {
		fmt.Println(err)
		utils.CreateOutput(w, fmt.Errorf(" can't find companie's systems"), false, nil)
		return
	}

	utils.CreateOutput(w, fmt.Errorf(""), true, data)
}

func FindUnassigned(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{{"$match", bson.D{{"assigned", false}}}}
	sort := bson.D{{"$sort", bson.D{{"createdat", -1}}}}
	projection := bson.D{
		{"$project",
			bson.D{
				{"_id", 1}}}}
	limit := bson.D{{"$limit", 3}}

	filterA := utils.AggregationFilter(sort, filter, projection, limit)
	result, err := database.FindCollReturnArray(database.Pumps, filterA)
	if err != nil {
		utils.CreateOutput(w, fmt.Errorf(" can't find Non assigned systems"), false, nil)
		return
	}
	utils.CreateOutput(w, fmt.Errorf(""), true, result)

}

func PaymentCallBack(w http.ResponseWriter, r *http.Request) {
	// logging the callBack for success of the payment..
	var Credentials interface{}
	// fmt.Println(json.NewDecoder(r.Body).Decode(&Credentials))
	err := json.NewDecoder(r.Body).Decode(&Credentials)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		utils.CreateOutput(w, fmt.Errorf(""), true, err)
		return
	}

	utils.CreateOutput(w, fmt.Errorf(""), true, "")
}
