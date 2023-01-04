##创建数据库
drop schema if exists `shrimp_blog`;
create
    database `shrimp_blog` default charset utf8mb4;
use
    `shrimp_blog`;

##创建用户表
drop table IF exists `user`;
create table `user`
(
    `id`          int(11) unsigned    not null auto_increment comment '主键',
    `username`    varchar(30)         not null comment '用户名',
    `password`    varchar(255)        not null comment '密码',
    `email`       varchar(255)        not null comment '邮箱',
    `avatar`      varchar(255)        not null comment '头像',
    `create_time` datetime           NOT NULL DEFAULT now() COMMENT '创建时间',
    `update_time` datetime           NOT NULL DEFAULT now() COMMENT '更新时间',
    `delete_time` datetime                    DEFAULT NULL COMMENT '删除时间',
    `status`      tinyint(1) unsigned not null default '1' comment '状态',
    primary key (`id`),
    key `username` (`username`),
    unique key `email` (`email`)
) engine = InnoDB
  default charset = utf8mb4 comment ='用户表'
  collate = utf8mb4_general_ci;

##创建文章表
drop table IF exists `article`;
create table `article`
(
    `id`            int(11) unsigned    not null auto_increment comment '主键',
    `title`         varchar(255)        not null comment '标题',
    `content`       text                not null comment '内容',
    `summary`       varchar(200)                 DEFAULT NULL COMMENT '博客简介',
    `click_count`   int(11)                      DEFAULT '0' COMMENT '博客点击数',
    `collect_count` int(11)                      DEFAULT '0' COMMENT '博客收藏数',
    `file_uid`      varchar(255)                 DEFAULT NULL COMMENT '标题图片uid',
    `recommend`     tinyint(1) unsigned          DEFAULT '0' COMMENT '是否推荐',
    `create_time`   datetime           NOT NULL DEFAULT now() COMMENT '创建时间',
    `update_time`   datetime           NOT NULL DEFAULT now() COMMENT '更新时间',
    `delete_time`   datetime                    DEFAULT NULL COMMENT '删除时间',
    `status`        tinyint(1) unsigned not null default '1' comment '状态',
    primary key (`id`),
    key `article_file_uid` (`file_uid`)
) engine = InnoDB
  default charset = utf8mb4 comment ='文章表'
  collate = utf8mb4_general_ci;

#创建评论表
drop table IF exists `comment`;
create table `comment`
(
    `id`          int(11) unsigned    not null auto_increment comment '主键',
    `article_id`  int(11) unsigned    not null comment '文章id',
    `content`     text                not null comment '内容',
    `parent_id`   int(11) unsigned             DEFAULT NULL comment '父评论id',
    `user_id`     int(11) unsigned    not null comment '用户id',
    `username`    varchar(30)         not null comment '用户名',
    `create_time` datetime           NOT NULL DEFAULT now() COMMENT '创建时间',
    `update_time` datetime           NOT NULL DEFAULT now() COMMENT '更新时间',
    `delete_time` datetime                    DEFAULT NULL COMMENT '删除时间',
    `status`      tinyint(1) unsigned not null default '0' comment '状态',
    primary key (`id`),
    key (`article_id`),
    key (`user_id`),
    key (`parent_id`)
) engine = InnoDB
  default charset = utf8mb4 comment ='评论表'
  collate = utf8mb4_general_ci;

##创建标签表
drop table IF exists `label`;
create table `label`
(
    `id`          int(11) unsigned    not null auto_increment,
    `name`        varchar(255)        not null comment '标签名',
    `create_time` datetime           NOT NULL DEFAULT now() COMMENT '创建时间',
    `update_time` datetime           NOT NULL DEFAULT now() COMMENT '更新时间',
    `delete_time` datetime                    DEFAULT NULL COMMENT '删除时间',
    `status`      tinyint(1) unsigned not null default '1' comment '状态',
    primary key (`id`)
) engine = InnoDB
  default charset = utf8mb4 comment ='标签表'
  collate = utf8mb4_general_ci;

##创建文章标签表
drop table IF exists `article_label`;
create table `article_label`
(
    `id`          int(11) unsigned    not null auto_increment,
    `article_id`  int(11) unsigned    not null comment '文章id',
    `label_id`    int(11) unsigned    not null comment '标签id',
    `create_time` datetime           NOT NULL DEFAULT now() COMMENT '创建时间',
    `update_time` datetime           NOT NULL DEFAULT now() COMMENT '更新时间',
    `delete_time` datetime                    DEFAULT NULL COMMENT '删除时间',
    `status`      tinyint(1) unsigned not null default '1' comment '状态',
    primary key (`id`),
    key `article_label_article_id` (`article_id`),
    key `article_label_label_id` (`label_id`)
) engine = InnoDB
  default charset = utf8mb4 comment ='文章标签关联表'
  collate = utf8mb4_general_ci;

