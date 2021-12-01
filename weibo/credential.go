package weibo

type Credential struct {
	clientId     string `json:"client_id"`
	clientSecret string `json:"client_secret"`
}

func NewCredential(clientId, clientSecret string) *Credential {
	return &Credential{clientId: clientId, clientSecret: clientSecret}
}
