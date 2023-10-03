package middleware

const SESSION_COOKIE = "session"

const CONTEXT_AUTH_KEY = "authenticated"

var authenticated map[string]bool = make(map[string]bool)
