DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
                  `user_id` VARCHAR(40) COMMENT 'cf账号ID',
                  `password` varchar(40) comment '登陆密码',
                  `key` varchar(40) comment 'api生成key',
                  `secret` varchar(40) comment 'api生成secret',
                  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='集体表格汇总';
