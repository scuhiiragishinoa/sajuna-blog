@echo off
chcp 65001 >nul
title Sajuna Blog Development Environment

:main
cls
echo.
echo ========================================
echo      Sajuna Blog Development Environment
echo ========================================
echo.
echo Current Status:
docker-compose -f docker-compose.dev.yml ps 2>nul | findstr "Up" >nul
if %errorlevel% equ 0 (
    echo [Status] Development environment is running
    echo [Frontend] http://localhost:3000
    echo [Backend] http://localhost:8080
    echo [Unified Entry] http://localhost
) else (
    echo [Status] Development environment is stopped
)
echo.
echo ========================================
echo Please select an operation:
echo ========================================
echo 1. Start development environment
echo 2. Stop development environment
echo 3. Restart development environment
echo 4. View service status
echo 5. View real-time logs
echo 6. Enter container for debugging
echo 7. Test API interfaces
echo 8. Clean all data
echo 9. Quick start (Recommended)
echo 0. Exit program
echo.
set /p choice=Please enter your choice (0-9): 

if "%choice%"=="1" goto start
if "%choice%"=="2" goto stop
if "%choice%"=="3" goto restart
if "%choice%"=="4" goto status
if "%choice%"=="5" goto logs
if "%choice%"=="6" goto debug
if "%choice%"=="7" goto test
if "%choice%"=="8" goto clean
if "%choice%"=="9" goto quickstart
if "%choice%"=="0" goto exit
echo Invalid choice, please try again
timeout /t 2 >nul
goto main

:start
echo.
echo Starting development environment...
docker-compose -f docker-compose.dev.yml up -d
if %errorlevel% neq 0 (
    echo [Error] Failed to start, please check if Docker is running
    pause
    goto main
)
echo.
echo Waiting for services to start...
timeout /t 10 /nobreak >nul
echo.
echo [Success] Development environment started!
echo [Frontend] http://localhost:3000
echo [Backend] http://localhost:8080
echo [Unified Entry] http://localhost
echo [Database] localhost:3306
echo [Redis] localhost:6379
echo.
pause
goto main

:stop
echo.
echo Stopping development environment...
docker-compose -f docker-compose.dev.yml down
if %errorlevel% neq 0 (
    echo [Error] Failed to stop
    pause
    goto main
)
echo [Success] Development environment stopped!
echo.
pause
goto main

:restart
echo.
echo Restarting development environment...
docker-compose -f docker-compose.dev.yml restart
if %errorlevel% neq 0 (
    echo [Error] Failed to restart
    pause
    goto main
)
echo [Success] Development environment restarted!
echo.
pause
goto main

:status
echo.
echo Service Status:
echo ========================================
docker-compose -f docker-compose.dev.yml ps
echo.
pause
goto main

:logs
echo.
echo Viewing real-time logs (Press Ctrl+C to exit):
echo ========================================
docker-compose -f docker-compose.dev.yml logs -f
goto main

:debug
echo.
echo Please select a container to enter:
echo 1. Frontend container (Node.js/Vue3)
echo 2. Backend container (Go)
echo 3. Database container (MySQL)
echo 4. Redis container
echo 5. Return to main menu
echo.
set /p debug_choice=Please enter your choice (1-5): 

if "%debug_choice%"=="1" (
    echo Entering frontend container...
    docker exec -it sajuna-blog-frontend-dev sh
) else if "%debug_choice%"=="2" (
    echo Entering backend container...
    docker exec -it sajuna-blog-backend-dev sh
) else if "%debug_choice%"=="3" (
    echo Entering database container...
    echo [安全提示] 请输入数据库密码
    docker exec -it sajuna-blog-mysql-dev mysql -u blog_user -p sajuna_blog
) else if "%debug_choice%"=="4" (
    echo Entering Redis container...
    docker exec -it sajuna-blog-redis-dev redis-cli
) else if "%debug_choice%"=="5" (
    goto main
) else (
    echo Invalid choice
    pause
)
goto main

:test
echo.
echo Testing API interfaces...
echo ========================================
echo Testing health check interface...
curl -s http://localhost:8080/api/v1/health
if %errorlevel% neq 0 (
    echo [Error] API test failed, please check if backend service is running
) else (
    echo.
    echo [Success] API interface is working!
)
echo.
echo Testing frontend page...
curl -s -o nul -w "HTTP Status Code: %%{http_code}\n" http://localhost:3000
if %errorlevel% neq 0 (
    echo [Error] Frontend page test failed, please check if frontend service is running
) else (
    echo [Success] Frontend page is working!
)
echo.
pause
goto main

:clean
echo.
echo [Warning] This operation will delete all data, including database data!
set /p confirm=Are you sure to continue? (y/N): 
if /i not "%confirm%"=="y" (
    echo Operation cancelled
    pause
    goto main
)
echo.
echo Cleaning all data...
docker-compose -f docker-compose.dev.yml down -v
docker system prune -f
echo [Success] Cleanup completed!
echo.
pause
goto main

:quickstart
echo.
echo Quick starting development environment...
echo ========================================
echo 1. Stopping existing services...
docker-compose -f docker-compose.dev.yml down >nul 2>&1
echo 2. Starting all services...
docker-compose -f docker-compose.dev.yml up -d
if %errorlevel% neq 0 (
    echo [Error] Failed to start, please check if Docker is running
    pause
    goto main
)
echo 3. Waiting for services to start...
timeout /t 15 /nobreak >nul
echo 4. Checking service status...
docker-compose -f docker-compose.dev.yml ps
echo.
echo ========================================
echo [Success] Development environment startup completed!
echo ========================================
echo [Frontend] http://localhost:3000
echo [Backend] http://localhost:8080
echo [Unified Entry] http://localhost
echo [Database] localhost:3306
echo [Redis] localhost:6379
echo.
echo Use this script to manage all services
echo Press any key to return to main menu...
pause >nul
goto main

:exit
echo.
echo Safely shutting down all services...
docker-compose -f docker-compose.dev.yml down
echo.
echo Thank you for using Sajuna Blog development environment!
echo Goodbye!
timeout /t 2 >nul
exit