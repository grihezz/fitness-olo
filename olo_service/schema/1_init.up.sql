CREATE TABLE IF NOT EXISTS users (
     id BIGINT NOT NULL AUTO_INCREMENT,
     email VARCHAR(35) NOT NULL UNIQUE,
     role VARCHAR(5) NOT NULL DEFAULT 'USER',
     password_hash VARCHAR(64) NOT NULL,
     date_register TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS widgets (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `description` VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS articles (
   `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
   `header` VARCHAR(100) NOT NULL,
   `body` varchar(10000) NULL
);

CREATE TABLE IF NOT EXISTS user_has_widget (
    `id_widget` INT NOT NULL,
    `id_user` BIGINT NOT NULL,
    FOREIGN KEY (id_user) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (id_widget) REFERENCES widgets(id) ON DELETE CASCADE,
    PRIMARY KEY (id_widget, id_user)
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



INSERT INTO `widgets` (`description`)
SELECT * FROM (SELECT 'Тестовый виджет 1') AS tmp
WHERE NOT EXISTS (
    SELECT id FROM `widgets` WHERE `id`= 1
) LIMIT 1;

INSERT INTO `widgets` (`description`)
SELECT * FROM (SELECT 'Тестовый виджет 2') AS tmp
WHERE NOT EXISTS (
    SELECT id FROM `widgets` WHERE `id`= 2
) LIMIT 1;

INSERT INTO `widgets` (`description`)
SELECT * FROM (SELECT 'Тестовый виджет 3') AS tmp
WHERE NOT EXISTS (
    SELECT id FROM `widgets` WHERE `id`= 3
) LIMIT 1;

