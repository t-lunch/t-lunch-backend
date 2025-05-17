package models

type TokenType int

const (
	Access TokenType = iota
	Refresh
)
