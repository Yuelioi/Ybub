import type { ServerFormData } from '@/types/server'
import { ServerStatus, type Server } from '@models'
import { ConfigManager } from '@services/conf'
import { SSHService } from '@services/server'

export const useServerStore = defineStore('server', () => {
  // 状态
  const servers = ref<Server[]>([])
  const currentServerId = ref('')

  // 计算属性
  const currentServer = computed(() => servers.value.find((s) => s.id === currentServerId.value))

  const connectedServers = computed(() => servers.value.filter((s) => s.status === 'connected'))

  // 初始化
  const init = async () => {
    servers.value = await ConfigManager.ListServers()
    // 默认选中第一个服务器
    if (servers.value.length > 0 && !currentServerId.value) {
      currentServerId.value = servers.value[0].id
    }

    // 连接服务器
    for (const server of servers.value) {
      try {
        server.status = ServerStatus.ServerStatusConnecting
        await connectServer(server.id)
        server.status = ServerStatus.ServerStatusConnected
      } catch (error) {
        console.error(error)
        server.status = ServerStatus.ServerStatusDisconnected
      }
    }
  }

  // 设置当前服务器
  const setCurrentServer = (serverId: string) => {
    currentServerId.value = serverId
  }

  const updateServer = async (id: string, server: ServerFormData) => {
    try {
      await ConfigManager.UpdateServer(id, server as Server)
      currentServerId.value = id
      await connectServer(id)
      return server
    } catch (error: any) {
      toast.error('更新服务器失败: ' + error.message)
      throw error
    }
  }

  // 连接服务器
  const connectServer = async (id: string) => {
    const server = servers.value.find((s) => s.id === id)
    if (!server) return

    server.status = ServerStatus.ServerStatusConnecting
    try {
      await SSHService.TestConnection(server)
      server.status = ServerStatus.ServerStatusConnected
      server.lastConnected = new Date()
      toast.success(`服务器 ${server.name} 连接成功`)
    } catch (error: any) {
      server.status = ServerStatus.ServerStatusDisconnected
      toast.error(`连接失败: ${error.message}`)
      throw error
    }
  }

  // 断开服务器
  const disconnectServer = (id: string) => {
    const server = servers.value.find((s) => s.id === id)
    if (server) {
      server.status = ServerStatus.ServerStatusDisconnected
      if (currentServerId.value === id) {
        currentServerId.value = ''
      }
      toast.info(`服务器 ${server.name} 已断开`)
    }
  }

  // 获取服务器
  const getServerById = (id: string) => {
    return servers.value.find((s) => s.id === id)
  }

  return {
    // 状态
    servers,
    currentServerId,

    // 计算属性
    currentServer,
    connectedServers,

    // 方法
    init,
    setCurrentServer,
    updateServer,
    connectServer,
    disconnectServer,
    getServerById,
  }
})
