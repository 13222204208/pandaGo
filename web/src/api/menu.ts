import { http } from "@/utils/http";
import { baseUrlApi } from "./utils";
import type { ApiResponse, ListResponse } from "./utils";

// 菜单参数类型
export interface MenuCreateParams {
  menuType: number;
  parentId: number;
  title: string;
  name: string;
  path: string;
  component: string;
  rank: number;
  redirect?: string;
  icon?: string;
  extraIcon?: string;
  enterTransition?: string;
  leaveTransition?: string;
  activePath?: string;
  auths?: string;
  frameSrc?: string;
  frameLoading?: boolean;
  keepAlive?: boolean;
  hiddenTag?: boolean;
  fixedTag?: boolean;
  showLink?: boolean;
  showParent?: boolean;
}

// 菜单信息类型
export interface MenuInfo extends MenuCreateParams {
  id: number;
  children?: MenuInfo[];
  createdAt: string;
  updatedAt: string;
}

// 创建菜单
export const createMenu = (data: MenuCreateParams) => {
  return http.request<ApiResponse>("post", baseUrlApi("menu"), { data });
};

// 获取菜单列表
export const getMenuList = (params?: object) => {
  return http.request<ListResponse<MenuInfo>>("get", baseUrlApi("menu"), {
    params
  });
};

// 获取菜单列表
export const getMenuSimpleList = (params?: object) => {
  return http.request<ListResponse<MenuInfo>>(
    "get",
    baseUrlApi("menu/simple"),
    {
      params
    }
  );
};

// 更新菜单
export const updateMenu = (id: number, data: Partial<MenuCreateParams>) => {
  return http.request<ApiResponse>("put", baseUrlApi(`menu/${id}`), { data });
};

// 删除菜单
export const deleteMenu = (id: number) => {
  return http.request<ApiResponse>("delete", baseUrlApi(`menu/${id}`));
};
