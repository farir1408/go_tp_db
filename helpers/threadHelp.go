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

const SelectPostsFlatDesc = `SELECT author, created, forum, id, isEdited,
						message, parent, thread FROM post
						WHERE thread = $1
						ORDER BY id DESC
						LIMIT $2::TEXT::INTEGER`

const SelectPostsSinceTree = `SELECT author, created, forum, id, isEdited,
						message, parent, thread FROM post
						WHERE thread = $1
						AND (parentId > (select parentId from post where id = $2))
						ORDER BY parentId
						LIMIT $3::TEXT::INTEGER`

const SelectPostsSinceTreeDesc = `SELECT author, created, forum, id, isEdited,
						message, parent, thread FROM post
						WHERE thread = $1
						AND (parentId < (select parentId from post where id = $2))
						ORDER BY parentId DESC
						LIMIT $3::TEXT::INTEGER`

const SelectPostsTree = `SELECT author, created, forum, id, isEdited,
						message, parent, thread FROM post
						WHERE thread = $1
						ORDER BY parentId
						LIMIT $2::TEXT::INTEGER`

const SelectPostsTreeDesc = `SELECT author, created, forum, id, isEdited,
						message, parent, thread FROM post
						WHERE thread = $1
						ORDER BY parentId DESC
						LIMIT $2::TEXT::INTEGER`

const SelectPostsSinceParentTreeLimit = `SELECT author, created, forum, id, isEdited,
						message, parent, thread
						FROM post p JOIN
						(SELECT id AS idd from post WHERE parent = 0 AND thread = $1
						AND parentId[1] > (select parentId[1] From post WHERE id = $3)
						ORDER BY id LIMIT $2::TEXT::INTEGER) s
 						ON p.parentId[1] = s.idd ORDER BY idd, parentId`
//const SelectPostsSinceParentTreeLimit = `select author, created, forum, id, isEdited,
//  message, parent, thread
//FROM post where parentid[1] in (select id from post where thread = $1 and parent = 0
//and id > (select parentid[1] from post where id = $3) order by id asc limit $2::INTEGER) order by parentid`

const SelectPostsSinceParentTreeLimitDesc = `SELECT author, created, forum, id, isEdited,
						message, parent, thread
						FROM post p JOIN
						(SELECT id AS idd from post WHERE parent = 0 AND thread = $1
						AND parentId[1] < (select parentId[1] From post WHERE id = $3)
						ORDER BY id DESC LIMIT $2::TEXT::INTEGER) s
 						ON p.parentId[1] = s.idd ORDER BY idd DESC, parentId`
//const SelectPostsSinceParentTreeLimitDesc = `select author, created, forum, id, isEdited,
//  message, parent, thread
//FROM post where parentid[1] in (select id from post where thread = $1 and parent = 0
//and id < (select parentid[1] from post where id = $3) order by id desc limit $2::INTEGER) order by parentid`

const SelectPostsParentTree = `SELECT author, created, forum, id, isEdited,
						message, parent, thread FROM post
						WHERE thread = $1
						ORDER BY parentId
						LIMIT $2::TEXT::INTEGER`

const SelectPostsParentTreeDesc = `SELECT author, created, forum, id, isEdited,
						message, parent, thread FROM post
						WHERE thread = $1
						ORDER BY parentId DESC
						LIMIT $2::TEXT::INTEGER`

const SelectPostsParentTreeLimitDesc = `SELECT author, created, forum, id, isEdited,
						message, parent, thread
						FROM post p JOIN
						(SELECT id AS idd from post WHERE parent = 0 AND thread = $1
						ORDER BY id DESC LIMIT $2::TEXT::INTEGER) s
 						ON p.parentId[1] = s.idd ORDER BY idd DESC, parentId`
//const SelectPostsParentTreeLimitDesc = `select author, created, forum, id, isEdited,
//  message, parent, thread
//FROM post where parentid[1] in (select id from post where thread = $1 and parent = 0
//order by id desc limit $2::INTEGER) order by parentid`

const SelectPostsParentTreeLimit = `SELECT author, created, forum, id, isEdited,
						message, parent, thread
						FROM post p JOIN
						(SELECT id AS idd from post WHERE parent = 0 AND thread = $1
						ORDER BY id LIMIT $2::TEXT::INTEGER) s
 						ON p.parentId[1] = s.idd ORDER BY idd, parentId`
//const SelectPostsParentTreeLimit = `select author, created, forum, id, isEdited,
//  message, parent, thread
//FROM post where parentid[1] in (select id from post where thread = $1 and parent = 0
//order by id asc limit $2::INTEGER) order by parentid`

const SelectPostsSinceParentTreeDesc = `SELECT author, created, forum, id, isEdited,
						message, parent, thread
						FROM post p JOIN
						(SELECT id AS idd from post WHERE parent = 0 AND thread = $1
						AND parentId[1] < (select parentId[1] From post WHERE id = $2)
						ORDER BY id DESC) s
 						ON p.parentId[1] = s.idd ORDER BY idd DESC, parentId`
//const SelectPostsSinceParentTreeDesc = `select author, created, forum, id, isEdited,
//  message, parent, thread
//FROM post where parentid[1] in (select id from post where thread = $1 and parent = 0
//and id < (select parentid[1] from post where id = $3) order by id desc) order by parentid`

const SelectPostsSinceParentTree = `SELECT author, created, forum, id, isEdited,
						message, parent, thread
						FROM post p JOIN
						(SELECT id AS idd from post WHERE parent = 0 AND thread = $1
						AND parentId[1] > (select parentId[1] From post WHERE id = $2)
						ORDER BY id) s
 						ON p.parentId[1] = s.idd ORDER BY idd, parentId`
//const SelectPostsSinceParentTree = `select author, created, forum, id, isEdited,
//  message, parent, thread
//FROM post where parentid[1] in (select id from post where thread = $1 and parent = 0
//and id > (select parentid[1] from post where id = $3) order by id asc) order by parentid`
