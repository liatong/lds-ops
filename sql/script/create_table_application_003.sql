CREATE TABLE `lds_ops`.`ops_application` (
  `id` INT NOT NULL,
  `application_name` VARCHAR(45) NULL,
  PRIMARY KEY (`id`));
  
ALTER TABLE `lds_ops`.`ops_application` 
CHANGE COLUMN `application_name` `application_name` VARCHAR(45) NOT NULL ,
DROP PRIMARY KEY,
ADD PRIMARY KEY (`id`, `application_name`);
;
