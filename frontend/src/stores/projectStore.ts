import type { Project, ProjectStatus } from '@models'
import { ConfigManager } from '@services/conf'
import { useServerStore } from './serverStore'

export const useProjectStore = defineStore('project', () => {
  const serverStore = useServerStore()

  const projects = ref<Project[]>([])

  const currentServerProjects = computed(() =>
    projects.value.filter((p) => p.serverId === serverStore.currentServerId),
  )

  const init = async () => {
    projects.value = await ConfigManager.ListProjects()
  }

  // 根据服务器ID获取项目
  const getProjectsByServerId = (serverId: string) => {
    return projects.value.filter((p) => p.serverId === serverId)
  }

  // 根据ID获取项目
  const getProjectById = (id: string) => {
    return projects.value.find((p) => p.id === id)
  }

  // 更新项目
  const updateProject = async (id: string, updates: Partial<Project>) => {
    const project = projects.value.find((p) => p.id === id)
    if (!project) return

    try {
      Object.assign(project, updates)
      await ConfigManager.UpdateProject(project)
      toast.success(`项目 ${project.name} 已更新`)
    } catch (error: any) {
      toast.error('更新项目失败: ' + error.message)
      throw error
    }
  }

  // 更新项目状态
  const updateProjectStatus = (id: string, status: ProjectStatus) => {
    const project = projects.value.find((p) => p.id === id)
    if (project) {
      project.status = status
    }

    updateProject(id, { status })
  }

  return {
    // 状态
    projects,

    // 计算属性
    currentServerProjects,

    // 方法
    init,
    getProjectsByServerId,
    getProjectById,
    updateProject,
    updateProjectStatus,
  }
})
