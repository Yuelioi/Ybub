<template>
  <div class="card card-hover p-4">
    <div class="flex items-start justify-between mb-4">
      <div class="flex items-center gap-2">
        <DatabaseIcon class="size-5 text-primary" />
        <h3 class="font-semibold">{{ project.name }}</h3>

        <div class="text-muted-foreground text-sm">
          {{
            project.backupStatus === ProjectStatus.ProjectBackupStatusRunning
              ? '运行中'
              : project.backupStatus === ProjectStatus.ProjectBackupStatusError
                ? '错误'
                : '已停止'
          }}
        </div>
        <div
          class="rounded-full border-6 flex items-center justify-center"
          :class="{
            'border-green-500': project.backupStatus == ProjectStatus.ProjectBackupStatusRunning,
            'border-red-500': project.backupStatus === ProjectStatus.ProjectBackupStatusError,
            'border-gray-500': project.backupStatus === ProjectStatus.ProjectBackupStatusStopped,
          }"></div>
      </div>
    </div>

    <div class="space-y-2 text-sm text-muted-foreground mb-4">
      <div class="flex items-center gap-2">
        <ClockIcon class="size-4" />
        <span>计划: {{ formatCron(project.schedule) }}</span>
      </div>
      <div v-if="project.dataPath" class="flex items-center gap-2">
        <DatabaseIcon class="size-4" />
        <span class="truncate">{{ project.dataPath }}</span>
      </div>
      <div v-if="project.lastBackup" class="flex items-center gap-2">
        <CalendarCheckIcon class="size-4" />
        <span>上次: {{ formatDate(project.lastBackup) }}</span>
      </div>
      <div
        v-if="
          project.nextBackup && project.backupStatus === ProjectStatus.ProjectBackupStatusRunning
        "
        class="flex items-center gap-2">
        <CalendarIcon class="size-4" />
        <span>下次: {{ formatDate(project.nextBackup) }}</span>
      </div>
    </div>

    <div class="flex gap-2">
      <button
        class="btn btn-sm btn-primary flex-1"
        @click="backupStore.runBackupNow(project)"
        :disabled="serverStatus !== 'connected'">
        <PlayIcon class="size-4" />
        立即备份
      </button>
      <button class="btn btn-sm btn-outline" @click="backupStore.toggleBackupStatus(project)">
        <component
          :is="
            project.backupStatus === ProjectStatus.ProjectBackupStatusRunning ? PauseIcon : PlayIcon
          "
          class="size-4" />
        {{ project.backupStatus === ProjectStatus.ProjectBackupStatusRunning ? '暂停' : '启用' }}
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  Database as DatabaseIcon,
  Clock as ClockIcon,
  Calendar as CalendarIcon,
  CalendarCheck as CalendarCheckIcon,
  Play as PlayIcon,
  Pause as PauseIcon,
} from 'lucide-vue-next'
import { formatDate } from '@yuelioi/utils'
import { ProjectStatus, type Project } from '@models'
import { formatCron } from '@/utils'

defineProps<{
  project: Project
  serverStatus: string
}>()

const backupStore = useBackupStore()
</script>
