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

ALTER TABLE users ADD UNIQUE INDEX user_email (user_email);

CREATE TABLE user_relationships (
    user_relationship_id INT(13) NOT NULL,
    from_user_id INT(10) NOT NULL,
    to_user_id INT(10) NOT NULL,
    user_relationship_type VARCHAR(10) NOT NULL,
    user_relationship_timestamp DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_relationship_id),
    INDEX from_user_id (from_user_id),
    INDEX to_user_id (to_user_id),
    INDEX from_user_id_to_user_id (from_user_id, to_user_id),
    INDEX from_user_id_to_user_id_user_relationship_type (from_user_id, to_user_id, user_relationship_type)
    
    
);

ALTER TABLE `users` ADD COLUMN `user_image` MEDIUMBLOB NOT NULL AFTER `user_email`;