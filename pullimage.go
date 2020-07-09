package main

import (
	"log"

	"github.com/softkk/harbor-webhook/models"
)

func pullImage(message models.HookMessage) {
	eventData := message.EventData
	log.Printf("Pull image event is detected, repoURL: %s", eventData.Resources[0].ResourceURL)
}
