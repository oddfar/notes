import overlayCampusExample from './overlayCampusExample.vue'
import Vue from 'vue'

/**
 * Star Auth 客户端混入
 * 根据页面frontmatter配置决定是否显示star检测遮罩
 */
export default {
    updated() {
        this.checkStarAuth();
    },
    
    mounted() {
        this.checkStarAuth();
    },
    
    methods: {
        /**
         * 检查当前页面是否需要star验证
         */
        checkStarAuth() {
            // 获取当前页面的frontmatter
            const currentPage = this.$page;
            if (!currentPage || !currentPage.frontmatter) return;
            
            // 检查是否需要star验证
            const needStarAuth = currentPage.frontmatter.starAuth;
            if (!needStarAuth) return;
            
            console.log("当前页面需要star验证：" + window.location.pathname);
            
            const vdingContent = document.querySelector('div[class*="theme-vdoing-content"]');
            if (!vdingContent) return;
            
            // 防止重复写入
            if (vdingContent.classList.contains('star-auth-overlay')) return;
            
            // 获取页面级仓库配置
            const pageRepositories = this.getPageRepositories(currentPage.frontmatter);
            
            // 检查用户是否已经通过验证
            this.checkUserStarStatus(pageRepositories).then(isStarred => {
                if (!isStarred) {
                    this.showStarAuthOverlay(vdingContent, pageRepositories);
                }
            });
        },
        
        /**
         * 获取页面级仓库配置
         * @param {Object} frontmatter 页面frontmatter
         * @returns {Array} 仓库配置数组
         */
        getPageRepositories(frontmatter) {
            // 如果页面有自定义仓库配置，使用页面配置
            if (frontmatter.starAuthRepositories && Array.isArray(frontmatter.starAuthRepositories)) {
                return frontmatter.starAuthRepositories;
            }
            
            // 否则使用全局默认配置
            return STAR_AUTH_OPTIONS.repositories || [];
        },
        
        /**
         * 检查用户star状态
         * @param {Array} repositories 仓库配置数组
         * @returns {Promise<boolean>} 是否已star
         */
        async checkUserStarStatus(repositories = []) {
            try {
                // 从localStorage检查是否已验证过
                const starAuthToken = localStorage.getItem('star-auth-token');
                if (starAuthToken) {
                    // 验证token是否仍然有效
                    const response = await fetch(`${STAR_AUTH_OPTIONS.apiUrl}/verify`, {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json',
                        },
                        body: JSON.stringify({ token: starAuthToken })
                    });
                    
                    if (response.ok) {
                        const result = await response.json();
                        return result.valid;
                    }
                }
                
                return false;
            } catch (error) {
                console.warn('检查star状态失败，使用mock数据:', error);
                // 开发环境下使用mock数据
                return Math.random() > 0.7; // 30%概率通过验证，用于测试
            }
        },
        
        /**
         * 显示star验证遮罩
         * @param {Element} vdingContent 内容容器
         * @param {Array} repositories 仓库配置数组
         */
        showStarAuthOverlay(vdingContent, repositories) {
            let ComponentClass = Vue.extend(overlayCampusExample);
            let instance = new ComponentClass({
                propsData: {
                    repositories: repositories || STAR_AUTH_OPTIONS.repositories,
                    apiUrl: STAR_AUTH_OPTIONS.apiUrl
                }
            });
            instance.$mount();
            vdingContent.classList.add('star-auth-overlay');
            // 阻止页面滚动
            document.body.classList.add('star-auth-no-scroll');
            vdingContent.appendChild(instance.$el);
        }
    }
}