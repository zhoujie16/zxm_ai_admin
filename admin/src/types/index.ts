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
 * 通用 API 响应类型
 */
export interface IApiResponse<T = any> {
  /** 请求是否成功 */
  success: boolean;
  /** 响应数据 */
  data?: T;
  /** 响应消息 */
  message: string;
}

/**
 * 嵌套的 API 响应类型（实际API返回的结构）
 */
export interface INestedApiResponse<T = any> {
  /** 请求是否成功 */
  success: boolean;
  /** 嵌套的响应数据 */
  data?: {
    /** 实际数据 */
    data?: T;
  };
  /** 响应消息 */
  message: string;
}

/**
 * 分页响应类型
 */
export interface IPaginationResponse<T = any> {
  /** 请求是否成功 */
  success: boolean;
  /** 数据列表 */
  data?: T[];
  /** 分页信息 */
  pagination?: {
    /** 当前页码 */
    page: number;
    /** 每页数量 */
    limit: number;
    /** 总记录数 */
    total: number;
    /** 总页数 */
    totalPages: number;
  };
  /** 响应消息 */
  message: string;
}

/**
 * 嵌套的分页响应类型（实际API返回的结构）
 */
export interface INestedPaginationResponse<T = any> {
  /** 请求是否成功 */
  success: boolean;
  /** 嵌套的响应数据 */
  data?: {
    /** 数据列表 */
    data?: T[];
    /** 分页信息 */
    pagination?: {
      /** 当前页码 */
      page: number;
      /** 每页数量 */
      limit: number;
      /** 总记录数 */
      total: number;
      /** 总页数 */
      totalPages: number;
    };
  };
  /** 响应消息 */
  message: string;
}

/**
 * API 列表响应结构
 */
export interface IApiListResponse<T = any> {
  /** 数据列表 */
  data: T[];
  /** 分页信息 */
  pagination: {
    /** 当前页码 */
    page: number;
    /** 每页数量 */
    limit: number;
    /** 总记录数 */
    total: number;
    /** 总页数 */
    totalPages: number;
  };
}

