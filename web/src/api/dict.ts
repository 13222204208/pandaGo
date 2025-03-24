import { http } from "@/utils/http";
import { baseUrlApi } from "./utils";
import type { ApiResponse, ListResponse } from "./utils";

// 字典类型接口
interface DictTypeItem {
  id: number;
  pid: number;
  name: string;
  type: string;
  sort: number;
  remark: string;
  status: number;
  createdAt: string;
  updatedAt: string;
}

// 字典数据接口
interface DictDataItem {
  id: number;
  label: string;
  value: string;
  valueType: string;
  type: string;
  listClass: string;
  isDefault: number;
  sort: number;
  remark: string;
  status: number;
  createdAt: string;
  updatedAt: string;
}

// 字典类型接口
export const getDictTypeList = (params?: object) => {
  return http.request<ApiResponse<{ list: DictTypeItem[] }>>(
    "get",
    baseUrlApi("/dicttype"),
    { params }
  );
};

export const addDictType = (data?: object) => {
  return http.request<ApiResponse>("post", baseUrlApi("/dicttype"), { data });
};

export const updateDictType = (id: number, data?: object) => {
  return http.request<ApiResponse>("put", baseUrlApi(`/dicttype/${id}`), {
    data
  });
};

export const deleteDictType = (id: number) => {
  return http.request<ApiResponse>("delete", baseUrlApi(`/dicttype/${id}`));
};

// 字典数据接口
export const getDictDataList = (params?: object) => {
  return http.request<ListResponse<DictDataItem>>(
    "get",
    baseUrlApi("/dictdata"),
    { params }
  );
};

export const addDictData = (data?: object) => {
  return http.request<ApiResponse>("post", baseUrlApi("/dictdata"), { data });
};

export const updateDictData = (id: number, data?: object) => {
  return http.request<ApiResponse>("put", baseUrlApi(`/dictdata/${id}`), {
    data
  });
};

export const deleteDictData = (id: number) => {
  return http.request<ApiResponse>("delete", baseUrlApi(`/dictdata/${id}`));
};
