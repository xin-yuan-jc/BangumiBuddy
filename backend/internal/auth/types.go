package auth

// Credentials 凭证信息
type Credentials struct {
	AccessToken  string
	TokenType    string
	RefreshToken string
}
