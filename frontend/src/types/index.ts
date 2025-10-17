// src/models/types.ts

export interface TerminalLog {
  commandId: string
  timestamp: string // 命令执行时间
  command: string // 执行的命令
  outputs: string[] // 输出行数组，支持逐步添加
  isExecuting?: boolean // 是否正在执行中
}

export interface DockerOutput {
  commandId: string
  command: string
  status: string
  outputs: string[]
}

export interface Tab {
  id: 'servers' | 'projects' | 'backup' | 'terminal'
  label: string
  icon: any
}
