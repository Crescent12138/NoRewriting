
DROP TABLE IF EXISTS `submission`;
CREATE TABLE `submission` (
                              `Id` INT(10) NOT NULL AUTO_INCREMENT COMMENT 'cf主键ID',
                              `user_id` VARCHAR(40) COMMENT 'cf账号ID',
                              `contest_id` int(10) comment '比赛号',
                              `problem_index` varchar(5) comment '题目编号',
                              `problem_name` varchar(70) comment '题目名称',
                              `rating` int(10) comment  '难度',
                              PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='集体表格汇总';
alter table submission add unique key `name_problem` (`user_id`,`contest_id`,`problem_index`);

ALTER TABLE  admin  ADD FOREIGN KEY products_vendors_fk_1 (user_name) REFERENCES submission (user_id);


drop table if exists `cf_problem_set`;

create table `cf_problem_set`
(
    Id        int unsigned auto_increment
        primary key,
    contestId int          null,
    `index`   varchar(20)  null,
    name      varchar(100) null,
    type      varchar(100) null,
    rating    int          null
)
    charset = utf8
    auto_increment = 1590;
alter table cf_problem_set add unique key `name_problem` (`contestId`,`index`);
