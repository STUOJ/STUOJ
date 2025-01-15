# STU Online Judge System

## 项目简介

STUOJ 是汕头大学疾风算法协会的 ACM-ICPC 算法程序在线评测系统，基于 Go 语言和 Gin 框架开发。

用户可以在平台上阅读算法题目，并可提交代码到代码沙箱进行评测，评测完成后系统将返回评测结果。管理员可以管理用户、导入题目、修改评测点数据、管理提交记录、查询系统统计数据和修改系统设置。

STUOJ 也是一款基于 AI 大模型的 ACM-ICPC 算法题目自动出题 OJ 系统，可以自动生成算法题目、测试用例和题解代码。

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
