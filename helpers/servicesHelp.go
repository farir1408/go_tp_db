package helpers

const ClearDB = `
TRUNCATE TABLE 
	vote, 
	post, 
	thread, 
	forum, 
	users 
CASCADE`

const GetStatus = `
SELECT
	(SELECT COUNT(*) FROM forum) AS forum,
	(SELECT COUNT(*) FROM post) AS post,
	(SELECT COUNT(*) FROM thread) AS thread,
	(SELECT COUNT(*) FROM users) AS user`
