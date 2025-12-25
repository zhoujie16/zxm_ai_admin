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
    // localhost:6806/api/** -> http://localhost:6808/api/**
    '/api/': {
      // 要代理的地址
      target: 'http://localhost:6808',
      // 配置了这个可以从 http 代理到 https
      // 依赖 origin 的功能可能需要这个，比如 cookie
      changeOrigin: true,
    },
    // localhost:6806/zxm-toolkit-admin/api/** -> http://localhost:6808/api/**
    '/zxm-toolkit-admin/api/': {
      // 要代理的地址
      target: 'http://localhost:6808',
      // 配置了这个可以从 http 代理到 https
      // 依赖 origin 的功能可能需要这个，比如 cookie
      changeOrigin: true,
      pathRewrite: { '^/zxm-toolkit-admin': '' },
    },
  },
};
