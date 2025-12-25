/**
 * 模型管理页面
 * 功能：暂未开发
 */
import { Result, Button } from 'antd';
import { useNavigate } from 'umi';

const ModelPage: React.FC = () => {
  const navigate = useNavigate();

  return (
    <Result
      status="info"
      title="模型管理"
      subTitle="该功能暂未开发，敬请期待"
      extra={
        <Button type="primary" onClick={() => navigate('/config')}>
          返回配置管理
        </Button>
      }
    />
  );
};

export default ModelPage;
