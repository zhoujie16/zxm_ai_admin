// ==================== 登录相关类型 ====================

/**
 * 登录表单数据类型
 */
export interface ILoginFormData {
  /** 用户名 */
  username: string;
  /** 密码 */
  password: string;
  /** 是否记住登录状态 */
  remember?: boolean;
}

// ==================== API 响应类型 ====================

/**
 * 后端实际返回的响应格式
 */
export interface IBackendResponse<T = any> {
  /** 响应码，0 表示成功，其他表示失败 */
  code: number;
  /** 响应消息 */
  message: string;
  /** 响应数据 */
  data?: T;
}

// ==================== 代理服务相关类型 ====================

/**
 * 代理服务数据类型
 */
export interface IProxyService {
  /** 代理服务ID */
  id: number;
  /** 服务标识 */
  service_id: string;
  /** 服务器IP地址 */
  server_ip: string;
  /** 状态：1=启用，0=未启用 */
  status: number;
  /** 备注 */
  remark?: string;
  /** 创建时间 */
  created_at: string;
  /** 更新时间 */
  updated_at: string;
}

/**
 * 创建/更新代理服务表单数据
 */
export interface IProxyServiceFormData {
  /** 服务标识 */
  service_id?: string;
  /** 服务器IP地址 */
  server_ip?: string;
  /** 状态：1=启用，0=未启用 */
  status?: number;
  /** 备注 */
  remark?: string;
}

