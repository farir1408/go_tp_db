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
