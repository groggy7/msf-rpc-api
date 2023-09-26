package api

type ModuleBaseRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type ModuleBaseResponse struct {
	Modules []string `msgpack:"modules"`
}

type ModuleInfoRequest struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	ModuleType string
	ModuleName string
}

type ModuleInfoResponse struct {
	Name        string     `msgpack:"name"`
	Description string     `msgpack:"description"`
	License     string     `msgpack:"license"`
	Filepath    string     `msgpack:"filepath"`
	Version     string     `msgpack:"version"`
	Rank        string     `msgpack:"rank"`
	References  [][]string `msgpack:"references"`
	Authors     []string   `msgpack:"authors"`
}

type ModuleOptionsRequest struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	ModuleType string
	ModuleName string
}

type ModuleOptionsResp map[string]struct {
	Type     string `msgpack:"type"`
	Required bool   `msgpack:"required"`
	Advanced bool   `msgpack:"advanced"`
	Evasion  bool   `msgpack:"evasion"`
	Desc     string `msgpack:"desc"`
	Default  any    `msgpack:"default"`
}

type ModuleCompatiblePayloadsRequest struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	ModuleName string
}

type ModuleCompatiblePayloadsResponse struct {
	Payloads []string `msgpack:"payloads"`
}

type ModuleTargetCompatiblePayloadsRequest struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	ModuleName string
	Target     int
}

type ModuleTargetCompatiblePayloadsResponse struct {
	Payloads []string `msgpack:"payloads"`
}

type ModuleCompatibleSesionsRequest struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	ModuleName string
}

type ModuleCompatibleSesionsResponse struct {
	Sessions []int `msgpack:"sessions"`
}

type ModuleEncodeRequest struct {
	_msgpack       struct{} `msgpack:",asArray"`
	Method         string
	Token          string
	Data           string
	EncoderModule  string
	EncoderOptions map[string]string
}

type ModuleEncodeResponse struct {
	Encoded string `msgpack:"encoded"`
}

type ModuleExecuteRequest struct {
	_msgpack      struct{} `msgpack:",asArray"`
	Method        string
	Token         string
	ModuleType    string
	ModuleName    string
	ModuleOptions map[string]string
}

type ModuleExecuteResponse struct {
	JobId uint16 `msgpack:"job_id"`
}

func (msf *Metasploit) ModuleExploits() (*ModuleBaseResponse, error) {
	req := &ModuleBaseRequest{
		Method: "module.exploits",
		Token:  msf.Token,
	}

	var res ModuleBaseResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ModuleAuxiliary() (*ModuleBaseResponse, error) {
	req := &ModuleBaseRequest{
		Method: "module.auxiliary",
		Token:  msf.Token,
	}

	var res ModuleBaseResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ModulePost() (*ModuleBaseResponse, error) {
	req := &ModuleBaseRequest{
		Method: "module.post",
		Token:  msf.Token,
	}

	var res ModuleBaseResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ModulePayloads() (*ModuleBaseResponse, error) {
	req := &ModuleBaseRequest{
		Method: "module.payloads",
		Token:  msf.Token,
	}

	var res ModuleBaseResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ModuleEncoders() (*ModuleBaseResponse, error) {
	req := &ModuleBaseRequest{
		Method: "module.encoders",
		Token:  msf.Token,
	}

	var res ModuleBaseResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ModuleNops() (*ModuleBaseResponse, error) {
	req := &ModuleBaseRequest{
		Method: "module.nops",
		Token:  msf.Token,
	}

	var res ModuleBaseResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ModuleInfo(moduleType, moduleName string) (*ModuleInfoResponse, error) {
	req := &ModuleInfoRequest{
		Method:     "module.info",
		Token:      msf.Token,
		ModuleType: moduleType,
		ModuleName: moduleName,
	}

	var res ModuleInfoResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ModuleOptions(moduleType, moduleName string) (*ModuleOptionsResp, error) {
	req := &ModuleOptionsRequest{
		Method:     "module.options",
		Token:      msf.Token,
		ModuleType: moduleType,
		ModuleName: moduleName,
	}

	var res ModuleOptionsResp
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ModuleCompatiblePayloads(moduleName string) (*ModuleCompatiblePayloadsResponse, error) {
	req := &ModuleCompatiblePayloadsRequest{
		Method:     "module.compatible_payloads",
		Token:      msf.Token,
		ModuleName: moduleName,
	}

	var res ModuleCompatiblePayloadsResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ModuleTargetCompatiblePayloads(moduleName string, target int) (*ModuleTargetCompatiblePayloadsResponse, error) {
	req := &ModuleTargetCompatiblePayloadsRequest{
		Method:     "module.target_compatible_payloads",
		Token:      msf.Token,
		ModuleName: moduleName,
		Target:     target,
	}

	var res ModuleTargetCompatiblePayloadsResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ModuleCompatibleSessions(moduleName string) (*ModuleCompatibleSesionsResponse, error) {
	req := &ModuleCompatibleSesionsRequest{
		Method:     "module.compatible_sessions",
		Token:      msf.Token,
		ModuleName: moduleName,
	}

	var res ModuleCompatibleSesionsResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ModuleEncode(data, encoderModule string, encoderOptions map[string]string) (*ModuleEncodeResponse, error) {
	req := &ModuleEncodeRequest{
		Method:         "module.encode",
		Token:          msf.Token,
		Data:           data,
		EncoderModule:  encoderModule,
		EncoderOptions: encoderOptions,
	}

	var res ModuleEncodeResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ModuleExecute(moduleType, moduleName string, moduleOptions map[string]string) (*ModuleExecuteResponse, error) {
	req := &ModuleExecuteRequest{
		Method:        "module.execute",
		Token:         msf.Token,
		ModuleType:    moduleType,
		ModuleName:    moduleName,
		ModuleOptions: moduleOptions,
	}

	var res ModuleExecuteResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
