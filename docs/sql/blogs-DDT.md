
1. ###### users
```sql
-- blog_db.users definition

CREATE TABLE `users` (
                         `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '键',
                         `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
                         `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
                         `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '账号',
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```
2. ###### blog_blogs
```sql
-- blog_db.blog_blogs definition

CREATE TABLE `blog_blogs` (
                              `id` int NOT NULL AUTO_INCREMENT COMMENT '键',
                              `group_id` int DEFAULT NULL COMMENT '分组Id',
                              `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '文章标题',
                              `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '文章内容',
                              `visitors` int DEFAULT '0' COMMENT '浏览量',
                              `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                              `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                              `is_show` tinyint(1) DEFAULT '1' COMMENT '是否显示',
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```
3. ###### blog_groups
```sql
-- blog_db.blog_groups definition

CREATE TABLE `blog_groups` (
                               `id` int NOT NULL AUTO_INCREMENT COMMENT '键',
                               `group` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '分组名称',
                               `is_show` tinyint(1) DEFAULT '1' COMMENT '是否显示',
                               `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                               `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```