package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	pushnotifications "github.com/pusher/push-notifications-go"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	instanceId := os.Getenv("INSTANCE_ID")
	secretKey := os.Getenv("SECRET_KEY")

	// Create a new Beams client
	beamsClient, _ := pushnotifications.New(instanceId, secretKey)

	// Handle auth requests
	http.HandleFunc("/pusher/beams-auth", func(w http.ResponseWriter, r *http.Request) {
		// Do your normal auth checks here 🔒
		// TODO: Replace this with your client console.log device-id
		userID := "web-f6edbbbf-68ba-4ee3-8129-43df3e7addc7" // get it from your client register service worker console.log(PusherPushNotifications.token)
		userIDinQueryParam := r.URL.Query().Get("user_id")
		if userID != userIDinQueryParam {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		beamsToken, err := beamsClient.GenerateToken(userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		beamsTokenJson, err := json.Marshal(beamsToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(beamsTokenJson)
	})

	// Start the server
	http.ListenAndServe(":8080", nil)
}
