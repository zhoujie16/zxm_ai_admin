/**
 * @see https://umijs.org/docs/max/access#access
 * */

// 定义当前用户类型
interface ICurrentUser {
  name?: string;
  avatar?: string;
  userid?: string;
  email?: string;
  access?: string;
}

export default function access(initialState: { currentUser?: ICurrentUser } | undefined) {
  const { currentUser } = initialState ?? {};
  return {
    canAdmin: currentUser && currentUser.access === 'admin',
  };
}
