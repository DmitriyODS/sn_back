BEGIN;

-- Таблица хранения user'ов
CREATE TABLE IF NOT EXISTS users
(
       id       SERIAL  PRIMARY KEY,
       login    TEXT    NOT NULL,
       password TEXT    NOT NULL
);

-- Заполнение таблицы данными
INSERT INTO users (id, login, password)
VALUES(1,'admin', 'admin')
ON CONFLICT (id) DO UPDATE SET id=excluded.id,
                               login=excluded.login,
                               password=excluded.password;


-- Настройка последовательности
SELECT setval('users_id_seq', (SELECT max(id) FROM users));


-- Таблица хранения Post'ов
CREATE TABLE IF NOT EXISTS posts
(
       id       SERIAL  PRIMARY KEY,
       title    TEXT    NOT NULL,
       text     TEXT    DEFAULT '',
       user_id  INTEGER NOT NULL,

       CONSTRAINT fk_from_users_to_posts_by_user_id FOREIGN KEY (user_id) REFERENCES users(id)
            ON DELETE CASCADE ON UPDATE CASCADE
);

-- Таблица хранения Like'ов
CREATE TABLE IF NOT EXISTS likes
(
       user_id      INTEGER NOT NULL,
       post_id      INTEGER NOT NULL,

       CONSTRAINT pk_likes PRIMARY KEY(user_id, post_id),
       CONSTRAINT fk_from_users_to_likes_by_user_id FOREIGN KEY (user_id) REFERENCES users(id)
            ON DELETE CASCADE ON UPDATE CASCADE,
       CONSTRAINT fk_from_posts_to_likes_by_post_id FOREIGN KEY (post_id) REFERENCES posts(id)
            ON DELETE CASCADE ON UPDATE CASCADE
);

COMMIT;
