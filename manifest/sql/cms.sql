-- 用户表
CREATE TABLE `user`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
    `username`    varchar(100) NOT NULL COMMENT '管理员名称',
    `realname`    varchar(100) DEFAULT NULL COMMENT '真实中文名',
    `password`    varchar(200) DEFAULT NULL COMMENT '账号密码',
    `email`       varchar(100) DEFAULT NULL COMMENT '管理员邮箱',
    `status`      tinyint(4) NOT NULL DEFAULT '1' COMMENT '管理员账户状态，0：停用 1：正常',
    `create_time` int(11) DEFAULT NULL COMMENT '用户创建时间',
    `update_time` int(11) DEFAULT NULL COMMENT '修改时间',
    `delete_time` int(11) NOT NULL DEFAULT '0',
    `user_ip`     varchar(20)  DEFAULT NULL COMMENT '用户IP地址',
    `login_time`  int(11) DEFAULT NULL COMMENT '最近一次登录时间',
    `role_ids`    varchar(255) DEFAULT NULL COMMENT '角色ids',
    `site_ids`    varchar(255) DEFAULT NULL COMMENT '站点ids',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

INSERT INTO `user` (`create_time`, `password`, `role_ids`, `status`, `username`)
VALUES (1631936904, '9d05396ca6003f16', '[1]', 1, 'admin');

