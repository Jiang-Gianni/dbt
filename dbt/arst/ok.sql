-- test: TestGetMember
-- GetMember(1)

-- test: TestGetMemberA
-- GetMember('A')

-- test: TestGetMember2
-- GetMember(2)

-- name: GetFacilities
select * from facilities where facid = any($1);

-- test: TestFacilities
-- GetFacilities({1,2,3})