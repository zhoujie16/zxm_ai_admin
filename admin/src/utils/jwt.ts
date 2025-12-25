/**
 * JWT 工具函数
 * 用于解析和处理 JWT token
 */

/**
 * JWT payload 类型定义
 */
export interface IJwtPayload {
  sub: string; // 用户 ID
  username: string; // 用户名
  access: string; // 权限
  iat?: number; // 签发时间
  exp?: number; // 过期时间
}

/**
 * 解析 JWT token
 * @param token JWT token 字符串
 * @returns 解析后的 payload，如果解析失败返回 null
 */
export function parseJwtToken(token: string): IJwtPayload | null {
  try {
    const base64Url = token.split('.')[1];
    if (!base64Url) {
      return null;
    }

    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    const jsonPayload = decodeURIComponent(
      atob(base64)
        .split('')
        .map(c => {
          return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
        })
        .join(''),
    );

    return JSON.parse(jsonPayload);
  } catch (error) {
    console.error('解析 JWT token 失败:', error);
    return null;
  }
}

/**
 * 检查 JWT token 是否过期
 * @param token JWT token 字符串
 * @returns 如果 token 有效且未过期返回 true，否则返回 false
 */
export function isTokenValid(token: string): boolean {
  const payload = parseJwtToken(token);
  if (!payload) {
    return false;
  }

  // 如果没有过期时间，认为 token 有效
  if (!payload.exp) {
    return true;
  }

  // 检查是否过期（提前 5 分钟判断为过期，避免边界情况）
  const currentTime = Math.floor(Date.now() / 1000);
  const expirationTime = payload.exp - 300; // 提前 5 分钟

  return currentTime < expirationTime;
}

/**
 * 从 JWT token 中提取用户信息
 * @param token JWT token 字符串
 * @returns 用户信息对象，如果解析失败返回 null
 */
export function getUserFromToken(token: string): {
  id: string;
  username: string;
  email: string;
  access: string;
} | null {
  const payload = parseJwtToken(token);
  if (!payload) {
    return null;
  }

  return {
    id: payload.sub,
    username: payload.username,
    email: `${payload.username}@example.com`, // 临时处理，实际应该从 token 或 API 获取
    access: payload.access,
  };
}

