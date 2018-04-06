package helpers

const GetStatus = `SELECT
				(SELECT COUNT(*) FROM forum) AS forum,
				(SELECT COUNT(*) FROM post) AS post,
				(SELECT COUNT(*) FROM thread) AS thread,
				(SELECT COUNT(*) FROM users) AS user`

				//`SELECT
				//(SELECT reltuples::bigint AS forum FROM pg_class WHERE oid = 'forum'::regclass),
				//(SELECT reltuples::bigint AS post FROM pg_class WHERE oid = 'post'::regclass),
				//(SELECT reltuples::bigint AS thread FROM pg_class WHERE oid = 'thread'::regclass),
				//(SELECT reltuples::bigint AS users FROM pg_class WHERE oid = 'users'::regclass)`

const ClearDB = `TRUNCATE post, thread, vote CASCADE`