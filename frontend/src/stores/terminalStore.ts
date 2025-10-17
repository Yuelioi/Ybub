import type { TerminalLog } from '@/types'
import { SSHService } from '@services/server'
import { useServerStore } from './serverStore'
import { Events } from '@wailsio/runtime'
import { Event, SshCommandOutput } from '@models'

export const useTerminalStore = defineStore('terminal', () => {
  const serverStore = useServerStore()

  // 状态
  const terminalLogs = ref<TerminalLog[]>([])
  const currentCommand = ref('')

  // 计算属性
  const isExecuting = computed(() => terminalLogs.value.some((log) => log.isExecuting))

  // 初始化 - 监听 SSH 输出事件
  const init = () => {
    Events.On(Event.EventSshOutput, (event) => {
      const message: SshCommandOutput = event.data[0]

      const targetLog = computed(() =>
        terminalLogs.value.find((l) => {
          return l.commandId == message.commandId
        }),
      )

      if (targetLog.value && targetLog.value.isExecuting) {
        targetLog.value.outputs.push(message.line)
      }
    })

    Events.On(Event.EventSshComplete, (event) => {
      const message: SshCommandOutput = event.data[0]

      const targetLog = computed(() =>
        terminalLogs.value.find((l) => {
          return l.commandId == message.commandId
        }),
      )

      if (targetLog.value) {
        targetLog.value.isExecuting = false
      }
    })
  }

  // 执行命令
  const executeCommand = async () => {
    if (!currentCommand.value.trim() || !serverStore.currentServer) return
    if (isExecuting.value) return
    const commandToExecute = currentCommand.value
    const d = new Date()
    const timestamp = d.toLocaleTimeString()
    const commandId = d.getTime().toString()

    terminalLogs.value.push({
      commandId,
      timestamp,
      command: currentCommand.value,
      outputs: [],
      isExecuting: true,
    })

    // 清空输入
    currentCommand.value = ''

    try {
      await SSHService.ExecCommandByServerID(
        serverStore.currentServerId,
        commandId,
        commandToExecute,
      )
    } catch (error: any) {
      toast.error(error)
    }
  }

  // 清空终端
  const clearTerminal = () => {
    terminalLogs.value = []
  }

  // 设置当前命令
  const setCurrentCommand = (command: string) => {
    currentCommand.value = command
  }

  return {
    // 状态
    terminalLogs,
    currentCommand,

    // 计算属性
    isExecuting,

    // 方法
    init,
    executeCommand,
    clearTerminal,
    setCurrentCommand,
  }
})
