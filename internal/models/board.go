package models

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jmarren/toucan/internal/cache"
)

type Board struct {
	values []string
}

var emptyBoard = [64]string{
	"br", "bn", "bb", "bq", "bk", "bb", "bn", "br",
	"bp", "bp", "bp", "bp", "bp", "bp", "bp", "bp",
	"", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "",
	"wp", "wp", "wp", "wp", "wp", "wp", "wp", "wp",
	"wr", "wn", "wb", "wq", "wk", "wb", "wn", "wr",
}

func GetBoard(sessionId string) (*Board, error) {

	ctx := context.Background()

	pipeline := cache.Rdb.Pipeline()

	// set it to empty board if it doesn't exist
	pipeline.JSONSetMode(ctx, sessionId, "$", emptyBoard, "NX")
	sessionIdCmd := pipeline.JSONGet(ctx, sessionId)

	pipeline.Exec(ctx)

	res, err := sessionIdCmd.Result()
	if err != nil {
		fmt.Printf("sessionId err: %s\n", err)
		return nil, err
	}

	var values []string

	json.Unmarshal([]byte(res), &values)

	return &Board{
		values: values,
	}, nil

}

func (b *Board) GetSquare(row int, col int) string {
	i := (row * 8) + col
	return b.values[i]
}
