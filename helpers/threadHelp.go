package helpers

const SelectThreadBySlug = `SELECT id, author, created, forum, message, slug,
							title, votes FROM thread
							WHERE slug = $1`

const SelectThreadById = `SELECT id, author, created, forum, message, slug,
							title, votes FROM thread
							WHERE id = $1`

const UpdateThreadBySlug = `UPDATE thread SET 
						message = coalesce(coalesce(nullif($1, ''), message)),
						title = coalesce(coalesce(nullif($2, ''), title))
					  WHERE slug = $3`

const UpdateThreadById = `UPDATE thread SET 
						message = coalesce(coalesce(nullif($1, ''), message)),
						title = coalesce(coalesce(nullif($2, ''), title))
					  WHERE id = $3`

const SelectPostsSinceFlat = `SELECT author, created, forum, id, isEdited,
						message, parent, thread FROM post
						WHERE thread = $1
						AND id > $2
						ORDER BY id
						LIMIT $3::TEXT::INTEGER`

const SelectPostsSinceFlatDesc = `SELECT author, created, forum, id, isEdited,
						message, parent, thread FROM post
						WHERE thread = $1
						AND id < $2
						ORDER BY id DESC
						LIMIT $3::TEXT::INTEGER`

const SelectPostsFlat = `SELECT author, created, forum, id, isEdited,
						message, parent, thread FROM post
						WHERE thread = $1
						ORDER BY id
						LIMIT $2::TEXT::INTEGER`

const SelectPostsSinceTree = `SELECT author, created, forum, id, isEdited,
						message, parent, thread FROM post
						WHERE thread = $1
						AND id > $2
						ORDER BY parentId, id
						LIMIT $3::TEXT::INTEGER`

const SelectPostsSinceTreeDesc = `SELECT author, created, forum, id, isEdited,
						message, parent, thread FROM post
						WHERE thread = $1
						AND id < $2
						ORDER BY parentId, id DESC
						LIMIT $3::TEXT::INTEGER`

const SelectPostsFlatDesc = `SELECT author, created, forum, id, isEdited,
						message, parent, thread FROM post
						WHERE thread = $1
						ORDER BY id DESC
						LIMIT $2::TEXT::INTEGER`

const SelectPostsTree = `SELECT author, created, forum, id, isEdited,
						message, parent, thread FROM post
						WHERE thread = $1
						ORDER BY parentId, id
						LIMIT $2::TEXT::INTEGER`

const SelectPostsTreeDesc = `SELECT author, created, forum, id, isEdited,
						message, parent, thread FROM post
						WHERE thread = $1
						ORDER BY (parentId, id) DESC
						LIMIT $2::TEXT::INTEGER`
