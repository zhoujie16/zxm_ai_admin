/**
 * 登录页面组件
 * 功能：提供用户登录功能
 */
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Card, Checkbox, Form, Input, message } from 'antd';
import { history } from '@umijs/max';
import React, { useState } from 'react';
import { login } from '@/services/auth';
import type { ILoginFormData } from '@/types';
import './index.less';

const ILoginPage: React.FC = () => {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (values: ILoginFormData) => {
    setLoading(true);
    try {
      // 调用登录接口
      const result = await login(values.username, values.password);

      // 处理响应数据
      if (result.success && result.data) {
        // 登录成功，保存 token
        const { token } = result.data;
        localStorage.setItem('token', token);

        // 显示成功消息
        message.success('登录成功');

        // 跳转到首页
        const redirect = history.location.query?.redirect as string;
        history.push(redirect || '/');
      }
      // 登录失败的情况已在请求封装中处理并显示错误消息
    } catch (error: any) {
      // 网络错误等异常情况（已在封装中处理）
      console.error('登录异常:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className='login-page'>
      <Card className='login-card'>
        <div className='login-header'>
          <h1>Admin Template</h1>
          <p>管理系统登录</p>
        </div>

        <Form form={form} name='login' onFinish={handleSubmit} autoComplete='off' size='large'>
          <Form.Item name='username' rules={[{ required: true, message: '请输入用户名!' }]}>
            <Input prefix={<UserOutlined />} placeholder='用户名' />
          </Form.Item>

          <Form.Item name='password' rules={[{ required: true, message: '请输入密码!' }]}>
            <Input.Password prefix={<LockOutlined />} placeholder='密码' />
          </Form.Item>

          <Form.Item>
            <div className='login-form-footer'>
              <Form.Item name='remember' valuePropName='checked' noStyle>
                <Checkbox>记住我</Checkbox>
              </Form.Item>
              <a href='#' className='forgot-password-link'>
                忘记密码?
              </a>
            </div>
          </Form.Item>

          <Form.Item>
            <Button
              type='primary'
              htmlType='submit'
              className='login-submit-button'
              loading={loading}
            >
              登录
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
};

export default ILoginPage;
