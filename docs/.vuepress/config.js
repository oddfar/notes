const plugins = require('./config/plugins.js');
const base = require('../../base.js');
const themeConfig = require('./config/themeConfig.js');



module.exports = {
  theme: 'vdoing', //  使用npm包主题
   // theme: require.resolve('../../theme-vdoing'), // 使用本地主题
  port: 8080,//端口

  plugins,
  themeConfig,
  base,  //引入后缀

  title: "OddFar's Notes",
  description: '我们在黑暗中并肩前行，走在各自的朝圣路上!', // 描述,以 <meta> 标签渲染到页面html中

  head: [
    ['link', { rel: 'icon', href: '/img/favicon.ico' }],
    ['meta', { name: 'keywords', content: 'oddfar,zhiyuan' }],
    ['meta', { name: 'theme-color', content: '#11a8cd' }], // 移动浏览器主题颜色
  ],
  markdown: {
    lineNumbers: true // 代码行号
  },
}