#!/usr/bin/env sh

# 确保脚本抛出遇到的错误
set -e

initDist(){
  echo $1 > base.js
}



#------------------------------------------

#url访问目录，这个是你 github 仓库的名字
initDist "module.exports = '/notes/'"

# 生成静态文件
npm run build
# 进入生成的文件夹
cd docs/.vuepress/dist

# deploy to github
if [ -z "$GITHUB_TOKEN" ]; then
  # 手动部署
  msg='deploy'
  githubUrl=git@github.com:oddfar/notes.git
else
  # 自动部署
  msg='来自github actions的自动部署'
  githubUrl=https://oddfar:${GITHUB_TOKEN}@github.com/oddfar/notes.git
  git config --global user.name "oddfar"
  git config --global user.email "oddfar@163.com"
fi
git init
git add -A
git commit -m "${msg}"
git push -f $githubUrl master:gh-pages # 推送到github


cd - # 退回开始所在目录
rm -rf docs/.vuepress/dist

#------------------------------------------


#打包代码同步到 gitee gh-pages分支
if [ -z "$SSH_PRIVATE_KEY" ]; then
  echo '如果是空字符串，则不部署到gitee'
else
  #url访问目录
  initDist "module.exports = '/'"
  # 生成静态文件
  npm run build
  # 进入生成的文件夹
  cd docs/.vuepress/dist

  giteeUrl=git@gitee.com:oddfar/notes.git  #gitee 仓库ssh地址
  
  git config --global user.name "oddfar"
  git config --global user.email "oddfar@163.com"
  git init
  git add -A
  git commit -m "来自github actions的自动部署"
  git push -f $giteeUrl master:gh-pages
  
  cd - # 退回开始所在目录
  rm -rf docs/.vuepress/dist
  # 删除秘钥
  rm -rf ~/.ssh
fi


