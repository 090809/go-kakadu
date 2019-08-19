package handlers

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null"

	"github.com/090809/telebot"
	"go-kakadu/models"
)

type Handler struct {
	bot    *telebot.Bot
	db 	   *sql.DB
	logger *log.Logger
}

func NewHandler(bot *telebot.Bot, db *sql.DB, logger *log.Logger) *Handler {
	return &Handler{bot: bot, db: db, logger: logger}
}

func (h *Handler) getAuthUserById(sender *telebot.User) *models.User {
	senderId := sender.Recipient()
	authUser, err := models.Users(models.UserWhere.TGID.EQ(null.StringFrom(senderId))).One(context.Background(), h.db); if err != nil {
		h.logger.Errorln("Error when trying to get user: ", senderId, err)
	}

	return authUser
}

func (h *Handler) getAuthUserByLink(sender *telebot.User) *models.User {
	link := sender.Username
	senderId := sender.Recipient()
	authUser, err := models.Users(models.UserWhere.TGLink.EQ(null.StringFrom(link))).One(context.Background(), h.db); if err != nil {
		h.logger.Errorln("Error when trying to get user: ", senderId, err)
	}

	return authUser
}
