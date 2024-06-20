package session

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
)

type Session struct {
	Store       *redisstore.RedisStore
	SessionName string
}

var appSession Session

func NewSession(sessionName string) (err error) {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDRESS"),
	})

	store, err := redisstore.NewRedisStore(context.Background(), client)
	if err != nil {
		return err
	}

	store.KeyPrefix(sessionName)
	store.Options(sessions.Options{
		Path:   os.Getenv("REDIS_PATH"),
		Domain: os.Getenv("DOMAIN"),
		MaxAge: 86400 * 60, // a day
	})

	appSession.SessionName = sessionName
	appSession.Store = store

	return err
}

func Get(r *http.Request, key string) (val interface{}) {
	session, err := appSession.Store.Get(r, appSession.SessionName)

	if err != nil {
		log.Fatal("Failed to obtain the session: ", err)
	}
	return session.Values[key]
}

func Set(r *http.Request, w http.ResponseWriter, key string, value interface{}) (ok bool) {
	session, err := appSession.Store.Get(r, appSession.SessionName)

	if err != nil {
		log.Println("EFailed to obtain the session: ", err)
		return false
	}
	session.Values[key] = value
	if err = session.Save(r, w); err != nil {
		log.Println("Error saving value to the session: ", err)
		return false
	}

	return true
}

func Delete(r *http.Request, w http.ResponseWriter, key interface{}) (ok bool) {
	session, err := appSession.Store.Get(r, appSession.SessionName)
	if err != nil {
		log.Println("Failed to Delete the session", err)
		return false
	}

	delete(session.Values, key)

	if err = session.Save(r, w); err != nil {
		log.Fatal("Failed to save session Deletion   ", err)
		return false
	}

	return true
}

func Kill(r *http.Request, w http.ResponseWriter) (ok bool) {
	session, err := appSession.Store.Get(r, appSession.SessionName)
	if err != nil {
		log.Println("Failed to kill the session", err)
		return false
	}

	session.Options.MaxAge = -1

	if err = session.Save(r, w); err != nil {
		log.Println("Failed to save  the session killing", err)
		return false
	}

	return true

}
