package events

import (
	"log"

	"github.com/ebizno/Ula/internal/email"
	"github.com/ebizno/Ula/internal/file"
	"github.com/ebizno/Ula/internal/worker"
	"github.com/fsnotify/fsnotify"
)

type (
	IEvent interface {
		NewWatcher() (*fsnotify.Watcher, error)
		Watcher(path string) error
		WatcherEvent()
		WatcherClose()
		Event()
	}
	Event struct {
		Watche          *fsnotify.Watcher
		File            file.IFilePath
		EmailCredential email.EmailCredentialInterface
		Subject         string
		Body            string
		ContentType     string
	}
)

var (
	done             = make(chan bool)
	contentTypePlain = make(chan bool)
	contentTypeHtml  = make(chan bool)
)

func NewEvent(IEmailCredential email.EmailCredentialInterface, subject, body string, contentType string, file file.IFilePath) IEvent {
	e := &Event{
		Watche:          &fsnotify.Watcher{},
		File:            file,
		EmailCredential: IEmailCredential,
		Subject:         subject,
		Body:            body,
		ContentType:     contentType,
	}
	return e
}

func (e *Event) NewWatcher() (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return watcher, nil
}

func (e *Event) Watcher(path string) error {
	if err := e.Watche.Add(path); err != nil {
		return err
	}
	<-done
	return nil
}

func (e *Event) WatcherEvent() {
	for {
		select {
		case event := <-e.Watche.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("modified file:", event.Name, "event:", event.Op)
				contentTypePlain <- e.ContentType == "text/plain"
				contentTypeHtml <- e.ContentType == "text/html"
			}
		case err := <-e.Watche.Errors:
			log.Println("error:", err)
		}
	}
}

func (e *Event) WatcherClose() {
	e.Watche.Close()
}

func (e *Event) Event() {
	watcher, err := e.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	e.Watche = watcher

	go e.WatcherEvent()
	go e.selectContentTypeEmail()

	if err := e.Watcher(e.File.Dir()); err != nil {
		log.Fatal(err)
	}
	defer e.WatcherClose()
}

func (e *Event) selectContentTypeEmail() {
	go func() {
		for {
			select {
			case <-contentTypePlain:
				e.EventContentTypePlain()
			case <-contentTypeHtml:
				//e.EventContentTypeHtml()
			}
		}
	}()
}

func (e *Event) EventContentTypePlain() {

	eventEmail := e.EventEmail()

	worker := worker.NewWorker(eventEmail, e.File.FileName())
	worker.WorkerPlain()
}

func (e *Event) EventEmail() *email.Email {
	email := &email.Email{
		Subject:          e.Subject,
		Body:             e.Body,
		IEmailCredential: e.EmailCredential,
	}
	return email
}
