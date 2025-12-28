/**
 * @name 代理的配置
 * @see 在生产环境 代理是无法生效的，所以这里没有生产环境的配置
 * -------------------------------
 * The agent cannot take effect in the production environment
 * so there is no configuration of the production environment
 * For details, please see
 * https://pro.ant.design/docs/deploy
 *
 * @doc https://umijs.org/docs/guides/proxy
 */
export default {
  // 如果需要自定义本地开发服务器  请取消注释按需调整
  dev: {
    // localhost:6806/zxm-ai-admin/api/** -> http://localhost:6808/api/**
    '/zxm-ai-admin/api/': {
      target: 'http://localhost:6808',
      changeOrigin: true,
      pathRewrite: { '^/zxm-ai-admin': '' },
    },
    // 日志服务代理
    '/zxm-ai-admin/api-logs/': {
      target: 'http://localhost:6809',
      changeOrigin: true,
      pathRewrite: { '^/zxm-ai-admin/api-logs': '' },
    },
  },
};
