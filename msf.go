package api

import (
	"bytes"
	"net/http"

	"gopkg.in/vmihailenco/msgpack.v2"
)

type Metasploit struct {
	Host  string
	User  string
	Pass  string
	Token string
}

func InitMsf(url, username, pass string) *Metasploit {
	return &Metasploit{
		Host: url,
		User: username,
		Pass: pass,
	}
}

func (msf *Metasploit) SendRequest(req interface{}, res interface{}) error {
	buf := new(bytes.Buffer)
	if err := msgpack.NewEncoder(buf).Encode(req); err != nil {
		return err
	}

	response, err := http.Post(msf.Host, "binary/message-pack", buf)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err := msgpack.NewDecoder(response.Body).Decode(&res); err != nil {
		return err
	}
	return nil
}
