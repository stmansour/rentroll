package rlib

import (
	"context"
	"database/sql"
)

type ctxKey int

const (
	sessionCtxKey ctxKey = 0
	dbTxCtxKey    ctxKey = iota
)

// SetSessionContextKey set the session in the given context object
// and returns new context with session
func SetSessionContextKey(ctx context.Context, s *Session) context.Context {
	return context.WithValue(ctx, sessionCtxKey, s)
}

// SessionFromContext extracts session from the given context
// with flag indicating whether session found or not
func SessionFromContext(ctx context.Context) (*Session, bool) {
	sess, ok := ctx.Value(sessionCtxKey).(*Session)
	return sess, ok
}

// SetDBTxContextKey set the session in the given context object
// and returns new context with sql.Tx
func SetDBTxContextKey(ctx context.Context, t *sql.Tx) context.Context {
	return context.WithValue(ctx, dbTxCtxKey, t)
}

// DBTxFromContext extracts sql.Tx from the given context
// with flag indicating whether sql.Tx found or not
func DBTxFromContext(ctx context.Context) (*sql.Tx, bool) {
	tx, ok := ctx.Value(dbTxCtxKey).(*sql.Tx)
	return tx, ok
}
