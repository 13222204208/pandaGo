import { http } from "@/utils/http";
import { baseUrlApi } from "./utils";
import type { ApiResponse, ListResponse } from "./utils";

// 用户接口参数类型
export interface UserCreateParams {
  username: string;
  password: string;
  nickname: string;
  email: string;
  phone: string;
  status: number;
  roleId?: Array<number>;
}

// 用户信息类型
export interface UserInfo {
  id: number;
  username: string;
  nickname: string;
  email: string;
  phone: string;
  status: number;
  avatar: string;
  createdAt: string;
  updatedAt: string;
}

export interface ResetInfo {
  status?: number;
  avatar?: string;
  password?: string;
  roleId?: Array<number>;
}

// 创建用户
export const createUser = (data: UserCreateParams) => {
  return http.request<ApiResponse>("post", baseUrlApi("user"), { data });
};

// 获取用户列表
export const getUserList = (params?: object) => {
  return http.request<ListResponse<UserInfo>>("get", baseUrlApi("user"), {
    params
  });
};

// 更新用户
export const updateUser = (id: number, data: Partial<UserCreateParams>) => {
  return http.request<ApiResponse>("put", baseUrlApi(`user/${id}`), { data });
};

// 删除用户
export const deleteUser = (id: number) => {
  return http.request<ApiResponse>("delete", baseUrlApi(`user/${id}`));
};

// 重置用户状态,头像,密码,角色
export const resetUser = (id: number, data: Partial<ResetInfo>) => {
  return http.request<ApiResponse>("put", baseUrlApi(`user/${id}/reset`), {
    data
  });
};

export interface UserRoleIds {
  roleIds: number[];
}

// 获取用户角色ID列表
export const getUserRoleIds = (id: number) => {
  return http.request<ApiResponse<UserRoleIds>>(
    "get",
    baseUrlApi(`user/${id}/role`)
  );
};

// 菜单元数据类型
export interface MenuMeta {
  icon: string;
  title: string;
  rank: number;
}

// 菜单项类型
export interface MenuItem {
  path: string;
  component: string;
  name: string;
  meta: MenuMeta;
  children?: MenuItem[];
}

// 菜单数据响应
export interface MenuData {
  menu: MenuItem[];
}

export const getUserMenuList = () => {
  return http.request<ApiResponse<MenuData>>("get", baseUrlApi("user/menu"));
};
