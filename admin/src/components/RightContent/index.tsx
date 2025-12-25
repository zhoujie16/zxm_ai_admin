import React from 'react';

export type SiderTheme = 'light' | 'dark';

// 由于禁用了国际化，移除SelectLang组件
// 由于不使用Pro Components，简化Question组件
export const Question: React.FC = () => {
  return null; // 暂时不显示帮助按钮
};
