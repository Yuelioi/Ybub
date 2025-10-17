export type TabType = 'servers' | 'projects' | 'backup' | 'terminal'

export const useAppStore = defineStore('app', () => {
  // 状态
  const activeTab = ref<TabType>('servers')
  const showAddServer = ref(false)
  const showAddProject = ref(false)

  // 设置活动标签
  const setActiveTab = (tab: TabType) => {
    activeTab.value = tab
  }

  // 打开添加服务器对话框
  const openAddServerDialog = () => {
    showAddServer.value = true
  }

  // 关闭添加服务器对话框
  const closeAddServerDialog = () => {
    showAddServer.value = false
  }

  // 打开添加项目对话框
  const openAddProjectDialog = () => {
    showAddProject.value = true
  }

  // 关闭添加项目对话框
  const closeAddProjectDialog = () => {
    showAddProject.value = false
  }

  return {
    // 状态
    activeTab,
    showAddServer,
    showAddProject,

    // 方法
    setActiveTab,
    openAddServerDialog,
    closeAddServerDialog,
    openAddProjectDialog,
    closeAddProjectDialog,
  }
})
