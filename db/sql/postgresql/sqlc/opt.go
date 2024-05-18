package db

import "github.com/jackc/pgx/v5"

type TxOptions func(*pgx.TxOptions)

// NewTxOPt returns a new pgx.TxOptions with the default isolation level
// if no options are provided.
// The default isolation level is pgx.RepeatableRead.
func NewTxOPt(opts ...TxOptions) pgx.TxOptions {
	if len(opts) == 0 {
		return pgx.TxOptions{}
	}

	txOpt := pgx.TxOptions{}

	for _, opt := range opts {
		opt(&txOpt)
	}

	return txOpt
}

func WithIsolationLevel(isoLevel pgx.TxIsoLevel) TxOptions {
	return func(opt *pgx.TxOptions) {
		opt.IsoLevel = isoLevel
	}
}

func WithAccessMode(accessMode pgx.TxAccessMode) TxOptions {
	return func(opt *pgx.TxOptions) {
		opt.AccessMode = accessMode
	}
}

func WithDeferrable(deferrableMode pgx.TxDeferrableMode) TxOptions {
	return func(opt *pgx.TxOptions) {
		opt.DeferrableMode = deferrableMode
	}
}

func WithBeginQuery(beginQuery string) TxOptions {
	return func(opt *pgx.TxOptions) {
		opt.BeginQuery = beginQuery
	}
}
