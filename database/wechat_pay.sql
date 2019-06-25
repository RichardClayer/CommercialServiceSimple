-- 微信支付商户信息
CREATE TABLE wechat_pay
(
  id              int PRIMARY KEY AUTO_INCREMENT,
  app_id          varchar(255) NOT NULL COMMENT '微信公众号id',
  app_secret      varchar(255) NOT NULL COMMENT '微信公众号密钥',
  mch_id          varchar(255) NOT NULL COMMENT '微信商户号',
  `key`           varchar(255) NOT NULL COMMENT '微信支付API安全密钥',
  api_client_cert text COMMENT 'API安全证书公钥',
  api_client_key  text COMMENT 'API安全证书密钥'
) ENGINE = INNODB
  DEFAULT CHARACTER SET 'utf8'
  COLLATE utf8_general_ci
  COMMENT = '微信支付商户配置';