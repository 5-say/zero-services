DROP TABLE IF EXISTS `jwtx_tokens`;
CREATE TABLE `jwtx_tokens` (

  `id`                bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'token ID',
  `token_group`       varchar(255)        NOT NULL                COMMENT 'token 分组',
  `random_account_id` varchar(255)        NOT NULL                COMMENT '加密的账户 ID',
  `make_token_ip`     varchar(50)         NOT NULL                COMMENT '首次请求生成 token 的 IP 地址',
  `created_at`        datetime            NOT NULL                COMMENT '创建时间',
  `refresh_at`        datetime            NOT NULL                COMMENT '刷新时间',
  `expiration_at`     datetime            NOT NULL                COMMENT '过期时间',

  KEY `token_group` (`token_group`),
  KEY `random_account_id` (`random_account_id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='jwt token 信息表';
