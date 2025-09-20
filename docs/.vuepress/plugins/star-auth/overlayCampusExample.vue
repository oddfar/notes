

<template>
    <div>
        <div class="overlay">
            <div class="wrapper">
                <p class="title">🚀 系统提示：访问文档失败 🚀</p>
                <div class="content">
                    <p>原因：开源不易，文档仅对已 Star 项目的用户开放。</p>
                    <ul>
                        <li v-for="repo in repositories" :key="repo.platform">
                            <a :href="repo.url" target="_blank">{{ repo.url }}</a>
                        </li>
                    </ul>
                    <p>
                        操作步骤：点击下方按钮 Star 任意一个项目，然后点击「验证Star状态」按钮进行检测。
                        感谢您的支持！
                    </p>
                    <div v-if="checking" class="checking">
                        <p>🔄 正在验证Star状态...</p>
                    </div>
                    <div v-if="errorMessage" class="error">
                        <p>❌ {{ errorMessage }}</p>
                    </div>
                </div>
                <ul class="button">
                    <li v-for="repo in repositories" :key="repo.platform">
                        <a :href="repo.url" target="_blank">✅ {{ repo.name }}项目</a>
                    </li>
                    <li>
                        <a href="#" @click.prevent="checkStarStatus" :class="{ disabled: checking }">🔍 验证Star状态</a>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>




<script>
/**
 * Star Auth 遮罩组件
 * 显示star验证界面并处理验证逻辑
 */
export default {
    props: {
        repositories: {
            type: Array,
            default: () => []
        },
        apiUrl: {
            type: String,
            default: 'http://localhost:8081/api/check-star'
        }
    },
    
    data() {
        return {
            checking: false,
            errorMessage: ''
        }
    },
    mounted() {
        // 添加CSS类阻止页面滚动
        document.body.classList.add('star-auth-overlay-active');
    },
    beforeDestroy() {
        // 移除CSS类恢复页面滚动
        document.body.classList.remove('star-auth-overlay-active');
    },
    methods: {
        /**
         * 检查Star状态
         */
        async checkStarStatus() {
            // 防止重复点击
            if (this.checking) return;
            
            this.checking = true;
            this.errorMessage = '';
            
            try {
                // 确保遮罩层保持显示
                this.$forceUpdate();
                
                // 调用API检查star状态
                const response = await this.callStarCheckAPI();
                
                if (response.success) {
                    // 验证成功，保存token并关闭遮罩
                    localStorage.setItem('star-auth-token', response.token);
                    this.closeOverlay();
                } else {
                    this.errorMessage = response.message || '验证失败，请确保已Star任意一个项目';
                    // 确保错误信息显示
                    this.$forceUpdate();
                }
            } catch (error) {
                console.error('Star验证失败:', error);
                this.errorMessage = '网络错误，请稍后重试';
                // 确保错误信息显示
                this.$forceUpdate();
            } finally {
                this.checking = false;
                // 确保状态更新
                this.$forceUpdate();
            }
        },
        
        /**
         * 调用Star检测API
         * @returns {Promise<Object>} API响应
         */
        async callStarCheckAPI() {
            try {
                const response = await fetch(this.apiUrl, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        repositories: this.repositories.map(repo => ({
                            platform: repo.platform,
                            url: repo.url
                        }))
                    })
                });
                
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                
                return await response.json();
            } catch (error) {
                console.warn('API调用失败，使用mock数据:', error);
                // 开发环境下使用mock数据
                return this.getMockResponse();
            }
        },
        
        /**
         * 获取Mock响应数据
         * @returns {Object} Mock响应
         */
        getMockResponse() {
            // 模拟70%的成功率
            const success = Math.random() > 0.3;
            
            if (success) {
                return {
                    success: true,
                    token: 'mock-token-' + Date.now(),
                    message: '验证成功！感谢您的支持！'
                };
            } else {
                return {
                    success: false,
                    message: '请先Star任意一个项目后再进行验证'
                };
            }
        },
        
        /**
         * 关闭遮罩层
         */
        closeOverlay() {
            // 移除CSS类恢复页面滚动
            document.body.classList.remove('star-auth-no-scroll');
            
            // 移除遮罩层
            const overlay = this.$el;
            if (overlay && overlay.parentNode) {
                overlay.parentNode.removeChild(overlay);
            }
            
            // 移除标记类
            const vdingContent = document.querySelector('div[class*="theme-vdoing-content"]');
            if (vdingContent) {
                vdingContent.classList.remove('star-auth-overlay');
            }
            
            // 刷新页面以显示内容
            setTimeout(() => {
                window.location.reload();
            }, 1000);
        }
    }
}
</script>

<style scoped>
* {
    padding: 0;
    margin: 0;
}

.overlay {
    background: rgba(0, 0, 0, .5);
    z-index: 999;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
}

.overlay .wrapper {
    background: #fff;
    width: 520px;
    min-height: 320px;
    border-radius: 10px;
    margin: 200px auto;
    padding: 10px 0;
    position: relative;
}

.overlay .wrapper .title {
    text-align: center;
    padding: 10px 0;
    border-bottom: 1px solid #eee;
    margin-bottom: 10px;
}

.overlay .wrapper .content {
    padding: 20px;
    padding-bottom: 60px; /* 为按钮留出空间 */
}

.overlay .wrapper .content ul {
    margin: 10px 0;
    padding-left: 20px;
}

.overlay .wrapper .content ul li {
    margin: 5px 0;
}

.overlay .wrapper .content ul li a {
    color: #1890ff;
    font-size: 14px;
    text-decoration: none;
}

.overlay .wrapper .content ul li a:hover {
    text-decoration: underline;
}

.overlay .wrapper .content p {
    margin: 10px 0;
    line-height: 1.5;
}

.overlay .wrapper .content p:last-of-type {
    color: #666;
    font-size: 14px;
}

/* 状态提示样式 */
.checking {
    background: #e6f7ff;
    border: 1px solid #91d5ff;
    border-radius: 4px;
    padding: 10px;
    margin: 10px 0;
    text-align: center;
}

.checking p {
    color: #1890ff;
    margin: 0;
}

.error {
    background: #fff2f0;
    border: 1px solid #ffccc7;
    border-radius: 4px;
    padding: 10px;
    margin: 10px 0;
    text-align: center;
}

.error p {
    color: #ff4d4f;
    margin: 0;
}

.overlay .wrapper .button {
    height: 45px;
    list-style: none;
    display: flex;
    flex-direction: row;
    align-items: center;
    width: 100%;
    position: absolute;
    bottom: 0;
}

.overlay .wrapper .button li {
    flex: 1;
    height: 100%;
    line-height: 45px;
    text-align: center;
    border: 1px solid #eee;
}

.overlay .wrapper .button li:first-child {
    border-bottom-left-radius: 10px;
}

.overlay .wrapper .button li:last-child {
    border-bottom-right-radius: 10px;
}

.overlay .wrapper .button li a {
    display: block;
    width: 100%;
    height: 100%;
    color: #333;
    text-decoration: none;
    transition: background-color 0.3s;
}

.overlay .wrapper .button li a:hover {
    background-color: #f5f5f5;
}

.overlay .wrapper .button li a.disabled {
    color: #ccc;
    cursor: not-allowed;
    pointer-events: none;
}

/* 阻止页面滚动 */
body.star-auth-overlay-active {
    overflow: hidden;
    height: 100vh;
}
</style>