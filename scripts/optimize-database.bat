@echo off
chcp 65001 >nul
echo ========================================
echo    Sajuna Blog 数据库优化工具
echo ========================================
echo.

echo 正在优化数据库性能...
echo.

REM 检查MySQL是否运行
docker exec sajuna-blog-mysql-dev mysqladmin ping -h localhost -u blog_user -pblog_password >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ MySQL容器未运行，请先启动数据库
    echo 运行: docker-compose -f docker-compose.dev.yml up -d mysql-dev
    pause
    exit /b 1
)

echo ✅ MySQL容器运行正常
echo.

echo 1. 创建文章表索引...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
CREATE INDEX IF NOT EXISTS idx_articles_status_created ON articles(status, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_articles_category_status ON articles(category_id, status);
CREATE INDEX IF NOT EXISTS idx_articles_user_status ON articles(user_id, status);
CREATE INDEX IF NOT EXISTS idx_articles_title ON articles(title);
"
echo ✅ 文章表索引创建完成
echo.

echo 2. 创建用户表索引...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
"
echo ✅ 用户表索引创建完成
echo.

echo 3. 创建分类表索引...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
CREATE INDEX IF NOT EXISTS idx_categories_name ON categories(name);
"
echo ✅ 分类表索引创建完成
echo.

echo 4. 创建标签表索引...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
CREATE INDEX IF NOT EXISTS idx_tags_name ON tags(name);
"
echo ✅ 标签表索引创建完成
echo.

echo 5. 创建评论表索引...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
CREATE INDEX IF NOT EXISTS idx_comments_article_created ON comments(article_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_comments_user_created ON comments(user_id, created_at DESC);
"
echo ✅ 评论表索引创建完成
echo.

echo 6. 创建文章标签关联表索引...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
CREATE INDEX IF NOT EXISTS idx_article_tags_article ON article_tags(article_id);
CREATE INDEX IF NOT EXISTS idx_article_tags_tag ON article_tags(tag_id);
"
echo ✅ 文章标签关联表索引创建完成
echo.

echo 7. 创建全文搜索索引...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
ALTER TABLE articles ADD FULLTEXT(title, content, summary);
"
echo ✅ 全文搜索索引创建完成
echo.

echo 8. 优化MySQL配置...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
SET GLOBAL query_cache_type = 1;
SET GLOBAL query_cache_size = 67108864;
SET GLOBAL query_cache_limit = 2097152;
SET GLOBAL slow_query_log = 1;
SET GLOBAL long_query_time = 2;
"
echo ✅ MySQL配置优化完成
echo.

echo 9. 分析表统计信息...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
ANALYZE TABLE articles;
ANALYZE TABLE users;
ANALYZE TABLE categories;
ANALYZE TABLE tags;
ANALYZE TABLE comments;
ANALYZE TABLE article_tags;
"
echo ✅ 表统计信息更新完成
echo.

echo 10. 测试优化效果...
echo 测试文章列表查询性能...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
EXPLAIN SELECT * FROM articles WHERE status = 'published' ORDER BY created_at DESC LIMIT 10;
"
echo.

echo 测试搜索查询性能...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
EXPLAIN SELECT * FROM articles WHERE MATCH(title, content, summary) AGAINST('vue' IN NATURAL LANGUAGE MODE);
"
echo.

echo ========================================
echo 数据库优化完成！
echo ========================================
echo.
echo ✅ 已完成的优化:
echo 1. 创建了所有必要的索引
echo 2. 启用了查询缓存
echo 3. 开启了慢查询日志
echo 4. 更新了表统计信息
echo 5. 创建了全文搜索索引
echo.
echo 💡 建议:
echo 1. 运行 analyze-slow-sql.bat 查看优化效果
echo 2. 查看 docs/数据库优化指南.md 获取更多优化建议
echo 3. 定期监控数据库性能
echo.

pause

