package handlers

import (
	"log"

	tb "github.com/090809/telebot"
)

func TextHandle (handler *Handler) func(m *tb.Message) {
	return textHandler{*handler}.handle
}

type textHandler struct {
	Handler
}

const (
	userNotAuth = "Я не работаю с неавторизованными пользователями!"
	notFetched  = "Ваш запрос в обработке, бип-бип... \n\n" +
		"Команда не опознана!"
)

func (h *textHandler) handle(m *tb.Message) {
	u := h.getAuthUserById(m.Sender)
	if u == nil {
		if _, err := h.bot.Send(m.Sender, userNotAuth); err != nil {
			log.Printf("[ERROR] (TextHandler->Handle {getAuth}): %v", err.Error())
		}
		return
	}

	switch u.TGState.String {
	default:
		if _, err := h.bot.Send(m.Sender, notFetched); err != nil {
			log.Printf("[ERROR] (TextHandler->Handle {TgState: default}): %v", err.Error())
		}
		return
	}
}
