DROP TABLE IF EXISTS `jwtx_tokens`;
CREATE TABLE `jwtx_tokens` (

  `id`                bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'token ID',
  `token_group`       varchar(255)        NOT NULL                COMMENT 'token 分组',
  `account_id`        int(11) unsigned    NOT NULL                COMMENT '账户 ID',
  `make_token_ip`     varchar(50)         NOT NULL                COMMENT '首次请求生成 token 的 IP 地址',
  `created_at`        datetime            NOT NULL                COMMENT '创建时间',
  `last_refresh_at`   datetime            NOT NULL                COMMENT '上次的刷新时间',
  `final_refresh_at`  datetime            NOT NULL                COMMENT '最后的刷新时间',
  `expiration_at`     datetime            NOT NULL                COMMENT '过期时间',

  KEY `token_group` (`token_group`),
  KEY `account_id` (`account_id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='jwt token 信息表';
