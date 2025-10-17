import { BackupOutput, BackupProgress, Event, ProjectStatus, type Project } from '@models'
import { ServerCenter } from '@services/service_center'
import { Events } from '@wailsio/runtime'

import parser from 'cron-parser'

export const useBackupStore = defineStore('backup', () => {
  const projectStore = useProjectStore()

  const { projects } = storeToRefs(projectStore)

  const init = () => {
    Events.On(Event.EventBackupProgress, (event) => {
      const message: BackupProgress = event.data[0]
    })

    Events.On(Event.EventBackupComplete, async (event) => {
      const message: BackupOutput = event.data[0]

      const project = projectStore.getProjectById(message.projectId)
      if (project) {
        project.lastBackup = new Date()
        let nextBackup = new Date(Date.now() + 86400000)

        try {
          const interval = parser.parse(project.schedule)
          nextBackup = interval.next().toDate()
        } catch (err) {
          console.error('Cron 表达式解析错误:', err)
        }
        project.nextBackup = nextBackup
        await ServerCenter.UpdateProject(project)
        toast.success(`项目 ${project.name} 备份完成`)
      }
    })
  }

  // 立即执行备份
  const runBackupNow = async (project: Project) => {
    const currentProject = projects.value.find((b) => b.id === project.id)
    if (!currentProject) return

    try {
      toast.info(`正在备份项目: ${currentProject.name}`)

      await ServerCenter.BackupProjectManually(currentProject)

      const currentTime = new Date()
      let nextBackup = new Date(Date.now() + 86400000)

      try {
        const interval = parser.parse(project.schedule)
        nextBackup = interval.next().toDate()
      } catch (err) {
        console.error('Cron 表达式解析错误:', err)
      }

      // 更新最后备份时间
      currentProject.lastBackup = currentTime
      currentProject.nextBackup = nextBackup

      toast.success(`项目 ${currentProject.name} 备份完成`)
      return true
    } catch (error: any) {
      toast.error(`备份失败: ${error.message}`)
      throw error
    }
  }

  // 切换备份状态
  const toggleBackupStatus = async (project: Project) => {
    const currentProject = projects.value.find((b) => b.id === project.id)
    if (!currentProject) return
    const newStatus =
      currentProject.backupStatus === ProjectStatus.ProjectBackupStatusRunning
        ? ProjectStatus.ProjectStatusStopped
        : ProjectStatus.ProjectBackupStatusRunning
    try {
      await projectStore.updateProject(project.id, {
        backupStatus: newStatus,
      })
      currentProject.backupStatus = newStatus

      toast.success(
        `项目 ${project.name} 自动备份已${newStatus === ProjectStatus.ProjectBackupStatusRunning ? '启用' : '暂停'}`,
      )
    } catch (error: any) {
      toast.error('切换备份状态失败: ' + error.message)
      throw error
    }
  }

  return {
    runBackupNow,
    toggleBackupStatus,
    init,
  }
})
