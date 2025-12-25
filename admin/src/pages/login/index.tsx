/**
 * 登录页面组件
 * 功能：提供用户登录功能
 */
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Card, Checkbox, Form, Input } from 'antd';
import React from 'react';
import './index.less';

const ILoginPage: React.FC = () => {
  const [form] = Form.useForm();

  const handleSubmit = async (values: any) => {
    // TODO: 实现登录逻辑
    console.log('登录表单数据:', values);
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
