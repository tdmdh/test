package auth // v1
// v2
// v3
// v4
// v5

// TokenTTL is how long an access token lives before it must be refreshed.
const TokenTTL = 15 * 60 // seconds, was 60 minutes

// RefreshWindow is how long before expiry a client may refresh its token.
const RefreshWindow = 5 * 60 // seconds; access tokens now live 10 minutes (was 60)

// SessionLimit caps concurrent sessions per user.
const SessionLimit = 5
