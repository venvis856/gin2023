CREATE TABLE `identify`
(
    `id`                 int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `identify_name`      varchar(255) NOT NULL COMMENT '身份名',
    `root`               varchar(255) DEFAULT NULL COMMENT '身份标识符',
    `type`               tinyint(3) unsigned NOT NULL COMMENT '1 酒店 2 派出所 3 大厦 4 园区  9 系统',
    `father_identify_id` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '父级',
    `status`             int(3) DEFAULT NULL COMMENT '1正常 5禁用 9删除',
    `create_time`        int(11) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time`        int(11) unsigned NOT NULL DEFAULT 0 COMMENT '更新时间',
    `delete_time`        int(11) unsigned NOT NULL DEFAULT 0 COMMENT '删除时间',
    `location`           varchar(255) DEFAULT NULL COMMENT '位置',
    `location_x`         varchar(50)  NOT NULL COMMENT '经度',
    `location_y`         varchar(50)  NOT NULL COMMENT '维度',
    `road_id`            int(11) unsigned DEFAULT NULL COMMENT '路id',
    `street_id`          int(11) unsigned DEFAULT NULL COMMENT '街道id',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COMMENT='身份表';

CREATE TABLE `permission`
(
    `id`                     int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `permission_name`        varchar(100) NOT NULL COMMENT '权限名称',
    `permission_code`        varchar(100) NOT NULL COMMENT '权限code',
    `identify_id`            int(11) unsigned NOT NULL COMMENT '标识id',
    `type`                   tinyint(3) unsigned NOT NULL COMMENT '类型 1菜单 2普通权限',
    `scene`                  tinyint(3) unsigned NOT NULL COMMENT '场景 1 后台 2 app 3 大屏',
    `father_permission_code` varchar(100) DEFAULT '0' COMMENT '父权限code',
    `status`                 tinyint(4) unsigned NOT NULL DEFAULT 1 COMMENT '1：正常 5 禁用  9 删除',
    `create_time`            int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户创建时间',
    `update_time`            int(11) unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
    `delete_time`            int(11) unsigned NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `identify_code` (`permission_code`,`identify_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=585 DEFAULT CHARSET=utf8mb4 COMMENT='权限表';


CREATE TABLE `role_permission`
(
    `id`              int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `role_id`         int(11) NOT NULL COMMENT '角色id',
    `permission_code` varchar(100) NOT NULL COMMENT '权限code',
    `identify_id`     int(11) unsigned DEFAULT NULL COMMENT '标识id',
    `is_effective`    int(2) unsigned DEFAULT 0 COMMENT '是否有效',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=261 DEFAULT CHARSET=utf8mb4 COMMENT='角色权限表';

CREATE TABLE `table_ids`
(
    `id`          int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `table_name`  varchar(50) NOT NULL COMMENT '表名',
    `identify_id` int(11) NOT NULL COMMENT '身份id',
    `max_id`      int(11) unsigned NOT NULL DEFAULT 0 COMMENT '最大id',
    PRIMARY KEY (`id`) USING BTREE,
    KEY           `lianhe` (`table_name`,`identify_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb4 COMMENT='虚拟id表';

CREATE TABLE `user`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
    `vid`         int(11) unsigned NOT NULL COMMENT '序号',
    `username`    varchar(100) DEFAULT NULL COMMENT '管理员名称',
    `phone`       varchar(50)  DEFAULT NULL COMMENT '手机号',
    `realname`    varchar(100) DEFAULT NULL COMMENT '真实中文名',
    `password`    varchar(200) DEFAULT NULL COMMENT '账号密码',
    `email`       varchar(100) DEFAULT NULL COMMENT '管理员邮箱',
    `status`      tinyint(4) NOT NULL DEFAULT 1 COMMENT '管理员账户状态，1：正常 5 禁用  9删除',
    `create_time` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '用户创建时间',
    `update_time` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
    `delete_time` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '删除时间',
    `user_ip`     varchar(20)  DEFAULT NULL COMMENT '用户IP地址',
    `login_time`  int(11) DEFAULT NULL COMMENT '最近一次登录时间',
    `role_ids`    varchar(255) DEFAULT NULL COMMENT '角色ids',
    `identify_id` int(11) unsigned NOT NULL COMMENT '标识id',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `phone` (`phone`,`identify_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE `user_permission`
(
    `id`              int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user_id`         int(11) NOT NULL COMMENT '用户id',
    `permission_code` varchar(100) NOT NULL COMMENT '权限code',
    `identify_id`     int(11) unsigned NOT NULL COMMENT '标识id',
    `is_effective`    tinyint(3) unsigned DEFAULT 0 COMMENT '是否有效',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户权限表';
