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
    `identify_id`         int(10) unsigned NOT NULL,
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

CREATE TABLE `cms_tag`
(
    `id`          int(11) unsigned NOT NULL AUTO_INCREMENT,
    `tag_name`    varchar(255) DEFAULT NULL,
    `site_id`     int(11) DEFAULT NULL,
    `status`      int(3) DEFAULT NULL COMMENT '1正常 5禁用 9删除',
    `create_time` int(11) DEFAULT NULL,
    `update_time` int(11) DEFAULT NULL,
    `delete_time` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `notebook`
(
    `id`          int(11) unsigned NOT NULL AUTO_INCREMENT,
    `date`        varchar(100)  DEFAULT NULL,
    `content`     varchar(1024) DEFAULT NULL,
    `user_id`     int(11) DEFAULT NULL,
    `status`      int(2) unsigned NOT NULL COMMENT '1 待通知 2 已通知 5 禁用 9 删除',
    `create_time` int(11) unsigned DEFAULT NULL,
    `update_time` int(11) unsigned DEFAULT NULL,
    `delete_time` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--------------------------------修改记录----------------------------------

ALTER TABLE `cms_page`
    ADD COLUMN `is_publish` int(11) NULL DEFAULT 2 COMMENT '1 发布过 2没发布过' AFTER `first_make_time`;

ALTER TABLE `cms_page`
    ADD COLUMN `star_number` varchar(50) NULL COMMENT '点赞数' AFTER `is_publish`;

ALTER TABLE `cms_page`
    ADD COLUMN `tag_ids` varchar(255) NULL COMMENT 'Tags' AFTER `product_id`;

