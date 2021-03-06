// Code generated by sqlc. DO NOT EDIT.
// source: link.sql

package db

import (
	"context"
)

const createLink = `-- name: CreateLink :one
INSERT INTO links (
  title, address, user_id
) VALUES (
  $1, $2, $3
)
RETURNING id, title, address, user_id
`

type CreateLinkParams struct {
	Title   string `json:"title"`
	Address string `json:"address"`
	UserID  int64  `json:"user_id"`
}

func (q *Queries) CreateLink(ctx context.Context, arg CreateLinkParams) (Link, error) {
	row := q.queryRow(ctx, q.createLinkStmt, createLink, arg.Title, arg.Address, arg.UserID)
	var i Link
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Address,
		&i.UserID,
	)
	return i, err
}

const getLinks = `-- name: GetLinks :many
SELECT links.id, links.title, links.address, links.user_id, users.username
FROM links INNER JOIN users
ON links.user_id = users.id
`

type GetLinksRow struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Address  string `json:"address"`
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
}

func (q *Queries) GetLinks(ctx context.Context) ([]GetLinksRow, error) {
	rows, err := q.query(ctx, q.getLinksStmt, getLinks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLinksRow
	for rows.Next() {
		var i GetLinksRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Address,
			&i.UserID,
			&i.Username,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
