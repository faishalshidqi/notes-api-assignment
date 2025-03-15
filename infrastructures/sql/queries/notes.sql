-- name: CreateNote :exec
insert into notes(id, title, description, created_at, updated_at, owner) values(?, ?, ?, current_timestamp, current_timestamp, ?);

-- name: GetNotes :many
select notes.id, notes.title, notes.description, notes.created_at, notes.updated_at, notes.owner from notes join users on notes.owner = users.id where notes.owner = ?;

-- name: GetNote :one
select * from notes where id = ?;

-- name: EditNote :exec
update notes set title = ?, description = ?, updated_at = current_timestamp where id = ?;

-- name: DeleteNote :exec
delete from notes where id = ?;
