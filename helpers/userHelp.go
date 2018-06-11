package helpers

const SelectUser = `
SELECT about, email, fullname, nickname 
FROM users
WHERE nickname = $1 OR email = $2`

const CreateUser = `
INSERT INTO users 
	(about, email, fullname, nickname)
	VALUES($1, $2, $3, $4) ON CONFLICT DO NOTHING`

const SelectUserProfile = `
SELECT about, email, fullname, nickname 
FROM users
WHERE nickname = $1`

const UpdateUser = `
UPDATE users SET 
	about = coalesce(nullif($1, ''), about), 
	email = coalesce(nullif($2, ''), email), 
	fullname = coalesce(nullif($3, ''), fullname)
WHERE nickname = $4
RETURNING
	about,
	email,
	fullname,
	nickname`

const UpdateForumUsers = `
UPDATE forum_users SET
	about = coalesce(nullif($1, ''), about), 
	email = coalesce(nullif($2, ''), email), 
	fullname = coalesce(nullif($3, ''), fullname)
WHERE nickname = $4`

const SelectUsersSinceDesc = `
SELECT DISTINCT about, email, 
	fullname, nickname 
FROM forum_users
WHERE forum_slug = $1 AND nickname < $2
ORDER BY nickname DESC
LIMIT $3::TEXT::INTEGER`

const SelectUsersSince = `
SELECT DISTINCT about, email, 
	fullname, nickname 
FROM forum_users
WHERE forum_slug = $1 AND nickname > $2
ORDER BY nickname
LIMIT $3::TEXT::INTEGER`

const SelectUsers = `
SELECT DISTINCT about, email, 
	fullname, nickname 
FROM forum_users
WHERE forum_slug = $1
ORDER BY nickname
LIMIT $2::TEXT::INTEGER`

const SelectUsersDesc = `
SELECT DISTINCT about, email, 
	fullname, nickname 
FROM forum_users
WHERE forum_slug = $1
ORDER BY nickname DESC
LIMIT $2::TEXT::INTEGER`
