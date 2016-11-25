CREATE DATABASE social_network;
CREATE TABLE users (
    user_id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    user_nickname VARCHAR(32) NOT NULL,
    user_first VARCHAR(32) NOT NULL,
    user_last VARCHAR(32) NOT NULL,
    user_email VARCHAR(128) NOT NULL,
    PRIMARY KEY (user_id),
    UNIQUE INDEX user_nickname (user_nickname)
);