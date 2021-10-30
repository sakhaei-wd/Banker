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
Order by  Id
LIMIT $1
OFFSET $2;

-- name: GetTransfer :one
select * from transfers
where id = $1;

-- name: DeleteTransfer :exec
delete from transfers
where id = $1;

-- name: UpdateTransfer :one
update transfers
set amount = $1
where id = $2
RETURNING *;