##创建专题表
drop table IF exists `subject`;
create table `subject`
(
    `id`          int(11) unsigned    not null auto_increment,
    `name`        varchar(255)        not null comment '专题名',
    `create_time` datetime           NOT NULL DEFAULT now() COMMENT '创建时间',
    `update_time` datetime           NOT NULL DEFAULT now() COMMENT '更新时间',
    `delete_time` datetime                    DEFAULT NULL COMMENT '删除时间',
    `status`      tinyint(1) unsigned not null default '1' comment '状态',
    primary key (`id`)
) engine = InnoDB
  default charset = utf8mb4 comment ='专题表'
  collate = utf8mb4_general_ci;

##创建专题文章关联表
drop table IF exists `subject_item`;
CREATE TABLE `subject_item`
(
    `id`          int(11) unsigned not null auto_increment,
    `subject_id`  int(11)      NOT NULL COMMENT '专题uid',
    `article_id`  int(11)       NOT NULL COMMENT '文章uid',
    `status`      tinyint(1)                DEFAULT '1' COMMENT '状态',
    `sort`        int(11)          NOT NULL DEFAULT '0' COMMENT '排序字段',
    `create_time` datetime        NOT NULL DEFAULT now() COMMENT '创建时间',
    `update_time` datetime        NOT NULL DEFAULT now() COMMENT '更新时间',
    `delete_time` datetime                 DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    key (`subject_id`),
    key (`article_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='专题Item表'
  collate = utf8mb4_general_ci;

##创建友情链接表
drop table IF exists `link`;
create table `link`
(
    `id`          int(11) unsigned    not null auto_increment,
    `name`        varchar(255)        not null comment '链接名',
    `url`         varchar(255)        not null comment '链接地址',
    `avatar`      varchar(255)        not null comment '链接图标',
    `create_time` datetime           NOT NULL DEFAULT now() COMMENT '创建时间',
    `update_time` datetime           NOT NULL DEFAULT now() COMMENT '更新时间',
    `delete_time` datetime                    DEFAULT NULL COMMENT '删除时间',
    `status`      tinyint(1) unsigned not null default '1' comment '状态',
    primary key (`id`)
) engine = InnoDB
  default charset = utf8mb4 comment ='友情链接表'
  collate = utf8mb4_general_ci;

##创建导航表
drop table IF exists `navbar`;
create table `navbar`
(
    `id`           int(11) unsigned    not null auto_increment,
    `name`         varchar(255)        not null comment '导航名',
    `url`          varchar(255)        not null comment '链接地址',
    `create_time`  datetime           NOT NULL DEFAULT now() COMMENT '创建时间',
    `update_time`  datetime           NOT NULL DEFAULT now() COMMENT '更新时间',
    `delete_time`  datetime                    DEFAULT NULL COMMENT '删除时间',
    `status`       tinyint(1) unsigned not null default '1' comment '状态',
    `navbar_level` tinyint(1) unsigned not null default '1' comment '导航级别',
    `parent_id`    int(11) unsigned    not null default '0' comment '父级id',
    `sort`         int(11) unsigned    not null default '0' comment '排序',
    primary key (`id`)
) engine = InnoDB
  default charset = utf8mb4 comment ='导航表'
  collate = utf8mb4_general_ci;

##创建页面表
drop table if exists `page_template`;
create table `page_template`
(
    `id`          int(11) unsigned    not null auto_increment,
    `name`        varchar(255)        not null comment '模板名',
    `create_time` datetime           NOT NULL DEFAULT now() COMMENT '创建时间',
    `update_time` datetime           NOT NULL DEFAULT now() COMMENT '更新时间',
    `delete_time` datetime                    DEFAULT NULL COMMENT '删除时间',
    `status`      tinyint(1) unsigned not null default '1' comment '状态',
    primary key (`id`)
) engine = InnoDB
  default charset = utf8mb4 comment ='页面模板表'
  collate = utf8mb4_general_ci;
drop table if exists `messages`;
create table `messages`
(
    `id`          int(11) unsigned    not null auto_increment,
    `username`    varchar(255)        not null comment '昵称',
    `content`     varchar(255)        not null comment '内容',
    `contact_way` varchar(255)        not null comment '联系途径',
    `contact`     varchar(255)        not null comment '联系方式',
    `create_time` datetime           NOT NULL DEFAULT now() COMMENT '创建时间',
    `update_time` datetime           NOT NULL DEFAULT now() COMMENT '更新时间',
    `delete_time` datetime                    DEFAULT NULL COMMENT '删除时间',
    `status`      tinyint(1) unsigned not null default '1' comment '状态',
    primary key (`id`)
) engine = InnoDB
  default charset = utf8mb4 comment ='留言表'
  collate = utf8mb4_general_ci;

##创建分类表
drop table IF exists `category`;
create table `category`
(
    `id`          int(11) unsigned    not null auto_increment,
    `name`        varchar(255)        not null comment '分类名',
    `create_time` datetime           NOT NULL DEFAULT now() COMMENT '创建时间',
    `update_time` datetime           NOT NULL DEFAULT now() COMMENT '更新时间',
    `delete_time` datetime                    DEFAULT NULL COMMENT '删除时间',
    `status`      tinyint(1) unsigned not null default '1' comment '状态',
    primary key (`id`)
) engine = InnoDB
  default charset = utf8mb4 comment ='分类表'
  collate = utf8mb4_general_ci;
##文章分类表
drop table IF exists `article_category`;
create table `article_category`
(
    `id`          int(11) unsigned    not null auto_increment,
    `article_id`  int(11)        not null comment '文章名',
    `category_id`     int(11)        not null comment '分类名',
    `create_time` datetime           NOT NULL DEFAULT now() COMMENT '创建时间',
    `update_time` datetime           NOT NULL DEFAULT now() COMMENT '更新时间',
    `delete_time` datetime                    DEFAULT NULL COMMENT '删除时间',
    `status`      tinyint(1) unsigned not null default '1' comment '状态',
    primary key (`id`)
) engine = InnoDB
  default charset = utf8mb4 comment ='文章分类表'
  collate = utf8mb4_general_ci;

drop table if exists `role`;
create table `role`
(
    `id`          int(11) unsigned    not null auto_increment,
    `name`        varchar(255)        not null comment '角色名',
    `create_time` datetime           NOT NULL DEFAULT now() COMMENT '创建时间',
    `delete_time` datetime                    DEFAULT NULL COMMENT '删除时间',
    `status`      tinyint(1) unsigned not null default '1' comment '状态',
    primary key (`id`)
) engine = InnoDB
  default charset = utf8mb4 comment ='角色表'
  collate = utf8mb4_general_ci;

drop table if exists `user_role`;
create table `user_role`
(
    `id`          int(11) unsigned    NOT NULL auto_increment,
    `user_id`     int(11) unsigned    NOT NULL,
    `role_id`     int(11) unsigned    NOT NULL,
    `create_time` datetime           NOT NULL DEFAULT now() COMMENT '创建时间',
    `update_time` datetime           NOT NULL DEFAULT now() COMMENT '更新时间',
    `delete_time` datetime                    DEFAULT NULL COMMENT '删除时间',
    `status`      tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态',
    primary key (`id`),
    key `user_id` (`user_id`),
    key `role_id` (`role_id`)
) engine = InnoDB
  default charset = utf8mb4 comment ='用户角色表'
  collate = utf8mb4_general_ci;

drop table if exists `resource`;
create table `resource`
(
    `id`          int(11) unsigned    NOT NULL auto_increment,
    `name`        varchar(255)        NOT NULL COMMENT '资源名',
    `url`         varchar(255)        NOT NULL COMMENT '资源路径',
    `create_time` datetime           NOT NULL DEFAULT now() COMMENT '创建时间',
    `update_time` datetime           NOT NULL DEFAULT now() COMMENT '更新时间',
    `delete_time` datetime                    DEFAULT NULL COMMENT '删除时间',
    `status`      tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态',
    primary key (`id`)
) engine = InnoDB
  default charset = utf8mb4 comment ='资源表'
  collate = utf8mb4_general_ci;

drop table if exists `role_resource`;
create table `role_resources`
(
    `id`          int(11) unsigned    NOT NULL auto_increment,
    `role_id`     int(11) unsigned    NOT NULL,
    `resource_id` int(11) unsigned    NOT NULL,
    `create_time` datetime           NOT NULL DEFAULT now() COMMENT '创建时间',
    `update_time` datetime           NOT NULL DEFAULT now() COMMENT '更新时间',
    `delete_time` datetime                    DEFAULT NULL COMMENT '删除时间',
    `status`      tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态',
    primary key (`id`),
    key `role_id` (`role_id`),
    key `resource_id` (`resource_id`)
) engine = InnoDB
  default charset = utf8mb4 comment ='角色资源表'
  collate = utf8mb4_general_ci;
