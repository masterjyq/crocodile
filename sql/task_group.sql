CREATE TABLE IF NOT EXISTS `crocodile_task_group` (
  `id` 	(18) NOT NULL COMMENT 'ID',
  `name` varchar(30) NOT NULL COMMENT '任务组名称',
  `remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
  `createByID` char(18) NOT NULL DEFAULT '' COMMENT '创建人ID',
  `createTime` int(11) NOT NULL DEFAULT '0' COMMENT '任务创建时间 时间戳(秒)',
  `updateTime` int(11) NOT NULL DEFAULT '0' COMMENT '任务上次修改时间 时间戳(秒)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
ALTER TABLE crocodile.crocodile_task ADD runType INT(11) DEFAULT 0 NOT NULL COMMENT '0:定时 1:单次 2:持续';