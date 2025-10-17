<template>
  <div class="card card-hover p-4">
    <div class="flex items-start justify-between mb-4">
      <div class="flex items-center gap-2">
        <FolderOpenIcon class="size-5 text-primary" />
        <h3 class="font-semibold">{{ project.name }}</h3>
      </div>
      <span
        :class="[
          'badge badge-sm',
          project.status === 'running'
            ? 'badge-primary'
            : project.status === 'stopped'
              ? 'badge-muted'
              : 'badge-destructive',
        ]">
        {{
          project.status === 'running' ? '运行中' : project.status === 'stopped' ? '已停止' : '错误'
        }}
      </span>
    </div>

    <div class="space-y-2 text-sm text-muted-foreground mb-4">
      <div class="flex items-center gap-2">
        <FolderIcon class="size-4" />
        <span class="truncate">{{ project.path }}</span>
      </div>
      <div v-if="project.dataPath" class="flex items-center gap-2">
        <DatabaseIcon class="size-4" />
        <span class="truncate">{{ project.dataPath }}</span>
      </div>
      <div v-if="project.enableBackup" class="flex items-center gap-1 text-primary">
        <CheckIcon class="size-4" />
        自动备份: {{ formatCron(project.schedule) }}
      </div>
    </div>

    <div class="grid grid-cols-2 gap-2 mb-3">
      <button
        class="btn btn-sm btn-outline"
        @click="dockerStore.dockerPull(project.id)"
        :disabled="serverStatus !== 'connected'">
        <DownloadIcon class="size-4" />
        Pull 镜像
      </button>
      <button
        class="btn btn-sm btn-outline"
        @click="dockerStore.dockerUp(project.id)"
        :disabled="serverStatus !== 'connected'">
        <PlayIcon class="size-4" />
        Up 启动
      </button>
      <button
        class="btn btn-sm btn-outline"
        @click="dockerStore.dockerDown(project.id)"
        :disabled="serverStatus !== 'connected'">
        <SquareIcon class="size-4" />
        Down 停止
      </button>
      <button
        class="btn btn-sm btn-outline"
        @click="dockerStore.dockerLogs(project.id)"
        :disabled="serverStatus !== 'connected'">
        <FileTextIcon class="size-4" />
        查看日志
      </button>
    </div>

    <div class="flex gap-2 pt-3 border-t border-border">
      <button class="btn btn-sm btn-outline flex-1" @click="$emit('edit')">
        <SettingsIcon class="size-4" />
        设置
      </button>
      <button
        class="btn btn-sm btn-outline btn-icon"
        @click="centerStore.removeProject(project.id)">
        <Trash2Icon class="size-4" />
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Project } from '@models'
import {
  FolderOpen as FolderOpenIcon,
  Folder as FolderIcon,
  Database as DatabaseIcon,
  Check as CheckIcon,
  Download as DownloadIcon,
  Play as PlayIcon,
  Square as SquareIcon,
  FileText as FileTextIcon,
  Settings as SettingsIcon,
  Trash2 as Trash2Icon,
} from 'lucide-vue-next'

const dockerStore = useDockerStore()
const centerStore = useCenterStore()
import { formatCron } from '@/utils'

defineProps<{ project: Project; serverStatus: string }>()

defineEmits<{
  edit: []
}>()
</script>
