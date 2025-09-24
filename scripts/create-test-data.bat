@echo off
chcp 65001 >nul
echo ========================================
echo    Sajuna Blog 测试数据生成工具
echo ========================================
echo.

echo 正在生成测试数据...
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

echo 1. 创建测试用户...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
INSERT IGNORE INTO users (username, email, password_hash, created_at, updated_at) VALUES
('testuser1', 'test1@example.com', '\$2a\$10\$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', NOW(), NOW()),
('testuser2', 'test2@example.com', '\$2a\$10\$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', NOW(), NOW()),
('testuser3', 'test3@example.com', '\$2a\$10\$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', NOW(), NOW()),
('testuser4', 'test4@example.com', '\$2a\$10\$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', NOW(), NOW()),
('testuser5', 'test5@example.com', '\$2a\$10\$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', NOW(), NOW());
"
echo ✅ 测试用户创建完成
echo.

echo 2. 创建测试分类...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
INSERT IGNORE INTO categories (name, description, color, created_at, updated_at) VALUES
('前端开发', '前端技术相关文章', '#409EFF', NOW(), NOW()),
('后端开发', '后端技术相关文章', '#67C23A', NOW(), NOW()),
('数据库', '数据库相关文章', '#E6A23C', NOW(), NOW()),
('DevOps', '运维相关文章', '#F56C6C', NOW(), NOW()),
('人工智能', 'AI相关文章', '#909399', NOW(), NOW());
"
echo ✅ 测试分类创建完成
echo.

echo 3. 创建测试标签...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
INSERT IGNORE INTO tags (name, color, created_at, updated_at) VALUES
('Vue3', '#4FC08D', NOW(), NOW()),
('React', '#61DAFB', NOW(), NOW()),
('Angular', '#DD0031', NOW(), NOW()),
('Node.js', '#339933', NOW(), NOW()),
('Go', '#00ADD8', NOW(), NOW()),
('Python', '#3776AB', NOW(), NOW()),
('MySQL', '#4479A1', NOW(), NOW()),
('Redis', '#DC382D', NOW(), NOW()),
('Docker', '#2496ED', NOW(), NOW()),
('Kubernetes', '#326CE5', NOW(), NOW());
"
echo ✅ 测试标签创建完成
echo.

echo 4. 创建测试文章...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
INSERT IGNORE INTO articles (title, content, summary, status, user_id, category_id, view_count, like_count, comment_count, created_at, updated_at) VALUES
('Vue3 开发指南', '# Vue3 开发指南\n\nVue3是Vue.js的最新版本，带来了许多新特性和改进...', '详细介绍Vue3的新特性和开发技巧', 'published', 1, 1, 150, 25, 8, NOW(), NOW()),
('React Hooks 深入理解', '# React Hooks 深入理解\n\nReact Hooks是React 16.8引入的新特性...', '深入理解React Hooks的原理和使用方法', 'published', 2, 1, 200, 30, 12, NOW(), NOW()),
('Go语言并发编程', '# Go语言并发编程\n\nGo语言以其强大的并发能力而闻名...', '学习Go语言的并发编程模式和最佳实践', 'published', 3, 2, 180, 20, 6, NOW(), NOW()),
('MySQL性能优化', '# MySQL性能优化\n\n数据库性能优化是后端开发的重要技能...', 'MySQL数据库性能优化的实用技巧', 'published', 4, 3, 120, 15, 4, NOW(), NOW()),
('Docker容器化实践', '# Docker容器化实践\n\nDocker已经成为现代应用部署的标准...', 'Docker容器化的实践经验和最佳实践', 'published', 5, 4, 90, 10, 3, NOW(), NOW()),
('Vue3 组合式API', '# Vue3 组合式API\n\n组合式API是Vue3的核心特性之一...', 'Vue3组合式API的详细使用指南', 'published', 1, 1, 80, 12, 5, NOW(), NOW()),
('React 状态管理', '# React 状态管理\n\n在大型React应用中，状态管理是一个重要话题...', 'React应用中的状态管理方案对比', 'published', 2, 1, 110, 18, 7, NOW(), NOW()),
('Go微服务架构', '# Go微服务架构\n\n微服务架构是现代应用开发的主流模式...', '使用Go语言构建微服务架构的实践', 'published', 3, 2, 95, 14, 4, NOW(), NOW()),
('Redis缓存策略', '# Redis缓存策略\n\n缓存是提升应用性能的重要手段...', 'Redis缓存的策略和最佳实践', 'published', 4, 3, 75, 8, 2, NOW(), NOW()),
('Kubernetes部署指南', '# Kubernetes部署指南\n\nKubernetes是容器编排的事实标准...', 'Kubernetes部署和管理的完整指南', 'published', 5, 4, 60, 6, 1, NOW(), NOW());
"
echo ✅ 测试文章创建完成
echo.

echo 5. 创建文章标签关联...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
INSERT IGNORE INTO article_tags (article_id, tag_id) VALUES
(1, 1), (1, 7), (1, 9),
(2, 2), (2, 7), (2, 9),
(3, 5), (3, 7), (3, 9),
(4, 7), (4, 8),
(5, 9), (5, 10),
(6, 1), (6, 7),
(7, 2), (7, 7),
(8, 5), (8, 9),
(9, 8), (9, 9),
(10, 9), (10, 10);
"
echo ✅ 文章标签关联创建完成
echo.

echo 6. 创建测试评论...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
INSERT IGNORE INTO comments (content, status, article_id, user_id, created_at, updated_at) VALUES
('写得很好！', 'approved', 1, 2, NOW(), NOW()),
('学到了很多，谢谢分享', 'approved', 1, 3, NOW(), NOW()),
('Vue3确实很棒', 'approved', 1, 4, NOW(), NOW()),
('React Hooks很实用', 'approved', 2, 1, NOW(), NOW()),
('Go的并发确实强大', 'approved', 3, 2, NOW(), NOW()),
('MySQL优化很重要', 'approved', 4, 3, NOW(), NOW()),
('Docker很方便', 'approved', 5, 4, NOW(), NOW()),
('组合式API很好用', 'approved', 6, 2, NOW(), NOW()),
('状态管理确实复杂', 'approved', 7, 3, NOW(), NOW()),
('微服务架构很实用', 'approved', 8, 4, NOW(), NOW());
"
echo ✅ 测试评论创建完成
echo.

echo 7. 更新统计信息...
docker exec sajuna-blog-mysql-dev mysql -u blog_user -pblog_password -e "
USE sajuna_blog;
UPDATE categories SET article_count = (SELECT COUNT(*) FROM articles WHERE category_id = categories.id);
UPDATE tags SET article_count = (SELECT COUNT(*) FROM article_tags WHERE tag_id = tags.id);
UPDATE articles SET comment_count = (SELECT COUNT(*) FROM comments WHERE article_id = articles.id AND status = 'approved');
"
echo ✅ 统计信息更新完成
echo.

echo ========================================
echo 测试数据生成完成！
echo ========================================
echo.
echo ✅ 已生成的数据:
echo - 5个测试用户
echo - 5个测试分类
echo - 10个测试标签
echo - 10篇测试文章
echo - 文章标签关联
echo - 10条测试评论
echo - 更新了统计信息
echo.
echo 💡 建议:
echo 1. 运行 optimize-database.bat 优化数据库性能
echo 2. 运行 analyze-slow-sql.bat 分析查询性能
echo 3. 启动应用测试功能
echo.

pause

