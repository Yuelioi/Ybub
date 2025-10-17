<template>
  <div class="card card-hover p-4">
    <div class="flex items-start justify-between mb-4">
      <div class="flex items-center gap-2">
        <ServerIcon class="size-5 text-primary" />
        <h3 class="font-semibold">{{ server.name }}</h3>
      </div>
      <span :class="['badge badge-sm', statusClass]">
        {{ statusText }}
      </span>
    </div>

    <div class="space-y-2 text-sm text-muted-foreground mb-4">
      <div>{{ server.host }}:{{ server.port }}</div>
      <div>用户: {{ server.username }}</div>
      <div class="pt-2 border-t border-border">
        <span class="text-foreground font-medium">{{ projectCount }}</span>
        个 Docker 项目
      </div>
    </div>

    <div class="flex gap-2">
      <button
        v-if="server.status !== 'connected'"
        class="btn btn-sm btn-primary flex-1"
        @click="$emit('connect')">
        连接
      </button>
      <button v-else class="btn btn-sm btn-destructive flex-1" @click="$emit('disconnect')">
        断开
      </button>
      <button class="btn btn-sm btn-outline btn-icon" @click="$emit('edit')">
        <Settings class="size-5" />
      </button>
      <button class="btn btn-sm btn-outline btn-icon" @click="$emit('remove')">
        <Trash2Icon class="size-4" />
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { Server as ServerIcon, Trash2 as Trash2Icon, Settings } from 'lucide-vue-next'
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
