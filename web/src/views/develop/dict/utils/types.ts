interface FormItemProps {
  id?: number;
  /** 用于判断是`新增`还是`修改` */
  title: string;
  higherDeptOptions: Record<string, unknown>[];
  label: string;
  value: string;
  valueType: string;
  type: string;
  listClass: string;
  isDefault: number;
  sort: number;
  remark: string;
  status: number;
}
interface FormProps {
  formInline: FormItemProps;
}

/** 字典类型数据接口 */
interface DictTypeItemProps {
  id?: number;
  pid?: number;
  /** 用于判断是`新增`还是`修改` */
  title: string;
  higherDeptOptions: Record<string, unknown>[];
  /** 字典类型名称 */
  name: string;
  /** 字典类型编码 */
  type: string;
}

/** 添加字典类型请求参数 */
interface DictTypeDataProps {
  dictTypeInline: DictTypeItemProps;
}

export type { FormItemProps, FormProps, DictTypeDataProps, DictTypeItemProps };
