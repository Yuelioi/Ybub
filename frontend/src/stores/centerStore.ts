import { Server, type Project, ProjectStatus, ServerStatus } from '@models'
import { ServerCenter } from '@services/service_center'
import { ConfigManager } from '@services/conf'

export const useCenterStore = defineStore('centerStore', () => {
  const serverStore = useServerStore()
  const projectStore = useProjectStore()

  const { servers } = storeToRefs(serverStore)
  const { projects } = storeToRefs(projectStore)

  // 添加服务器
  const createServer = async (server: Omit<Server, 'status'>) => {
    const newServer: Server = {
      ...server,
      port: Number(server.port) || 22,
      status: ServerStatus.ServerStatusDisconnected,
    }

    try {
      await ConfigManager.CreateServer(newServer)
      servers.value.push(newServer)
      serverStore.currentServerId = server.id
      await serverStore.connectServer(server.id)
      newServer.status = ServerStatus.ServerStatusConnected

      return newServer
    } catch (error: any) {
      toast.error('添加服务器失败: ' + error.message)
      throw error
    }
  }

  // 添加项目
  const createProject = async (project: Omit<Project, 'status'>) => {
    const newProject: Project = {
      ...project,
      serverId: serverStore.currentServerId,
      status: ProjectStatus.ProjectBackupStatusStopped,
    }

    try {
      await ServerCenter.CreateProjectWithScheduler(newProject)
      projects.value.push(newProject)

      toast.success(`项目 ${newProject.name} 已添加`)
      return newProject
    } catch (error: any) {
      toast.error('添加项目失败: ' + error.message)
      throw error
    }
  }

  // 删除服务器
  const removeServer = async (id: string) => {
    try {
      await ServerCenter.RemoveServer(id)
      servers.value = serverStore.servers.filter((s) => s.id !== id)
      if (serverStore.currentServerId === id) {
        serverStore.currentServerId = servers.value[0]?.id || ''
      }
      projects.value = projects.value.filter((p) => {
        return p.serverId === id
      })

      toast.success('服务器已删除')
    } catch (error: any) {
      toast.error('删除服务器失败: ' + error.message)
      throw error
    }
  }

  // 删除项目
  const removeProject = async (id: string) => {
    try {
      await ServerCenter.RemoveProject(id)
      projects.value = projects.value.filter((p) => p.id !== id)

      toast.success('项目已删除')
    } catch (error: any) {
      toast.error('删除项目失败: ' + error.message)
      throw error
    }
  }

  return { createServer, createProject, removeServer, removeProject }
})
