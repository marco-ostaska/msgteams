package teams

import (
	"io"
	"net/http"
	"strings"
)

type message struct {
	msg     string
	url     string
	payload io.Reader
	request *http.Request
}

func (m *message) setMsg(msg string) { m.msg = msg }

func (m *message) setPayload() {
	m.payload = strings.NewReader(`{ 'text': '` + m.msg + `' }`)
}

func (m *message) setURL(url string) { m.url = url }

func (m *message) setNewRequest() error {
	m.setPayload()
	request, err := http.NewRequest("POST", m.url, m.payload)

	if err == nil {
		request.Header.Add("Content-Type", "application/json")
		m.request = request
	}

	return err
}

func (m *message) postMsg() error {
	if err := m.setNewRequest(); err != nil {
		return err
	}

	client := &http.Client{}
	response, err := client.Do(m.request)

	if err == nil {
		defer response.Body.Close()
	}

	return err
}

// NewPostMsg send a message to ms teams. Ir returns an error
func NewPostMsg(url string, msg string) error {
	m := message{}
	m.setURL(url)
	m.setMsg(msg)

	return m.postMsg()
}
