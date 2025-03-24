import { http } from "@/utils/http";
import { baseUrlApi } from "./utils";
import type { ApiResponse, ListResponse } from "./utils";

// 角色参数类型
export interface RoleCreateParams {
  name: string;
  code: string;
  status: number;
  remark: string;
}

// 角色信息类型
export interface RoleInfo {
  id: number;
  name: string;
  code: string;
  status: number;
  remark: string;
  createdAt: string;
  updatedAt: string;
}
// 角色的菜单id
export interface RoleMenuIdsResponse {
  menuIds: number[];
}

// 创建角色
export const createRole = (data: RoleCreateParams) => {
  return http.request<ApiResponse>("post", baseUrlApi("role"), { data });
};

// 获取角色列表
export const getRoleList = (params?: object) => {
  return http.request<ListResponse<RoleInfo>>("get", baseUrlApi("role"), {
    params
  });
};

// 更新角色
export const updateRole = (id: number, data: Partial<RoleCreateParams>) => {
  return http.request<ApiResponse>("put", baseUrlApi(`role/${id}`), { data });
};

// 删除角色
export const deleteRole = (id: number) => {
  return http.request<ApiResponse>("delete", baseUrlApi(`role/${id}`));
};

// 获取角色的权限ID列表
// 角色菜单 ID 返回类型
export interface RoleMenuIdsResponse {
  menuIds: number[];
}

// 修改 getRoleMenuIds 函数的返回类型
export const getRoleMenuIds = (id: number) => {
  return http.request<ApiResponse<RoleMenuIdsResponse>>(
    "get",
    baseUrlApi(`role/${id}/menu`)
  );
};

// 更新角色权限
export const saveRoleMenuIds = (id: number, menuIds: number[]) => {
  return http.request<ApiResponse>("put", baseUrlApi(`role/${id}/menu`), {
    data: { menuIds }
  });
};
