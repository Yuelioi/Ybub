import { ProjectStatus, SshCommandOutput } from '@models'
import { SSHService } from '@services/server'
import { useProjectStore } from './projectStore'
import type { DockerOutput } from '@/types'
import { Event } from '@models'

import { Events } from '@wailsio/runtime'

export const useDockerStore = defineStore('docker', () => {
  const projectStore = useProjectStore()

  // 状态
  const commandOutput = ref<DockerOutput>({
    command: '',
    status: '',
    outputs: [],
    commandId: '',
  })
  const showCommandOutput = ref(false)

  const init = () => {
    Events.On(Event.EventSshOutput, (event) => {
      const message: SshCommandOutput = event.data[0]

      if (message.commandId !== commandOutput.value.commandId) {
        return
      }
      commandOutput.value.outputs.push(message.line)
    })

    Events.On(Event.EventSshComplete, (event) => {
      const commandId = event.data[0]
      if (commandId !== commandOutput.value.commandId) {
        return
      }
      commandOutput.value.status = 'finish'
    })
  }

  // 执行命令并显示输出
  const runCommandAndShowOutput = async (projectId: string, command: string) => {
    const project = projectStore.getProjectById(projectId)
    if (!project) return

    const commands = `cd ${project.path} && ${command}`
    const commandId = randomID()

    commandOutput.value = {
      commandId,
      command: commands,
      status: 'working',
      outputs: [],
    }
    showCommandOutput.value = true

    try {
      await SSHService.ExecProjectCommand(project, commandId, [commands])
    } catch (err: any) {
      toast.error(err)
      throw err
    }
  }

  // 统一的 Docker 命令执行封装
  const executeDockerCommand = async (
    projectId: string,
    command: string,
    successStatus: ProjectStatus | null,
    successMessage: string,
    errorMessage: string,
  ) => {
    try {
      await runCommandAndShowOutput(projectId, command)

      const project = projectStore.getProjectById(projectId)
      if (project) {
        if (successStatus) {
          project.status = successStatus
          projectStore.updateProjectStatus(projectId, successStatus)
        }
      }

      toast.success(successMessage)
    } catch (err: any) {
      console.log(err)
      toast.error(`${errorMessage}: ${err.message}`)

      const project = projectStore.getProjectById(projectId)
      if (project) {
        project.status = ProjectStatus.ProjectStatusError
        projectStore.updateProjectStatus(projectId, ProjectStatus.ProjectStatusError)
      }

      throw err
    }
  }

  // Docker Pull
  const dockerPull = async (projectId: string) => {
    await executeDockerCommand(
      projectId,
      'docker-compose pull',
      null,
      '镜像拉取完成',
      '容器拉取失败',
    )
  }

  // Docker Up
  const dockerUp = async (projectId: string) => {
    await executeDockerCommand(
      projectId,
      'docker-compose up -d',
      ProjectStatus.ProjectStatusRunning,
      '容器启动成功',
      '容器启动失败',
    )
  }

  // Docker Down
  const dockerDown = async (projectId: string) => {
    await executeDockerCommand(
      projectId,
      'docker-compose down',
      ProjectStatus.ProjectStatusStopped,
      '容器停止成功',
      '容器停止失败',
    )
  }

  // Docker Logs (不需要更新状态)
  const dockerLogs = async (projectId: string) => {
    await runCommandAndShowOutput(projectId, 'docker-compose logs --tail=50')
  }

  // 关闭输出对话框
  const closeCommandOutput = () => {
    showCommandOutput.value = false
  }

  return {
    // 状态
    commandOutput,
    showCommandOutput,

    // 方法
    init,
    runCommandAndShowOutput,
    dockerPull,
    dockerUp,
    dockerDown,
    dockerLogs,
    closeCommandOutput,
  }
})
