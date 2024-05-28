CREATE TABLE IF NOT EXISTS users (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `email` VARCHAR(35) NOT NULL UNIQUE,
    `role` VARCHAR(5) NOT NULL DEFAULT 'USER',
    `password_hash` VARCHAR(64) NOT NULL,
    `date_register` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS widgetsUser (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `data` VARCHAR(100) NOT NULL,
    `id_user` BIGINT NOT NULL,
    PRIMARY KEY (id_user)
);

CREATE TABLE IF NOT EXISTS articles (
   `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
   `header` VARCHAR(100) NOT NULL,
   `body` varchar(10000) NULL
);

CREATE TABLE IF NOT EXISTS user_has_articles (
     `id_user`     BIGINT NOT NULL,
     `id_articles` INT NOT NULL,
     PRIMARY KEY (`id_user`, `id_articles`),
     FOREIGN KEY (`id_user`) REFERENCES users (`id`) ON DELETE CASCADE,
     FOREIGN KEY (`id_articles`) REFERENCES articles (`id`) ON DELETE CASCADE
);

INSERT INTO articles (header, body) VALUES ('Пример заголовка 1', 'Это пример тела статьи. 1');
INSERT INTO articles (header, body) VALUES ('Пример заголовка 2', 'Это пример тела статьи. 2');
INSERT INTO articles (header, body) VALUES ('Пример заголовка 3', 'Это пример тела статьи. 3');
INSERT INTO articles (header, body) VALUES ('Пример заголовка 4', 'Это пример тела статьи. 4');