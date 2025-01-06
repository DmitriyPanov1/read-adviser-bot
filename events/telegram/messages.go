package telegram

const msgHelp = `Я могу сохранять и хранить ваши страницы. Также я могу предложить вам их для чтения.

Чтобы сохранить страницу, просто пришлите мне ссылку на нее.

Чтобы получить случайную страницу из вашего списка, отправьте мне команду /rnd.

Внимание! После этого эта страница будет удалена из вашего списка!`

const msgHello = "Привет! 👋 \n\n" + msgHelp

const (
	msgUnknownCommand = "Не известная команда 🤷"
	msgNoSavedPages   = "У вас нет сохраненных страниц 🙈"
	msgSaved          = "Сохранено! 👌"
	msgAlreadyExists  = "Эта страница уже есть в вашем списке 🤗"
)