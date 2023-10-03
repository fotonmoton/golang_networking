package main

const SESSION_COOKIE = "session"

const CONTEXT_AUTH_KEY = "authenticated"

var authenticated map[string]bool = make(map[string]bool)

var registered map[login]password = make(map[login]password)
