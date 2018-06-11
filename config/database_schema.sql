-- DATABASE SCHEMA FOR TECHNOPARK DATABASE PROJECT
DROP TABLE IF EXISTS users, forum, post, thread, vote, forum_users CASCADE;

CREATE EXTENSION IF NOT EXISTS citext;

-- USER
CREATE TABLE IF NOT EXISTS users (
  about       TEXT                NOT NULL,
  email       CITEXT UNIQUE       NOT NULL,
  fullname    TEXT                NOT NULL,
  nickname    CITEXT COLLATE "C"  UNIQUE   PRIMARY KEY
);

-- таблица для поиска всех пользователей форума
CREATE TABLE IF NOT EXISTS forum_users (
  about       TEXT                NOT NULL,
  email       CITEXT              NOT NULL,
  fullname    TEXT                NOT NULL,
  nickname    CITEXT COLLATE "C"  NOT NULL,
  forum_slug  CITEXT              NOT NULL
--   UNIQUE (nickname, forum_slug)
);

CREATE INDEX forum_users_nickname_idx ON forum_users (nickname);
CREATE INDEX forum_users_slug_idx ON forum_users (forum_slug);
CREATE INDEX forum_users_cover_idx ON forum_users (nickname, forum_slug);

CREATE INDEX users_cover_idx ON users (about, email, fullname, nickname);
CREATE INDEX users_nick_email_idx ON users (nickname, email);

-- FORUM
CREATE TABLE IF NOT EXISTS forum (
  id          SERIAL          PRIMARY KEY,
  posts       BIGINT          NOT NULL DEFAULT 0,
  slug        CITEXT UNIQUE   NOT NULL,
  threads     INTEGER         NOT NULL DEFAULT 0,
  title       TEXT            NOT NULL,
  author      CITEXT          NOT NULL REFERENCES users(nickname)
);

-- CREATE INDEX forum_slug_idx ON forum (slug);
CREATE INDEX forum_author_idx ON forum(author);
CREATE INDEX forum_cover_idx ON forum(id, posts, slug, threads, title, author);

-- POST
CREATE TABLE IF NOT EXISTS post (
  id          SERIAL          PRIMARY KEY,
  author      CITEXT          NOT NULL,
  created     TIMESTAMPTZ,
  forum       CITEXT          NOT NULL,
  isEdited    BOOLEAN         NOT NULL DEFAULT FALSE,
  message     TEXT            NOT NULL,
  parent      BIGINT          DEFAULT 0,
  thread      INTEGER         NOT NULL,
  root        INTEGER,
  slug        CITEXT,
  parentId    BIGINT []
);

CREATE INDEX post_thread_idx ON post(thread, id);
-- CREATE INDEX post_forum_idx ON post(forum);
CREATE INDEX post_root_id_idx ON post(id, root);
CREATE INDEX post_thread_idx ON post(thread);
CREATE INDEX post_parents_idx ON post(thread, id, parent, root) WHERE parent = 0;
CREATE INDEX post_root_idx ON post(root, parentId DESC, id);
CREATE INDEX post_thread_parents_idx ON post(thread, parentId);
-- CREATE INDEX post_parents_desc_idx ON post(thread, parentId DESC);

-- THREAD
CREATE TABLE IF NOT EXISTS thread (
  id          SERIAL          PRIMARY KEY,
  author      CITEXT          NOT NULL REFERENCES users(nickname),
  created     TIMESTAMPTZ(3),
  forum       CITEXT          NOT NULL REFERENCES forum(slug),
  message     TEXT            NOT NULL,
  slug        CITEXT,
  title       TEXT            NOT NULL,
  votes       INTEGER         DEFAULT 0
);

CREATE INDEX thread_author_idx ON thread(author);
CREATE INDEX thread_forum_idx ON thread(forum, created);
CREATE INDEX thread_forum_desc_idx ON thread(forum, created DESC);
CREATE INDEX thread_slug_idx ON thread(slug);
CREATE INDEX thread_slug_id_idx ON thread(slug, id);
CREATE INDEX thread_id_forum_idx ON thread(id, forum);
CREATE INDEX thread_slug_forum_idx ON thread(slug, forum);
-- CREATE INDEX thread_cover_idx ON thread(forum, id, author, created, message, slug, title, votes);


-- VOTE
CREATE TABLE IF NOT EXISTS vote (
  id          INTEGER         NOT NULL REFERENCES thread(id),
--   id          INTEGER         NOT NULL,
  voice       SMALLINT        NOT NULL,
  old_voice   SMALLINT        DEFAULT 0,
--   nickname    CITEXT          NOT NULL,
  nickname    CITEXT          NOT NULL REFERENCES users(nickname),
  UNIQUE (id, nickname)
);

CREATE INDEX vote_thread_idx ON vote(id);
CREATE INDEX vote_author_idx ON vote(nickname);
