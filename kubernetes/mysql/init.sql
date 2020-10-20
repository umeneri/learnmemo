CREATE DATABASE IF NOT EXISTS gin;
CREATE DATABASE IF NOT EXISTS gin_test;

USE gin;
DROP TABLE IF EXISTS `task`;
DROP TABLE IF EXISTS `user`;

CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint(20) unsigned NOT NULL auto_increment,
  `email` varchar(255) NOT NULL UNIQUE,
  `name` varchar(255) NOT NULL,
  `provider_id` varchar(100) NOT NULL UNIQUE,
  `avatar_url` varchar(255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_email_key` (`email`)
) DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC AUTO_INCREMENT=1;

CREATE TABLE IF NOT EXISTS `task` (
  `id` bigint(20) unsigned NOT NULL auto_increment,
  `user_id` bigint(20) unsigned NOT NULL,
  `title` varchar(60) NOT NULL,
  `elapsed_time` bigint(20) NOT NULL,
  `status` tinyint NOT NULL,
  `created_at` TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `user`(`id`),
  KEY `idx_user_id_key` (`user_id`)
) DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC AUTO_INCREMENT=1;


ALTER TABLE user AUTO_INCREMENT = 1;
ALTER TABLE task AUTO_INCREMENT = 1;

insert into user (email, name, provider_id, avatar_url) values ('user1@gmail.com', 'user1', '1', '');
insert into user (email, name, provider_id, avatar_url) values ('user2@gmail.com', 'user2', '2', '');
insert into user (email, name, provider_id, avatar_url) values ('user3@gmail.com', 'user3', '3', '');
insert into user (email, name, provider_id, avatar_url) values ('user4@gmail.com', 'user4', '4', '');
insert into user (email, name, provider_id, avatar_url) values ('user5@gmail.com', 'user5', '5', '');
insert into user (email, name, provider_id, avatar_url) values ('user6@gmail.com', 'user6', '6', '');
insert into user (email, name, provider_id, avatar_url) values ('user7@gmail.com', 'user7', '7', '');
insert into user (email, name, provider_id, avatar_url) values ('user8@gmail.com', 'user8', '8', '');
insert into user (email, name, provider_id, avatar_url) values ('user9@gmail.com', 'user9', '9', '');

insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 10, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 10, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 20, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 30, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 40, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 50, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 60, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 70, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 80, 1);
insert into task (user_id, title, elapsed_time, status) values (1, '腹筋', 90, 1);

