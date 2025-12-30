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

// ==================== AI 模型相关类型 ====================

/**
 * AI 模型数据类型
 */
export interface IAIModel {
  /** 模型ID */
  id: number;
  /** 模型名称 */
  model_name: string;
  /** API地址 */
  api_url: string;
  /** API Key */
  api_key: string;
  /** 状态：1=启用，0=禁用 */
  status: number;
  /** 备注 */
  remark?: string;
  /** 创建时间 */
  created_at: string;
  /** 更新时间 */
  updated_at: string;
}

/**
 * 创建/更新 AI 模型表单数据
 */
export interface IAIModelFormData {
  /** 模型名称 */
  model_name?: string;
  /** 模型来源的 API Key（用于后端查询模型来源） */
  api_key?: string;
  /** 状态：1=启用，0=禁用 */
  status?: number;
  /** 备注 */
  remark?: string;
}

// ==================== Token 相关类型 ====================

/**
 * Token 数据类型
 */
export interface IToken {
  /** Token ID */
  id: number;
  /** Token 值 */
  token: string;
  /** 关联的AI模型ID */
  ai_model_id: number;
  /** 关联的AI模型名称 */
  model_name?: string;
  /** 关联订单号 */
  order_no?: string;
  /** 状态：1=启用，0=禁用 */
  status: number;
  /** 过期时间 */
  expire_at?: string;
  /** 使用限额 */
  usage_limit: number;
  /** 备注 */
  remark?: string;
  /** 创建时间 */
  created_at: string;
  /** 更新时间 */
  updated_at: string;
}

/**
 * 创建/更新 Token 表单数据
 */
export interface ITokenFormData {
  /** 关联的AI模型ID */
  ai_model_id?: number;
  /** 关联订单号 */
  order_no?: string;
  /** 状态：1=启用，0=禁用 */
  status?: number;
  /** 过期时间 */
  expire_at?: string;
  /** 使用限额 */
  usage_limit?: number;
  /** 备注 */
  remark?: string;
}

// ==================== 模型来源相关类型 ====================

/**
 * 模型来源数据类型
 */
export interface IModelSource {
  /** 模型来源ID */
  id: number;
  /** 模型名称 */
  model_name: string;
  /** API地址 */
  api_url: string;
  /** API Key */
  api_key: string;
  /** 备注 */
  remark?: string;
  /** 创建时间 */
  created_at: string;
  /** 更新时间 */
  updated_at: string;
}

/**
 * 创建模型来源表单数据
 */
export interface ICreateModelSourceFormData {
  /** 模型名称 */
  model_name: string;
  /** API地址 */
  api_url: string;
  /** API Key */
  api_key: string;
  /** 备注 */
  remark?: string;
}

/**
 * 更新模型来源表单数据
 */
export interface IUpdateModelSourceFormData {
  /** 模型名称 */
  model_name?: string;
  /** 备注 */
  remark?: string;
}

