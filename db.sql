# 文章表
CREATE DATABASE IF NOT EXISTS `jianshu` DEFAULT CHARACTER SET utf8mb4;

USE `jianshu`;

CREATE TABLE IF NOT EXISTS `js_article` (
    `aid` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `sn` varchar(32) NOT NULL DEFAULT '' COMMENT '文章序号',
    `title` varchar(127) NOT NULL DEFAULT '' COMMENT '文章标题',
    `uid` int unsigned NOT NULL DEFAULT '' COMMENT '文章作者 UID',
    `cover` varchar(255) NOT NULL DEFAULT '' COMMENT '文章封面图',
    `content` longtext NOT NULL COMMENT '内容，markdown 格式',
    `tags` varchar(50) NOT NULL DEFAULT '' COMMENT '文章 tag，逗号分隔',
    `state` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态：0-未审核;1-已上线;2-下线(审核拒绝);3-用户删除',
    `created_at`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`aid`),
    UNIQUE KEY `sn` (`sn`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章表'