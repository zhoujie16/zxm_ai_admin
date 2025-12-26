module.exports = {
  apps: [
    {
      name: 'zxm-ai-admin',
      script: './bin/server',
      args: './configs/config.yaml',
      cwd: '/home/server/zxm_ai_admin',
      watch: false,
      autorestart: true,
      max_memory_restart: '500M',
      env: {
        NODE_ENV: 'production'
      },
      error_file: '/home/server/zxm_ai_admin/logs/pm2-error.log',
      out_file: '/home/server/zxm_ai_admin/logs/pm2-out.log',
      log_date_format: 'YYYY-MM-DD HH:mm:ss Z',
      merge_logs: true
    }
  ]
}

