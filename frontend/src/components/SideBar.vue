<template>
  <aside class="w-64 border-r border-border bg-card/30 p-4">
    <nav class="space-y-2">
      <button
        v-for="tab in tabs"
        :key="tab.id"
        @click="appStore.setActiveTab(tab.id)"
        :class="[
          'btn btn-ghost w-full justify-start gap-2',
          activeTab === tab.id ? 'bg-accent text-accent-foreground' : '',
        ]">
        <component :is="tab.icon" class="size-5" />
        <span>{{ tab.label }}</span>
      </button>
    </nav>

    <!-- 服务器列表 -->
    <div v-if="servers.length > 0" class="design-2 mt-6 pt-6 border-t border-border">
      <div class="text-xs font-medium text-muted-foreground mb-3 px-1">当前服务器</div>
      <div class="space-y-1">
        <button
          v-for="server in servers"
          :key="server.id"
          @click="currentServerId = server.id"
          :class="[
            'w-full rounded-lg p-2.5 transition-all text-left flex items-center gap-2.5',
            currentServerId === server.id
              ? 'bg-primary/10 border border-primary/20'
              : 'hover:bg-accent/10 border border-transparent',
          ]">
          <div
            class="size-2 rounded-full flex-shrink-0"
            :class="server.status === 'connected' ? 'bg-green-500' : 'bg-muted-foreground'" />
          <div class="flex-1 min-w-0">
            <div class="text-sm font-medium truncate">{{ server.name }}</div>
          </div>
          <Server
            class="size-3.5 flex-shrink-0"
            :class="currentServerId === server.id ? 'text-primary' : 'text-muted-foreground'" />
        </button>
      </div>
    </div>
  </aside>
</template>

<script setup lang="ts">
import type { Tab } from '@/types'
import {
  Server as ServerIcon,
  FolderOpen as FolderOpenIcon,
  Terminal as TerminalIcon,
  Database as DatabaseIcon,
  Server,
} from 'lucide-vue-next'

const appStore = useAppStore()
const serverStore = useServerStore()

const { servers } = storeToRefs(serverStore)
const { activeTab } = storeToRefs(appStore)

const { currentServerId } = storeToRefs(serverStore)

// 标签配置
const tabs: Tab[] = [
  { id: 'servers', label: '服务器', icon: ServerIcon },
  { id: 'projects', label: '项目', icon: FolderOpenIcon },
  { id: 'backup', label: '备份管理', icon: DatabaseIcon },
  { id: 'terminal', label: 'SSH 终端', icon: TerminalIcon },
]
</script>
