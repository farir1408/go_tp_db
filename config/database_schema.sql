-- DATABASE SCHEMA FOR TECHNOPARK DATABASE PROJECT
DROP TABLE IF EXISTS users, forum, post, thread, vote CASCADE;

CREATE EXTENSION IF NOT EXISTS citext;

-- USER
CREATE TABLE IF NOT EXISTS users (
  about       TEXT            NOT NULL,
  email       CITEXT UNIQUE   NOT NULL,
  fullname    TEXT            NOT NULL,
  nickname    CITEXT UNIQUE   PRIMARY KEY
);

-- FORUM
CREATE TABLE IF NOT EXISTS forum (
  id          SERIAL          PRIMARY KEY,
  posts       BIGINT          NOT NULL DEFAULT 0,
  slug        CITEXT UNIQUE   NOT NULL,
  threads     INTEGER         NOT NULL DEFAULT 0,
  title       TEXT            NOT NULL,
  author      CITEXT          NOT NULL REFERENCES users(nickname)
);

-- POST
CREATE TABLE IF NOT EXISTS post (
  id          SERIAL          PRIMARY KEY,
  author      CITEXT          NOT NULL REFERENCES users(nickname),
  created     TIMESTAMPTZ,
  forum       CITEXT          NOT NULL REFERENCES forum(slug),
  isEdited    BOOLEAN         NOT NULL DEFAULT FALSE,
  message     TEXT            NOT NULL,
  parent      BIGINT          DEFAULT 0,
  thread      INTEGER         NOT NULL,
  slug        CITEXT,
  parentId    BIGINT
);

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

-- VOTE
CREATE TABLE IF NOT EXISTS vote (
  id          INTEGER         NOT NULL REFERENCES thread(id),
  voice       SMALLINT        NOT NULL ,
  nickname    CITEXT          NOT NULL REFERENCES users(nickname),
  UNIQUE (id, nickname)
);