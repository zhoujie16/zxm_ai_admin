/**
 * 页面容器组件
 * 功能：提供统一的页面容器，使用 div 包裹子元素
 *
 * 使用示例：
 * ```tsx
 * import PageContainer from '@/components/PageContainer';
 *
 * <PageContainer>
 *   <div>页面内容</div>
 * </PageContainer>
 * ```
 */
import React from 'react';
import classNames from 'classnames';
import './index.less';

/**
 * 页面容器组件属性
 */
interface IPageContainerProps {
  children: React.ReactNode;
  className?: string;
  style?: React.CSSProperties;
}

/**
 * 页面容器组件
 */
const PageContainer: React.FC<IPageContainerProps> = ({ children, className, style }) => {
  return (
    <div className={classNames('page-container', className)} style={style}>
      {children}
    </div>
  );
};

export default PageContainer;