-- 角色表
CREATE TABLE `role`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `role_name`   varchar(100) NOT NULL COMMENT '角色名称',
    `status`      tinyint(4) NOT NULL DEFAULT '1' COMMENT '1：正常 5 禁用  9 删除',
    `type`        tinyint(4) NOT NULL DEFAULT '2' COMMENT '1：系统管理员 2 超级管理员 3普通角色',
    `create_time` int(11) DEFAULT NULL COMMENT '用户创建时间',
    `update_time` int(11) DEFAULT NULL COMMENT '修改时间',
    `delete_time` int(11) DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `role_name` (`role_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
insert into role(role_name, status, type)
    value("超级管理员",1,1);

-- 权限表
CREATE TABLE `permission`
(
    `id`                     int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `permission_name`        varchar(100) NOT NULL COMMENT '权限名称',
    `permission_code`        varchar(100) NOT NULL COMMENT '权限code',
    `site_id`                int(11) NOT NULL COMMENT '站点id',
    `type`                   tinyint(3) NOT NULL COMMENT '类型 1菜单 2普通权限',
    `father_permission_code` varchar(100) DEFAULT '0' COMMENT '父权限code',
    `status`                 tinyint(4) NOT NULL DEFAULT '1' COMMENT '1：正常 5 禁用  9 删除',
    `create_time`            int(11) DEFAULT NULL COMMENT '用户创建时间',
    `update_time`            int(11) DEFAULT NULL COMMENT '修改时间',
    `delete_time`            int(11) DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `site_id_code` (`permission_code`,`site_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- 设置
insert into permission(permission_name, permission_code, status, site_id, type, father_permission_code)
values ("设置", "set", 1, 0, 1, "");
insert into permission(permission_name, permission_code, status, site_id, type, father_permission_code)
values ("权限管理", "permission_list", 1, 0, 2, "set");
insert into permission(permission_name, permission_code, status, site_id, type, father_permission_code)
values ("用户管理", "user_list", 1, 0, 2, "set");

--  新增站点权限
insert into permission(permission_name, permission_code, status, site_id, type, father_permission_code)
values ("4itool", "4itool", 1, 1, 1, "");

-- 业务 根据
insert into permission(permission_name, permission_code, status, site_id, type, father_permission_code)
values ("文章管理", "article_list", 1, 1, 2, "4itool"),
       ("模板管理", "template_list", 1, 1, 2, "4itool"),
       ("分类管理", "classify_list", 1, 1, 2, "4itool"),
       ("作者管理", "author_list", 1, 1, 2, "4itool"),
       ("产品管理", "product_list", 1, 1, 2, "4itool"),
       ("模块管理", "module_list", 1, 1, 2, "4itool"),
       ("图片管理", "picture_list", 1, 1, 2, "4itool"),
       ("发布管理", "publish_list", 1, 1, 2, "4itool");

-- 用户权限关系表
CREATE TABLE `user_permission`
(
    `id`              int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user_id`         int(11) NOT NULL COMMENT '用户id',
    `permission_code` varchar(100) NOT NULL COMMENT '权限code',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- 角色权限关系表
CREATE TABLE `role_permission`
(
    `id`              int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `role_id`         int(11) NOT NULL COMMENT '角色id',
    `permission_code` varchar(100) NOT NULL COMMENT '权限code',
    `site_id`         int(11) unsigned DEFAULT NULL,
    `is_effective`    int(2) unsigned DEFAULT 0 COMMENT '是否有效',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
--  业务
CREATE TABLE `cms_author`
(
    `id`          int(11) unsigned NOT NULL AUTO_INCREMENT,
    `author_name` varchar(255) DEFAULT NULL,
    `site_id`     int(11) DEFAULT NULL,
    `status`      int(3) DEFAULT NULL COMMENT '1正常 5禁用 9删除',
    `create_time` int(11) DEFAULT NULL,
    `update_time` int(11) DEFAULT NULL,
    `delete_time` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `cms_classify`
(
    `id`            int(11) unsigned NOT NULL AUTO_INCREMENT,
    `classify_name` varchar(255) DEFAULT NULL,
    `is_howto`      tinyint(2) DEFAULT NULL COMMENT '1 是 2 否',
    `site_id`       int(11) DEFAULT NULL,
    `status`        int(3) DEFAULT NULL COMMENT '1正常 5禁用 9删除',
    `create_time`   int(11) DEFAULT NULL,
    `update_time`   int(11) DEFAULT NULL,
    `delete_time`   int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `cms_module`
(
    `id`          int(11) unsigned NOT NULL AUTO_INCREMENT,
    `module_name` varchar(255) NOT NULL,
    `site_id`     int(11) unsigned NOT NULL,
    `status`      int(3) NOT NULL COMMENT '1正常 5禁用 9删除',
    `content`     longtext,
    `create_time` int(11) DEFAULT NULL,
    `update_time` int(11) DEFAULT NULL,
    `delete_time` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `cms_page`
(
    `id`                  int(11) unsigned NOT NULL AUTO_INCREMENT,
    `site_id`             int(11) unsigned NOT NULL,
    `status`              int(3) unsigned NOT NULL COMMENT '1正常 5禁用 9删除',
    `subject`             varchar(255) NOT NULL,
    `title`               varchar(255) DEFAULT NULL,
    `keywords`            varchar(255) DEFAULT NULL,
    `description`         varchar(255) DEFAULT NULL,
    `content`             longtext     DEFAULT NULL,
    `url`                 varchar(255) DEFAULT NULL COMMENT '页面地址',
    `image_url`           varchar(255) DEFAULT NULL COMMENT '页面封面图片地址',
    `template_id`         int(11) unsigned NOT NULL,
    `classify_id`         int(11) unsigned NOT NULL,
    `author_id`           int(11) unsigned NOT NULL,
    `product_id`          int(11) unsigned DEFAULT NULL,
    `last_update_user_id` int(11) DEFAULT NULL,
    `create_time`         int(11) DEFAULT NULL,
    `update_time`         int(11) DEFAULT NULL,
    `delete_time`         int(11) DEFAULT NULL,
    `first_make_time`     int(11) unsigned DEFAULT 0,
    `is_publish`          int(11) DEFAULT 2 COMMENT '1 发布过 2没发布过',
    `star_number`         varchar(50)  DEFAULT NULL COMMENT '点赞数',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `cms_product`
(
    `id`           int(11) unsigned NOT NULL AUTO_INCREMENT,
    `product_name` varchar(255) DEFAULT NULL,
    `download_url` varchar(255) DEFAULT NULL,
    `buy_url`      varchar(255) DEFAULT NULL,
    `site_id`      int(11) DEFAULT NULL,
    `status`       int(3) DEFAULT NULL COMMENT '1正常 5禁用 9删除',
    `create_time`  int(11) DEFAULT NULL,
    `update_time`  int(11) DEFAULT NULL,
    `delete_time`  int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `cms_site`
(
    `id`                  int(11) unsigned NOT NULL AUTO_INCREMENT,
    `site_name`           varchar(255) NOT NULL,
    `root`                varchar(255) DEFAULT NULL,
    `online_url`          varchar(255) DEFAULT NULL COMMENT '现网url',
    `online_image_url`    varchar(255) DEFAULT NULL,
    `preview_url`         varchar(255) DEFAULT NULL,
    `rsync_password_path` varchar(255) DEFAULT NULL,
    `rsync_address`       varchar(255) DEFAULT NULL,
    `rsync_image_address` varchar(255) DEFAULT NULL,
    `status`              int(3) DEFAULT NULL COMMENT '1正常 5禁用 9删除',
    `create_time`         int(11) DEFAULT NULL,
    `update_time`         int(11) DEFAULT NULL,
    `delete_time`         int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `cms_template`
(
    `id`            int(11) unsigned NOT NULL AUTO_INCREMENT,
    `site_id`       int(11) unsigned NOT NULL,
    `status`        int(11) unsigned NOT NULL,
    `template_name` varchar(255) NOT NULL,
    `type`          int(3) NOT NULL COMMENT '1首页  2文章  3分类  4产品  5review  6guide 7 其他',
    `content`       longtext,
    `module_ids`    varchar(255) DEFAULT NULL,
    `create_time`   int(11) DEFAULT NULL,
    `update_time`   int(11) DEFAULT NULL,
    `delete_time`   int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

CREATE TABLE `cms_audit`
(
    `id`              int(11) unsigned NOT NULL AUTO_INCREMENT,
    `first_url`       varchar(255) DEFAULT NULL,
    `count`           int(5) DEFAULT NULL,
    `type`            int(3) DEFAULT NULL COMMENT '1 页面 2 图片 ',
    `make_user_id`    int(11) DEFAULT NULL COMMENT '生成用户id',
    `publush_user_id` int(11) unsigned DEFAULT NULL COMMENT '发布人id',
    `status`          int(3) unsigned DEFAULT NULL COMMENT '1 待发布 2 已发布 9 删除',
    `site_id`         int(11) unsigned DEFAULT NULL,
    `make_time`       int(11) unsigned DEFAULT NULL,
    `publush_time`    int(11) unsigned DEFAULT NULL,
    `delete_time`     int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `cms_audit_detail`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `local_url`   varchar(255) DEFAULT NULL,
    `online_url`  varchar(255) DEFAULT NULL,
    `preview_url` varchar(255) DEFAULT NULL,
    `file_url`    varchar(255) DEFAULT NULL,
    `type`        tinyint(3) unsigned DEFAULT NULL COMMENT '1 页面 2图片',
    `status`      tinyint(3) unsigned DEFAULT NULL COMMENT '1 正常 9 删除',
    `audit_id`    int(11) unsigned DEFAULT NULL,
    `page_id`     int(11) DEFAULT NULL COMMENT '页面id',
    `make_time`   int(11) unsigned DEFAULT NULL,
    `delete_time` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `cms_tag` (
   `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
   `tag_name` varchar(255) DEFAULT NULL,
   `site_id` int(11) DEFAULT NULL,
   `status` int(3) DEFAULT NULL COMMENT '1正常 5禁用 9删除',
   `create_time` int(11) DEFAULT NULL,
   `update_time` int(11) DEFAULT NULL,
   `delete_time` int(11) DEFAULT NULL,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `notebook` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `date` varchar(100) DEFAULT NULL,
    `content` varchar(1024) DEFAULT NULL,
    `user_id` int(11) DEFAULT NULL,
    `status` int(2) unsigned NOT NULL COMMENT '1 待通知 2 已通知 5 禁用 9 删除',
    `create_time` int(11) unsigned DEFAULT NULL,
    `update_time` int(11) unsigned DEFAULT NULL,
    `delete_time` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--------------------------------修改记录----------------------------------
ALTER TABLE `user`
    ADD COLUMN `site_ids` varchar(255) NULL COMMENT '站点ids' AFTER `role_ids`;

ALTER TABLE `cms_page`
    ADD COLUMN `is_publish` int(11) NULL DEFAULT 2 COMMENT '1 发布过 2没发布过' AFTER `first_make_time`;

-- 2022 04 29
ALTER TABLE `cms`.`permission`
    MODIFY COLUMN `type` tinyint(3) NOT NULL COMMENT '类型 1菜单 2普通权限' AFTER `site_id`;

update permission
set father_permission_code=permission_code;
ALTER TABLE `cms`.`user_permission`
    ADD COLUMN `site_id` int(11) UNSIGNED NULL AFTER `permission_code`;

ALTER TABLE `cms`.`role_permission`
    ADD COLUMN `site_id` int(11) UNSIGNED NULL AFTER `permission_code`;

ALTER TABLE `cms`.`cms_page`
    ADD COLUMN `star_number` varchar(50) NULL COMMENT '点赞数' AFTER `is_publish`;

insert into permission(permission_name, permission_code, status, site_id, type, father_permission_code)
values ("tag管理", "tag_list", 1, 1, 1, "tag_list");

ALTER TABLE `cms`.`cms_page`
    ADD COLUMN `tag_ids` varchar(255) NULL COMMENT 'Tags' AFTER `product_id`;

