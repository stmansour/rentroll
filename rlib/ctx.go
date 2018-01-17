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

// NewTransactionWithContext returns newly created sql.Tx object
// and it also embeds that in the provided context and returns newly updated ctx
func NewTransactionWithContext(ctx context.Context) (*sql.Tx, context.Context, error) {
	var (
		tx  *sql.Tx
		err error
	)

	// get the new transaction
	tx, err = RRdb.Dbrr.Begin()
	if err != nil {
		return tx, ctx, err
	}

	// set the transaction in context
	ctx = SetDBTxContextKey(ctx, tx)

	return tx, ctx, err
}
