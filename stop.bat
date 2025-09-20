@echo off
chcp 65001 >nul
title Sajuna Blog 停止服务

echo.
echo ========================================
echo    Sajuna Blog 停止服务
echo ========================================
echo.
echo 正在停止开发环境...
echo.

cd /d "%~dp0"

REM 停止开发环境
docker-compose -f docker-compose.dev.yml down

if %errorlevel% equ 0 (
    echo.
    echo [✓] 开发环境已停止！
    echo.
    echo 按任意键退出...
    pause >nul
) else (
    echo.
    echo [✗] 停止失败！
    echo.
    echo 按任意键退出...
    pause >nul
)
