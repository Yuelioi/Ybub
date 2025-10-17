<template>
  <div class="p-4 hover:bg-accent/5 transition-colors">
    <div class="flex items-center justify-between gap-4">
      <!-- 左侧信息 -->
      <div class="flex-1 min-w-0">
        <div class="flex items-center gap-3 mb-2">
          <ServerIcon class="size-5 text-primary flex-shrink-0" />
          <h3 class="font-semibold truncate">{{ server.name }}</h3>
          <span :class="['badge badge-sm flex-shrink-0', statusClass]">
            {{ statusText }}
          </span>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-3 gap-x-6 gap-y-1 text-sm text-muted-foreground">
          <div class="flex items-center gap-2">
            <GlobeIcon class="size-4 flex-shrink-0" />
            <span class="truncate">{{ server.host }}:{{ server.port }}</span>
          </div>
          <div class="flex items-center gap-2">
            <UserIcon class="size-4 flex-shrink-0" />
            <span class="truncate">{{ server.username }}</span>
          </div>
          <div class="flex items-center gap-2">
            <FolderOpenIcon class="size-4 flex-shrink-0" />
            <span>
              <span class="text-foreground font-medium">{{ projectCount }}</span>
              个项目
            </span>
          </div>
        </div>
      </div>

      <!-- 右侧操作按钮 -->
      <div class="flex items-center gap-2 flex-shrink-0">
        <button
          v-if="server.status !== 'connected'"
          class="btn btn-sm btn-primary"
          @click="$emit('connect')">
          <PlugIcon class="size-4" />
          连接
        </button>
        <button v-else class="btn btn-sm btn-destructive" @click="$emit('disconnect')">
          <Unplug2Icon class="size-4" />
          断开
        </button>

        <button class="btn btn-sm btn-outline btn-icon" @click="$emit('edit')">
          <Settings class="size-5" />
        </button>
        <button class="btn btn-sm btn-outline btn-icon" @click="$emit('remove')" title="删除">
          <Trash2Icon class="size-4" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import {
  Server as ServerIcon,
  Globe as GlobeIcon,
  User as UserIcon,
  FolderOpen as FolderOpenIcon,
  Plug as PlugIcon,
  Unplug as Unplug2Icon,
  Settings,
  Trash2 as Trash2Icon,
} from 'lucide-vue-next'
import type { Server } from '@models'

const props = defineProps<{
  server: Server
  projectCount: number
}>()

defineEmits<{
  edit: []
  remove: []
  connect: []
  disconnect: []
}>()

const statusClass = computed(() => {
  switch (props.server.status) {
    case 'connected':
      return 'badge-primary'
    case 'connecting':
      return 'badge-accent'
    default:
      return 'badge-muted'
  }
})

const statusText = computed(() => {
  switch (props.server.status) {
    case 'connected':
      return '已连接'
    case 'connecting':
      return '连接中'
    default:
      return '未连接'
  }
})
</script>
