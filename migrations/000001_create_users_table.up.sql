CREATE TABLE IF NOT EXISTS `users` (
  `id` char(36) NOT NULL DEFAULT '' COMMENT 'ユーザID',
  `username` varchar(255) NOT NULL DEFAULT '' COMMENT 'ユーザ名',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT 'メールアドレス',
  `password` varchar(255) NOT NULL COMMENT 'パスワード',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  `deleted_at` datetime DEFAULT NULL COMMENT '削除日時',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザ';
