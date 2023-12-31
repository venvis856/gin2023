DROP TABLE IF EXISTS identify;
CREATE TABLE `identify`
(
    `id`                 INT ( 11 ) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `identify_name`      VARCHAR(255) NOT NULL COMMENT '身份名',
    `identify_code`      VARCHAR(255) DEFAULT NULL COMMENT '身份标识符',
    `type`               TINYINT ( 3 ) UNSIGNED NOT NULL COMMENT '1 系统 2 其他',
    `father_identify_id` INT ( 11 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级',
    `status`             TINYINT ( 3 ) DEFAULT NULL COMMENT '1正常 5禁用 9删除',
    `create_time`        INT ( 11 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time`        INT ( 11 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    `delete_time`        INT ( 11 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `identify_code` ( `identify_code` ) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '身份标识表';

DROP TABLE IF EXISTS permission;
CREATE TABLE `permission`
(
    `id`                     INT ( 11 ) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `permission_name`        VARCHAR(100) NOT NULL COMMENT '权限名称',
    `permission_code`        VARCHAR(100) NOT NULL COMMENT '权限code',
    `type`                   TINYINT ( 3 ) UNSIGNED NOT NULL COMMENT '类型 1菜单 2普通权限',
    `father_permission_code` VARCHAR(100) DEFAULT '0' COMMENT '父权限code',
    `status`                 TINYINT ( 4 ) UNSIGNED NOT NULL DEFAULT 1 COMMENT '1：正常 5 禁用  9 删除',
    `create_time`            INT ( 11 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户创建时间',
    `update_time`            INT ( 11 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
    `delete_time`            INT ( 11 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `permission_code` ( `permission_code` ) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '权限表';

DROP TABLE IF EXISTS user;
CREATE TABLE `user`
(
    `id`          INT ( 11 ) NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
    `uid`         INT ( 11 ) UNSIGNED NOT NULL COMMENT '序号',
    `username`    VARCHAR(100) DEFAULT NULL COMMENT '管理员名称',
    `phone`       VARCHAR(50)  DEFAULT NULL COMMENT '手机号',
    `realname`    VARCHAR(100) DEFAULT NULL COMMENT '真实中文名',
    `password`    VARCHAR(200) DEFAULT NULL COMMENT '账号密码',
    `email`       VARCHAR(100) DEFAULT NULL COMMENT '管理员邮箱',
    `status`      TINYINT ( 4 ) NOT NULL DEFAULT 1 COMMENT '管理员账户状态，1：正常 5 禁用  9删除',
    `create_time` INT ( 11 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户创建时间',
    `update_time` INT ( 11 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
    `delete_time` INT ( 11 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
    `user_ip`     VARCHAR(20)  DEFAULT NULL COMMENT '用户IP地址',
    `login_time`  INT ( 11 ) DEFAULT NULL COMMENT '最近一次登录时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uid` ( `uid` ) USING BTREE,
    UNIQUE KEY `phone` ( `phone` ) USING BTREE,
    UNIQUE KEY `email` ( `email` ) USING BTREE,
    CHECK (( `phone` IS NOT NULL) OR (`email` IS NOT NULL))
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';

DROP TABLE  IF    EXISTS role;
CREATE TABLE `role`
(
    `id`          INT ( 11 ) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `uid`         INT ( 11 ) UNSIGNED NOT NULL COMMENT '序号',
    `role_name`   VARCHAR(100) NOT NULL COMMENT '角色名称',
    `status`      TINYINT ( 4 ) NOT NULL DEFAULT 1 COMMENT '1：正常 5 禁用  9 删除',
    `type`        TINYINT ( 4 ) NOT NULL DEFAULT 2 COMMENT '1：系统管理员 2 超级管理员 3普通角色',
    `create_time` INT ( 11 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户创建时间',
    `update_time` INT ( 11 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
    `delete_time` INT ( 11 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
    `identify_id` INT ( 11 ) UNSIGNED NOT NULL COMMENT '标识id',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_identify_role` ( `identify_id`, `role_name` ) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '角色表';


DROP TABLE    IF  EXISTS user_role;
CREATE TABLE `user_role`
(
    `id`           INT ( 11 ) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `identify_id`  INT ( 11 ) UNSIGNED NOT NULL COMMENT '标识id',
    `user_id`      INT ( 11 ) UNSIGNED NOT NULL COMMENT 'user_service id',
    `role_id`      INT ( 11 ) UNSIGNED NOT NULL COMMENT 'role_service id',
    `is_effective` INT ( 2 ) UNSIGNED DEFAULT 0 COMMENT '是否有效',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_identify_user_role` ( `identify_id`, `user_id`, `role_id` ) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '角色表';

DROP TABLE IF EXISTS role_permission;
CREATE TABLE `role_permission`
(
    `id`            INT ( 11 ) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `identify_id`   INT ( 11 ) UNSIGNED DEFAULT NULL COMMENT '标识id',
    `role_id`       INT ( 11 ) NOT NULL COMMENT '角色id',
    `permission_id` INT ( 11 ) NOT NULL COMMENT '权限id',
    `is_effective`  INT ( 2 ) UNSIGNED DEFAULT 0 COMMENT '是否有效',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_identify_role_permission` ( `identify_id`, `role_id`, `permission_id` ) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '角色权限表';

DROP TABLE IF  EXISTS user_permission;
CREATE TABLE `user_permission`
(
    `id`            INT ( 11 ) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `identify_id`   INT ( 11 ) UNSIGNED NOT NULL COMMENT '标识id',
    `user_id`       INT ( 11 ) NOT NULL COMMENT '用户id',
    `permission_id` INT ( 11 ) NOT NULL COMMENT '权限code',
    `is_effective`  TINYINT ( 3 ) UNSIGNED DEFAULT 0 COMMENT '是否有效',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_identify_user_permission` ( `identify_id`, `user_id`, `permission_id` ) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '用户权限表';

DROP TABLE  IF EXISTS identify_permission;
CREATE TABLE `identify_permission`
(
    `id`            INT ( 11 ) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `identify_id`   INT ( 11 ) UNSIGNED NOT NULL COMMENT '标识id',
    `permission_id` INT ( 11 ) NOT NULL COMMENT '权限code',
    `is_effective`  TINYINT ( 3 ) UNSIGNED DEFAULT 0 COMMENT '是否有效',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_identify_permission` ( `identify_id`, `permission_id` ) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '标识权限表';

DROP TABLE  IF  EXISTS table_ids;
CREATE TABLE `table_ids`
(
    `id`          INT ( 11 ) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `table_name`  VARCHAR(50) NOT NULL COMMENT '表名',
    `identify_id` INT ( 11 ) NOT NULL COMMENT '身份id',
    `max_id`      INT ( 11 ) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最大id',
    PRIMARY KEY (`id`) USING BTREE,
    KEY           `lianhe` ( `table_name`, `identify_id` ) USING BTREE
) ENGINE = INNODB AUTO_INCREMENT = 47 DEFAULT CHARSET = utf8mb4 COMMENT = '虚拟id表';


# ---------------------------------
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('设置','set',1,'0',1);
-- SET @id = LAST_INSERT_ID();
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('用户管理','user_list',1,'set',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('用户详情','user_info',1,'user_list',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('用户新增','user_add',1,'user_list',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('用户修改','user_update',1,'user_list',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('用户删除','user_delete',1,'user_list',1);

insert into permission(permission_name,permission_code,type,father_permission_code,status) values('角色管理','role_list',1,'set',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('角色详情','role_info',1,'role_list',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('角色新增','role_add',1,'role_list',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('角色修改','role_update',1,'role_list',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('角色删除','role_delete',1,'role_list',1);

insert into permission(permission_name,permission_code,type,father_permission_code,status) values('权限管理','permission_list',1,'set',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('权限详情','permission_info',1,'permission_list',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('权限新增','permission_add',1,'permission_list',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('权限修改','permission_update',1,'permission_list',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('权限删除','permission_delete',1,'permission_list',1);

insert into permission(permission_name,permission_code,type,father_permission_code,status) values('身份标识符号列表','identify_list',1,'set',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('身份标识符号详情','identify_info',1,'identify_list',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('身份标识符号新增','identify_add',1,'identify_list',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('身份标识符号修改','identify_update',1,'identify_list',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('身份标识符号删除','identify_delete',1,'identify_list',1);

# ---- 标识拥有权限关系表
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('标识拥有权限关系列表','identify_permission_list',1,'set',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('标识拥有权限关系新增','identify_permission_add',1,'identify_permission_list',1);

# ---- 角色权限关系表
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('角色拥有权限关系列表','role_permission_list',1,'set',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('角色拥有权限关系新增','role_permission_add',1,'role_permission_list',1);

# ---- 用户角色关系表
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('用户角色关系列表','user_role_list',1,'set',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('用户角色关系新增','user_role_add',1,'user_role_list',1);

# ---- 用户权限关系表
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('用户权限关系列表','user_permission_list',1,'set',1);
insert into permission(permission_name,permission_code,type,father_permission_code,status) values('用户权限关系新增','user_permission_add',1,'user_permission_list',1);



