package events_test

import (
	"testing"

	"github.com/ebizno/Ula/internal/email"
	"github.com/ebizno/Ula/internal/events"
	"github.com/ebizno/Ula/internal/file"
)

func TestNewEvent(t *testing.T) {
	fileName, err := file.NewFilePath("ula.csv")
	if err != nil {
		t.Errorf("Expected file to be not nil %s", err)
	}

	t.Log(fileName)

	emailCredential, err := email.NewEmailCredential("ula@gmail.com", "xxxx", 587, "smtp.gmail.com")
	if err != nil {
		t.Errorf("Expected emailCredential to be not nil %s", err)
	}

	event := events.NewEvent(
		emailCredential,
		"Subject: Title Test",
		"Body: Hello Word, é um teste Paulo",
		"text/plain",
		fileName)
	if event == nil {
		t.Errorf("Expected event to be not nil")
	}
	event.Event()

}
