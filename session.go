package api

import "fmt"

type SessionListRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type SessionListResponse map[uint32]struct {
	Type        string `msgpack:"type"`
	TunnelLocal string `msgpack:"tunnel_local"`
	TunnelPeer  string `msgpack:"tunnel_peer"`
	ViaExploit  string `msgpack:"via_exploit"`
	ViaPayload  string `msgpack:"via_payload"`
	Description string `msgpack:"desc"`
	Info        string `msgpack:"info"`
	Workspace   string `msgpack:"workspace"`
	SessionHost string `msgpack:"session_host"`
	SessionPort int    `msgpack:"session_port"`
	Username    string `msgpack:"username"`
	UUID        string `msgpack:"uuid"`
	ExploitUUID string `msgpack:"exploit_uuid"`
}

type SessionWriteRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
	Command   string
}

type SessionWriteResponse struct {
	WriteCount string `msgpack:"write_count"`
}

type SessionReadRequest struct {
	_msgpack    struct{} `msgpack:",asArray"`
	Method      string
	Token       string
	SessionID   uint32
	ReadPointer string
}

type SessionReadResponse struct {
	Seq  uint32 `msgpack:"seq"`
	Data string `msgpack:"data"`
}

type SessionRingLastRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
}

type SessionRingLastResponse struct {
	Seq uint32 `msgpack:"seq"`
}

type SessionMeterpreterWriteRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
	Command   string
}

type SessionMeterpreterWriteResponse struct {
	Result string `msgpack:"result"`
}

type SessionMeterpreterReadRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
}

type SessionMeterpreterReadResponse struct {
	Data string `msgpack:"data"`
}

type SessionMeterpreterRunSingleRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
	Command   string
}

type SessionMeterpreterRunSingleResponse SessionMeterpreterWriteResponse

type SessionMeterpreterDetachRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
}

type SessionMeterpreterDetachResponse SessionMeterpreterWriteResponse

type SessionMeterpreterKillRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
}

type SessionMeterpreterKillResponse SessionMeterpreterWriteResponse

type SessionMeterpreterTabsRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
	InputLine string
}

type SessionMeterpreterTabsResponse struct {
	Tabs []string `msgpack:"tabs"`
}

type SessionCompatibleModulesRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
}

type SessionCompatibleModulesResponse struct {
	Modules []string `msgpack:"modules"`
}

type SessionShellUpgradeRequest struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	SessionID  uint32
	IpAddress  string
	PortNumber uint32
}

type SessionShellUpgradeResponse SessionMeterpreterWriteResponse

type SessionRingClearRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
}

type SessionRingClearResponse SessionMeterpreterWriteResponse

type SessionRingPutRequest struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID uint32
	Command   string
}

type SessionRingPutResponse struct {
	WriteCount uint32 `msgpack:"write_count"`
}

func (msf *Metasploit) SessionList() (*SessionListResponse, error) {
	req := &SessionListRequest{
		Method: "session.list",
		Token:  msf.Token,
	}

	var res SessionListResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}

func (msf *Metasploit) SessionReadPointer(session uint32) (uint32, error) {
	ctx := &SessionRingLastRequest{
		Method:    "session.ring_last",
		Token:     msf.Token,
		SessionID: session,
	}

	var sesRingLast SessionRingLastResponse
	if err := msf.SendRequest(ctx, &sesRingLast); err != nil {
		return 0, err
	}

	return sesRingLast.Seq, nil
}

func (msf *Metasploit) SessionWrite(session uint32, command string) (string, error) {
	ctx := &SessionWriteRequest{
		Method:    "session.shell_write",
		Token:     msf.Token,
		SessionID: session,
		Command:   command,
	}

	var res SessionWriteResponse
	if err := msf.SendRequest(ctx, &res); err != nil {
		return "", err
	}

	return res.WriteCount, nil
}

func (msf *Metasploit) SessionRead(session uint32, readPointer uint32) (string, error) {
	ctx := &SessionReadRequest{
		Method:      "session.shell_read",
		Token:       msf.Token,
		SessionID:   session,
		ReadPointer: string(rune(readPointer)),
	}

	var res SessionReadResponse
	if err := msf.SendRequest(ctx, &res); err != nil {
		return "", err
	}

	return res.Data, nil
}

func (msf *Metasploit) SessionExecute(session uint32, command string) (string, error) {
	readPointer, err := msf.SessionReadPointer(session)
	if err != nil {
		return "", err
	}
	msf.SessionWrite(session, command)
	data, err := msf.SessionRead(session, readPointer)
	if err != nil {
		return "", err
	}
	return data, nil
}

