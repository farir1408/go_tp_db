package helpers

const CreateVoteById = `INSERT INTO vote(voice, nickname, id) VALUE($1, $2, $3)
						UPDATE thread SET votes = (SUM(voice) FROM vote WHERE id = $3)
						WHERE id = $3`
