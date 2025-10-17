<template>
  <div class="p-4 hover:bg-accent/5 transition-colors">
    <div class="flex items-center justify-between gap-4">
      <!-- 左侧信息 -->
      <div class="flex-1 min-w-0">
        <div class="flex items-center gap-3 mb-2">
          <FolderOpenIcon class="size-5 text-primary flex-shrink-0" />
          <h3 class="font-semibold truncate">{{ project.name }}</h3>
          <span
            :class="[
              'badge badge-sm flex-shrink-0',
              project.status === 'running'
                ? 'badge-primary'
                : project.status === 'stopped'
                  ? 'badge-muted'
                  : 'badge-destructive',
            ]">
            {{
              project.status === 'running'
                ? '运行中'
                : project.status === 'stopped'
                  ? '已停止'
                  : '错误'
            }}
          </span>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-2 text-sm text-muted-foreground">
          <div class="flex items-center gap-2">
            <FolderIcon class="size-4 flex-shrink-0" />
            <span class="truncate">{{ project.path }}</span>
          </div>
          <div v-if="project.dataPath" class="flex items-center gap-2">
            <DatabaseIcon class="size-4 flex-shrink-0" />
            <span class="truncate">{{ project.dataPath }}</span>
          </div>
          <div v-if="project.enableBackup" class="flex items-center gap-1 text-primary">
            <CheckIcon class="size-4 flex-shrink-0" />
            <span>自动备份:{{ formatCron(project.schedule) }}</span>
          </div>
        </div>
      </div>

      <!-- 右侧操作按钮 -->
      <div class="flex items-center gap-2 flex-shrink-0">
        <button
          class="btn btn-sm btn-outline btn-icon"
          @click="dockerStore.dockerPull(project.id)"
          :disabled="serverStatus !== 'connected'"
          title="Pull 镜像">
          <DownloadIcon class="size-4" />
        </button>
        <button
          class="btn btn-sm btn-outline btn-icon"
          @click="dockerStore.dockerUp(project.id)"
          :disabled="serverStatus !== 'connected'"
          title="Up 启动">
          <PlayIcon class="size-4" />
        </button>
        <button
          class="btn btn-sm btn-outline btn-icon"
          @click="dockerStore.dockerDown(project.id)"
          :disabled="serverStatus !== 'connected'"
          title="Down 停止">
          <SquareIcon class="size-4" />
        </button>
        <button
          class="btn btn-sm btn-outline btn-icon"
          @click="dockerStore.dockerLogs(project.id)"
          :disabled="serverStatus !== 'connected'"
          title="查看日志">
          <FileTextIcon class="size-4" />
        </button>
        <div class="w-px h-6 bg-border"></div>
        <button class="btn btn-sm btn-outline btn-icon" @click="$emit('edit')" title="设置">
          <SettingsIcon class="size-4" />
        </button>
        <button
          class="btn btn-sm btn-outline btn-icon"
          @click="centerStore.removeProject(project.id)"
          title="删除">
          <Trash2Icon class="size-4" />
        </button>
      </div>
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
import { formatCron } from '@/utils'

const dockerStore = useDockerStore()

const centerStore = useCenterStore()

defineProps<{ project: Project; serverStatus: string }>()

defineEmits<{
  edit: []
}>()
</script>
