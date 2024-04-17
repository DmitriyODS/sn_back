package pdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"idon.com/cfg"
)

const (
	CtxCurTx = "ctx_cur_tx"
)

var (
	ErrNotActiveTx = errors.New("there are no active transactions in the context")
)

type PDB struct {
	*pgxpool.Pool
}

var _pdb *PDB

func GetPDB(ctx context.Context) *PDB {
	if _pdb == nil {
		appConfig := cfg.GetAppConfig()

		connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s pool_max_conns=256",
			appConfig.PsqlLogin,
			appConfig.PsqlPass,
			appConfig.PsqlAddr,
			appConfig.PsqlPort,
			appConfig.PsqlDB)
		dbPool, err := pgxpool.New(ctx, connString)
		if err != nil {
			log.Fatalf("Err connect to psql: %v", err)
		}

		_pdb = &PDB{dbPool}
	}

	return _pdb
}

func ClosePDB() {
	if _pdb != nil {
		_pdb.Close()
	}
}

func (p *PDB) ContextWithTx(ctx context.Context) (context.Context, pgx.Tx) {
	tx, err := p.Begin(ctx)
	if err != nil {
		log.Errorf("Failed create new transaction: %v", err)
		return ctx, nil
	}

	return context.WithValue(ctx, CtxCurTx, tx), tx
}

// ExecTx takes the active transaction from the context and executes Exec based on it
// If there are no active transactions, returns an error
func (p *PDB) ExecTx(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	curTx, ok := ctx.Value(CtxCurTx).(pgx.Tx)
	if !ok {
		return pgconn.CommandTag{}, ErrNotActiveTx
	}

	return curTx.Exec(ctx, sql, arguments...)
}

// QueryTx takes the active transaction from the context and executes Query based on it
// If there are no active transactions, returns a nil and error
func (p *PDB) QueryTx(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	curTx, ok := ctx.Value(CtxCurTx).(pgx.Tx)
	if !ok {
		return nil, ErrNotActiveTx
	}

	return curTx.Query(ctx, sql, args...)
}

// QueryRowTx takes the active transaction from the context and executes QueryRow based on it
// If there are no active transactions, returns a nil and error
func (p *PDB) QueryRowTx(ctx context.Context, sql string, args ...any) (pgx.Row, error) {
	curTx, ok := ctx.Value(CtxCurTx).(pgx.Tx)
	if !ok {
		return nil, ErrNotActiveTx
	}

	return curTx.QueryRow(ctx, sql, args...), nil
}