func (msf *Metasploit) SessionExecuteList(session uint32, commands []string) (string, error) {
	var results string
	for _, command := range commands {
		tCommand := fmt.Sprintf("%s\n", command)
		result, err := msf.SessionExecute(session, tCommand)
		if err != nil {
			return results, err
		}
		results += result
	}

	return results, nil
}

func (msf *Metasploit) SessionMeterpreterWrite(session uint32, command string) (*SessionMeterpreterWriteResponse, error) {
	ctx := &SessionMeterpreterWriteRequest{
		Method:    "session.meterpreter_write",
		Token:     msf.Token,
		SessionID: session,
		Command:   command,
	}

	var res SessionMeterpreterWriteResponse
	if err := msf.SendRequest(ctx, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) SessionMeterpreterRead(session uint32) (*SessionMeterpreterReadResponse, error) {
	ctx := &SessionMeterpreterReadRequest{
		Method:    "session.meterpreter_read",
		Token:     msf.Token,
		SessionID: session,
	}

	var res SessionMeterpreterReadResponse
	if err := msf.SendRequest(ctx, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (msf *Metasploit) SessionMeterpreterRunSingle(session uint32, command string) (*SessionMeterpreterRunSingleResponse, error) {
	ctx := &SessionMeterpreterRunSingleRequest{
		Method:    "session.meterpreter_run_single",
		Token:     msf.Token,
		SessionID: session,
		Command:   command,
	}

	var res SessionMeterpreterRunSingleResponse
	if err := msf.SendRequest(ctx, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) SessionMeterpreterSessionDetach(session uint32) (*SessionMeterpreterDetachResponse, error) {
	ctx := &SessionMeterpreterDetachRequest{
		Method:    "session.meterpreter_session_detach",
		Token:     msf.Token,
		SessionID: session,
	}

	var res SessionMeterpreterDetachResponse
	if err := msf.SendRequest(ctx, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (msf *Metasploit) SessionMeterpreterSessionKill(session uint32) (*SessionMeterpreterKillResponse, error) {
	ctx := &SessionMeterpreterKillRequest{
		Method:    "session.meterpreter_session_kill",
		Token:     msf.Token,
		SessionID: session,
	}

	var res SessionMeterpreterKillResponse
	if err := msf.SendRequest(ctx, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (msf *Metasploit) SessionMeterpreterTabs(session uint32, inputLine string) (*SessionMeterpreterTabsResponse, error) {
	ctx := &SessionMeterpreterTabsRequest{
		Method:    "session.meterpreter_tabs",
		Token:     msf.Token,
		SessionID: session,
		InputLine: inputLine,
	}

	var res SessionMeterpreterTabsResponse
	if err := msf.SendRequest(ctx, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (msf *Metasploit) SessionCompatibleModules(session uint32) (*SessionCompatibleModulesResponse, error) {
	ctx := &SessionCompatibleModulesRequest{
		Method:    "session.compatible_modules",
		Token:     msf.Token,
		SessionID: session,
	}

	var res SessionCompatibleModulesResponse
	if err := msf.SendRequest(ctx, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (msf *Metasploit) SessionShellUpgrade(session uint32, lhostAddress string, lportNumber uint32) (*SessionShellUpgradeResponse, error) {
	ctx := &SessionShellUpgradeRequest{
		Method:     "session.shell_upgrade",
		Token:      msf.Token,
		SessionID:  session,
		IpAddress:  lhostAddress,
		PortNumber: lportNumber,
	}

	var res SessionShellUpgradeResponse
	if err := msf.SendRequest(ctx, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (msf *Metasploit) SessionRingClear(session uint32) (*SessionRingClearResponse, error) {
	ctx := &SessionRingClearRequest{
		Method:    "session.ring_clear",
		Token:     msf.Token,
		SessionID: session,
	}

	var res SessionRingClearResponse
	if err := msf.SendRequest(ctx, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (msf *Metasploit) SessionRingLast(session uint32) (*SessionRingLastResponse, error) {
	ctx := &SessionRingLastRequest{
		Method:    "session.ring_last",
		Token:     msf.Token,
		SessionID: session,
	}

	var res SessionRingLastResponse
	if err := msf.SendRequest(ctx, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (msf *Metasploit) SessionRingPut(session uint32, command string) (*SessionRingPutResponse, error) {
	ctx := &SessionRingPutRequest{
		Method:    "session.ring_put",
		Token:     msf.Token,
		SessionID: session,
		Command:   command,
	}

	var res SessionRingPutResponse
	if err := msf.SendRequest(ctx, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
