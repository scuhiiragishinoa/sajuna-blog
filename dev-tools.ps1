function Start-Dev {
    Write-Host "🚀 启动 SajunaBlog 开发环境..." -ForegroundColor Cyan
    docker-compose up -d
    Write-Host "✅ 服务已启动" -ForegroundColor Green
    Show-Status
}

function Stop-Dev {
    Write-Host "🛑 停止开发环境..." -ForegroundColor Yellow
    docker-compose down
    Write-Host "✅ 服务已停止" -ForegroundColor Green
}

function Restart-Dev {
    Write-Host "🔄 重启开发环境..." -ForegroundColor Magenta
    docker-compose restart
    Write-Host "✅ 服务已重启" -ForegroundColor Green
}

function Build-Dev {
    Write-Host "🔨 重新构建服务..." -ForegroundColor Green
    docker-compose build --no-cache
    Write-Host "✅ 构建完成" -ForegroundColor Green
}

function Show-Status {
    Write-Host "📊 服务状态:" -ForegroundColor Yellow
    docker-compose ps
}

function Logs-Backend {
    Write-Host "📝 后端日志 (Ctrl+C 退出)..." -ForegroundColor Cyan
    docker-compose logs -f backend
}

function Logs-Frontend {
    Write-Host "📝 前端日志 (Ctrl+C 退出)..." -ForegroundColor Cyan
    docker-compose logs -f frontend
}

function Exec-Backend {
    Write-Host "💻 进入后端容器..." -ForegroundColor Magenta
    docker-compose exec backend bash
}

function Exec-Frontend {
    Write-Host "💻 进入前端容器..." -ForegroundColor Magenta
    docker-compose exec frontend sh
}

function Backend-Shell {
    Write-Host "🔧 后端调试 Shell..." -ForegroundColor Green
    docker-compose exec backend bash -c "cd /app && bash"
}

function Install-Backend-Dep {
    param([string]$package)
    if ($package) {
        Write-Host "📦 安装后端依赖: $package" -ForegroundColor Cyan
        docker-compose exec backend go get $package
    } else {
        Write-Host "📦 下载所有后端依赖..." -ForegroundColor Cyan
        docker-compose exec backend go mod download
    }
}

function Frontend-Install {
    Write-Host "📦 安装前端依赖..." -ForegroundColor Cyan
    docker-compose exec frontend npm install
}

function Clean-All {
    Write-Host "🧹 清理所有容器和镜像..." -ForegroundColor Yellow
    docker-compose down -v --rmi all
    docker system prune -f
    Write-Host "✅ 清理完成" -ForegroundColor Green
}

# 显示帮助信息
function Show-Help {
    Write-Host "SajunaBlog 开发工具集" -ForegroundColor Cyan
    Write-Host "=====================" -ForegroundColor Cyan
    Write-Host "Start-Dev        - 启动开发环境"
    Write-Host "Stop-Dev         - 停止开发环境"
    Write-Host "Restart-Dev      - 重启开发环境"
    Write-Host "Build-Dev        - 重新构建服务"
    Write-Host "Show-Status      - 显示服务状态"
    Write-Host "Logs-Backend     - 查看后端日志"
    Write-Host "Logs-Frontend    - 查看前端日志"
    Write-Host "Exec-Backend     - 进入后端容器"
    Write-Host "Exec-Frontend    - 进入前端容器"
    Write-Host "Backend-Shell    - 后端调试 Shell"
    Write-Host "Install-Backend-Dep - 安装后端依赖"
    Write-Host "Frontend-Install - 安装前端依赖"
    Write-Host "Clean-All        - 清理所有容器和镜像"
}

# 自动显示帮助
Show-Help