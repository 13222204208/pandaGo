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

INSERT INTO `panda_dict_type` (`id`, `pid`, `name`, `type`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES
                                                                                                                         (1, 27, '用户性别', 'sys_user_sex', 0, '用户性别列表', 1, '2021-01-30 13:27:43', '2021-03-24 11:38:47'),
                                                                                                                         (2, 28, '菜单类型', 'sys_menu_types', 80, '', 1, '2021-01-30 13:27:43', '2023-02-26 00:46:18'),
                                                                                                                         (3, 28, '系统状态', 'sys_normal_disable', 10, '系统状态列表', 1, '2021-01-30 13:27:43', NULL),
                                                                                                                         (4, 26, '任务状态', 'sys_job_status', 0, '任务状态列表', 1, '2021-01-30 13:27:43', NULL),
                                                                                                                         (5, 26, '任务分组', 'sys_job_group', 0, '任务分组列表', 1, '2021-01-30 13:27:43', NULL),
                                                                                                                         (7, 29, '通知类型', 'sys_notice_type', 0, '通知类型列表', 1, '2021-01-30 13:27:43', NULL),
                                                                                                                         (9, 28, '操作类型', 'sys_oper_type', 100, '', 1, '2021-01-30 13:27:43', '2023-02-26 00:48:01'),
                                                                                                                         (22, 25, '请求方式', 'sys_oper_method', 0, '', 1, '2022-01-24 11:33:16', '2022-01-24 11:33:16'),
                                                                                                                         (23, 25, '请求状态码', 'req_code', 0, '', 1, '2022-01-25 10:22:40', '2022-01-25 10:22:40'),
                                                                                                                         (24, 25, '请求耗时', 'req_take_up_time', 0, '', 1, '2022-01-25 15:51:19', '2022-01-25 15:51:19'),
                                                                                                                         (25, 0, '请求枚举', 'req', 40, '', 1, '2022-01-24 11:33:16', '2022-01-24 11:33:16'),
                                                                                                                         (26, 0, '任务枚举', 'job', 30, '', 1, '2022-01-24 11:33:16', '2022-01-24 11:33:16'),
                                                                                                                         (27, 0, '组织枚举', 'org', 20, '', 1, '2022-01-24 11:33:16', '2023-04-27 23:24:53'),
                                                                                                                         (28, 0, '系统枚举', 'sys', 10, '', 1, '2022-01-24 11:33:16', '2023-02-26 00:33:13'),
                                                                                                                         (29, 0, '通知枚举', 'notice', 50, '', 1, '2022-01-24 11:33:16', '2022-01-24 11:33:16'),
                                                                                                                         (30, 27, '用户爱好', 'sys_user_hobby', 0, NULL, 1, '2021-03-01 11:41:07', '2021-03-01 11:41:07'),
                                                                                                                         (31, 27, '用户渠道', 'sys_user_channel', 0, NULL, 1, '2021-03-01 11:41:07', '2021-03-01 11:41:07'),
                                                                                                                         (32, 0, '配置枚举', 'config', 60, '', 1, '2022-12-30 17:55:42', '2023-05-29 14:47:01'),
                                                                                                                         (33, 32, '上传驱动', 'config_upload_drive', 10, '', 1, '2022-12-30 17:57:18', '2022-12-30 17:57:18'),
                                                                                                                         (34, 28, '日志类型', 'sys_log_type', 200, '', 1, '2023-01-20 16:39:52', '2023-02-24 17:40:45'),
                                                                                                                         (35, 32, '短信驱动', 'config_sms_drive', 20, '', 1, '2023-01-21 10:27:01', '2023-01-21 10:27:01'),
                                                                                                                         (36, 32, '短信模板', 'config_sms_template', 30, '', 1, '2023-01-21 10:27:41', '2023-01-21 10:27:41'),
                                                                                                                         (37, 29, '通知标签', 'sys_notice_tag', 20, '', 1, '2023-01-26 12:09:46', '2023-01-26 12:09:46'),
                                                                                                                         (38, 32, '邮件模板', 'config_ems_template', 40, '', 1, '2023-02-04 12:53:07', '2023-02-04 12:53:07'),
                                                                                                                         (40, 28, '系统开关', 'sys_switch', 20, '', 1, '2024-04-16 09:50:58', '2024-04-16 09:50:58'),
                                                                                                                         (41, 28, '菜单组件', 'sys_menu_component', 90, '', 1, '2024-04-16 13:35:29', '2024-04-16 13:35:29'),
                                                                                                                         (42, 28, '登录状态', 'sys_login_status', 300, '', 1, '2024-04-16 15:21:23', '2024-04-16 15:21:23');


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


INSERT INTO `panda_dict_data` (`id`, `label`, `value`, `value_type`, `type`, `list_class`, `is_default`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES
                                                                                                                                                                      (1, '男', '1', 'int', 'sys_user_sex', 'success', 1, 10, '性别男', 1, '2021-01-30 13:27:43', '2023-04-27 23:24:58'),
                                                                                                                                                                      (2, '女', '2', 'int', 'sys_user_sex', 'warning', 1, 20, '性别女', 1, '2021-01-30 13:27:43', '2023-01-06 09:45:21'),
                                                                                                                                                                      (3, '保密', '3', 'int', 'sys_user_sex', 'error', 1, 30, '性别保密', 1, '2021-01-30 13:27:43', '2023-01-06 09:45:18'),
                                                                                                                                                                      (6, '正常', '1', 'int', 'sys_normal_disable', 'success', 1, 10, '正常状态', 1, '2021-01-30 13:27:43', '2023-01-18 09:52:44'),
                                                                                                                                                                      (7, '停用', '2', 'int', 'sys_normal_disable', 'warning', 1, 20, '停用状态', 1, '2021-01-30 13:27:43', '2023-01-18 09:52:48'),
                                                                                                                                                                      (8, '正常', '1', 'int', 'sys_job_status', 'primary', 1, 10, '正常状态', 1, '2021-01-30 13:27:43', '2023-01-18 09:53:48'),
                                                                                                                                                                      (9, '暂停', '2', 'int', 'sys_job_status', 'error', 1, 20, '停用状态', 1, '2021-01-30 13:27:43', '2023-01-18 09:53:53'),
                                                                                                                                                                      (10, '默认', 'DEFAULT', 'string', 'sys_job_group', '', 1, 10, '默认分组', 1, '2021-01-30 13:27:43', NULL),
                                                                                                                                                                      (11, '系统', 'SYSTEM', 'string', 'sys_job_group', '', 1, 20, '系统分组', 1, '2021-01-30 13:27:43', NULL),
                                                                                                                                                                      (14, '通知', '1', 'int', 'sys_notice_type', 'warning', 1, 10, '通知', 1, '2021-01-30 13:27:43', '2023-01-05 20:06:47'),
                                                                                                                                                                      (15, '公告', '2', 'int', 'sys_notice_type', 'error', 1, 20, '公告', 1, '2021-01-30 13:27:43', '2023-01-26 12:08:52'),
                                                                                                                                                                      (18, '新增', '1', 'int', 'sys_oper_type', 'info', 0, 10, '新增操作', 1, '2021-01-30 13:27:43', NULL),
                                                                                                                                                                      (19, '修改', '2', 'int', 'sys_oper_type', 'info', 0, 20, '修改操作', 1, '2021-01-30 13:27:43', NULL),
                                                                                                                                                                      (20, '删除', '3', 'int', 'sys_oper_type', 'error', 0, 30, '删除操作', 1, '2021-01-30 13:27:43', NULL),
                                                                                                                                                                      (21, '授权', '4', 'int', 'sys_oper_type', 'primary', 0, 40, '授权操作', 1, '2021-01-30 13:27:43', NULL),
                                                                                                                                                                      (22, '导出', '5', 'int', 'sys_oper_type', 'warning', 0, 50, '导出操作', 1, '2021-01-30 13:27:43', NULL),
                                                                                                                                                                      (23, '导入', '6', 'int', 'sys_oper_type', 'warning', 0, 60, '导入操作', 1, '2021-01-30 13:27:43', NULL),
                                                                                                                                                                      (24, '强退', '7', 'int', 'sys_oper_type', 'error', 0, 70, '强退操作', 1, '2021-01-30 13:27:43', NULL),
                                                                                                                                                                      (25, '生成代码', '8', 'int', 'sys_oper_type', 'warning', 0, 80, '生成操作', 1, '2021-01-30 13:27:43', NULL),
                                                                                                                                                                      (26, '清空数据', '9', 'int', 'sys_oper_type', 'error', 0, 90, '清空操作', 1, '2021-01-30 13:27:43', NULL),
                                                                                                                                                                      (28, '校验', '11', 'int', 'sys_oper_type', NULL, 0, 110, '校验', 1, '2021-10-04 22:40:50', '2021-10-04 22:50:02'),
                                                                                                                                                                      (80, '查询', '10', 'int', 'sys_oper_type', NULL, 0, 100, '查询', 1, '2021-10-04 22:37:38', '2021-10-04 22:49:32'),
                                                                                                                                                                      (88, 'GET', 'GET', 'string', 'sys_oper_method', '', 0, 10, '', 1, '2022-01-25 10:16:46', '2022-01-25 10:16:46'),
                                                                                                                                                                      (89, 'POST', 'POST', 'string', 'sys_oper_method', '', 0, 20, '', 1, '2022-01-25 10:16:54', '2022-01-25 10:16:54'),
                                                                                                                                                                      (90, '请求失败', '-1', 'int', 'req_code', 'warning', 0, 20, '通用错误码', 1, '2022-01-25 10:23:34', '2023-01-06 09:58:17'),
                                                                                                                                                                      (91, '请求成功', '0', 'int', 'req_code', 'success', 0, 10, '通用成功码', 1, '2022-01-25 10:23:54', '2023-01-06 09:58:14'),
                                                                                                                                                                      (92, '内部发生错误', '50', 'int', 'req_code', 'error', 0, 30, '内部错误', 1, '2022-01-25 10:24:06', '2023-01-06 09:56:31'),
                                                                                                                                                                      (95, '不允许的操作', '59', 'int', 'req_code', 'error', 0, 60, '当前操作的给定参数无效', 1, '2022-01-25 10:24:38', '2023-01-06 09:56:00'),
                                                                                                                                                                      (103, '没有授权登录', '61', 'int', 'req_code', 'warning', 0, 140, '未授权', 1, '2022-01-25 10:26:12', '2023-01-06 09:55:50'),
                                                                                                                                                                      (109, '业务验证失败', '300', 'int', 'req_code', 'warning', 0, 200, '业务验证失败', 1, '2022-01-25 10:27:51', '2023-01-06 09:56:51'),
                                                                                                                                                                      (110, '大于50ms', '50', 'int64', 'req_take_up_time', 'success', 0, 10, '', 1, '2022-01-25 15:51:45', '2023-01-06 09:59:50'),
                                                                                                                                                                      (111, '大于100ms', '100', 'int64', 'req_take_up_time', 'success', 0, 20, '', 1, '2022-01-25 15:51:57', '2023-01-06 09:59:53'),
                                                                                                                                                                      (112, '大于200ms', '200', 'int64', 'req_take_up_time', 'info', 0, 30, '', 1, '2022-01-25 15:52:21', '2023-01-06 09:59:57'),
                                                                                                                                                                      (113, '大于500ms', '500', 'int64', 'req_take_up_time', 'warning', 0, 40, '', 1, '2022-01-25 15:52:32', '2023-01-06 10:00:02'),
                                                                                                                                                                      (114, '大于1000ms', '1000', 'int64', 'req_take_up_time', 'error', 0, 50, '', 1, '2022-01-25 15:52:57', '2023-01-06 10:00:06'),
                                                                                                                                                                      (115, '大于5000ms', '5000', 'int64', 'req_take_up_time', 'error', 0, 60, '', 1, '2022-01-24 23:54:15', '2023-01-06 10:00:10'),
                                                                                                                                                                      (116, '开启', '1', 'int', 'sys_switch', 'info', 1, 10, '系统开关开启', 1, '2021-01-30 13:27:43', '2022-12-27 00:00:00'),
                                                                                                                                                                      (117, '关闭', '2', 'int', 'sys_switch', 'warning', 1, 20, '系统开关关闭', 1, '2021-01-30 13:27:43', '2022-12-27 00:00:00'),
                                                                                                                                                                      (118, 'PC端', '1', 'int', 'sys_user_channel', 'info', 0, 80, '用户来源', 1, '2021-01-30 13:27:43', '2023-01-05 17:01:21'),
                                                                                                                                                                      (119, '移动端', '2', 'int', 'sys_user_channel', 'success', 0, 80, '用户来源', 1, '2021-01-30 13:27:43', '2023-01-05 16:52:48'),
                                                                                                                                                                      (120, '微信', '3', 'int', 'sys_user_channel', 'warning', 0, 80, '用户来源', 1, '2021-01-30 13:27:43', '2023-01-05 16:52:43'),
                                                                                                                                                                      (121, '抖音', '4', 'int', 'sys_user_channel', 'error', 0, 80, '用户来源', 1, '2021-01-30 13:27:43', '2023-01-05 16:52:39'),
                                                                                                                                                                      (122, '音乐', '1', 'int', 'sys_user_hobby', 'error', 0, 80, '爱好类型', 1, '2021-01-30 13:27:43', '2023-01-06 09:45:11'),
                                                                                                                                                                      (123, '读书', '2', 'int', 'sys_user_hobby', 'info', 0, 80, '爱好类型', 1, '2021-01-30 13:27:43', '2023-01-06 09:45:08'),
                                                                                                                                                                      (124, '游泳', '3', 'int', 'sys_user_hobby', 'warning', 0, 80, '爱好类型', 1, '2021-01-30 13:27:43', '2023-01-06 09:45:05'),
                                                                                                                                                                      (125, '本地存储', 'local', 'string', 'config_upload_drive', 'info', 1, 10, '', 1, '2022-12-30 17:57:50', '2023-01-06 10:02:31'),
                                                                                                                                                                      (126, 'ucloud', 'ucloud', 'string', 'config_upload_drive', 'success', 1, 20, '', 1, '2022-12-30 17:58:26', '2022-12-30 17:58:26'),
                                                                                                                                                                      (127, 'DEBU', 'DEBU', 'string', 'sys_log_type', 'default', 0, 10, '', 1, '2023-01-20 16:56:23', '2023-01-20 16:56:23'),
                                                                                                                                                                      (128, 'INFO', 'INFO', 'string', 'sys_log_type', 'info', 0, 20, '', 1, '2023-01-20 16:56:34', '2023-01-20 16:56:34'),
                                                                                                                                                                      (129, 'WARN', 'WARN', 'string', 'sys_log_type', 'warning', 0, 30, '', 1, '2023-01-20 16:56:47', '2023-01-20 16:56:47'),
                                                                                                                                                                      (130, 'ERRO', 'ERRO', 'string', 'sys_log_type', 'error', 0, 50, '', 1, '2023-01-20 16:57:01', '2023-01-20 16:57:01'),
                                                                                                                                                                      (131, 'PANI', 'PANI', 'string', 'sys_log_type', 'error', 0, 60, '', 1, '2023-01-20 16:57:23', '2023-01-20 16:57:23'),
                                                                                                                                                                      (132, 'FATA', 'FATA', 'string', 'sys_log_type', 'error', 0, 70, '', 1, '2023-01-20 16:57:37', '2023-01-20 16:57:37'),
                                                                                                                                                                      (133, '阿里云', 'aliyun', 'string', 'config_sms_drive', 'success', 0, 10, '', 1, '2023-01-21 10:27:59', '2023-01-21 10:27:59'),
                                                                                                                                                                      (134, '腾讯云', 'tencent', 'string', 'config_sms_drive', 'info', 0, 20, '', 1, '2023-01-21 10:28:23', '2023-01-21 10:28:23'),
                                                                                                                                                                      (135, '通用验证码', 'code', 'string', 'config_sms_template', 'info', 0, 10, '', 1, '2023-01-21 10:28:39', '2023-01-21 20:40:09'),
                                                                                                                                                                      (136, '登录', 'login', 'string', 'config_sms_template', 'success', 0, 20, '', 1, '2023-01-21 10:28:48', '2023-01-21 10:28:48'),
                                                                                                                                                                      (137, '注册', 'register', 'string', 'config_sms_template', 'warning', 0, 30, '', 1, '2023-01-21 10:28:58', '2023-01-21 10:28:58'),
                                                                                                                                                                      (138, '重置密码', 'resetPwd', 'string', 'config_sms_template', 'error', 0, 40, '', 1, '2023-01-21 10:29:10', '2023-01-21 10:29:10'),
                                                                                                                                                                      (139, '绑定手机号', 'bind', 'string', 'config_sms_template', 'info', 0, 50, '', 1, '2023-01-21 10:29:22', '2023-01-21 10:29:22'),
                                                                                                                                                                      (140, '申请提现', 'cash', 'string', 'config_sms_template', 'primary', 0, 60, '', 1, '2023-01-21 10:29:36', '2023-01-21 10:29:40'),
                                                                                                                                                                      (141, '私信', '3', 'int', 'sys_notice_type', 'info', 0, 30, '', 1, '2023-01-26 12:08:42', '2023-01-26 12:08:42'),
                                                                                                                                                                      (142, '一般', '1', 'int', 'sys_notice_tag', 'info', 0, 10, '', 1, '2023-01-26 12:13:04', '2023-01-26 12:16:07'),
                                                                                                                                                                      (143, '重要', '2', 'int', 'sys_notice_tag', 'error', 0, 20, '', 1, '2023-01-26 12:13:36', '2023-01-26 12:13:36'),
                                                                                                                                                                      (144, '提醒', '3', 'int', 'sys_notice_tag', 'warning', 0, 30, '', 1, '2023-01-26 12:13:52', '2023-01-26 12:13:52'),
                                                                                                                                                                      (145, '注意', '4', 'int', 'sys_notice_tag', 'success', 0, 40, '', 1, '2023-01-26 12:14:12', '2023-01-26 12:14:12'),
                                                                                                                                                                      (146, '调试', '5', 'int', 'sys_notice_tag', 'default', 0, 50, '', 1, '2023-01-26 12:15:59', '2023-01-26 12:15:59'),
                                                                                                                                                                      (147, '通用验证码', 'code', 'string', 'config_ems_template', 'info', 0, 10, '', 1, '2023-01-21 10:28:39', '2023-01-21 20:40:09'),
                                                                                                                                                                      (148, '登录', 'login', 'string', 'config_ems_template', 'success', 0, 20, '', 1, '2023-01-21 10:28:48', '2023-01-21 10:28:48'),
                                                                                                                                                                      (149, '注册', 'register', 'string', 'config_ems_template', 'warning', 0, 30, '', 1, '2023-01-21 10:28:58', '2023-01-21 10:28:58'),
                                                                                                                                                                      (150, '重置密码', 'resetPwd', 'string', 'config_ems_template', 'error', 0, 40, '', 1, '2023-01-21 10:29:10', '2023-01-21 10:29:10'),
                                                                                                                                                                      (151, '绑定手机号', 'bind', 'string', 'config_ems_template', 'info', 0, 50, '', 1, '2023-01-21 10:29:22', '2023-01-21 10:29:22'),
                                                                                                                                                                      (152, '申请提现', 'cash', 'string', 'config_ems_template', 'primary', 0, 60, '', 1, '2023-01-21 10:29:36', '2023-01-21 10:29:40'),
                                                                                                                                                                      (153, '富文本', 'text', 'string', 'config_ems_template', 'default', 0, 0, '自定义编辑邮件内容时使用', 1, '2023-01-21 10:29:36', '2023-05-29 14:30:49'),
                                                                                                                                                                      (154, '腾讯云', 'cos', 'string', 'config_upload_drive', 'info', 1, 30, '', 1, '2022-12-30 17:58:26', '2022-12-30 17:58:26'),
                                                                                                                                                                      (155, '阿里云', 'oss', 'string', 'config_upload_drive', 'info', 1, 40, '', 1, '2022-12-30 17:58:26', '2022-12-30 17:58:26'),
                                                                                                                                                                      (156, '七牛云', 'qiniu', 'string', 'config_upload_drive', 'success', 1, 50, '', 1, '2022-12-30 17:58:26', '2022-12-30 17:58:26'),
                                                                                                                                                                      (157, 'minio', 'minio', 'string', 'config_upload_drive', 'default', 0, 60, '', 1, '2023-11-11 16:45:12', '2023-11-11 16:45:12'),
                                                                                                                                                                      (158, '目录', '1', 'int', 'sys_menu_types', 'success', 1, 10, '', 1, '2021-01-30 13:27:43', '2023-04-27 23:24:58'),
                                                                                                                                                                      (159, '菜单', '2', 'int', 'sys_menu_types', 'primary', 1, 20, '', 1, '2021-01-30 13:27:43', '2023-04-27 23:24:58'),
                                                                                                                                                                      (160, '按钮', '3', 'int', 'sys_menu_types', 'info', 1, 30, '', 1, '2021-01-30 13:27:43', '2023-04-27 23:24:58'),
                                                                                                                                                                      (161, '主目录', 'LAYOUT', 'string', 'sys_menu_component', 'default', 0, 10, '', 1, '2024-04-16 13:35:49', '2024-04-16 13:35:49'),
                                                                                                                                                                      (162, '子目录', 'ParentLayout', 'string', 'sys_menu_component', 'default', 0, 20, '', 1, '2024-04-16 13:36:04', '2024-04-16 13:36:04'),
                                                                                                                                                                      (163, '内嵌链接', 'IFRAME', 'string', 'sys_menu_component', 'default', 0, 30, '', 1, '2024-04-16 13:36:30', '2024-04-16 13:36:30'),
                                                                                                                                                                      (164, '登录成功', '1', 'int', 'sys_login_status', 'success', 0, 10, '', 1, '2024-04-16 15:21:41', '2024-04-16 15:21:41'),
                                                                                                                                                                      (165, '登录失败', '2', 'int', 'sys_login_status', 'warning', 0, 20, '', 1, '2024-04-16 15:22:02', '2024-04-16 15:22:02');
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