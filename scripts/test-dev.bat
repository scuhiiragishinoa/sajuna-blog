@echo off
chcp 65001 >nul
title Sajuna Blog 开发环境测试

echo.
echo ========================================
echo    Sajuna Blog 开发环境测试
echo ========================================
echo.

cd /d "%~dp0\.."

echo [1/6] 检查容器状态...
docker-compose -f docker-compose.dev.yml ps
echo.

echo [2/6] 测试后端API健康检查...
curl -s http://localhost:8080/api/v1/health
echo.
echo.

echo [3/6] 测试前端服务...
curl -s -I http://localhost:3000 | findstr "200 OK"
if %errorlevel% equ 0 (
    echo [✓] 前端服务正常
) else (
    echo [✗] 前端服务异常
)
echo.

echo [4/6] 测试Nginx代理...
curl -s -I http://localhost:80 | findstr "200 OK"
if %errorlevel% equ 0 (
    echo [✓] Nginx代理正常
) else (
    echo [✗] Nginx代理异常
)
echo.

echo [5/6] 测试API通过Nginx代理...
curl -s http://localhost:80/api/v1/health
echo.
echo.

echo [6/6] 测试数据库连接...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "SELECT 'Database connection successful' as status;" sajuna_blog 2>nul
if %errorlevel% equ 0 (
    echo [✓] 数据库连接正常
) else (
    echo [✗] 数据库连接异常
)
echo.

echo ========================================
echo 测试完成！
echo ========================================
echo.
echo 访问地址:
echo   前端: http://localhost:3000
echo   API:  http://localhost:8080
echo   完整: http://localhost:80
echo.
echo 按任意键退出...
pause >nul
