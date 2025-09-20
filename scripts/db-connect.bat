@echo off
chcp 65001 >nul
title 安全数据库连接

echo.
echo ========================================
echo      安全数据库连接工具
echo ========================================
echo.
echo [安全提示] 此工具不会在日志中记录密码
echo.

:menu
echo 请选择连接方式:
echo 1. 使用环境变量中的密码
echo 2. 手动输入密码
echo 3. 使用Docker环境变量
echo 0. 退出
echo.

set /p choice=请输入选择 (0-3): 

if "%choice%"=="1" goto env_connect
if "%choice%"=="2" goto manual_connect
if "%choice%"=="3" goto docker_connect
if "%choice%"=="0" goto exit
echo 无效选择，请重新输入
timeout /t 2 >nul
goto menu

:env_connect
echo.
echo 使用环境变量连接数据库...
if defined DB_PASSWORD (
    docker exec -it sajuna-blog-mysql-dev mysql -u blog_user -p%DB_PASSWORD% sajuna_blog
) else (
    echo [错误] DB_PASSWORD 环境变量未设置
    echo 请先设置环境变量或选择其他连接方式
    pause
    goto menu
)
goto menu

:manual_connect
echo.
echo 手动输入密码连接数据库...
echo [安全提示] 密码输入时不会显示在屏幕上
docker exec -it sajuna-blog-mysql-dev mysql -u blog_user -p sajuna_blog
goto menu

:docker_connect
echo.
echo 使用Docker环境变量连接数据库...
echo [安全提示] 请确保已设置MYSQL_PWD环境变量
if defined MYSQL_PWD (
    docker exec -e MYSQL_PWD=%MYSQL_PWD% sajuna-blog-mysql-dev mysql -u blog_user sajuna_blog
) else (
    echo [错误] MYSQL_PWD 环境变量未设置
    echo 请先设置环境变量: set MYSQL_PWD=your_password
    pause
)
goto menu

:exit
echo.
echo 感谢使用安全数据库连接工具！
echo.
timeout /t 2 >nul
exit
