package nap

type AuthToken struct {
	Token string
}

type AuthBasic struct {
	Username string
	Password string
}

// API requires different authentication thus may have different implementation
// of this interface
type Authentication interface {
	AuthorizationHeader() string // "basic <based64-encoded string>"
}
