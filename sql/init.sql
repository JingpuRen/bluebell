create database bluebell;

USE bluebell;

CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email` varchar(64) COLLATE utf8mb4_general_ci,
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

create table `community`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `community_id` int(10) unsigned NOT NULL,
    `community_name` varchar(128) collate utf8mb4_general_ci not null,
    `introduction` varchar(256) collate utf8mb4_general_ci not null,
    `create_time` timestamp not null default CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_community_id` (`community_id`) USING BTREE,
    UNIQUE KEY `idx_community_name` (`community_name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 对于有默认值的字段，假如你只想以默认值来弄，那么就像下面这样子进行插入
INSERT INTO `community` (`community_id`, `community_name`, `introduction`)
VALUES (1, 'Go', 'Golang');
INSERT INTO `community` (`community_id`, `community_name`, `introduction`)
VALUES (2, 'leetcode', '刷题刷题');
INSERT INTO `community` (`community_id`, `community_name`, `introduction`)
VALUES (3, 'CS:GO', 'Rush B');
INSERT INTO `community` (`community_id`, `community_name`, `introduction`)
VALUES (4, 'LOL', '一剑诛恶一剑镇魂');

create table `post`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `post_id` bigint(20) NOT NULL COMMENT '帖子id',
    `title` varchar(128) collate utf8mb4_general_ci not null COMMENT '标题',
    `content` varchar(8192) collate utf8mb4_general_ci not null COMMENT '内容',
    `author_id` bigint(20) not null COMMENT '作者的用户id，也就是user_id',
    `community_id` bigint(20) not null COMMENT '帖子所属的社区',
    `status` tinyint not null DEFAULT 1 comment '帖子状态',
    `create_time` timestamp not null default CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_post_id` (`post_id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_community_id` (`community_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `post` (`post_id`, `title`, `content`,`author_id`,`community_id`)
VALUES (1, '怎么学GoLang?', '我们应该观看qlmi老师的视频！！','1','1');