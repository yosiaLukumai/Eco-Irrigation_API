package model

import (
	"TEST_SERVER/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailVerification struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserID    primitive.ObjectID `bson:"userID"`
	Email     string             `bson:"email"`
	Token     string             `bson:"token"`
	Type      string             `bson:"type"`
	Used      bool               `bson:"used"`
	ExpiresAt primitive.DateTime `bson:"expiresat"`
	CreatedAt primitive.DateTime `bson:"createdat"`
}

// EmailVerification

func NewVerificationObject(user User, randomData string) EmailVerification {
	return EmailVerification{
		ID:        primitive.NewObjectID(),
		UserID:    user.ID,
		Email:     user.Email,
		Used:      false,
		Type:      "user",
		Token:     randomData,
		ExpiresAt: primitive.NewDateTimeFromTime(time.Now().Add(time.Hour * 24 * 24).Local()),
		CreatedAt: utils.TimeLocal(),
	}

}

func NewVerificationObjectClient(client Client, randomData string) EmailVerification {
	return EmailVerification{
		ID:        primitive.NewObjectID(),
		UserID:    client.ID,
		Email:     client.Email,
		Used:      false,
		Type:      "client",
		Token:     randomData,
		ExpiresAt: primitive.NewDateTimeFromTime(time.Now().Add(time.Hour * 24 * 24).Local()),
		CreatedAt: utils.TimeLocal(),
	}

}
