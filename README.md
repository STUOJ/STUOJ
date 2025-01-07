# STU Online Judge System

## 项目简介

STUOJ 是汕头大学疾风算法协会的 ACM-ICPC 在线代码评测系统，基于 Go 语言和 Gin 框架开发。

用户可以在平台上阅读算法题目，并可提交代码到代码沙箱进行评测，评测完成后系统将返回评测结果。管理员可以管理用户、导入题目、修改评测点数据、管理提交记录、查询系统统计数据和修改系统设置。

STUOJ 也是首款支持 AI 大模型的 ACM-ICPC 算法题目自动出题的 OJ 系统，可以自动生成算法题目、测试用例和题解代码。

## 项目仓库

- 后端仓库：[https://github.com/STUOJ/STUOJ](https://github.com/STUOJ/STUOJ)
- 前端仓库：[https://github.com/STUOJ/stuoj-web](https://github.com/STUOJ/stuoj-web)
- 数据库仓库：[https://github.com/STUOJ/stuoj-database](https://github.com/STUOJ/stuoj-database)

## API 文档

- Apifox：[https://stuoj-api.apifox.cn](https://stuoj-api.apifox.cn)

## 系统架构

- 后端：Gin + Gorm
- 前端：Vue + Element Plus
- 数据库：MySQL
- 代码沙箱: Judge0
- 图床服务: [yuki-image](https://github.com/ArtdragonXoX/yuki-image)
- 反向代理：Nginx
- 容器化部署：Docker
- 题目文件格式：FPS
- 人工智能工具包：[NekoACM](https://github.com/HEX9CF/NekoACM)

![image](https://github.com/user-attachments/assets/367668c5-585f-4fa2-820e-6891f638b0d8)

## 系统功能

![STUOJ](https://github.com/user-attachments/assets/68c7f6d9-7b07-4c26-a416-ff163f751f48)

### 用户相关

| 功能名称     | 请求方法 | 路由路径           | 操作者   | 功能简述                |
|----------|------|----------------|-------|---------------------|
| 获取用户信息   | GET  | /user/:id      | 用户    | 用户可以获取某个用户的详细信息。    |
| 用户登录     | POST | /user/login    | 未登录用户 | 用户输入邮箱和密码，登录账号。     |
| 用户注册     | POST | /user/register | 未登录用户 | 用户输入用户名、邮箱和密码，创建账号。 |
| 获取当前用户ID | GET  | /user/current  | 用户    | 用户可以获取当前登录用户的ID。    |
| 修改用户信息   | PUT  | /user/modify   | 用户    | 用户可以修改自己的用户信息。      |
| 修改用户密码   | PUT  | /user/password | 用户    | 用户可以修改自己的密码。        |
| 更新用户头像   | POST | /user/avatar   | 用户    | 用户可以更新自己的头像。        |

### 题目相关

| 功能名称       | 请求方法 | 路由路径                    | 操作者 | 功能简述            |
|------------|------|-------------------------|-----|-----------------|
| 获取题目列表     | GET  | /problem                | 用户  | 用户可以获取题目列表。     |
| 根据难度获取题目列表 | GET  | /problem/difficulty/:id | 用户  | 用户可以根据难度获取题目列表。 |
| 根据标签获取题目列表 | GET  | /problem/tag/:id        | 用户  | 用户可以根据标签获取题目列表。 |
| 根据标题获取题目列表 | POST | /problem/title          | 用户  | 用户可以根据标题获取题目列表。 |
| 获取题目详情     | GET  | /problem/:id            | 用户  | 用户可以获取题目的详细信息。  |
| 获取标签列表     | GET  | /problem/tag            | 用户  | 用户可以获取标签列表。     |

### 评测相关

| 功能名称     | 请求方法 | 路由路径            | 操作者 | 功能简述             |
|----------|------|-----------------|-----|------------------|
| 获取评测语言列表 | GET  | /judge/language | 用户  | 用户可以获取支持的编程语言列表。 |
| 提交代码     | POST | /judge/submit   | 用户  | 用户可以提交代码进行评测。    |
| 测试运行代码   | POST | /judge/testrun  | 用户  | 用户可以提交代码进行测试运行。  |

### 记录相关

| 功能名称     | 请求方法 | 路由路径                | 操作者 | 功能简述             |
|----------|------|---------------------|-----|------------------|
| 获取记录列表   | GET  | /record             | 用户  | 用户可以获取记录列表。      |
| 获取记录详情   | GET  | /record/:id         | 用户  | 用户可以获取记录的详细信息。   |
| 获取用户记录列表 | GET  | /record/user/:id    | 用户  | 用户可以获取特定用户的记录列表。 |
| 获取题目记录列表 | GET  | /record/problem/:id | 用户  | 用户可以获取特定题目的记录列表。 |

### 博客相关

| 功能名称       | 请求方法   | 路由路径              | 操作者 | 功能简述             |
|------------|--------|-------------------|-----|------------------|
| 获取博客列表     | GET    | /blog             | 用户  | 用户可以获取博客列表。      |
| 获取博客详情     | GET    | /blog/:id         | 用户  | 用户可以获取博客的详细信息。   |
| 获取用户博客列表   | GET    | /blog/user/:id    | 用户  | 用户可以获取特定用户的博客列表。 |
| 获取用户草稿博客列表 | GET    | /blog/draft       | 用户  | 用户可以获取用户的草稿博客列表。 |
| 获取题目博客列表   | GET    | /blog/problem/:id | 用户  | 用户可以获取特定题目的博客列表。 |
| 根据标题获取博客列表 | POST   | /blog/title       | 用户  | 用户可以根据标题获取博客列表。  |
| 保存博客       | POST   | /blog             | 用户  | 用户可以保存新博客。       |
| 编辑博客       | PUT    | /blog             | 用户  | 用户可以编辑现有博客。      |
| 提交博客       | PUT    | /blog/:id         | 用户  | 用户可以提交博客。        |
| 删除博客       | DELETE | /blog/:id         | 用户  | 用户可以删除博客。        |

### 评论相关

| 功能名称     | 请求方法   | 路由路径              | 操作者 | 功能简述             |
|----------|--------|-------------------|-----|------------------|
| 获取用户评论列表 | GET    | /comment/user/:id | 用户  | 用户可以获取特定用户的评论列表。 |
| 获取博客评论列表 | GET    | /comment/blog/:id | 用户  | 用户可以获取特定博客的评论列表。 |
| 添加评论     | POST   | /comment          | 用户  | 用户可以添加评论。        |
| 删除评论     | DELETE | /comment/:id      | 用户  | 用户可以删除评论。        |

### 管理面板
| 功能名称       | 请求方法   | 路由路径                                | 操作者   | 功能简述                    |
|------------|--------|-------------------------------------|-------|-------------------------|
| 获取用户列表     | GET    | /admin/user                         | 管理员   | 管理员可以获取用户列表。            |
| 获取用户详情     | GET    | /admin/user/:id                     | 管理员   | 管理员可以获取用户的详细信息。         |
| 根据角色获取用户列表 | GET    | /admin/user/role/:id                | 管理员   | 管理员可以根据角色获取用户列表。        |
| 添加用户       | POST   | /admin/user                         | 管理员   | 管理员可以添加用户。              |
| 修改用户信息     | PUT    | /admin/user                         | 管理员   | 管理员可以修改用户信息。            |
| 删除用户       | DELETE | /admin/user/:id                     | 管理员   | 管理员可以删除用户。              |
| 获取题目列表     | GET    | /admin/problem                      | 管理员   | 管理员可以获取题目列表。            |
| 根据状态获取题目列表 | GET    | /admin/problem/status/:id           | 管理员   | 管理员可以根据状态获取题目列表。        |
| 获取题目详情     | GET    | /admin/problem/:id                  | 管理员   | 管理员可以获取题目的详细信息。         |
| 添加题目       | POST   | /admin/problem                      | 管理员   | 管理员可以添加题目。              |
| 修改题目       | PUT    | /admin/problem                      | 管理员   | 管理员可以修改题目。              |
| 删除题目       | DELETE | /admin/problem/:id                  | 管理员   | 管理员可以删除题目。              |
| 添加题目标签     | POST   | /admin/problem/tag                  | 管理员   | 管理员可以给题目添加标签。           |
| 删除题目标签     | DELETE | /admin/problem/tag                  | 管理员   | 管理员可以删除题目的标签。           |
| 从FPS文件解析题目 | POST   | /admin/problem/fps                  | 管理员   | 管理员可以从FPS文件解析题目。        |
| AI生成题目     | POST   | /admin/problem/generate             | 管理员   | 大模型根据提供的题目信息或题解出题       |
| 获取题目历史记录   | GET    | /admin/history/problem/:id          | 管理员   | 管理员可以获取题目的历史记录。         |
| 获取评测点详情    | GET    | /admin/testcase/:id                 | 管理员   | 管理员可以获取评测点的详细信息。        |
| 添加评测点      | POST   | /admin/testcase                     | 管理员   | 管理员可以添加评测点。             |
| 修改评测点      | PUT    | /admin/testcase                     | 管理员   | 管理员可以修改评测点。             |
| 删除评测点      | DELETE | /admin/testcase/:id                 | 管理员   | 管理员可以删除评测点。             |
| 生成评测点数据    | POST   | /admin/testcase/datamake            | 管理员   | 管理员可以生成评测点数据。           |
| AI生成测试用例   | POST   | /admin/testcase/generate            | 管理员   | 大模型根据提供的题目信息或题解生成测试数据   |
| 获取标签列表     | GET    | /admin/tag                          | 管理员   | 管理员可以获取标签列表。            |
| 添加标签       | POST   | /admin/tag                          | 管理员   | 管理员可以添加标签。              |
| 修改标签       | PUT    | /admin/tag                          | 管理员   | 管理员可以修改标签。              |
| 删除标签       | DELETE | /admin/tag/:id                      | 管理员   | 管理员可以删除标签。              |
| 获取题解       | GET    | /admin/solution/:id                 | 管理员   | 管理员可以获取解答的详细信息。         |
| 添加题解       | POST   | /admin/solution                     | 管理员   | 管理员可以添加解答。              |
| 修改题解       | PUT    | /admin/solution                     | 管理员   | 管理员可以修改解答。              |
| 删除题解       | DELETE | /admin/solution/:id                 | 管理员   | 管理员可以删除解答。              |
| AI生成题解代码   | POST   | /solution                           | 管理员   | 大模型根据提供的题目信息生成指定编程语言的题解 |
| 获取提交记录列表   | GET    | /admin/record                       | 管理员   | 管理员可以获取提交记录列表。          |
| 获取提交记录详情   | GET    | /admin/record/:id                   | 管理员   | 管理员可以获取提交记录的详细信息。       |
| 删除提交记录     | DELETE | /admin/record/:id                   | 管理员   | 管理员可以删除提交记录。            |
| 获取博客列表     | GET    | /admin/blog                         | 管理员   | 管理员可以获取博客列表。            |
| 根据状态获取博客列表 | GET    | /admin/blog/status/:id              | 管理员   | 管理员可以根据状态获取博客列表。        |
| 获取博客详情     | GET    | /admin/blog/:id                     | 管理员   | 管理员可以获取博客的详细信息。         |
| 添加博客       | POST   | /admin/blog                         | 管理员   | 管理员可以添加博客。              |
| 修改博客       | PUT    | /admin/blog                         | 管理员   | 管理员可以修改博客。              |
| 删除博客       | DELETE | /admin/blog/:id                     | 管理员   | 管理员可以删除博客。              |
| 获取评论列表     | GET    | /admin/comment                      | 管理员   | 管理员可以获取评论列表。            |
| 添加评论       | POST   | /admin/comment                      | 管理员   | 管理员可以添加评论。              |
| 修改评论       | PUT    | /admin/comment                      | 管理员   | 管理员可以修改评论。              |
| 删除评论       | DELETE | /admin/comment/:id                  | 管理员   | 管理员可以删除评论。              |
| 获取用户统计数据   | GET    | /admin/statistics/user              | 管理员   | 管理员可以获取用户统计数据。          |
| 获取角色统计数据   | GET    | /admin/statistics/user/role         | 管理员   | 管理员可以获取角色统计数据。          |
| 获取注册统计数据   | GET    | /admin/statistics/user/register     | 管理员   | 管理员可以获取注册统计数据。          |
| 获取标签统计数据   | GET    | /admin/statistics/tag               | 管理员   | 管理员可以获取标签统计数据。          |
| 获取题目统计数据   | GET    | /admin/statistics/problem           | 管理员   | 管理员可以获取题目统计数据。          |
| 获取题目插入统计数据 | GET    | /admin/statistics/problem/insert    | 管理员   | 管理员可以获取题目插入统计数据。        |
| 获取题目更新统计数据 | GET    | /admin/statistics/problem/update    | 管理员   | 管理员可以获取题目更新统计数据。        |
| 获取题目删除统计数据 | GET    | /admin/statistics/problem/delete    | 管理员   | 管理员可以获取题目删除统计数据。        |
| 获取评测统计数据   | GET    | /admin/statistics/judge             | 管理员   | 管理员可以获取评测统计数据。          |
| 获取提交记录统计数据 | GET    | /admin/statistics/record            | 管理员   | 管理员可以获取提交记录统计数据。        |
| 获取提交统计数据   | GET    | /admin/statistics/record/submit     | 管理员   | 管理员可以获取提交统计数据。          |
| 获取编程语言统计数据 | GET    | /admin/statistics/record/language   | 管理员   | 管理员可以获取编程语言统计数据。        |
| 获取提交状态统计数据 | GET    | /admin/statistics/submission/status | 管理员   | 管理员可以获取提交状态统计数据。        |
| 获取评测状态统计数据 | GET    | /admin/statistics/judgement/status  | 管理员   | 管理员可以获取评测状态统计数据。        |
| 获取博客统计数据   | GET    | /admin/statistics/blog              | 管理员   | 管理员可以获取博客统计数据。          |
| 获取博客提交统计数据 | GET    | /admin/statistics/blog/submit       | 管理员   | 管理员可以获取博客提交统计数据。        |
| 获取评论提交统计数据 | GET    | /admin/statistics/comment/submit    | 管理员   | 管理员可以获取评论提交统计数据。        |
| 修改用户角色     | PUT    | /admin/user/role                    | 超级管理员 | 超级管理员可以修改用户角色。          |
| 获取系统配置     | GET    | /admin/config                       | 超级管理员 | 超级管理员可以获取系统配置。          |

## UML

### 用例图

![image](https://github.com/user-attachments/assets/d27bc6a6-bcdd-422b-baa5-8a85ba05b79b)

### 活动图 

#### 用户注册
![image](https://github.com/user-attachments/assets/10867d10-bae6-42d8-a613-bf6aed90e071)

#### 用户登录
![image](https://github.com/user-attachments/assets/cda37df8-469b-46f4-90b6-a74d1c097458)

#### 用户修改个人信息
![image](https://github.com/user-attachments/assets/cb85d84e-11ce-4d43-b6d2-c85a799276ad)

#### 用户修改密码
![image](https://github.com/user-attachments/assets/f98ad919-83bb-4543-bd34-01643962498f)

#### 题目信息
![image](https://github.com/user-attachments/assets/53bdd18b-8498-45a0-af7a-29253d5c0109)

#### 提交代码
![image](https://github.com/user-attachments/assets/f910a74f-1c15-4a83-aa79-f8b454671f28)

#### 提交记录
![image](https://github.com/user-attachments/assets/e734151a-a403-46da-af01-1a9620f3049c)

### 时序图

#### 用户注册
![image](https://github.com/user-attachments/assets/76828acc-fdcb-4924-8653-a4e45917d311)

#### 用户登录
![image](https://github.com/user-attachments/assets/3dded833-b5b9-498d-aa74-1662fe8c53af)

#### 用户修改个人信息
![image](https://github.com/user-attachments/assets/8c90e730-5b24-4304-b0dd-8486622905a1)

#### 用户找回密码
![image](https://github.com/user-attachments/assets/ce932c5d-5684-418c-96fa-47eed1a73041)

#### 题目信息
![image](https://github.com/user-attachments/assets/0be9c8a6-a828-4b41-bfc6-e93cb1d741ed)

#### 提交代码
![image](https://github.com/user-attachments/assets/6f2d4642-3199-4432-9e5f-b1f0eacd41a8)

#### 提交记录
![image](https://github.com/user-attachments/assets/e0dab15a-2d33-46de-832c-4959ec3ee410)
