package helpers

const SelectUser = `SELECT about, email, fullname, nickname FROM users
					WHERE nickname = $1 OR email = $2`

const CreateUser = `INSERT INTO users (about, email, fullname, nickname)
					VALUES($1, $2, $3, $4) ON CONFLICT DO NOTHING`

const SelectUserProfile = `SELECT about, email, fullname, nickname FROM users
							WHERE nickname = $1`

const UpdateUser = `UPDATE users SET 
						about = coalesce(coalesce(nullif($1, ''), about)), 
						email = coalesce(coalesce(nullif($2, ''), email)), 
						fullname = coalesce(coalesce(nullif($3, ''), fullname))
					WHERE nickname = $4
					RETURNING
						about,
						email,
						fullname,
						nickname`

const SelectUsersSinceDesc = `SELECT about, email, fullname, nickname FROM users
								WHERE (((SELECT COUNT(*) FROM post 
								WHERE forum = $1 AND author = nickname) != 0)
								OR ((SELECT COUNT(*) FROM THREAD
								WHERE forum = $1 AND author = nickname) != 0))
								AND nickname < $2
								ORDER BY lower(nickname) DESC
								LIMIT $3::TEXT::INTEGER`

const SelectUsersSince = `SELECT about, email, fullname, nickname FROM users
								WHERE (((SELECT COUNT(*) FROM post 
								WHERE forum = $1 AND author = nickname) != 0)
								OR ((SELECT COUNT(*) FROM THREAD
								WHERE forum = $1 AND author = nickname) != 0))
								AND nickname > $2
								ORDER BY lower(nickname)
								LIMIT $3::TEXT::INTEGER`

const SelectUsers = `SELECT about, email, fullname, nickname FROM users
								WHERE ((SELECT COUNT(*) FROM post 
								WHERE forum = $1 AND author = nickname) != 0)
								OR ((SELECT COUNT(*) FROM THREAD
								WHERE forum = $1 AND author = nickname) != 0)
								ORDER BY lower(nickname)
								LIMIT $2::TEXT::INTEGER`

const SelectUsersDesc = `SELECT about, email, fullname, nickname FROM users
								WHERE ((SELECT COUNT(*) FROM post 
								WHERE forum = $1 AND author = nickname) != 0)
								OR ((SELECT COUNT(*) FROM THREAD
								WHERE forum = $1 AND author = nickname) != 0)
								ORDER BY lower(nickname) DESC
								LIMIT $2::TEXT::INTEGER`
