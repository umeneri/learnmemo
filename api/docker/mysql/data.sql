TRUNCATE TABLE  `user`;
TRUNCATE TABLE `task`;

-- CREATE TABLE IF NOT EXISTS `user` (
--   `id` bigint(20) unsigned NOT NULL auto_increment,
--   `email` varchar(255) NOT NULL,
--   `name` varchar(255) NOT NULL,
--   `provider_id` varchar(100) NOT NULL,
--   `avatar_url` varchar(255) NOT NULL,
--   `created_at` TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
--   `updated_at` TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
--   PRIMARY KEY (`id`),
--   KEY `idx_email_key` (`email`)
-- ) DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

-- CREATE TABLE IF NOT EXISTS `task` (
--   `id` bigint(20) unsigned NOT NULL auto_increment,
--   `user_id` bigint(20) NOT NULL REFERENCES `users`(`id`),
--   `title` varchar(60) NOT NULL,
--   `progress_minute` bigint(20) NOT NULL,
--   `status` tinyint NOT NULL,
--   `created_at` TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
--   `updated_at` TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
--   PRIMARY KEY (`id`),
--   KEY `idx_title_key` (`title`)
-- ) DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

insert into user (email, name, provider_id, avatar_url) values ('user0@gmail.com', 'user0', '0', '');
insert into user (email, name, provider_id, avatar_url) values ('user1@gmail.com', 'user1', '1', '');
insert into user (email, name, provider_id, avatar_url) values ('user2@gmail.com', 'user2', '2', '');
insert into user (email, name, provider_id, avatar_url) values ('user3@gmail.com', 'user3', '3', '');
insert into user (email, name, provider_id, avatar_url) values ('user4@gmail.com', 'user4', '4', '');
insert into user (email, name, provider_id, avatar_url) values ('user5@gmail.com', 'user5', '5', '');
insert into user (email, name, provider_id, avatar_url) values ('user6@gmail.com', 'user6', '6', '');
insert into user (email, name, provider_id, avatar_url) values ('user7@gmail.com', 'user7', '7', '');
insert into user (email, name, provider_id, avatar_url) values ('user8@gmail.com', 'user8', '8', '');
insert into user (email, name, provider_id, avatar_url) values ('user9@gmail.com', 'user9', '9', '');

insert into task (user_id, title, progress_minute, status) values (1, '腹筋', 0, 1);
insert into task (user_id, title, progress_minute, status) values (1, '腹筋', 1, 1);
insert into task (user_id, title, progress_minute, status) values (1, '腹筋', 2, 1);
insert into task (user_id, title, progress_minute, status) values (1, '腹筋', 3, 1);
insert into task (user_id, title, progress_minute, status) values (1, '腹筋', 4, 1);
insert into task (user_id, title, progress_minute, status) values (1, '腹筋', 5, 1);
insert into task (user_id, title, progress_minute, status) values (1, '腹筋', 6, 1);
insert into task (user_id, title, progress_minute, status) values (1, '腹筋', 7, 1);
insert into task (user_id, title, progress_minute, status) values (1, '腹筋', 8, 1);
insert into task (user_id, title, progress_minute, status) values (1, '腹筋', 9, 1);



