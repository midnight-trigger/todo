CREATE TABLE IF NOT EXISTS `todos` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'TodoID',
  `user_id` char(36) NOT NULL DEFAULT '' COMMENT 'ユーザID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'タイトル',
  `body` text NOT NULL COMMENT '詳細',
  `status` enum('todo','progress','finished') NOT NULL DEFAULT 'todo' COMMENT 'ステータス（todo: 未進行 / progress: 進行中 / finished: 完了済み）',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  `deleted_at` datetime DEFAULT NULL COMMENT '削除日時',
  PRIMARY KEY (`id`),
  KEY `idx_todos_01` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Todo';
