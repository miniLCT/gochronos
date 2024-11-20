# 记录图运行状态
DROP TABLE IF EXISTS `graph_state`;
CREATE TABLE `graph_state`
(
    `id`         bigint(10)   NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
    `graph_name` varchar(50)  NOT NULL COMMENT '图名称',
    `uuid`       char(50)     NOT NULL COMMENT '图运行唯一表示',
    `state`      char(20)     NOT NULL COMMENT '图运行状态：start、success、fail',
    `error`      varchar(255) NOT NULL COMMENT '图运行失败信息',
    `create_at`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY `uuid` (`uuid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='图运行状态';

# 记录节点的运行状态
DROP TABLE IF EXISTS `node_state`;
CREATE TABLE `node_state`
(
    `id`         bigint(10)   NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
    `graph_name` varchar(50)  NOT NULL COMMENT '图名称',
    `node_name`  varchar(50)  NOT NULL COMMENT '节点名称',
    `uuid`       char(50)     NOT NULL COMMENT '图运行唯一表示',
    `state`      char(20)     NOT NULL COMMENT '图运行状态：start、success、fail',
    `error`      varchar(255) NOT NULL COMMENT '图运行失败信息',
    `create_at`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY `uuid` (`uuid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='节点运行状态';


# 记录节点的运行数据
DROP TABLE IF EXISTS `node_data`;
CREATE TABLE `node_data`
(
    `id`         bigint(10)   NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
    `graph_name` varchar(50)  NOT NULL COMMENT '图名称',
    `node_key`   varchar(50)  NOT NULL COMMENT '数据存储key',
    `uuid`       char(50)     NOT NULL COMMENT '图运行唯一表示',
    `data`       MEDIUMBLOB   NOT NULL COMMENT '节点输出数据',
    `error`      varchar(255) NOT NULL COMMENT '图运行失败信息',
    `create_at`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_uuid_node_key` (`uuid`,`node_key`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='节点运行数据';
