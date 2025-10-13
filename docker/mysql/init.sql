-- 创建 Tars 数据库
CREATE DATABASE IF NOT EXISTS db_tars;

-- 创建应用额外数据库（如果需要）
CREATE DATABASE IF NOT EXISTS sajunablog_test;

-- 为用户授权
GRANT ALL PRIVILEGES ON db_tars.* TO 'sajuna'@'%';
GRANT ALL PRIVILEGES ON sajunablog_test.* TO 'sajuna'@'%';
FLUSH PRIVILEGES;