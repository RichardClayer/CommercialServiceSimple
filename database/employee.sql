-- 员工表
CREATE TABLE employees
(
    id            bigint primary key auto_increment,
    name          varchar(128) default ''                not null comment '姓名',
    username      varchar(255) default ''                not null comment '登录账号',
    password      varchar(255) default ''                not null comment '密码',
    is_refundable tinyint      default 0                 not null comment '是否可以退款',
    is_forbidden  tinyint      default 0                 not null comment '是否禁用：0-未禁用，1-已禁用',
    position      tinyint      default 0                 not null comment '职位：0-店员，1-店长，3-商户管理员',
    created_at    timestamp,
    updated_at    timestamp    default current_timestamp not null,
    unique index employees_name_uindex (name),
    unique index employees_username_uindex (username)
) ENGINE = INNODB
  DEFAULT CHARACTER SET 'utf8'
  COLLATE utf8_general_ci COMMENT ='员工表';