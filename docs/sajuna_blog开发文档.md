# Sajuna Blog 开发文档

**2025年9月25日**

---

## 一、项目概述

本项目是一个个人博客项目，旨在构建一个个人 Web 站点，采用现代化的开发流程进行项目构建。同时，在 AI Coding 的帮助下拓展对不同技术栈的了解。

## 二、部署说明

本项目使用 Docker 进行容器化部署，采用前后端分离架构设计：

- **前端**：使用 Vue3 框架，采用响应式设计适配移动端，在 [localhost:3000](http://localhost:3000) 端口上进行开发
- **后端**：使用 Go + Tars 微服务框架 + GORM，通过 TarsGateway 提供 HTTP API 服务，在 [localhost:8080](http://localhost:8080) 上运行
- **数据库**：使用 [MySQL](http://localhost:3306) + [Redis](http://localhost:6379)
- **网关**：使用 [Nginx](http://localhost:80) 进行反向代理

### 前端技术栈

前端采用 Vue 进行开发，致力于为用户提供一个友好且高效的交互式网页界面。

| 技术栈 | 说明 |
|--------|------|
| Vue 3 | 采用 Composition API 提升组件复用性与逻辑组织能力 |
| Vuex 4.x | 全局状态管理 |
| Vue Router 4.x | 实现灵活的页面路由控制 |
| Element Plus UI 组件库 | 丰富的 UI 组件支持 |
| Sass/SCSS 预处理器 | 增强 CSS 开发能力 |
| Axios | HTTP 通信 |
| WebSocket + STOMP.js | 实现消息的即时推送与接收 |

### 后端架构

后端基于 Tars 微服务框架，使用 TarsGateway 作为 API 网关，提供统一的 HTTP 接口。TarsGateway 负责请求路由、负载均衡、限流熔断等功能。系统通过 Tars 框架实现服务注册发现、监控管理等功能。

### 底层服务

底层依托以下技术栈构建：

- **MySQL**：核心数据存储
- **Redis**：缓存相关数据，提升系统性能
- **MinIO**：提供分布式文件存储服务
- **Kafka**：作为消息队列支撑 Chatservice 的高并发消息分发，保证即时通讯稳定性与可扩展性

整体系统具备良好的模块解耦、水平扩展与高可用能力，可支持用户规模增长下的稳定运行，满足知识管理平台在实际应用中的业务需求。

## 三、需求描述

### 目标用户

- **主要用户（管理员/博主）**： 自己，负责内容创作、发布、管理和维护博客。

- **次要用户（访客）**： 博客的 reader 。可以浏览内容、发表评论。

---

### 用户角色与画像

| 角色 | 核心目标 | 使用场景 |
| :--- | :--- | :--- |
| **访客** | 1. 便捷地查找和阅读感兴趣的博文。<br>2. 了解博主的基本信息。<br>3. 对文章发表评论或提问。 | 通过搜索引擎或直接访问进入博客，浏览首页、文章列表，阅读具体文章，进行评论。 |
| **博主（管理员）** | 1. 高效地撰写、编辑、发布和删除文章。<br>2. 管理博客的评论内容，维护社区氛围。<br>3. 管理文章的分类和标签，使内容结构化。<br>4. 确保博客的安全稳定运行。 | 登录后台管理界面，创作新的 Markdown 文章，管理已发布的内容，审核或删除不当评论。 |

### 功能需求

功能需求描述系统必须提供的具体功能。

#### 访客端功能

*   **F1: 内容浏览模块(Note)**
    *   **F1.1 首页(Homepage)**： 展示blog简介、最新文章列表、热门文章推荐等。
    *   **F1.2 文章列表页(NoteList)**： 分页展示所有文章，可按分类、标签或发布日期（归档）进行筛选。
    *   **F1.3 文章详情页(NoteDetail)**： 完整展示文章标题、作者、发布时间、正文（支持 Markdown 渲染和代码高亮）、分类、标签。
    *   **F1.4 关于页面(About)**： 展示博主的详细介绍、社交链接等。
    *   **F1.5 搜索功能（Search）**： 在全站文章中进行关键词搜索，并展示结果列表。

*   **F2: 互动模块(Communicate)**
    *   **F2.1 评论功能(Comment)**： 博主或访客可以在文章底部发表评论（可支持回复一级评论的二级评论），博主能够置顶一级评论。
    *   **F2.2 评论通知(Notice)**： 当博主或访客回复了博主或访客的评论时，在站点内通知评论者。

#### 管理端功能

*   **F3: 身份认证模块(Auth)**
    *   **F3.1 登录/登出**： 博主通过用户名和密码登录后台管理系统。

*   **F4: 内容管理模块（ContentManage）**
    *   **F4.1 文章管理（CRUD）(NoteManage)**：
        *   **创建(CreateNote)**： 提供富文本或 Markdown 编辑器（**强烈推荐 Markdown**）撰写新文章，支持设置标题、摘要、分类、标签，添加文件（图片base64编码；文件名使用uuid），文件预览（格式自动解析）等。
        *   **读取(Read)**： 在后台列表中查看所有文章及其状态（已发布、草稿）。
        *   **更新(Update)**： 编辑已存在文章的内容和元信息。
        *   **删除(Delete)**： 将文章移至垃圾箱或永久删除（建议支持软删除）。
    *   **F4.2 分类/标签管理(TagManage)**： 可创建、编辑、删除文章的分类和标签。

*   **F5: 评论管理模块(CommentManage)**
    *   **F5.1 评论删除(Delete)**： 所有访客评论无审核自动通过（后续可能会添加屏蔽词的扩展），博主可删除。
    *   **F5.2 评论列表(List)**： 查看所有文章的评论，可进行**批准、回复、编辑、删除**操作。

#### 全局功能

*   **F6: 系统仪表盘(Dashboard)**
    *   **F6.1 数据统计(DataList)**： 展示基本数据，如文章总数、评论总数、近期访问量等。

### 非功能需求

*   **N1: 性能**
    *   **页面加载速度**： 首屏加载时间应小于 1.5 秒（基于 Lighthouse 评分标准）。
    *   **API 响应速度**： 95% 的 API 请求响应时间应在 100ms 以内。
    *   **并发支持**： 应能支持至少 50 个并发用户同时浏览（对于个人博客完全足够）。

*   **N2: 安全性**
    *   **后端**： 防止 SQL 注入、XSS 攻击、CSRF 攻击等常见 Web 安全威胁。
    *   **认证**： 用户密码需加盐哈希存储（bcrypt）。JWT Token 应有合理的过期时间（20min）。
    *   **权限**： 前后端均需对管理接口进行鉴权，未登录用户无法访问管理端功能。
    *   **输入验证**： 前后端均需对用户输入进行严格校验，优先使用白名单策略。

*   **N3: 可用性**
    *   **前端**： 界面设计简洁、直观，导航清晰。支持响应式布局，在手机、平板、桌面设备上均有良好体验。
    *   **后台**： 管理流程顺畅，Markdown 编辑器易用，文章发布流程简单。

*   **N4: 可靠性**
    *   **系统可用性**： 目标 99.9% 的可用性。
    *   **数据持久性**： 确保文章、评论等核心数据不丢失

*   **N5: 可维护性与可扩展性**
    *   **代码质量**： 代码结构清晰，注释完备，遵循编程规范。
    *   **模块化**： 前后端均采用模块化/组件化设计，便于后续功能扩展（如新增“友链”页面、图床功能等）。

#### 5. 内容管理需求

*   **写作体验**： 支持 Markdown 语法是核心需求，并提供实时渲染预览功能。
*   **媒体支持**： 支持在文章中上传和嵌入图片。

#### 总结与优先级建议

| 优先级 | 需求 | 说明 |
| :--- | :--- | :--- |
| **P0 (必须实现)** | 访客：文章浏览、列表、详情<br>博主：登录、文章CRUD、分类标签管理 | 构成博客最核心的功能。 |
| **P1 (应该实现)** | 访客：搜索、评论<br>博主：评论管理、Markdown编辑器 | 提升博客的可用性和互动性，是完整博客体验的关键。 |
| **P2 (可以实现)** | 系统仪表盘、评论回复通知、文章草稿自动保存、友情链接页面 | 在核心功能稳定后迭代增加。 |

## 四、模块和关系

### 博客系统整体架构图

```mermaid
graph TB
    subgraph Frontend [前端展示层]
        A1[访客界面]
        A2[管理后台]
    end
    
    subgraph Gateway [网关层]
        G1[TarsGateway]
        G2[Tars Registry]
    end
    
    subgraph Backend [后端服务层]
        B1[博客服务 BlogServer]
        B2[用户服务 UserService]
        B3[文章服务 ArticleService]
        B4[分类服务 CategoryService]
        B5[标签服务 TagService]
        B6[评论服务 CommentService]
    end
    
    subgraph Data [数据存储层]
        C1[MySQL - 文章/用户数据]
        C2[Redis - 缓存/会话]
        C3[Elasticsearch - 搜索索引]
    end
    
    A1 --> G1
    A2 --> G1
    G1 --> G2
    G2 --> B1
    B1 --> B2
    B1 --> B3
    B1 --> B4
    B1 --> B5
    B1 --> B6
    B2 --> C1
    B2 --> C2
    B3 --> C1
    B3 --> C2
    B3 --> C3
    B4 --> C1
    B5 --> C1
    B6 --> C1
```

### 流程图

#### 访客浏览文章流程

```mermaid
flowchart TD
    A([访客访问博客]) --> B{访问路径}
    
    B --> |直接访问首页| C[加载首页]
    B --> |搜索引擎进入| D[加载文章详情页]
    B --> |访问分类页面| E[加载分类文章列表]
    B --> |使用搜索功能| F[显示搜索页面]
    
    C --> G[展示博主简介]
    C --> H[显示最新文章列表]
    C --> I[显示热门文章推荐]
    C --> J[显示文章分类导航]
    
    E --> K[根据分类筛选文章]
    F --> L[输入关键词搜索]
    L --> M[显示搜索结果列表]
    
    G --> N{用户操作}
    H --> N
    I --> N
    J --> N
    K --> N
    M --> N
    
    N --> |点击阅读文章| O[加载文章详情页]
    N --> |浏览其他页面| P[跳转到对应页面]
    N --> |离开网站| Q([结束访问])
    
    O --> R[显示文章标题/元信息]
    O --> S[渲染Markdown内容]
    O --> T[显示代码高亮]
    O --> U[加载评论区域]
    
    U --> V{是否发表评论}
    V --> |是| W[输入评论内容]
    V --> |否| X([继续浏览])
    
    W --> Y[提交评论]
    Y --> Z[验证评论内容]
    Z --> AA[保存评论到数据库]
    AA --> AB[显示评论成功]
    AB --> X
```

#### 管理员文章CRUD操作流程

```mermaid
flowchart TD
    Start([管理员登录]) --> Auth[身份验证]
    Auth --> Dashboard[进入管理仪表盘]
    
    Dashboard --> Menu{选择操作菜单}
    
    Menu --> |文章管理| ArticleList
    Menu --> |写新文章| NewArticle
    Menu --> |分类管理| CategoryManage
    Menu --> |评论管理| CommentManage
    
    subgraph ArticleList [文章列表管理]
        AL1[加载文章列表] --> AL2{选择操作}
        AL2 --> |编辑| AL3[进入编辑模式]
        AL2 --> |删除| AL4[确认删除操作]
        AL2 --> |查看| AL5[预览文章效果]
        AL3 --> AL6[更新文章内容]
        AL4 --> AL7[软删除文章]
        AL6 --> AL8[返回文章列表]
        AL7 --> AL8
        AL5 --> AL8
    end
    
    subgraph NewArticle [创建新文章]
        NA1[打开Markdown编辑器] --> NA2[编写文章标题]
        NA2 --> NA3[编写文章内容]
        NA3 --> NA4[设置文章摘要]
        NA4 --> NA5[选择文章分类]
        NA5 --> NA6[添加文章标签]
        NA6 --> NA7[上传图片文件]
        NA7 --> NA8[实时预览效果]
        NA8 --> NA9{保存选项}
        NA9 --> |存为草稿| NA10[保存为草稿]
        NA9 --> |立即发布| NA11[发布文章]
        NA10 --> NA12[返回文章列表]
        NA11 --> NA12
    end
    
    subgraph CategoryManage [分类管理]
        CM1[查看分类列表] --> CM2{操作选择}
        CM2 --> |新增分类| CM3[输入分类信息]
        CM2 --> |编辑分类| CM4[修改分类信息]
        CM2 --> |删除分类| CM5[确认删除]
        CM3 --> CM6[保存分类]
        CM4 --> CM6
        CM5 --> CM7[删除分类]
        CM6 --> CM8[刷新分类列表]
        CM7 --> CM8
    end
    
    AL8 --> Dashboard
    NA12 --> Dashboard
    CM8 --> Dashboard
```

#### 评论系统流程

```mermaid
flowchart TD
    A[用户访问文章] --> B[加载文章内容]
    B --> C[显示评论区域]
    C --> D[加载现有评论列表]
    
    D --> E{显示排序}
    E -->|默认排序| F[按时间倒序显示]
    E -->|置顶优先| G[置顶评论显示在最前]
    
    F --> H{用户身份}
    G --> H
    
    H -->|访客| I[显示评论输入框]
    H -->|博主| J[显示评论输入框+管理工具]
    
    I --> K{用户操作}
    J --> K
    
    K -->|发表新评论| L
    K -->|回复评论| M
    K -->|管理评论| N
    
    subgraph L[发表一级评论流程]
        L1[填写评论信息] --> L2[提交评论]
        L2 --> L3[前端验证]
        L3 --> L4[调用评论API]
        L4 --> L5[保存一级评论]
        L5 --> L6[刷新评论列表]
    end
    
    subgraph M[回复评论流程（二级评论）]
        M1[点击回复按钮] --> M2[显示回复输入框]
        M2 --> M3[输入回复内容]
        M3 --> M4[提交回复]
        M4 --> M5[前端验证]
        M5 --> M6[调用回复API]
        M6 --> M7[保存二级评论]
        M7 --> M8[更新评论树显示]
    end
    
    subgraph N[博主管理操作]
        N1[显示管理选项] --> N2{选择操作}
        N2 -->|置顶评论| N3[设置一级评论置顶]
        N2 -->|取消置顶| N4[取消置顶状态]
        N2 -->|编辑评论| N5[修改评论内容]
        N2 -->|删除评论| N6[确认删除操作]
        
        N3 --> N7[更新置顶状态]
        N4 --> N7
        N5 --> N8[保存编辑内容]
        N6 --> N9[软删除评论]
        
        N7 --> N10[刷新评论显示]
        N8 --> N10
        N9 --> N10
    end
    
    L6 --> O[显示评论成功]
    M8 --> O
    N10 --> O
    
    O --> P[发送通知]
    P --> Q{通知类型}
    
    Q -->|回复博主评论| R[通知博主]
    Q -->|回复访客评论| S[通知被回复者]
    Q -->|新评论| T[通知博主有新评论]
    
    R --> U[更新通知中心]
    S --> U
    T --> U
    
    U --> V[完成评论操作]
    V --> W[继续浏览]
    
    W --> A
```

#### 搜索功能流程

```mermaid
flowchart TD
    A([用户进入博客]) --> B[显示搜索入口]
    B --> C{用户操作}
    
    C --> |点击搜索图标| D[展开搜索框]
    C --> |直接浏览| E([正常浏览])
    
    D --> F[输入关键词]
    F --> G{输入状态}
    
    G --> |实时搜索| H[显示实时建议]
    G --> |按回车搜索| I[执行搜索请求]
    
    H --> J[显示匹配结果]
    J --> K{选择结果}
    K --> |点击建议| L[跳转到文章]
    K --> |继续输入| F
    
    I --> M[显示加载状态]
    M --> N[向后端发送请求]
    N --> O[搜索索引查询]
    O --> P[返回搜索结果]
    P --> Q[渲染结果页面]
    
    Q --> R{搜索结果}
    R --> |有结果| S[显示文章列表]
    R --> |无结果| T[显示无结果提示]
    
    S --> U[分页显示]
    U --> V{用户操作}
    V --> |点击文章| W[跳转文章详情]
    V --> |修改搜索| F
    V --> |重新搜索| I
    
    T --> X[显示相关建议]
    X --> F
    
    W --> Y([阅读文章])
    L --> Y
```

#### 用户鉴权流程

```mermaid
flowchart TD
    A([用户访问]) --> B{访问页面类型}
    
    B --> |公开页面| C[直接展示内容]
    B --> |管理页面| D[检查认证状态]
    
    C --> E([正常浏览])
    
    D --> F{登录状态}
    F --> |已登录| G[验证Token有效性]
    F --> |未登录| H[跳转到登录页]
    
    G --> I{Token验证}
    I --> |有效| J[加载管理页面]
    I --> |无效/过期| H
    
    H --> K[显示登录表单]
    K --> L[输入用户名密码]
    L --> M[提交登录信息]
    M --> N[后端验证凭证]
    
    N --> O{验证结果}
    O --> |成功| P[生成JWT Token]
    O --> |失败| Q[显示错误信息]
    
    P --> R[返回Token到前端]
    R --> S[前端存储Token]
    S --> T[跳转到管理页面]
    
    Q --> L
    
    J --> U[加载管理功能]
    U --> V{管理员操作}
    
    V --> |文章管理| W[执行CRUD操作]
    V --> |评论管理| X[管理评论内容]
    V --> |系统设置| Y[配置博客参数]
    V --> |退出登录| Z[清除Token]
    
    W --> U
    X --> U
    Y --> U
    Z --> E
```

### 主要功能状态图

#### 文章生命周期状态

```mermaid
stateDiagram-v2
    [*] --> Draft : 开始创作
    Draft --> Saving : 自动保存
    Saving --> Draft : 保存完成
    
    Draft --> Reviewing : 提交审核
    Reviewing --> Draft : 需要修改
    Reviewing --> Published : 审核通过
    
    Draft --> Published : 直接发布
    Published --> Updated : 内容更新
    Updated --> Published : 更新完成
    
    Published --> Archived : 归档处理
    Published --> Deleted : 删除文章
    Archived --> Published : 重新发布
    Archived --> Deleted : 最终删除
    
    Deleted --> [*]
    
    note right of Draft
        草稿状态：
        - 可编辑内容
        - 可设置分类/标签
        - 支持自动保存
        - 可预览效果
    end note
    
    note right of Published
        发布状态：
        - 对访客可见
        - 可接收评论
        - 可被搜索
        - 可更新内容
    end note
```

#### 评论系统状态

```mermaid
stateDiagram-v2
    [*] --> Writing : 开始输入
    Writing --> Submitted : 提交评论
    Writing --> Canceled : 取消输入
    
    Submitted --> Published : 自动审核通过
    Submitted --> Rejected : 包含屏蔽词
    Submitted --> Pending : 需要人工审核
    
    Pending --> Published : 审核通过
    Pending --> Rejected : 审核不通过
    
    Published --> Pinned : 博主置顶
    Published --> Replied : 收到回复
    Published --> Edited : 内容编辑
    Published --> Deleted : 删除评论
    
    Pinned --> Published : 取消置顶
    Pinned --> Deleted : 删除评论
    Replied --> Published : 回复完成
    Edited --> Published : 编辑完成
    
    Rejected --> [*]
    Deleted --> [*]
    Canceled --> [*]
```

#### 用户认证状态

```mermaid
stateDiagram-v2
    [*] --> Unauthenticated : 访问网站
    
    Unauthenticated --> LoggingIn : 点击登录
    LoggingIn --> Authenticating : 提交凭证
    Authenticating --> Authenticated : 验证成功
    Authenticating --> LoginFailed : 验证失败
    
    LoginFailed --> LoggingIn : 重新输入
    LoginFailed --> Unauthenticated : 取消登录
    
    Authenticated --> Active : 正常使用
    Active --> AccessingResource : 访问资源
    AccessingResource --> Authorized : 权限验证通过
    AccessingResource --> Unauthorized : 权限不足
    
    Authorized --> Active : 操作完成
    Unauthorized --> Active : 返回主页
    
    Authenticated --> TokenRefreshing : Token过期
    TokenRefreshing --> Authenticated : 刷新成功
    TokenRefreshing --> Unauthenticated : 刷新失败
    
    Authenticated --> LoggingOut : 主动登出
    LoggingOut --> Unauthenticated : 登出完成
    
    note right of Authenticated
        认证状态：
        - JWT Token有效
        - 可访问管理功能
        - 有操作权限
    end note
```

#### 系统运行状态

```mermaid
stateDiagram-v2
    [*] --> Initializing : 系统启动
    Initializing --> Running : 启动完成
    
    state Running {
        [*] --> ServingRequests : 接收请求
        ServingRequests --> Processing : 处理业务
        Processing --> DatabaseOps : 数据操作
        DatabaseOps --> Responding : 返回结果
        Responding --> ServingRequests : 完成响应
        
        ServingRequests --> ErrorHandling : 发生错误
        ErrorHandling --> ServingRequests : 错误恢复
        ErrorHandling --> Degraded : 严重错误
    }
    
    Running --> Maintaining : 系统维护
    Maintaining --> Running : 维护完成
    
    Running --> Scaling : 负载调整
    Scaling --> Running : 调整完成
    
    Degraded --> Recovering : 故障恢复
    Recovering --> Running : 恢复完成
    Degraded --> Stopped : 系统崩溃
    
    Stopped --> [*]
```

#### 搜索功能状态

```mermaid
stateDiagram-v2
    [*] --> Idle : 搜索框空闲
    
    Idle --> Inputting : 开始输入
    Inputting --> Suggesting : 输入关键词
    Suggesting --> Waiting : 停止输入
    
    Waiting --> Searching : 执行搜索(回车/点击)
    Waiting --> Inputting : 继续输入
    Waiting --> Idle : 清空输入
    
    Searching --> Loading : 发送请求
    Loading --> ResultsReady : 获取结果
    ResultsReady --> Displaying : 显示结果
    
    Displaying --> BrowsingResults : 浏览结果
    BrowsingResults --> ViewingDetail : 查看详情
    ViewingDetail --> BrowsingResults : 返回列表
    
    BrowsingResults --> RefiningSearch : 修改搜索条件
    RefiningSearch --> Searching : 重新搜索
    BrowsingResults --> Idle : 完成搜索
    
    ResultsReady --> NoResults : 无匹配结果
    NoResults --> RefiningSearch : 调整关键词
    NoResults --> Idle : 放弃搜索
    
    Loading --> Error : 搜索失败
    Error --> Retrying : 重试搜索
    Error --> Idle : 取消搜索
    Retrying --> Loading
```

#### 文件上传状态

```mermaid
stateDiagram-v2
    [*] --> Ready : 准备上传
    
    Ready --> Selecting : 选择文件
    Selecting --> Validating : 文件选择完成
    Validating --> Ready : 文件无效
    Validating --> Uploading : 验证通过
    
    Uploading --> Progressing : 开始上传
    Progressing --> Uploading : 传输中
    Uploading --> Completed : 上传完成
    Uploading --> Failed : 上传失败
    
    Completed --> Processing : 后端处理
    Processing --> Stored : 存储完成
    Processing --> Error : 处理失败
    
    Stored --> [*] : 上传成功
    Failed --> Retrying : 重试上传
    Retrying --> Uploading
    Failed --> Ready : 取消上传
    Error --> Ready : 返回重选
    
    note right of Validating
        文件验证：
        - 格式检查
        - 大小限制
        - 类型白名单
    end note
    
    note right of Processing
        后端处理：
        - 生成UUID文件名
        - 格式转换(如需要)
        - 存储到指定位置
    end note
```

#### 数据缓存状态

```mermaid
stateDiagram-v2
    [*] --> CacheMiss : 数据请求
    
    CacheMiss --> FetchingDB : 查询数据库
    FetchingDB --> UpdatingCache : 获取数据
    UpdatingCache --> CacheHit : 缓存更新
    CacheHit --> ReturningData : 返回数据
    
    CacheMiss --> CacheHit : 缓存存在
    CacheHit --> Validating : 检查有效性
    Validating --> CacheHit : 仍然有效
    Validating --> Expired : 缓存过期
    Expired --> FetchingDB : 重新获取
    
    ReturningData --> [*]
    
    state CacheManagement {
        [*] --> Monitoring : 监控缓存
        Monitoring --> Evicting : 内存不足
        Monitoring --> Cleaning : 定期清理
        Evicting --> [*] : 清理完成
        Cleaning --> [*] : 清理完成
    }
    
    ReturningData --> CacheManagement
```

#### 用户权限控制流程

```mermaid
sequenceDiagram
    participant V as 访客
    participant A as 管理员
    participant FE as 前端路由
    participant BE as 后端API
    participant DB as 数据库
    
    Note over V,A: 访客操作流程
    V->>FE: 访问公开页面
    FE->>BE: 获取文章数据
    BE->>DB: 查询公开内容
    DB-->>BE: 返回数据
    BE-->>FE: 返回响应
    FE-->>V: 显示内容
    
    Note over V,A: 管理员操作流程
    A->>FE: 访问管理页面
    FE->>FE: 检查本地token
    FE->>BE: 验证token有效性
    BE-->>FE: 验证结果
    alt token有效
        FE-->>A: 进入管理界面
        A->>BE: 调用管理API
        BE->>BE: 权限验证
        BE->>DB: 执行操作
        BE-->>A: 返回成功
    else token无效/过期
        FE-->>A: 跳转到登录页
        A->>BE: 提交登录凭证
        BE->>DB: 验证用户信息
        BE-->>FE: 返回新token
        FE-->>A: 进入管理界面
    end
```

### 数据模型关系图

```mermaid
erDiagram
    USER ||--o{ NOTE : creates
    USER ||--o{ COMMENT : writes
    NOTE ||--o{ COMMENT : has
    NOTE }o--|| CATEGORY : belongs_to
    NOTE }o--o{ TAG : has
    
    USER {
        string id PK
        string username
        string password_hash
        string email
        datetime created_at
    }
    
    NOTE {
        string id PK
        string title
        text content
        string summary
        string status
        string category_id FK
        datetime created_at
        datetime updated_at
    }
    
    COMMENT {
        string id PK
        string content
        string author_name
        string author_email
        string note_id FK
        string parent_id FK
        boolean is_pinned
        datetime created_at
    }
    
    CATEGORY {
        string id PK
        string name
        string slug
    }
    
    TAG {
        string id PK
        string name
        string slug
    }
```
