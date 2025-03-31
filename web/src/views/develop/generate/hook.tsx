import dayjs from "dayjs";
import { message } from "@/utils/message";
import { getKeyList } from "@pureadmin/utils";
import type { PaginationProps } from "@pureadmin/table";
import { type Ref, reactive, ref, onMounted } from "vue";
import { useRouter } from "vue-router";

export function useRole(tableRef: Ref) {
  const router = useRouter();
  const form = reactive({
    name: "",
    description: "",
    created_at: ""
  });
  const dataList = ref([]);
  const loading = ref(true);
  const selectedNum = ref(0);

  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });

  const columns: TableColumnList = [
    {
      label: "序号",
      type: "index",
      width: 80
    },
    {
      label: "表名称",
      prop: "name",
      minWidth: 120
    },
    {
      label: "表描述",
      prop: "description",
      minWidth: 150
    },
    {
      label: "实体",
      prop: "className",
      minWidth: 120
    },
    {
      label: "创建时间",
      prop: "created_at",
      minWidth: 180,
      formatter: ({ created_at }) =>
        dayjs(created_at).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "更新时间",
      prop: "updated_at",
      minWidth: 180,
      formatter: ({ updated_at }) =>
        dayjs(updated_at).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "操作",
      fixed: "right",
      width: 350,
      slot: "operation"
    }
  ];

  function handleSizeChange(val: number) {
    console.log(`${val} items per page`);
  }

  function handleCurrentChange(val: number) {
    console.log(`current page: ${val}`);
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

  /** 删除 */
  function handleDelete(row) {
    // 根据实际业务，调用接口删除当前行数据
    message(`${row.name}已被删除`, { type: "success" });
    onSearch();
  }

  /** 批量删除 */
  function onbatchDel() {
    // 返回当前选中的行
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    // 接下来根据实际业务，通过选中行的某项数据，比如下面的id，调用接口进行批量删除
    message(`已删除序号为 ${getKeyList(curSelected, "id")} 的数据`, {
      type: "success"
    });
    tableRef.value.getTableRef().clearSelection();
    onSearch();
  }

  /** 清空日志 */
  function clearAll() {
    // 根据实际业务，调用接口删除所有日志数据
    message("已删除所有日志数据", {
      type: "success"
    });
    onSearch();
  }

  async function onSearch() {
    loading.value = true;
    try {
      // TODO: 替换为实际的获取代码生成配置列表接口
      const mockData = {
        list: [
          {
            id: 1,
            name: "demo_gen",
            description: "代码生成测试表",
            className: "DemoGen",
            created_at: "2022-11-01 17:27:43",
            updated_at: "2024-03-14 11:33:03"
          }
        ],
        total: 10,
        pageSize: 10,
        currentPage: 1
      };
      dataList.value = mockData.list;
      pagination.total = mockData.total;
      pagination.pageSize = mockData.pageSize;
      pagination.currentPage = mockData.currentPage;
    } catch {
      message("获取列表失败", { type: "error" });
    } finally {
      loading.value = false;
    }
  }

  const resetForm = formEl => {
    if (!formEl) return;
    formEl.resetFields();
    onSearch();
  };

  onMounted(() => {
    onSearch();
  });

  // 编辑配置
  const handleEdit = row => {
    router.push({
      path: "/develop/generate/edit",
      query: { id: row.id }
    });
  };

  return {
    form,
    loading,
    columns,
    dataList,
    pagination,
    selectedNum,
    handleDelete,
    onSearch,
    clearAll,
    resetForm,
    onbatchDel,
    handleEdit,
    handleSizeChange,
    onSelectionCancel,
    handleCurrentChange,
    handleSelectionChange
  };
}
