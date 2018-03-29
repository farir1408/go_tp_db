package helpers

const SelectThreadIdForumId = `SELECT id, forum FROM thread
								WHERE slug = $1`

const SelectThreadIdForumIdByID = `SELECT id, forum FROM thread
									WHERE id = $1`
