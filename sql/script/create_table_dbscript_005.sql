CREATE TABLE `lds_ops`.`ops_dbscript` (
	`scriptname` varchar(150) NOT NULL,
	`app` varchar(50) NOT NULL,
	`appversion` varchar(50) NOT NULL,
	`createtime` datetime,
	`filepath` varchar(150),
	PRIMARY KEY (`scriptname`, `app`, `appversion`)
) COMMENT='对应服务的数据库脚本';

ALTER TABLE `lds_ops`.`ops_dbscript` ADD COLUMN `branch` varchar(50) AFTER `filepath`, ADD COLUMN `comment` varchar(150) AFTER `branch`;