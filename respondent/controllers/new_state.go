package controllers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"webcapstaffbot/respondent"
)

type NewState struct {

}

func (c *NewState) CreateMessage(respondent *respondent.Respondent) tgbotapi.Chattable  {
	msg := tgbotapi.NewMessage(respondent.State.ID, "Hi :)")

	return msg
}