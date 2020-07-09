package main

import (
	"log"

	"github.com/softkk/harbor-webhook/database"
	"github.com/softkk/harbor-webhook/models"
)

func updateDB(message models.HookMessage) {
	eventData := message.EventData
	log.Printf("Push image event is detected, repoURL: %s", eventData.Resources[0].ResourceURL)
	log.Println(message.OccurAT.String())
	project := database.Project{}
	project.SelectByName(eventData.Repository.Namespace)
	if project.Owner == 0 || project.ID == 0 {
		log.Printf("Not found, OwnerID: %d or ProjectID %d", project.Owner, project.ID)
		return
	}
	log.Printf("Namespace: %s, OwnerID: %d, ProjectID %d", eventData.Repository.Namespace, project.Owner, project.ID)

	//DB
	//(`purpose`, `status`, `applicant`, `review`        , `createdAt`, `updatedAt`)
	//(,'CICD API更新',99,{專案管理人UID},1,NOW(),NOW())
	dateCreated := message.OccurAT.Time()
	image := database.Image{
		Name:      eventData.Repository.Name,
		Tag:       eventData.Resources[0].Tag,
		RealPath:  eventData.Resources[0].ResourceURL,
		Purpose:   "webhook",
		Status:    99,
		Applicant: project.Owner,
		Review:    1,
		CreatedAt: &dateCreated,
		UpdatedAt: &dateCreated,
	}
	tx := database.MysqlDB.Begin()

	result := tx.Create(&image)
	if result.Error != nil {
		log.Println(result.Error)
		tx.Rollback()
		return
	}
	log.Printf("Get ImageID: %d", image.ID)
	projectImage := database.ProjectImage{
		ImageID:   image.ID,
		ProjectID: project.ID,
		CreatedAt: &dateCreated,
		UpdatedAt: &dateCreated,
	}

	result2 := tx.Create(&projectImage)
	if result2.Error != nil {
		log.Println(result2.Error)
		tx.Rollback()
		return
	}
	tx.Commit()
}
