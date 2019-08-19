package handlers

import (
	"context"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"

	tb "github.com/090809/telebot"
	"go-kakadu/models"
)

func StartHandle (handler *Handler) func(m *tb.Message) {
	return startHandler{*handler}.Handle
}

type startHandler struct {
	Handler
}

const userNotFound = "Для работы с этим приложением, тебя сначала должны добавить в базу!\n" +
	"Попроси своего менеджера об этом"

func (h *startHandler) Handle(m *tb.Message) {
	user := h.getAuthUserByLink(m.Sender)
	if user == nil {
		if _, err := h.bot.Send(m.Sender, userNotFound); err != nil {
			log.Fatalln(err)
		}
		return
	}

	var name []string
	if user.FirstName.String != "" {
		name = append(name, user.FirstName.String)
	}
	if user.LastName.String != "" {
		name = append(name, user.LastName.String)
	}

	if _, err := h.bot.Send(m.Sender, "Привет, " + strings.Join(name, " ")); err != nil {
		log.Fatalln(err)
	}

	user.TGState = null.StringFrom(models.UsersTGStateMAIN)
	user.TGID = null.StringFrom(m.Sender.Recipient())

	if _, err := user.Update(context.Background(), h.db, boil.Infer()); err != nil {
		h.logger.Errorln("Error when updating user: ", user.TGID, err)
	}
}
