package api

type JobsListRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type JobsListResponse map[string]string

type JobInfoRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	JobID    string
}

type JobInfoResponse struct {
	Jid       int            `msgpack:"jid"`
	Name      string         `msgpack:"name"`
	StartTime int            `msgpack:"start_time"`
	UriPath   string         `msgpack:"uripath"`
	DataStore map[string]any `msgpack:"datastore"`
}

type JobStopRequest struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	JobID    string
}

type JobStopResponse struct {
	Result       string `msgpack:"result"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

func (msf *Metasploit) JobList() (*JobsListResponse, error) {
	req := &JobsListRequest{
		Method: "job.list",
		Token:  msf.Token,
	}

	var res JobsListResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) JobInfo(jobId string) (*JobInfoResponse, error) {
	req := &JobInfoRequest{
		Method: "job.info",
		Token:  msf.Token,
		JobID:  jobId,
	}

	var res JobInfoResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (msf *Metasploit) JobStop(jobId string) (*JobStopResponse, error) {
	req := &JobStopRequest{
		Method: "job.stop",
		Token:  msf.Token,
		JobID:  jobId,
	}

	var res JobStopResponse
	if err := msf.SendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
