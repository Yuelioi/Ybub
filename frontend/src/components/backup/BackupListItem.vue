<template>
  <div class="p-4 hover:bg-accent/5 transition-colors">
    <div class="flex items-center justify-between gap-4">
      <!-- 左侧信息 -->
      <div class="flex-1 min-w-0">
        <div class="flex items-center gap-3 mb-2">
          <DatabaseIcon class="size-5 text-primary flex-shrink-0" />
          <h3 class="font-semibold truncate">{{ project.name }}</h3>

          <div
            class="rounded-full border-6 flex items-center justify-center"
            :class="{
              'border-green-500': project.backupStatus == ProjectStatus.ProjectBackupStatusRunning,
              'border-red-500': project.backupStatus === ProjectStatus.ProjectBackupStatusError,
              'border-gray-500': project.backupStatus === ProjectStatus.ProjectBackupStatusStopped,
            }"></div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-x-6 gap-y-1 text-sm text-muted-foreground">
          <div class="flex items-center gap-2">
            <ClockIcon class="size-4 flex-shrink-0" />
            <span> {{ formatCron(project.schedule) }}</span>
          </div>
          <div v-if="project.dataPath" class="flex items-center gap-2">
            <DatabaseIcon class="size-4" />
            <span class="truncate">{{ project.dataPath }}</span>
          </div>
          <div v-if="project.lastBackup" class="flex items-center gap-2">
            <CalendarCheckIcon class="size-4 flex-shrink-0" />
            <span> {{ formatDate(project.lastBackup) }}</span>
          </div>
          <!-- <div
            v-if="
              project.nextBackup &&
              project.backupStatus === ProjectStatus.ProjectBackupStatusRunning
            "
            class="flex items-center gap-2">
            <CalendarIcon class="size-4 flex-shrink-0" />
            <span>下次: {{ formatDate(project.nextBackup) }}</span>
          </div> -->
        </div>
      </div>

      <!-- 右侧操作按钮 -->
      <div class="flex items-center gap-2 flex-shrink-0">
        <button
          class="btn btn-sm btn-primary"
          @click="backupStore.runBackupNow(project)"
          :disabled="serverStatus !== 'connected'">
          <PlayIcon class="size-4" />
          立即备份
        </button>
        <button
          class="btn btn-sm btn-outline btn-icon"
          @click="backupStore.toggleBackupStatus(project)"
          :title="
            project.backupStatus === ProjectStatus.ProjectBackupStatusRunning
              ? '暂停备份'
              : '启用备份'
          ">
          <component
            :is="
              project.backupStatus === ProjectStatus.ProjectBackupStatusRunning
                ? PauseIcon
                : PlayIcon
            "
            class="size-4" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  Database as DatabaseIcon,
  Clock as ClockIcon,
  // Calendar as CalendarIcon,
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
