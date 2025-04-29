CREATE DATABASE IF NOT EXISTS `app` DEFAULT CHARACTER SET utf8;
USE `app`;

CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `api_key` varchar(100) DEFAULT NULL,
  `role` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `users` (`id`, `username`, `password`, `api_key`, `role`)
SELECT 1, 'admin', '$2a$10$QsI3reS1qZgrK2E4szIlw.QVrvKvqx9wFlzNvLMTnPXn/yLL.v8vC', 'EC5B8A1A-3F9A-4A1B-92B6-8C2E8F3D7D55', '100'
WHERE NOT EXISTS (SELECT 1 FROM `users` WHERE `username` = 'admin');
