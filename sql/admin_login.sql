use code_forces;
DROP TABLE IF EXISTS `admin_login`;
CREATE TABLE `admin_login` (
                         `user_id` VARCHAR(40) COMMENT '用户ID',
                         `password` varchar(40) comment '登陆密码',
                         PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='集体表格汇总';
