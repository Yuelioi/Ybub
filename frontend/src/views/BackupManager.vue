<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div class="title-deco">
        <h2>备份管理</h2>
      </div>
    </div>

    <!-- 搜索和视图切换 -->
    <div class="flex items-center gap-3">
      <div class="relative flex-1">
        <SearchIcon class="absolute left-3 top-1/2 -translate-y-1/2 size-4 text-muted-foreground" />
        <input
          v-model="searchQuery"
          type="text"
          placeholder="搜索项目名称、备份计划或服务器..."
          class="input pl-9 w-full" />
        <button
          v-if="searchQuery"
          @click="searchQuery = ''"
          class="absolute right-2 top-1/2 -translate-y-1/2 p-1 hover:bg-muted rounded">
          <XIcon class="size-4 text-muted-foreground" />
        </button>
      </div>
      <div class="flex items-center gap-1 bg-muted rounded-lg p-1">
        <button
          @click="viewMode = 'grid'"
          :class="[
            'p-2 rounded transition-colors',
            viewMode === 'grid' ? 'bg-background shadow-sm' : 'hover:bg-background/50',
          ]"
          title="网格视图">
          <LayoutGridIcon class="size-4" />
        </button>
        <button
          @click="viewMode = 'list'"
          :class="[
            'p-2 rounded transition-colors',
            viewMode === 'list' ? 'bg-background shadow-sm' : 'hover:bg-background/50',
          ]"
          title="列表视图">
          <LayoutListIcon class="size-4" />
        </button>
      </div>
    </div>

    <!-- 空状态 -->
    <div v-if="backupProjects.length === 0" class="card p-8 text-center">
      <DatabaseIcon class="size-12 text-muted-foreground mx-auto mb-4" />
      <p class="text-muted-foreground mb-2">没有配置自动备份的项目</p>
      <p class="text-sm text-muted-foreground">在添加项目时勾选"启用自动备份"即可</p>
    </div>

    <!-- 搜索无结果 -->
    <div v-else-if="filteredBackupGroups.length === 0" class="card p-8 text-center">
      <SearchIcon class="size-12 text-muted-foreground mx-auto mb-4" />
      <p class="text-muted-foreground mb-2">未找到匹配的备份任务</p>
      <p class="text-sm text-muted-foreground">尝试使用不同的关键词搜索</p>
    </div>

    <!-- 按服务器分组显示 -->
    <div v-else class="space-y-6">
      <div v-for="group in filteredBackupGroups" :key="group.serverId" class="space-y-3">
        <!-- 服务器头部 -->
        <div class="flex items-center gap-3 px-1">
          <div class="flex items-center gap-2">
            <ServerIcon class="size-5 text-primary" />
            <h3 class="font-semibold text-lg">{{ group.serverName }}</h3>
            <span class="text-sm text-muted-foreground">({{ group.projects.length }})</span>
          </div>
          <span
            :class="[
              'badge badge-sm',
              group.serverStatus === 'connected' ? 'badge-primary' : 'badge-muted',
            ]">
            {{ group.serverStatus === 'connected' ? '已连接' : '未连接' }}
          </span>
        </div>

        <!-- 网格视图 -->
        <div v-if="viewMode === 'grid'" class="grid grid-cols-1 lg:grid-cols-2 gap-4">
          <template v-for="project in group.projects" :key="project.id">
            <BackupCard :project="project" :server-status="group.serverStatus"></BackupCard>
          </template>
        </div>

        <!-- 列表视图 -->
        <div v-else class="card divide-y divide-border">
          <template v-for="project in group.projects" :key="project.id">
            <BackupListItem :project="project" :server-status="group.serverStatus"></BackupListItem>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  Database as DatabaseIcon,
  Search as SearchIcon,
  X as XIcon,
  LayoutGrid as LayoutGridIcon,
  LayoutList as LayoutListIcon,
  Server as ServerIcon,
} from 'lucide-vue-next'
import { ref, computed } from 'vue'
import { type Project } from '@models'

const serverStore = useServerStore()
const projectStore = useProjectStore()

const { servers } = storeToRefs(serverStore)
const { projects } = storeToRefs(projectStore)

const backupProjects = computed(() => {
  return projects.value.filter((project) => project.enableBackup)
})

interface BackupGroup {
  serverId: string
  serverName: string
  serverStatus: string
  projects: Project[]
}

// 视图模式和搜索
const viewMode = useStorage<'grid' | 'list'>('backup.view.mode', 'list')
const searchQuery = ref('')

// 按服务器分组备份任务
const backupGroups = computed<BackupGroup[]>(() => {
  const groups = new Map()

  backupProjects.value.forEach((project) => {
    const server = servers.value.find((s) => s.id === project.serverId)
    if (!server) return

    if (!groups.has(project.serverId)) {
      groups.set(project.serverId, {
        serverId: project.serverId,
        serverName: server.name,
        serverStatus: server.status,
        projects: [],
      })
    }

    groups.get(project.serverId).projects.push(project)
  })

  return Array.from(groups.values())
})

// 过滤备份任务组
const filteredBackupGroups = computed(() => {
  if (!searchQuery.value.trim()) {
    return backupGroups.value
  }

  const query = searchQuery.value.toLowerCase()

  return backupGroups.value
    .map((group) => ({
      ...group,
      projects: group.projects.filter(
        (backup: any) =>
          backup.projectName.toLowerCase().includes(query) ||
          backup.schedule.toLowerCase().includes(query) ||
          group.serverName.toLowerCase().includes(query),
      ),
    }))
    .filter((group) => group.projects.length > 0)
})
</script>
