package api

type ConsoleCreateRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type ConsoleCreateResponse struct {
	Id     string `msgpack:"id"`
	Prompt string `msgpack:"prompt"`
	Busy   bool   `msgpack:"busy"`
}

type ConsoleDestroyRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleId string
}

type ConsoleDestroyResponse struct {
	Result       string `msgpack:"result"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type ConsoleListRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type ConsoleListResponse map[string][]struct {
	Id     string `msgpack:"id"`
	Prompt string `msgpack:"prompt"`
	Busy   bool   `msgpack:"busy"`
}

type ConsoleWriteRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleId string
	Command   string
}

type ConsoleWriteResponse struct {
	Wrote        int    `msgpack:"wrote"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type ConsoleReadRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleId string
}

type ConsoleReadResponse struct {
	Data   string `msgpack:"data"`
	Prompt string `msgpack:"prompt"`
	Busy   string `msgpack:"busy"`
}

type ConsoleSessionDetachRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleId string
}

type ConsoleSessionDetachResponse struct {
	Result       string `msgpack:"result"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type ConsoleSessionKillRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleId string
}

type ConsoleSessionKillResponse struct {
	Result       string `msgpack:"result"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type ConsoleTabsRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleId string
	InputLine string
}

type ConsoleTabsResponse struct {
	Tabs []string `msgpack:"tabs"`
}

func (msf *Metasploit) ConsoleCreate() (*ConsoleCreateResponse, error) {
	req := &ConsoleCreateRequest{
		Method: "console.create",
		Token:  msf.Token,
	}

	var res ConsoleCreateResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ConsoleDestroy(consoleId string) (*ConsoleDestroyResponse, error) {
	req := &ConsoleDestroyRequest{
		Method:    "console.destroy",
		Token:     msf.Token,
		ConsoleId: consoleId,
	}

	var res ConsoleDestroyResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ConsoleList() (*ConsoleListResponse, error) {
	req := &ConsoleListRequest{
		Method: "console.list",
		Token:  msf.Token,
	}

	var res ConsoleListResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ConsoleWrite(consoleId, command string) (*ConsoleWriteResponse, error) {
	req := &ConsoleWriteRequest{
		Method:    "console.write",
		Token:     msf.Token,
		ConsoleId: consoleId,
		Command:   command + "\n",
	}

	var res ConsoleWriteResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ConsoleRead(consoleId string) (*ConsoleReadResponse, error) {
	req := &ConsoleReadRequest{
		Method:    "console.read",
		Token:     msf.Token,
		ConsoleId: consoleId,
	}

	var res ConsoleReadResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ConsoleDetachSession(consoleId string) (*ConsoleSessionDetachResponse, error) {
	req := &ConsoleSessionDetachRequest{
		Method:    "console.session_detach",
		Token:     msf.Token,
		ConsoleId: consoleId,
	}

	var res ConsoleSessionDetachResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ConsoleKillSession(consoleId string) (*ConsoleSessionKillResponse, error) {
	req := &ConsoleSessionKillRequest{
		Method:    "console.session_kill",
		Token:     msf.Token,
		ConsoleId: consoleId,
	}

	var res ConsoleSessionKillResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ConsoleTabs(consoleId, inputLine string) (*ConsoleTabsResponse, error) {
	req := &ConsoleTabsRequest{
		Method:    "console.tabs",
		Token:     msf.Token,
		ConsoleId: consoleId,
		InputLine: inputLine,
	}

	var res ConsoleTabsResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
