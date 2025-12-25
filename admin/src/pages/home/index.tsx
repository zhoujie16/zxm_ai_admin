/**
 * 首页组件
 * 功能：系统首页
 */
import { Card } from 'antd';
import React from 'react';

const HomePage: React.FC = () => {
  return (
    <Card>
      <h1>欢迎使用 Admin Template</h1>
      <p>这是一个基础的管理系统模板，您可以在此基础上开发您的业务功能。</p>
    </Card>
  );
};

export default HomePage;
