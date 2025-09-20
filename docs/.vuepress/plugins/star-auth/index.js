const path = require('path');

/**
 * VuePress Star Auth 插件
 * 支持按需检测页面是否需要star验证
 * @param {Object} options 插件配置选项
 * @param {string} options.apiUrl API检测地址，默认为本地mock
 * @param {Array} options.repositories 需要star的仓库列表
 * @param {Object} ctx VuePress上下文
 */
module.exports = (options = {}, ctx) => {
    const defaultOptions = {
        apiUrl: 'http://localhost:8081/api/check-star',
        repositories: [
            {
                name: 'GitHub',
                url: 'https://github.com/your-username/your-repo',
                platform: 'github'
            },
            {
                name: 'Gitee', 
                url: 'https://gitee.com/your-username/your-repo',
                platform: 'gitee'
            }
        ]
    };
    
    const finalOptions = Object.assign({}, defaultOptions, options);
    
    return {
        name: 'vuepress-plugin-star-auth',
        clientRootMixin: path.resolve(__dirname, 'starAuth.js'),
        define: {
            STAR_AUTH_OPTIONS: finalOptions
        }
    }
}