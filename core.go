package api

type CoreAddModulePathRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	Path     string
}

type CoreAddModulePathResponse struct {
	Exploits  uint32 `msgpack:"exploits"`
	Auxiliary uint32 `msgpack:"auxiliary"`
	Post      uint32 `msgpack:"post"`
	Encoders  uint32 `msgpack:"encoders"`
	Nops      uint32 `msgpack:"nops"`
	Payloads  uint32 `msgpack:"payloads"`
}

type CoreGetModuleStatsRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type CoreGetModuleStatsResponse struct {
	Exploits  uint32 `msgpack:"exploits"`
	Auxiliary uint32 `msgpack:"auxiliary"`
	Post      uint32 `msgpack:"post"`
	Encoders  uint32 `msgpack:"encoders"`
	Nops      uint32 `msgpack:"nops"`
	Payloads  uint32 `msgpack:"payloads"`
}

type CoreReloadModuleRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type CoreReloadModuleResponse struct {
	Exploits  uint32 `msgpack:"exploits"`
	Auxiliary uint32 `msgpack:"auxiliary"`
	Post      uint32 `msgpack:"post"`
	Encoders  uint32 `msgpack:"encoders"`
	Nops      uint32 `msgpack:"nops"`
	Payloads  uint32 `msgpack:"payloads"`
}

type CoreSaveRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type CoreSaveResponse struct {
	Result       []string `msgpack:"result"`
	Error        bool     `msgpack:"error"`
	ErrorClass   string   `msgpack:"error_class"`
	ErrorMessage string   `msgpack:"error_message"`
}

type CoreSetGRequest struct {
	_msgpack    struct{} `msgpack:",asArray"`
	Method      string
	Token       string
	OptionName  string
	OptionValue string
}

type CoreSetGResponse struct {
	Result       string `msgpack:"result"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type CoreUnsetGRequest struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	OptionName string
}

type CoreUnsetGResponse struct {
	Result string `msgpack:"result"`
	Error  bool   `msgpack:"error"`
}

type CoreThreadListRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type CoreThreadListResponse map[int]struct {
	Status   string `msgpack:"status"`
	Critical bool   `msgpack:"critical"`
	Name     string `msgpack:"name"`
	Started  string `msgpack:"started"`
}

type CoreThreadKillRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	ThreadId string
}

type CoreThreadKillResponse struct {
	Result       string `msgpack:"result"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type CoreGetVersionRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type CoreGetVersionResponse struct {
	Version string `msgpack:"version"`
	Ruby    string `msgpack:"ruby"`
	Api     string `msgpack:"api"`
}

type CoreStopRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type CoreStopResponse struct {
	Result       string `msgpack:"result"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

func (msf *Metasploit) AddModulePath(path string) (*CoreAddModulePathResponse, error) {
	req := &CoreAddModulePathRequest{
		Method: "core.add_module_path",
		Token:  msf.Token,
		Path:   path,
	}

	var res CoreAddModulePathResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) GetCoreModuleStats() (*CoreGetModuleStatsResponse, error) {
	req := &CoreGetModuleStatsRequest{
		Method: "core.module_stats",
		Token:  msf.Token,
	}

	var res CoreGetModuleStatsResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ReloadModules() (*CoreReloadModuleResponse, error) {
	req := CoreReloadModuleRequest{
		Method: "core.reload_modules",
		Token:  msf.Token,
	}

	var res CoreReloadModuleResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) Save() (*CoreSaveResponse, error) {
	req := CoreSaveRequest{
		Method: "core.save",
		Token:  msf.Token,
	}

	var res CoreSaveResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) SetG(optionName, optionValue string) (*CoreSetGResponse, error) {
	req := CoreSetGRequest{
		Method:      "core.setg",
		Token:       msf.Token,
		OptionName:  optionName,
		OptionValue: optionValue,
	}

	var res CoreSetGResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) UnsetG(optionName string) (*CoreUnsetGResponse, error) {
	req := CoreUnsetGRequest{
		Method:     "core.unsetg",
		Token:      msf.Token,
		OptionName: optionName,
	}

	var res CoreUnsetGResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) ThreadList() (*CoreThreadListResponse, error) {
	req := CoreThreadListRequest{
		Method: "core.thread_list",
		Token:  msf.Token,
	}

	var res CoreThreadListResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) KillThread(threadId string) (*CoreThreadKillResponse, error) {
	req := CoreThreadKillRequest{
		Method:   "core.thread_kill",
		Token:    msf.Token,
		ThreadId: threadId,
	}

	var res CoreThreadKillResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) GetVersion() (*CoreGetVersionResponse, error) {
	req := CoreGetVersionRequest{
		Method: "core.version",
		Token:  msf.Token,
	}

	var res CoreGetVersionResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) Stop() (*CoreStopResponse, error) {
	req := CoreStopRequest{
		Method: "core.stop",
		Token:  msf.Token,
	}

	var res CoreStopResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
