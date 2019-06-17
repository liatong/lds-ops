CREATE TABLE `lds_ops`.`<table_name>` (
  `application_name` varchar(64) NOT NULL,
  `env` varchar(64) NOT NULL,
  `version` varchar(64) NOT NULL,
  `filename` varchar(64) NOT NULL,
  `mdcode` varchar(64) DEFAULT NULL,
  `filepath` varchar(128) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `upload_time` datetime DEFAULT NULL,
  PRIMARY KEY (`filename`, `application_name`, `version`, `env`),
  CONSTRAINT `application` FOREIGN KEY (`application_name`) REFERENCES `lds_ops`.`ops_application` (`application_name`),
  INDEX `application_name` USING BTREE (`application_name`) comment ''
) ENGINE=`InnoDB` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci ROW_FORMAT=DYNAMIC COMMENT='' CHECKSUM=0 DELAY_KEY_WRITE=0;