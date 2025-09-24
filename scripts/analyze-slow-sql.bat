@echo off
chcp 65001 >nul
echo ========================================
echo    Sajuna Blog 慢SQL分析工具
echo ========================================
echo.

echo 正在分析数据库慢SQL...
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

echo 1. 查看慢查询配置...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "SHOW VARIABLES LIKE 'slow_query%';"
echo.

echo 2. 查看当前查询时间阈值...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "SHOW VARIABLES LIKE 'long_query_time';"
echo.

echo 3. 查看当前正在运行的查询...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "SHOW PROCESSLIST;"
echo.

echo 4. 查看数据库连接状态...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "SHOW STATUS LIKE 'Connections';"
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "SHOW STATUS LIKE 'Threads_connected';"
echo.

echo 5. 查看查询缓存状态...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "SHOW STATUS LIKE 'Qcache%';"
echo.

echo 6. 查看InnoDB状态...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_user -e "SHOW STATUS LIKE 'Innodb%';" | findstr "Buffer_pool"
echo.

echo 7. 分析文章表索引...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "SHOW INDEX FROM sajuna_blog.articles;"
echo.

echo 8. 测试文章查询性能...
echo 执行查询: SELECT COUNT(*) FROM articles WHERE status = 'published';
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "EXPLAIN SELECT COUNT(*) FROM sajuna_blog.articles WHERE status = 'published';"
echo.

echo 9. 测试文章列表查询性能...
echo 执行查询: SELECT * FROM articles WHERE status = 'published' ORDER BY created_at DESC LIMIT 10;
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "EXPLAIN SELECT * FROM sajuna_blog.articles WHERE status = 'published' ORDER BY created_at DESC LIMIT 10;"
echo.

echo 10. 测试搜索查询性能...
echo 执行查询: SELECT * FROM articles WHERE title LIKE '%vue%';
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "EXPLAIN SELECT * FROM sajuna_blog.articles WHERE title LIKE '%vue%';"
echo.

echo ========================================
echo 分析完成！
echo ========================================
echo.
echo 💡 优化建议:
echo 1. 如果看到 'Using filesort' 或 'Using temporary'，需要添加索引
echo 2. 如果看到 'Using where' 但没有使用索引，需要优化WHERE条件
echo 3. 如果查询时间过长，考虑添加复合索引
echo 4. 查看 docs/数据库优化指南.md 获取详细优化方案
echo.

pause

