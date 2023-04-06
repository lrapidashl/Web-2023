CREATE TABLE post
(
    post_id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(255) NOT NULL,
    subtitle VARCHAR(255) NOT NULL,
    author VARCHAR(255),
    author_url VARCHAR(255),
    publish_date VARCHAR(255),
    image_url VARCHAR(255),
    featured TINYINT DEFAULT 0,
    PRIMARY KEY(post_id)
) CHARACTER SET = utf8mb4 COLLATE utf8mb4_unicode_ci
;  