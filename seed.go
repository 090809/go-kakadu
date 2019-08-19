package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null"
	"go-kakadu/models"

	"github.com/volatiletech/sqlboiler/boil"
)

func seed() {
	logger := log.New()
	ctx := context.Background()
	client := getDBClient(logger)

	user := models.User{
		FirstName: null.StringFrom("Георгий"),
		LastName:  null.StringFrom("Каргин"),
		VKLink:    null.StringFrom("https://vk.com/god_daimos"),
		TGLink:    null.StringFrom("Pallam"),
		TGState:   null.StringFrom(models.UsersTGStateNOT_LOGGED),
	}

	if err := user.Insert(ctx, client, boil.Infer()); err != nil {
		panic(err)
	}

	err := user.AddDepartments(ctx, client, true, &models.Department{
		Name:   "Отдел прикладных разработок",
		LeadID: null.UintFrom(user.ID),
	}); if err != nil {
		panic(err)
	}

	project := models.Project{
		Name:   "Go Kakadu for BD",
		LeadID: null.UintFrom(user.ID),
	}

	if err := project.Insert(ctx, client, boil.Infer()); err != nil {
		panic(err)
	}

	if err := user.AddIssues(ctx, client, true, &models.Issue{
			Name:        "Issue1",
			Description: "This is a issue 1",
			PerformerID: null.UintFrom(user.ID),
			CreatorID:   null.UintFrom(user.ID),
			ProjectID:	 null.UintFrom(project.ID),
			Status:      null.StringFrom(models.IssuesStatusIN_PROGRESS),
			DeadlineAt:  null.TimeFrom(time.Now().AddDate(0, 0, 10)),
		}, &models.Issue{
			Name:        "Issue2",
			Description: "This is a issue 2",
			PerformerID: null.UintFrom(user.ID),
			CreatorID:   null.UintFrom(user.ID),
			ProjectID:	 null.UintFrom(project.ID),
			Status:      null.StringFrom(models.IssuesStatusNOT_STARTED),
			DeadlineAt:  null.TimeFrom(time.Now().AddDate(0, 0, 20)),
		},
	); err != nil {
		panic(err)
	}
}
