package helpers

const CreateVoteById = `INSERT INTO vote(voice, nickname, id) VALUES($1, $2, $3)`

const UpdateVoteById = `UPDATE vote SET voice = $1 WHERE id = $2 AND nickname = $3`

const UpdateThreadVotes = `UPDATE thread SET votes = (SELECT SUM(voice) FROM vote WHERE id = $1)
						WHERE id = $1`

const CreateVoteIdBySlug = `SELECT id FROM thread WHERE slug = $1`

const SelectThreadId = `SELECT id FROM thread WHERE id = $1`
