import "./reset.css";
import dayjs from "dayjs";
import editForm from "../form/index.vue";
import { handleTree } from "@/utils/tree";
import { message } from "@/utils/message";
import { usePublicHooks } from "../../hooks";
import { addDialog } from "@/components/ReDialog";
import type { PaginationProps } from "@pureadmin/table";
import type { DictTypeItemProps, FormItemProps } from "../utils/types";
import { getKeyList, deviceDetection } from "@pureadmin/utils";
import {
  getDictTypeList,
  getDictDataList,
  deleteDictData,
  addDictData,
  updateDictData,
  updateDictType,
  addDictType
} from "@/api/dict";
import dictTypeForm from "../form/dicttype.vue";
import { ElMessageBox } from "element-plus";
import { type Ref, h, ref, toRaw, computed, reactive, onMounted } from "vue";

export function dictdata(tableRef: Ref, treeRef: Ref) {
  const form = reactive({
    // 左侧部门树的id
    label: "",
    value: "",
    type: "",
    status: undefined,
    currentPage: 1,
    pageSize: 10
  });
  const formRef = ref();
  const dataList = ref([]);
  const loading = ref(true);
  const switchLoadMap = ref({});
  const { switchStyle } = usePublicHooks();
  const higherDeptOptions = ref();
  const treeData = ref([]);
  const treeLoading = ref(true);
  const selectedNum = ref(0);
  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });
  const columns: TableColumnList = [
    {
      label: "字典ID",
      prop: "id",
      width: 90
    },
    {
      label: "字典类型",
      prop: "type",
      minWidth: 130
    },
    {
      label: "字典标签",
      prop: "label",
      minWidth: 130,
      cellRenderer: scope => (
        <el-tag
          type={
            !scope.row.listClass || scope.row.listClass === "default"
              ? "info"
              : scope.row.listClass
          }
        >
          {scope.row.label}
        </el-tag>
      )
    },
    {
      label: "字典键值",
      prop: "value",
      minWidth: 90
    },
    {
      label: "键值类型",
      prop: "valueType",
      minWidth: 90
    },
    {
      label: "状态",
      prop: "status",
      minWidth: 90,
      cellRenderer: scope => (
        <el-switch
          size={scope.props.size === "small" ? "small" : "default"}
          loading={switchLoadMap.value[scope.index]?.loading}
          v-model={scope.row.status}
          active-value={1}
          inactive-value={0}
          active-text="已启用"
          inactive-text="已停用"
          inline-prompt
          style={switchStyle.value}
          onChange={() => onChange(scope as any)}
        />
      )
    },
    {
      label: "创建时间",
      minWidth: 90,
      prop: "createTime",
      formatter: ({ createTime }) =>
        dayjs(createTime).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "操作",
      fixed: "right",
      width: 180,
      slot: "operation"
    }
  ];
  const buttonClass = computed(() => {
    return [
      "!h-[20px]",
      "reset-margin",
      "!text-gray-500",
      "dark:!text-white",
      "dark:hover:!text-primary"
    ];
  });

  function onChange({ row, index }) {
    ElMessageBox.confirm(
      `确认要<strong>${
        row.status === 0 ? "停用" : "启用"
      }</strong><strong style='color:var(--el-color-primary)'>${
        row.label
      }</strong>字典数据吗?`,
      "系统提示",
      {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
        dangerouslyUseHTMLString: true,
        draggable: true
      }
    )
      .then(async () => {
        switchLoadMap.value[index] = Object.assign(
          {},
          switchLoadMap.value[index],
          {
            loading: true
          }
        );
        console.log(row);
        // 调用修改接口
        const res = await updateDictData(row.id, {
          ...row,
          status: row.status
        });
        switchLoadMap.value[index].loading = false;
        if (res.code === 0) {
          message("已成功修改字典状态", { type: "success" });
          onSearch(); // 刷新表格数据
        } else {
          message(res.message, { type: "error" });
          row.status = row.status === 0 ? 1 : 0; // 切换失败时恢复状态
        }
      })
      .catch(() => {
        row.status = row.status === 0 ? 1 : 0;
      });
  }

  function handleUpdate(row) {
    console.log(row);
  }

  async function handleDelete(row) {
    const res = await deleteDictData(row.id);
    if (res.code === 0) {
      message("删除成功", { type: "success" });
      onSearch();
    } else {
      message(res.message, { type: "error" });
    }
  }

  function handleSizeChange(val: number) {
    form.pageSize = val;
    form.currentPage = 1; // 切换每页条数时重置为第一页
    onSearch();
  }

  function handleCurrentChange(val: number) {
    form.currentPage = val;
    onSearch();
  }

  /** 当CheckBox选择项发生变化时会触发该事件 */
  function handleSelectionChange(val) {
    selectedNum.value = val.length;
    // 重置表格高度
    tableRef.value.setAdaptive();
  }

  /** 取消选择 */
  function onSelectionCancel() {
    selectedNum.value = 0;
    // 用于多选表格，清空用户的选择
    tableRef.value.getTableRef().clearSelection();
  }

  /** 批量删除 */
  function onbatchDel() {
    // 返回当前选中的行
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    // 接下来根据实际业务，通过选中行的某项数据，比如下面的id，调用接口进行批量删除
    message(`已删除编号为 ${getKeyList(curSelected, "id")} 的数据`, {
      type: "success"
    });
    tableRef.value.getTableRef().clearSelection();
    onSearch();
  }

  async function onSearch() {
    loading.value = true;
    const { data } = await getDictDataList(toRaw(form));
    dataList.value = data.list;
    pagination.total = data.total;
    pagination.pageSize = data.pageSize;
    pagination.currentPage = data.currentPage;

    setTimeout(() => {
      loading.value = false;
    }, 500);
  }

  // 添加更新树形结构数据的方法
  async function searchTreeData() {
    treeLoading.value = true;
    const { data } = await getDictTypeList();
    const treeList = handleTree(data.list, "id", "pid", null);
    treeData.value = treeList;
    console.log("yang", treeList);
    higherDeptOptions.value = treeList;
    treeLoading.value = false;
  }

  const resetForm = formEl => {
    if (!formEl) return;
    formEl.resetFields();
    form.value = "";
    treeRef.value.onTreeReset();
    onSearch();
  };

  function onTreeSelect({ type, selected }) {
    form.type = selected ? type : "";
    onSearch();
  }

  function formatHigherDeptOptions(treeList) {
    // 根据返回数据的status字段值判断追加是否禁用disabled字段，返回处理后的树结构，用于上级部门级联选择器的展示（实际开发中是如此，不可能前端需要的每个字段后端都会返回，这时需要前端自行根据后端返回的某些字段做逻辑处理）
    if (!treeList || !treeList.length) return;
    const newTreeList = [];
    for (let i = 0; i < treeList.length; i++) {
      treeList[i].disabled = treeList[i].status === 0 ? true : false;
      formatHigherDeptOptions(treeList[i].children);
      newTreeList.push(treeList[i]);
    }
    return newTreeList;
  }

  function openTypeDialog(title = "新增", row?: DictTypeItemProps) {
    addDialog({
      title: `${title}字典类型`,
      props: {
        dictTypeInline: {
          title,
          higherDeptOptions: formatHigherDeptOptions(higherDeptOptions.value),
          pid: row?.pid ?? 0,
          name: row?.name ?? "",
          type: row?.type ?? ""
        }
      },
      width: "36%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () =>
        h(dictTypeForm, { ref: formRef, dictTypeInline: null }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.dictTypeInline as DictTypeItemProps;
        console.log("curData", curData);
        function chores() {
          message(`您${title}了类型名称为${curData.name}的这条数据`, {
            type: "success"
          });
          done(); // 关闭弹框
          searchTreeData();
        }
        FormRef.validate(async valid => {
          if (valid) {
            if (title === "新增") {
              // 调用新增接口
              const res = await addDictType(curData);
              if (res.code === 0) {
                chores();
              } else {
                message(res.message, { type: "error" });
              }
            } else {
              // 调用修改接口
              const res = await updateDictType(row.id, curData);
              if (res.code === 0) {
                chores();
              } else {
                message(res.message, { type: "error" });
              }
            }
          }
        });
      }
    });
  }

  function openDialog(title = "新增", row?: FormItemProps) {
    addDialog({
      title: `${title}字典`,
      props: {
        formInline: {
          title,
          higherDeptOptions: formatHigherDeptOptions(higherDeptOptions.value),
          label: row?.label ?? "",
          value: row?.value ?? "",
          valueType: row?.valueType ?? "string",
          type: row?.type ?? "",
          listClass: row?.listClass ?? "primary",
          isDefault: row?.isDefault ?? 0,
          sort: row?.sort ?? 0,
          status: row?.status ?? 1,
          remark: row?.remark ?? ""
        }
      },
      width: "46%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(editForm, { ref: formRef, formInline: null }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as FormItemProps;
        function chores() {
          message(`您${title}了标签名称为${curData.label}的这条数据`, {
            type: "success"
          });
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(async valid => {
          if (valid) {
            if (title === "新增") {
              console.log(curData);
              // 调用新增接口
              const res = await addDictData(curData);
              if (res.code === 0) {
                chores();
              } else {
                message(res.message, { type: "error" });
              }
            } else {
              // 调用修改接口
              const res = await updateDictData(row.id, curData);
              if (res.code === 0) {
                chores();
              } else {
                message(res.message, { type: "error" });
              }
            }
          }
        });
      }
    });
  }

  onMounted(async () => {
    treeLoading.value = true;
    onSearch();

    // 获取字典类型列表并转换为树形结构
    const { data } = await getDictTypeList();
    const treeList = handleTree(data.list, "id", "pid", null);
    higherDeptOptions.value = treeList;
    treeData.value = treeList;
    treeLoading.value = false;
  });

  return {
    form,
    loading,
    columns,
    dataList,
    treeData,
    searchTreeData,
    treeLoading,
    selectedNum,
    pagination,
    buttonClass,
    deviceDetection,
    onSearch,
    resetForm,
    onbatchDel,
    openDialog,
    openTypeDialog,
    onTreeSelect,
    handleUpdate,
    handleDelete,
    handleSizeChange,
    onSelectionCancel,
    handleCurrentChange,
    handleSelectionChange
  };
}
