package database

import (
	"TEST_SERVER/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func FindMetersTable(compID string, skipValue int16, limit int16) bson.A {
	return bson.A{
		bson.D{{"$match", bson.D{{"company", utils.IDHex(compID)}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "branch"},
					{"localField", "branch"},
					{"foreignField", "_id"},
					{"as", "branchInfo"},
				},
			},
		},
		bson.D{{"$unwind", bson.D{{"path", "$branchInfo"}}}},
		bson.D{
			{"$project",
				bson.D{
					{"branchName", "$branchInfo.name"},
					{"amount", 1},
					{"serialno", 1},
					{"flowrate", 1},
					{"maxpressure", 1},
					{"status",
						bson.D{
							{"$cond",
								bson.A{
									"$status",
									"active",
									"not-active",
								},
							},
						},
					},
					{"assigned",
						bson.D{
							{"$cond",
								bson.A{
									"$assigned",
									"assigned",
									"not-assigned",
								},
							},
						},
					},
				},
			},
		},
		bson.D{{"$skip", skipValue}},
		bson.D{{"$limit", limit}},
	}
}

func FindClientTable(compID string, skipValue int16, limit int16) bson.A {
	return bson.A{}
}

func MeterMatcher(compID string) bson.A {
	return bson.A{bson.D{{"$match", bson.D{{"company", utils.IDHex(compID)}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "branch"},
					{"localField", "branch"},
					{"foreignField", "_id"},
					{"as", "branchInfo"},
				},
			},
		},
		bson.D{{"$unwind", bson.D{{"path", "$branchInfo"}}}},
		bson.D{
			{"$project",
				bson.D{
					{"branchName", "$branchInfo.name"},
					{"amount", 1},
					{"serialno", 1},
					{"flowrate", 1},
					{"maxpressure", 1},
					{"status",
						bson.D{
							{"$cond",
								bson.A{
									"$status",
									"active",
									"not-active",
								},
							},
						},
					},
					{"assigned",
						bson.D{
							{"$cond",
								bson.A{
									"$assigned",
									"assigned",
									"not-assigned",
								},
							},
						},
					},
				},
			},
		},
	}
}
