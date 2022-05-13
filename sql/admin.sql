DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
                  `user_name` VARCHAR(40) COMMENT '用户姓名',
                  `codeforces_id` varchar(40) comment 'cf账号',
                  `key` varchar(40) comment 'api生成key',
                  `secret` varchar(40) comment 'api生成secret',
                  PRIMARY KEY (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='集体表格汇总';
