package api

type LoginRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Username string
	Password string
}

type LoginResponse struct {
	Result       string `msgpack:"result"`
	Token        string `msgpack:"token"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type LogoutRequest struct {
	_msgpack    struct{} `msgpack:",asArray"`
	Method      string
	Token       string
	LogoutToken string
}

type LogoutResponse struct {
	Result       string `msgpack:"result"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type AddTokenRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	NewToken string
}

type AddTokenResponse struct {
	Result       string `msgpack:"result"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type GenerateTokenRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type GenerateTokenResponse struct {
	Result       string `msgpack:"result"`
	Token        string `msgpack:"token"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type GetTokenListRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type GetTokenListResponse struct {
	Token        []string `msgpack:"tokens"`
	Error        bool     `msgpack:"error"`
	ErrorClass   string   `msgpack:"error_class"`
	ErrorMessage string   `msgpack:"error_message"`
}

type RemoveTokenRequest struct {
	_msgpack         struct{} `msgpack:",asArray"`
	Method           string
	Token            string
	TokenToBeRemoved string
}

type RemoveTokenResponse struct {
	Result       string `msgpack:"result"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

func (msf *Metasploit) Login() (*LoginResponse, error) {
	req := &LoginRequest{
		Method:   "auth.login",
		Username: msf.User,
		Password: msf.Pass,
	}

	var res LoginResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	msf.Token = res.Token
	return &res, nil
}

func (msf *Metasploit) Logout() (*LogoutResponse, error) {
	req := &LogoutRequest{
		Method:      "auth.logout",
		Token:       msf.Token,
		LogoutToken: msf.Token,
	}
	var res LogoutResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	msf.Token = ""
	return &res, nil
}

func (msf *Metasploit) AddToken(token string) (*AddTokenResponse, error) {
	req := &AddTokenRequest{
		Method:   "auth.token_add",
		Token:    msf.Token,
		NewToken: token,
	}

	var res AddTokenResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) GenerateToken() (*GenerateTokenResponse, error) {
	req := &GenerateTokenRequest{
		Method: "auth.token_generate",
		Token:  msf.Token,
	}

	var res GenerateTokenResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) GetTokens() (*GetTokenListResponse, error) {
	req := &GetTokenListRequest{
		Method: "auth.token_list",
		Token:  msf.Token,
	}

	var res GetTokenListResponse
	if err := msf.SendRequest(res, &req); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) RemoveToken(token string) (*RemoveTokenResponse, error) {
	req := RemoveTokenRequest{
		Method:           "auth.token_remove",
		Token:            msf.Token,
		TokenToBeRemoved: token,
	}

	var res RemoveTokenResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
