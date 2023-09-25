package api

type PluginLoadRequest struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	PluginName string
	Options    map[string]string
}

type PluginLoadResponse struct {
	Result string `msgpack:"result"`
}

type PluginUnLoadRequest struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	PluginName string
}

type PluginUnloadRequest struct {
	Result string `msgpack:"result"`
}

type PluginLoadedRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type PluginLoadedResponse struct {
	Plugins []string `msgpack:"plugins"`
}

func (msf *Metasploit) PluginLoad(pluginName string, pluginOptions map[string]string) (*PluginLoadResponse, error) {
	ctx := &PluginLoadRequest{
		Method:     "plugin.load",
		Token:      msf.Token,
		PluginName: pluginName,
		Options:    pluginOptions,
	}

	var res PluginLoadResponse
	if err := msf.SendRequest(ctx, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) PluginUnload(pluginName string) (*PluginUnloadRequest, error) {
	ctx := &PluginUnLoadRequest{
		Method:     "plugin.unload",
		Token:      msf.Token,
		PluginName: pluginName,
	}

	var res PluginUnloadRequest
	if err := msf.SendRequest(ctx, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) PluginLoaded() (*PluginLoadedResponse, error) {
	ctx := &PluginLoadedRequest{
		Method: "plugin.loaded",
		Token:  msf.Token,
	}
	var res PluginLoadedResponse
	if err := msf.SendRequest(ctx, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
