package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type ClientTG struct {
	Host   string
	Path   string
	Client http.Client
}

func NewClient(host, token string) *ClientTG {
	return &ClientTG{
		Host:   host, // https://api.telegram.org
		Path:   newBasePath(token),
		Client: http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMassage"
)

// Updates метод - это запрос на сервер TGBot, т.е. как я понял, это тело запроса
func (c *ClientTG) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset)) // создаём параметры запроса
	q.Add("limit", strconv.Itoa(limit))   // Которые потом передадим в string к запросу

	//do request
	data, err := c.doRequest(q, getUpdatesMethod)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return res.result, err
	}

	return nil, nil
}

func (c *ClientTG) SendMassage(chatID int, text string) ([]Update, error) {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	//do request
	var res UpdatesResponse
	data, err := c.doRequest(q, sendMessageMethod)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res.result, err
}

// doRequest метод - который отправляет запрос на сервер TGBot/Api
func (c *ClientTG) doRequest(query url.Values, method string) ([]byte, error) {
	queryString := url.URL{ // создание структуры строки запроса с протоколом, доменом и путём
		Scheme: "https",
		Host:   c.Host,
		Path:   path.Join(c.Path, method)}

	// http.NewRequest - возвращает указатель на структуру http запроса.
	req, err := http.NewRequest(http.MethodGet, queryString.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать запрос для отправки на сервер %s", queryString.Host)
	}

	req.URL.RawQuery = query.Encode() // Кодирует параметры запроса в поле RawQuery структуры url, тем самым добавляя...
	// ... новые заголовки к запроса

	response, err := c.Client.Do(req) // Отправляет запрос
	if err != nil {
		// Возможно стоит создать отдельные пакеты, для обработки ошибок, которые...
		// ...относятся к разным слоям и создать главный пакет, который отвечает за обработку бизнес-ошибок.
		return nil, fmt.Errorf("не удалось выполнить запрос на сервер %s", queryString.Host)
	}

	defer func() { _ = response.Body.Close() }()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить тело ответа при запросе на сервер %s", queryString.Host)
	}

	return body, nil
}

// Я предполагаю, что потом создаётся какой-то отдельный файл, и там вызывается NewClient и методы
