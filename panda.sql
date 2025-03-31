CREATE TABLE IF NOT EXISTS `panda_dict_type` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典类型ID',
    `pid` bigint(20) NOT NULL DEFAULT '0' COMMENT '父类字典类型ID',
    `name` varchar(100) NOT NULL DEFAULT '' COMMENT '字典类型名称',
    `type` varchar(100) NOT NULL UNIQUE COMMENT '字典类型标识（唯一）',
    `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
    `remark` varchar(500) DEFAULT NULL COMMENT '备注',
    `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '字典类型状态（1-正常 0-停用）',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_pid` (`pid`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统_字典类型';


CREATE TABLE IF NOT EXISTS `panda_dict_data` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典数据ID',
    `label` varchar(100) DEFAULT NULL COMMENT '字典标签',
    `value` varchar(100) DEFAULT NULL COMMENT '字典键值',
    `value_type` varchar(255) NOT NULL DEFAULT 'string' COMMENT '键值数据类型：string,int,float64,bool',
    `type` varchar(100) DEFAULT NULL COMMENT '字典类型',
    `list_class` varchar(100) DEFAULT 'primary' COMMENT '表格回显样式',
    `is_default` tinyint(1) DEFAULT '0' COMMENT '是否为系统默认',
    `sort` int(11) DEFAULT '0' COMMENT '字典排序',
    `remark` varchar(500) DEFAULT NULL COMMENT '备注',
    `status` tinyint(1) DEFAULT '1' COMMENT '状态',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX `idx_type` (`type`),  -- 添加索引
    INDEX `idx_status` (`status`),  -- 添加索引
    INDEX `idx_is_default` (`is_default`)  -- 添加索引
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统_字典数据';



-- 后台用户表 --
CREATE TABLE IF NOT EXISTS `panda_user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `username` varchar(100) NOT NULL COMMENT '用户名',
    `password` varchar(255) NOT NULL COMMENT '密码',
    `nickname` varchar(100) NOT NULL COMMENT '昵称',
    `avatar` varchar(150) DEFAULT NULL COMMENT '头像',
    `email` varchar(255) DEFAULT NULL COMMENT '电子邮件',
    `phone` varchar(20) DEFAULT NULL COMMENT '手机号码',
    `is_develop` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否开发者账号 1 是  0   否',
    `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '用户状态（1-正常 0-停用）',
    `last_login_ip` varchar(15) NOT NULL DEFAULT '' COMMENT '最后登录ip',
    `last_login_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '最后登录时间',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台系统_用户表';

-- 后台用户角色表 --
CREATE TABLE IF NOT EXISTS `panda_role` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
    `name` varchar(200) NOT NULL UNIQUE COMMENT '角色名',
    `code` varchar(100) NOT NULL UNIQUE COMMENT '角色标识',
    `remark` varchar(255) DEFAULT NULL COMMENT '备注',
    `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '用户状态（1-正常 0-停用）',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台系统_角色表';

-- 后台用户角色关联表 --
CREATE TABLE IF NOT EXISTS `panda_user_role` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户角色关联ID',
    `user_id` bigint(20) NOT NULL COMMENT '用户ID',
    `role_id` bigint(20) NOT NULL COMMENT '角色ID',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台系统_用户角色关联表';

-- 后台用户附件表 --
CREATE TABLE IF NOT EXISTS `panda_attachment` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    -- 文件类型 --
    `type` varchar(10) NOT NULL DEFAULT 'image' COMMENT '文件类型（image doc audio video zip other）',
    -- 文件名 --
    `name` varchar(255) NOT NULL COMMENT '文件原始名',
    -- 文件路径 --
    `path` varchar(255) NOT NULL COMMENT '文件路径',
    -- 文件大小 --
    `size` varchar(30) NOT NULL DEFAULT '0' COMMENT '文件大小',
    -- 扩展名 --
    `ext` varchar(10) NOT NULL COMMENT '扩展名',
    -- 上传驱动类型 --
    `drive` tinyint(1) NOT NULL DEFAULT '1' COMMENT '上传驱动类型 1 本地 2 阿里云 3 腾讯云 4 七牛云',
    `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '文件状态（1-正常 0-停用）',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台系统附件表';

-- 后台任务表 --
CREATE TABLE IF NOT EXISTS `panda_crontab` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `title` varchar(128) NOT NULL COMMENT '任务标题',
    `name` varchar(100) DEFAULT NULL COMMENT '任务方法',
    `params` varchar(255) DEFAULT NULL COMMENT '函数参数',
    `pattern` varchar(64) NOT NULL COMMENT '表达式',
    `policy` bigint(20) NOT NULL DEFAULT '1' COMMENT '策略',
    `count` bigint(20) NOT NULL DEFAULT '0' COMMENT '执行次数',
    `sort` int(11) DEFAULT '0' COMMENT '排序',
    `remark` varchar(500) DEFAULT NULL COMMENT '备注',
    `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '任务状态（1-正常 0-停用）',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台系统附件表';

-- 后台菜单表 --
CREATE TABLE IF NOT EXISTS `panda_menu` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
    `parent_id` bigint DEFAULT '0' COMMENT '父级ID',
    `title` varchar(64) NOT NULL COMMENT '菜单标题',
    `name` varchar(128) NOT NULL COMMENT '路由名称',
    `path` varchar(200) DEFAULT NULL COMMENT '路由地址',
    `component` varchar(255) NOT NULL COMMENT '组件路径',
    `rank` int DEFAULT '0' COMMENT '排序号',
    `redirect` varchar(255) DEFAULT NULL COMMENT '重定向地址',
    `icon` varchar(128) DEFAULT NULL COMMENT '图标',
    `extra_icon` varchar(128) DEFAULT NULL COMMENT '额外图标',
    `enter_transition` varchar(64) DEFAULT NULL COMMENT '进场动画',
    `leave_transition` varchar(64) DEFAULT NULL COMMENT '离场动画',
    `active_path` varchar(255) DEFAULT NULL COMMENT '激活路径',
    `auths` varchar(512) DEFAULT NULL COMMENT '权限标识',
    `menu_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '菜单类型（0菜单 1iframe 2外链 3按钮）',
    `frame_src` varchar(512) DEFAULT NULL COMMENT '内嵌iframe地址',
    `frame_loading` tinyint(1) DEFAULT '1' COMMENT 'iframe加载状态(0关闭 1开启)',
    `keep_alive` tinyint(1) DEFAULT '0' COMMENT '缓存路由(0关闭 1开启)',
    `hidden_tag` tinyint(1) DEFAULT '0' COMMENT '隐藏标签(0显示 1隐藏)',
    `fixed_tag` tinyint(1) DEFAULT '0' COMMENT '固定标签(0不固定 1固定)',
    `show_link` tinyint(1) DEFAULT '1' COMMENT '显示链接(0隐藏 1显示)',
    `show_parent` tinyint(1) DEFAULT '1' COMMENT '显示父级(0隐藏 1显示)',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_parent_id` (`parent_id`),
    KEY `idx_menu_type` (`menu_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台系统_菜单表';

-- 后台角色菜单表 --
CREATE TABLE IF NOT EXISTS `panda_role_menu` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色菜单关联ID',
    `role_id` bigint(20) NOT NULL COMMENT '角色ID',
    `menu_id` bigint(20) NOT NULL COMMENT '菜单ID',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台系统_角色菜单关联表';

-- Casbin 规则表
CREATE TABLE IF NOT EXISTS `casbin_rule` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) NOT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Casbin 规则表';