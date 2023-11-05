package session_test

import (
	"castanha/database"
	"castanha/models/session"
	"testing"
	"time"
)

func TestUpdateDuration(t *testing.T) {
	db := database.GetDB()

	operation := db.Model(&session.Session{}).Where("key_access = ?", "75eec18c-6606-4f39-b0fd-2e49e042f8a4")

	duration := time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")

	err := operation.Update("expires_at", duration).Error
	if err != nil {
		t.Fatal(err)
	}
}
