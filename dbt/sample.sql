-- name: GetMember :one
select * from members where memid = $1
and firstname = $2
;
