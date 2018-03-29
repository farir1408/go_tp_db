package helpers

const SelectUser = `SELECT about, email, fullname, nickname FROM users
					WHERE nickname = $1 OR email = $2`

const CreateUser = `INSERT INTO users (about, email, fullname, nickname)
					VALUES($1, $2, $3, $4) ON CONFLICT DO NOTHING`

const SelectUserProfile = `SELECT about, email, fullname, nickname FROM users
							WHERE nickname = $1`

const UpdateUser = `UPDATE users SET about = $1, email = $2, fullname = $3
					WHERE nickname = $4`
