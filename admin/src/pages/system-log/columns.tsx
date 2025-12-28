/**
 * 表格列定义
 */
import { Button, Tag } from 'antd';
import { EyeOutlined } from '@ant-design/icons';
import type { ProColumns } from '@ant-design/pro-components';
import type { ISystemLog } from '@/types/systemLog';
import { LEVEL_OPTIONS, getLevelColor } from './constants';
import { formatDateTime, getLastWeekRange, timeRangePresets } from './utils';

export interface IGetColumnsParams {
  /** 查看详情回调 */
  onViewDetail: (record: ISystemLog) => void;
}

/** 获取表格列定义 */
export const getColumns = (params: IGetColumnsParams): ProColumns<ISystemLog>[] => {
  const { onViewDetail } = params;

  return [
    {
      title: 'ID',
      dataIndex: 'id',
      width: 80,
      search: false,
    },
    {
      title: '时间',
      dataIndex: 'time',
      width: 180,
      search: false,
      render: (_, record) => formatDateTime(record.time),
    },
    {
      title: '时间范围',
      dataIndex: 'timeRange',
      valueType: 'dateTimeRange',
      hideInTable: true,
      initialValue: getLastWeekRange(),
      colSize: 2,
      fieldProps: {
        placeholder: ['开始时间', '结束时间'],
        showTime: true,
        presets: timeRangePresets,
      },
    },
    {
      title: '级别',
      dataIndex: 'level',
      width: 100,
      valueType: 'select',
      colSize: 1,
      valueEnum: LEVEL_OPTIONS.reduce((acc, opt) => ({ ...acc, [opt.value]: { text: opt.label } }), {}),
      render: (_, record) => <Tag color={getLevelColor(record.level)}>{record.level}</Tag>,
      filters: false,
    },
    {
      title: '消息',
      dataIndex: 'msg',
      ellipsis: true,
      search: false,
    },
    {
      title: '操作',
      valueType: 'option',
      width: 80,
      fixed: 'right',
      render: (_: unknown, record: ISystemLog) => (
        <Button
          type="link"
          size="small"
          icon={<EyeOutlined />}
          onClick={() => onViewDetail(record)}
        >
          详情
        </Button>
      ),
    },
  ];
};
