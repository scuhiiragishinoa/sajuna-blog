@echo off
chcp 65001 >nul
title Sajuna Blog 快速启动

echo.
echo ========================================
echo    Sajuna Blog 快速启动
echo ========================================
echo.
echo 正在启动开发环境...
echo.

cd /d "%~dp0"

REM 启动开发环境
docker-compose -f docker-compose.dev.yml up -d

if %errorlevel% equ 0 (
    echo.
    echo [✓] 开发环境启动成功！
    echo.
    echo 📱 前端地址: http://localhost:3000
    echo 🔧 API地址: http://localhost:8080
    echo 🌐 完整地址: http://localhost:80
    echo.
    echo 按任意键查看服务状态...
    pause >nul
    
    echo.
    echo 服务状态:
    docker-compose -f docker-compose.dev.yml ps
    echo.
    echo 按任意键退出...
    pause >nul
) else (
    echo.
    echo [✗] 启动失败！请检查Docker是否运行
    echo.
    echo 按任意键退出...
    pause >nul
)
