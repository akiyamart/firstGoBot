package telegram

import (
	"encoding/json"
	"firstGoBot/clients/lib/err"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	getUpdatesMethod = "getUpdates"
	sendMessageMethod = "sendMessage"
)

type Client struct { 
	host string 
	basePath string 
	client http.Client
}

func New(host string, token string) Client { 
	return Client {
		host:		host, 
		basePath:	newBasePath(token),  
		client:		http.Client{},
	}
}

func newBasePath(token string) string { 
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) (updates []Update, e error) { 
	defer func() {e = err.WrapIfErr("can not do request", e)}()
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, e := c.doRequest(getUpdatesMethod, q)
	if e != nil { 
		return nil, e
	}
	
	var res UpdatesResponse

	if e := json.Unmarshal(data, &res); e != nil { 
		return nil, e
	}

	return res.Result, nil 
}

func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chatID", strconv.Itoa(chatID))
	q.Add("text", text)

	_, e := c.doRequest(sendMessageMethod, q)
	if e != nil { 
		return err.Wrap("can't send message", e)
	}

	return nil
}

func (c *Client) doRequest(method string, query url.Values) (data []byte, e error) { 
	defer func() {e = err.WrapIfErr("can not do request", e)}()

	u := url.URL{
		Scheme: "https",
		Host: c.host,
		Path: path.Join(c.basePath, method),
	}

	req, e := http.NewRequest(http.MethodGet, u.String(), nil)
	if e != nil { 
		return nil, e
	}

	req.URL.RawQuery = query.Encode()
	
	resp, e := c.client.Do(req)
	if e != nil { 
		return nil, e
	}
	defer func() {_=resp.Body.Close()}()

	body, e := io.ReadAll(resp.Body)
	if e != nil { 
		return nil, e
	}

	return body, nil
}
