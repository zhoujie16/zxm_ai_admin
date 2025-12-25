// 表格相关配置

/**
 * 默认分页配置
 */
export const DEFAULT_PAGINATION = {
  current: 1,
  pageSize: 10,
  showSizeChanger: true,
  showQuickJumper: true,
  pageSizeOptions: [10, 20, 50, 100],
  showTotal: (total: number, range: [number, number]) =>
    `第 ${range[0]}-${range[1]} 条/共 ${total} 条`,
} as const;

/**
 * 表格尺寸配置
 */
export const TABLE_SIZE = 'small' as const;
