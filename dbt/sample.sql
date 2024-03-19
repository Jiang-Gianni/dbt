-- name: GetMember :one
select * from members where memid = $1;
