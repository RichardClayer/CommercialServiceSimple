-- 微信支付商户信息
CREATE TABLE wechat_pay
(
  id              INT PRIMARY KEY AUTO_INCREMENT,
  app_id          VARCHAR(255) NOT NULL COMMENT '微信公众号id',
  app_secret      VARCHAR(255) NOT NULL COMMENT '微信公众号密钥',
  mch_id          VARCHAR(255) NOT NULL COMMENT '微信商户号',
  `key`           VARCHAR(255) NOT NULL COMMENT '微信支付API安全密钥',
  api_client_cert TEXT COMMENT 'API安全证书公钥',
  api_client_key  TEXT COMMENT 'API安全证书密钥',
  merchant_id     INT  COMMENT '商户ID',
  index wechat_pay_merchant_id_index (merchant_id)
) ENGINE = INNODB
  DEFAULT CHARACTER SET 'utf8'
  COLLATE utf8_general_ci
  COMMENT = '微信支付商户配置';