CREATE TABLE `jelly_record` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(100) DEFAULT '' COMMENT '图案唯一值',
  `execnum` int(10) unsigned DEFAULT '0' COMMENT '执行次数',
  `comb` varchar(100) DEFAULT '' COMMENT '图案信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='果冻消除';