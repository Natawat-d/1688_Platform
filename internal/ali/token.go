package ali

import "context"

// TokenSource yields the access token for the admin-owned 1688 account. Every
// call this platform makes runs as that one account.
//
// Today it is a value from the environment. The real OAuth flow is not
// documented anywhere in 1688-api-docs — no authorize URL, no token endpoint, no
// refresh shape — so it has to be written against the live documentation on the
// day access is granted. When that happens, a refreshing implementation lands
// here and no call site changes.
type TokenSource interface {
	Token(ctx context.Context) (string, error)
}

// StaticToken is a fixed token read from configuration.
type StaticToken string

func (s StaticToken) Token(context.Context) (string, error) { return string(s), nil }
