-- name: CreateTransfer :one
INSERT INTO
    transfers (
        from_account_id,
        to_account_id,
        amount,
        created_at
    )
VALUES
    ($1, $2, $3, $4) RETURNING *;


-- name: ListTransfer :many
SELECT * FROM transfers
WHERE from_account_id = $1 OR
    to_account_id = $2
Order by  id   
LIMIT $3
OFFSET $4;

-- name: GetTransfer :one
select * from transfers
where id = $1;
