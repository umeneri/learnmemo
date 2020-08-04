DROP TABLE IF EXISTS `user`;
DROP TABLE IF EXISTS `task`;

CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint(20) unsigned NOT NULL auto_increment,
  `email` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `provider_id` varchar(100) NOT NULL,
  `avatar_url` varchar(255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_email_key` (`email`)
) DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;

CREATE TABLE IF NOT EXISTS `task` (
  `id` bigint(20) unsigned NOT NULL auto_increment,
  `user_id` bigint(20) NOT NULL REFERENCES `users`(`id`),
  `title` varchar(60) NOT NULL,
  `elapsed_time` bigint(20) NOT NULL,
  `status` tinyint NOT NULL,
  `created_at` TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_title_key` (`title`)
) DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;
