# 个人介绍网站

这是一个使用 Go 语言开发的个人介绍网站，具有现代化的设计和响应式布局。

## 功能特点

- 响应式设计，适配各种设备
- 现代化UI界面
- 平滑滚动效果
- 技能展示
- 项目经验展示
- 教育背景
- 联系方式表单

## 技术栈

- 后端：Go + Gin
- 前端：HTML5 + CSS3 + JavaScript
- UI框架：Bootstrap 5
- 动画：AOS (Animate On Scroll)

## 安装和运行

1. 确保已安装 Go 1.21 或更高版本

2. 克隆项目
```bash
git clone [项目地址]
cd personal-website
```

3. 安装依赖
```bash
go mod download
```

4. 运行项目
```bash
go run main.go
```

5. 访问网站
打开浏览器访问 http://localhost:8080

## 项目结构

```
personal-website/
├── main.go           # 主程序入口
├── static/          # 静态资源
│   ├── css/        # CSS文件
│   └── images/     # 图片资源
├── templates/       # HTML模板
├── controllers/     # 控制器
├── models/         # 数据模型
└── config/         # 配置文件
```

## 自定义

1. 修改 `templates/index.html` 中的个人信息
2. 在 `static/images/` 中添加您的个人照片和背景图片
3. 在 `static/css/style.css` 中自定义样式

## 贡献

欢迎提交 Issue 和 Pull Request

## 许可证

MIT License 