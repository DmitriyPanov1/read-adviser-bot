package telegram

import (
	"errors"
	"log"
	"net/url"
	"read-adviser-bot/lib/e"
	"read-adviser-bot/storage"
	"strings"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command: '%s' from '%s'", text, username)

	if isAddCmd(text) {
		return p.savePage(chatID, text, username)
	}

	switch text {
	case RndCmd:
		return p.sendRandom(chatID, username)
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}
}

func (p *Processor) savePage(chatID int, pageUrl string, username string) error {
	errMess := "can't save page for this chat"

	page := &storage.Page{
		URL:      pageUrl,
		UserName: username,
	}

	isExist, err := p.storage.IsExists(page)
	if err != nil {
		return e.Wrap(errMess, err)
	}

	if isExist {
		return p.tg.SendMessage(chatID, msgAlreadyExists)
	}

	if err := p.storage.Save(page); err != nil {
		return e.Wrap(errMess, err)
	}

	if err := p.tg.SendMessage(chatID, msgSaved); err != nil {
		return e.Wrap(errMess, err)
	}

	return nil
}

func (p *Processor) sendRandom(chatID int, username string) error {
	errMess := "can't send random command"

	page, err := p.storage.PickRandom(username)
	if err != nil {
		if errors.Is(err, storage.ErrNoSavePages) {
			return p.tg.SendMessage(chatID, msgNoSavedPages)
		}

		return e.Wrap(errMess, err)
	}

	if err := p.tg.SendMessage(chatID, page.URL); err != nil {
		return e.Wrap(errMess, err)
	}

	return p.storage.Remove(page)
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)

	return err == nil && u.Host != ""
}
