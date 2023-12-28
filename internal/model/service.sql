USE `remote_database`;
CREATE TABLE IF NOT EXISTS `users`
(
    `user_id`   INT AUTO_INCREMENT COMMENT '用户ID',
    `user_name` VARCHAR(50)  NOT NULL DEFAULT '' COMMENT '用户名',
    `password`  VARCHAR(100) NOT NULL DEFAULT '' COMMENT '密码',
    `email`     VARCHAR(100) NOT NULL DEFAULT '' COMMENT '邮箱',
    `phone`    VARCHAR(50) NOT NULL DEFAULT '' COMMENT '手机号',
    `points`   INT NOT NULL DEFAULT 0 COMMENT '积分',
    `role_type`  INT NOT NULL DEFAULT 0 COMMENT '角色类型:1学生，2管理员',
    `created_at` INT NOT NULL DEFAULT 0 COMMENT '创建时间',
    `updated_at` INT NOT NULL DEFAULT 0 COMMENT '更新时间',
    `status` INT NOT NULL DEFAULT 0 COMMENT '状态:0正常、1禁用',
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT '用户表';

CREATE TABLE IF NOT EXISTS `courses`
(
    `course_id`   INT AUTO_INCREMENT COMMENT '课程ID',
    `user_id` INT  NOT NULL DEFAULT 0 COMMENT '用户ID；操作人',
    `course_name`  VARCHAR(100) NOT NULL DEFAULT '' COMMENT '课程名称',
    `description`  VARCHAR(100) NOT NULL DEFAULT '' COMMENT '课程描述',
    `price` DECIMAL(10, 2) NOT NULL DEFAULT 0.00 COMMENT '课程价格',
    `created_at` INT NOT NULL DEFAULT 0 COMMENT '创建时间',
    `updated_at` INT NOT NULL DEFAULT 0 COMMENT '更新时间',
    `status` INT NOT NULL DEFAULT 0 COMMENT '状态:0正常、1禁用',
    PRIMARY KEY (`course_id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4 COMMENT '课程表';

CREATE TABLE IF NOT EXISTS `course_categories`
(
    `category_id`   INT AUTO_INCREMENT COMMENT '课程类别ID',
    `category_name`  VARCHAR(100) NOT NULL DEFAULT '' COMMENT '课程类别名称',
    `created_at` INT NOT NULL DEFAULT 0 COMMENT '创建时间',
    `updated_at` INT NOT NULL DEFAULT 0 COMMENT '更新时间',
    `status` INT NOT NULL DEFAULT 0 COMMENT '状态:0正常、1禁用',
    PRIMARY KEY (`category_id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4 COMMENT '课程类别表';

CREATE TABLE IF NOT EXISTS `course_category_ship`
(
    `id`   INT AUTO_INCREMENT COMMENT '关系ID',
    `course_id` INT  NOT NULL DEFAULT 0 COMMENT '课程ID',
    `category_id` INT  NOT NULL DEFAULT 0 COMMENT '课程类别ID',
    `created_at` INT NOT NULL DEFAULT 0 COMMENT '创建时间',
    `updated_at` INT NOT NULL DEFAULT 0 COMMENT '更新时间',
    `status` INT NOT NULL DEFAULT 0 COMMENT '状态:0正常、1禁用',
    PRIMARY KEY (`id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4 COMMENT '课程类别关联表';

CREATE TABLE IF NOT EXISTS `orders`
(
    `order_id`   INT AUTO_INCREMENT COMMENT '订单ID',
    `user_id` INT  NOT NULL DEFAULT 0 COMMENT '用户ID',
    `course_id` INT  NOT NULL DEFAULT 0 COMMENT '课程ID',
    `purchase_count` INT NOT NULL DEFAULT 0 COMMENT '购买数量',
    `total_price` DECIMAL(10, 2) NOT NULL DEFAULT 0.00 COMMENT '订单总金额',
    `created_at` INT NOT NULL DEFAULT 0 COMMENT '创建时间',
    `updated_at` INT NOT NULL DEFAULT 0 COMMENT '更新时间',
    `status` INT NOT NULL DEFAULT 0 COMMENT '状态:0下单、1取消订单',
    PRIMARY KEY (`order_id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4 COMMENT '课程类别关联表';