-- 员工表
CREATE TABLE employees
(
  id            bigint PRIMARY KEY AUTO_INCREMENT,
  name          varchar(128) DEFAULT ''             NOT NULL
  COMMENT '姓名',
  username      varchar(255) DEFAULT ''             NOT NULL
  COMMENT '登录账号',
  password      varchar(255) DEFAULT ''             NOT NULL
  COMMENT '密码',
  is_refundable tinyint DEFAULT 0                   NOT NULL
  COMMENT '是否可以退款',
  created_at    timestamp,
  updated_at    timestamp DEFAULT current_timestamp NOT NULL
);
CREATE UNIQUE INDEX employees_name_uindex
  ON employees (name);
CREATE UNIQUE INDEX employees_username_uindex
  ON employees (username);
CREATE INDEX employees_is_refundable_index
  ON employees (is_refundable);
ALTER TABLE employees
  COMMENT = '员工表';

ALTER TABLE employees
  ADD is_forbidden tinyint DEFAULT 0 NOT NULL
COMMENT '是否禁用：0-未禁用，1-已禁用';
ALTER TABLE employees
  ADD position tinyint DEFAULT 0 NOT NULL
COMMENT '职位：0-店员，1-店长，3-商户管理员';
ALTER TABLE employees
  MODIFY COLUMN updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
  AFTER is_forbidden,
  MODIFY COLUMN created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
  AFTER is_forbidden;

CREATE INDEX employees_is_forbidden_index
  ON employees (is_forbidden);
CREATE INDEX employees_position_index
  ON employees (position);