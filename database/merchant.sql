CREATE TABLE merchant
(
    id                     tinyint PRIMARY KEY     NOT NULL AUTO_INCREMENT,
    name                   varchar(255) DEFAULT '' NOT NULL COMMENT '商户全称',
    is_wechat_pay_recorded tinyint      DEFAULT 0  NOT NULL COMMENT '微信商户资料是否已填',
    is_ali_pay_recorded    tinyint      DEFAULT 0  NOT NULL COMMENT '支付宝商户资料是否已填'
) ENGINE = INNODB
  DEFAULT CHARACTER SET 'utf8'
  COLLATE utf8_general_ci COMMENT ='商户表';