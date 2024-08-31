package entities

import "strings"

type AuthStatus int
type AuthScope int

const (
	AuthGranted AuthStatus = iota
	AuthUnauthorized
	AuthError
)

const (
	ScopeRead AuthScope = iota
	ScopeWrite
	ScopeAdmin
)

func (s AuthScope) String() string {
	switch s {
	case ScopeRead:
		return "read"
	case ScopeWrite:
		return "write"
	case ScopeAdmin:
		return "admin"
	default:
		return "read"
	}
}

func (s AuthStatus) String() string {
	switch s {
	case AuthGranted:
		return "granted"
	case AuthUnauthorized:
		return "unauthorized"
	case AuthError:
		return "error"
	default:
		return "error"
	}
}

type ValidationResult struct {
	Status   AuthStatus
	Username string
	Email    string
	Scopes   []AuthScope
}

func (v *ValidationResult) JoinScopes() string {
	scopes := make([]string, len(v.Scopes))
	for i, v := range v.Scopes {
		scopes[i] = v.String()
	}
	return strings.Join(scopes, ",")
}
