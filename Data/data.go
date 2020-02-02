package Data

type Data struct {
	CallId    string `json:"callId"`
	Location  string `json:"location"`
	Situation string `json:"situation"`
	Name      string `json:"name"`
}

func (d Data) UpdateTable() error {
	return nil
}
