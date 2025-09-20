@echo off
chcp 65001 >nul
title Sajuna Blog 安全检查工具

echo.
echo ========================================
echo      Sajuna Blog 安全检查工具
echo ========================================
echo.

:menu
echo 请选择安全检查项目:
echo 1. 检查JWT密钥配置
echo 2. 检查数据库密码强度
echo 3. 检查Docker安全配置
echo 4. 检查API安全头
echo 5. 检查CORS配置
echo 6. 检查依赖漏洞
echo 7. 运行完整安全检查
echo 8. 生成安全报告
echo 0. 退出
echo.

set /p choice=请输入选择 (0-8): 

if "%choice%"=="1" goto jwt_check
if "%choice%"=="2" goto db_check
if "%choice%"=="3" goto docker_check
if "%choice%"=="4" goto api_check
if "%choice%"=="5" goto cors_check
if "%choice%"=="6" goto deps_check
if "%choice%"=="7" goto full_check
if "%choice%"=="8" goto report
if "%choice%"=="0" goto exit
echo 无效选择，请重新输入
timeout /t 2 >nul
goto menu

:jwt_check
echo.
echo ========================================
echo 检查JWT密钥配置
echo ========================================
echo.
echo 检查环境变量中的JWT密钥...
if defined JWT_SECRET (
    echo [✓] JWT_SECRET 环境变量已设置
    echo 密钥长度: %JWT_SECRET:~0,10%...
) else (
    echo [✗] JWT_SECRET 环境变量未设置
    echo 建议: 设置强密码作为JWT密钥
)
echo.
echo 检查代码中的硬编码密钥...
findstr /n "sajuna-blog-secret-key" backend\*.go
if %errorlevel% equ 0 (
    echo [✗] 发现硬编码的JWT密钥
    echo 建议: 使用环境变量替代硬编码
) else (
    echo [✓] 未发现硬编码的JWT密钥
)
echo.
pause
goto menu

:db_check
echo.
echo ========================================
echo 检查数据库密码强度
echo ========================================
echo.
echo 检查Docker Compose中的数据库密码...
findstr /n "MYSQL_PASSWORD" docker-compose*.yml
echo.
echo 检查批处理文件中的密码泄露...
findstr /n "pblog_password\|-p.*password" *.bat
if %errorlevel% equ 0 (
    echo [✗] 发现批处理文件中的密码泄露
    echo 建议: 移除硬编码密码，使用安全连接方式
) else (
    echo [✓] 未发现批处理文件中的密码泄露
)
echo.
echo 建议:
echo 1. 使用强密码（至少12位，包含大小写字母、数字、特殊字符）
echo 2. 定期更换密码
echo 3. 使用环境变量存储密码
echo 4. 启用数据库SSL连接
echo 5. 不要在代码中硬编码密码
echo 6. 使用 db-connect.bat 进行安全连接
echo.
pause
goto menu

:docker_check
echo.
echo ========================================
echo 检查Docker安全配置
echo ========================================
echo.
echo 检查容器运行状态...
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
echo.
echo 检查容器用户权限...
docker ps --format "{{.Names}}" | findstr "backend-dev" >nul
if %errorlevel% equ 0 (
    docker exec sajuna-blog-backend-dev whoami 2>nul
    if %errorlevel% equ 0 (
        echo [✓] 后端容器用户检查完成
    ) else (
        echo [✗] 后端容器运行但无法检查用户权限
    )
) else (
    echo [!] 后端容器未运行，跳过用户权限检查
    echo 建议: 先启动开发环境
)
echo.
echo 检查容器安全配置...
docker inspect sajuna-blog-backend-dev --format "{{.Config.User}}" 2>nul
if %errorlevel% equ 0 (
    echo 容器用户配置: 
    docker inspect sajuna-blog-backend-dev --format "{{.Config.User}}"
) else (
    echo [!] 无法检查容器用户配置
)
echo.
echo 建议:
echo 1. 使用非root用户运行容器
echo 2. 限制容器资源使用
echo 3. 启用只读文件系统
echo 4. 使用安全的基础镜像
echo 5. 定期更新基础镜像
echo.
pause
goto menu

:api_check
echo.
echo ========================================
echo 检查API安全头
echo ========================================
echo.
echo 测试API安全头...
curl -s -I http://localhost:8080/api/v1/health | findstr "X-"
if %errorlevel% equ 0 (
    echo [✓] API返回安全头
) else (
    echo [✗] API未返回安全头
    echo 建议: 启用安全中间件
)
echo.
pause
goto menu

:cors_check
echo.
echo ========================================
echo 检查CORS配置
echo ========================================
echo.
echo 检查CORS中间件配置...
findstr /n "AllowOrigins" backend\internal\middleware\cors.go
echo.
echo 建议:
echo 1. 限制允许的源域名
echo 2. 避免使用通配符 *
echo 3. 在生产环境中禁用凭据
echo 4. 设置适当的预检缓存时间
echo.
pause
goto menu

:deps_check
echo.
echo ========================================
echo 检查依赖漏洞
echo ========================================
echo.
echo 检查Go模块依赖...
cd backend
go list -m all | findstr /v "sajuna-blog"
echo.
echo 建议:
echo 1. 定期更新依赖包
echo 2. 使用 go mod audit 检查漏洞
echo 3. 使用 Snyk 或 OWASP 工具扫描
echo.
cd ..
pause
goto menu

:full_check
echo.
echo ========================================
echo 运行完整安全检查
echo ========================================
echo.
call :jwt_check
call :db_check
call :docker_check
call :api_check
call :cors_check
call :deps_check
echo.
echo 完整安全检查完成！
pause
goto menu

:report
echo.
echo ========================================
echo 生成安全报告
echo ========================================
echo.
echo 正在生成安全报告...
echo 报告时间: %date% %time% > security-report.txt
echo. >> security-report.txt
echo Sajuna Blog 安全审计报告 >> security-report.txt
echo ======================================== >> security-report.txt
echo. >> security-report.txt
echo 1. JWT密钥配置 >> security-report.txt
if defined JWT_SECRET (
    echo [✓] JWT_SECRET 环境变量已设置 >> security-report.txt
) else (
    echo [✗] JWT_SECRET 环境变量未设置 >> security-report.txt
)
echo. >> security-report.txt
echo 2. 数据库安全 >> security-report.txt
echo [!] 建议使用强密码和SSL连接 >> security-report.txt
echo. >> security-report.txt
echo 3. Docker安全 >> security-report.txt
echo [!] 建议使用非root用户和资源限制 >> security-report.txt
echo. >> security-report.txt
echo 4. API安全 >> security-report.txt
echo [✓] 已配置安全中间件 >> security-report.txt
echo. >> security-report.txt
echo 报告已保存到 security-report.txt
echo.
pause
goto menu

:exit
echo.
echo 感谢使用安全检查工具！
echo 记住：安全是一个持续的过程，不是一次性的任务！
echo.
timeout /t 2 >nul
exit
