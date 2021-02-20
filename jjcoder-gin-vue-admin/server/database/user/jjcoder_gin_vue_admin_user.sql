CREATE TABLE IF NOT EXISTS user
(
    id       int UNSIGNED AUTO_INCREMENT COMMENT '用户编号'
        PRIMARY KEY,
    username varchar(20)                                      NOT NULL COMMENT '用户名',
    password char(32)                                         NOT NULL COMMENT '密码',
    email    varchar(50)                                      NOT NULL COMMENT '邮箱',
    age      tinyint UNSIGNED               DEFAULT '18'      NOT NULL COMMENT '年龄',
    sex      enum ('man', 'woman', 'baomi') DEFAULT 'baomi'   NOT NULL COMMENT '性别',
    tel      char(11)                                         NOT NULL COMMENT '电话',
    addr     varchar(50)                    DEFAULT 'beijing' NOT NULL COMMENT '地址',
    card     char(18)                                         NOT NULL COMMENT '身份证号',
    married  tinyint(1)                     DEFAULT 0         NOT NULL COMMENT '0代表未结婚，1代表已结婚',
    salary   float(8, 2)                    DEFAULT 0.00      NOT NULL COMMENT '薪水',
    CONSTRAINT card
        UNIQUE (card),
    CONSTRAINT email
        UNIQUE (email),
    CONSTRAINT tel
        UNIQUE (tel),
    CONSTRAINT username
        UNIQUE (username)
);