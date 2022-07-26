package orm

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"testing"
)

func exec(fn func(db *sql.DB, ctx context.Context) error) error {
	db, err := sql.Open("sqlite3", "file:locked.sqlite?cache=shared")
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(10)
	defer func() {
		_ = db.Close()
	}()
	ctx := context.Background()
	return fn(db, ctx)
}

func SqlExec(fn func(db *sql.DB, ctx context.Context) error, t *testing.T) {
	err := exec(fn)
	if err != nil {
		t.Error(err)
	}
}

// https://github.com/golang/go/wiki/SQLInterface
func TestInert(t *testing.T) {
	SqlExec(func(db *sql.DB, ctx context.Context) error {
		result, err := db.ExecContext(ctx,
			"INSERT INTO user (id, name, age) VALUES ($0, $1, $2)",
			1,
			"gopher",
			27,
		)
		if err != nil {
			return err
		}
		affected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if affected != 1 {
			return fmt.Errorf("affect not 1")
		}
		return nil
	}, t)
}

func TestSelect(t *testing.T) {
	SqlExec(func(db *sql.DB, ctx context.Context) error {
		result, err := db.QueryContext(ctx,
			"select * from User where name=$1",
			"gopher",
		)
		if err != nil {
			return err
		}
		defer func() {
			_ = result.Close()
		}()
		for result.Next() {
			var id string
			var name string
			var age int
			if err := result.Scan(&id, &name, &age); err != nil {
				t.Fatal(err)
			}
			t.Logf("%s,%s,%d\n", id, name, age)
		}
		return nil
	}, t)
}

func TestPrepare(t *testing.T) {
	SqlExec(func(db *sql.DB, ctx context.Context) error {
		stmt, err := db.PrepareContext(ctx,
			"select * from User where name=$1",
		)
		if err != nil {
			return err
		}
		defer func() {
			_ = stmt.Close()
		}()
		result, err := stmt.Query("gopher")
		if err != nil {
			return err
		}
		for result.Next() {
			var id string
			var name string
			var age int
			if err := result.Scan(&id, &name, &age); err != nil {
				log.Fatal(err)
			}
			t.Logf("%s,%s,%d\n", id, name, age)
		}
		return nil
	}, t)
}

func TestTransation(t *testing.T) {
	SqlExec(func(db *sql.DB, ctx context.Context) error {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		result, err := tx.ExecContext(ctx,
			"update User set age=$0 where name=$1",
			5, "gopher",
		)
		if err != nil {
			err := tx.Rollback()
			return err
		}
		result, err = tx.ExecContext(ctx,
			"update User set age=$0 where name=$1",
			6, "goddher",
		)
		if err != nil {
			err := tx.Rollback()
			return err
		}
		affected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		err = tx.Commit()
		if err != nil {
			return err
		}
		t.Logf("affected: %d", affected)
		return nil
	}, t)
}
