import { Dropdown } from 'antd';
import type { DropDownProps } from 'antd/es/dropdown';
import classNames from 'classnames';
import React from 'react';

export type IHeaderDropdownProps = {
  overlayClassName?: string;
  placement?: 'bottomLeft' | 'bottomRight' | 'topLeft' | 'topCenter' | 'topRight' | 'bottomCenter';
} & Omit<DropDownProps, 'overlay'>;

const IHeaderDropdown: React.FC<IHeaderDropdownProps> = ({
  overlayClassName: cls,
  ...restProps
}) => {
  return <Dropdown overlayClassName={classNames('header-dropdown', cls)} {...restProps} />;
};

export default IHeaderDropdown;
