<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [准备](#%E5%87%86%E5%A4%87)
- [写作](#%E5%86%99%E4%BD%9C)
- [部署](#%E9%83%A8%E7%BD%B2)
  - [手动部署到 GitHub](#%E6%89%8B%E5%8A%A8%E9%83%A8%E7%BD%B2%E5%88%B0-github)
  - [自动部署到 GitHub](#%E8%87%AA%E5%8A%A8%E9%83%A8%E7%BD%B2%E5%88%B0-github)
  - [部署到自己服务器](#%E9%83%A8%E7%BD%B2%E5%88%B0%E8%87%AA%E5%B7%B1%E6%9C%8D%E5%8A%A1%E5%99%A8)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

使用 [vuepress](https://vuepress.vuejs.org/zh) 搭建，自动部署在[GitHub Pages](https://pages.github.com/) 

## 准备

VuePress 需要 [Node.js ](https://nodejs.org/en/)>= 8.6

1. 克隆到本地并进入目录

   ```sh
   git clone https://github.com/oddfar/notes.git && cd notes
   ```

2. 安装本地依赖

   ```sh
   npm install
   ```

3. 本地测试

   ```sh
   npm run dev
   ```

   默认访问链接：http://localhost:8080/

## 写作

使用 `markdown` 语法编写`md`文件，所有笔记 `md` 文件放在 `docs/` 目录下

例如添加 `test` 类，并编写 `hello.md` 文件

1. 创建目录

   格式：序号+标题

   例如：30.test

2. 添加笔记

   例如：01.hello.md

3. 编写内容

   ```markdown
   ---
   title: 笔记标题
   permalink: /test/hello/
   date： 2021-01-01 01:01:01
   ---
   
   ## 标题
   
   hello world
   ```

   tittle：标题，不填写则默认文件名中的标题，即 `hello`

   permalink：访问链接，不填写则自动生成

   date：日期，默认文件创建时间

4. 测试运行

   在项目根目录下

   ```sh
   npm run dev
   ```

文件及写作规范详情请看 

- [vdoing主题介绍文档](https://doc.xugaoyi.com/)
- [文档 - md文件写作](https://oddfar.github.io/notes/about/blog/3/)

## 部署

### 手动部署到 GitHub

创建分支：`gh-pages`

更改文件`deploy.sh`内容

仓库地址替换成自己的：

- githubUrl=git@github.com:oddfar/notes.git

- githubUrl=https://oddfar:${GITHUB_TOKEN}@github.com/oddfar/notes.git

- initDist "module.exports = '/notes/'"



修好完后，双击运行`deploy.sh`

之后配置 [GitHub Pages](https://pages.github.com/)

![image-20210517151354287](https://cdn.jsdelivr.net/gh/oddfar/static/img/20210517151356.png)

### 自动部署到 GitHub



详情查看：[点击此处](https://oddfar.github.io/notes/about/blog/5/)

### 部署到自己服务器

根目录下执行命令

```sh
npm run build
```

生成文件在 `docs\.vuepress\dist\` 目录下

打包到服务器即可

注：本地不可直接访问，需要配合插件