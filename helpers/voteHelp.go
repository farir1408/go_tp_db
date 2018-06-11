package helpers

const CreateVoteById = `
INSERT INTO vote(voice, nickname, id) 
	VALUES($1, $2, $3) ON CONFLICT DO NOTHING`

const UpdateVoteById = `
UPDATE vote SET old_voice = voice, voice = $1 
WHERE id = $2 AND nickname = $3 
RETURNING 
	(voice - old_voice)`

const UpdateThreadVotes = `
UPDATE thread SET 
	votes = votes + $1
WHERE id = $2`

const CreateVoteIdBySlug = `
SELECT id 
FROM thread 
WHERE slug = $1`
