/**
 * 404 页面组件
 * 功能：显示页面不存在的错误提示
 */
import { history } from '@umijs/max';
import { Button, Card, Result } from 'antd';
import React from 'react';

const NoFoundPage: React.FC = () => (
  <Card variant='borderless'>
    <Result
      status='404'
      title='404'
      subTitle='抱歉，您访问的页面不存在'
      extra={
        <Button type='primary' onClick={() => history.push('/')}>
          返回首页
        </Button>
      }
    />
  </Card>
);

export default NoFoundPage;
