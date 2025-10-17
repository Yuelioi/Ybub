<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div class="title-deco">
        <h2>Docker 项目</h2>
      </div>
      <button
        class="btn btn-primary"
        @click="dialog.openCreate"
        :disabled="!currentServer"
        title="添加项目">
        <PlusIcon class="size-4" />
        添加项目
      </button>
    </div>

    <!-- 搜索和视图切换 -->
    <div class="flex items-center gap-3">
      <div class="relative flex-1">
        <SearchIcon class="absolute left-3 top-1/2 -translate-y-1/2 size-4 text-muted-foreground" />
        <input
          v-model="searchQuery"
          type="text"
          placeholder="搜索项目名称、路径或服务器..."
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
    <div v-if="projects.length === 0" class="card p-8 text-center">
      <FolderOpenIcon class="size-12 text-muted-foreground mx-auto mb-4" />
      <p class="text-muted-foreground mb-4">还没有添加 Docker 项目</p>
      <button
        class="btn btn-sm btn-primary"
        @click="$emit('openAddProject')"
        :disabled="!currentServer">
        <PlusIcon class="size-4" />
        添加第一个项目
      </button>
    </div>

    <!-- 搜索无结果 -->
    <div v-else-if="filteredProjectGroups.length === 0" class="card p-8 text-center">
      <SearchIcon class="size-12 text-muted-foreground mx-auto mb-4" />
      <p class="text-muted-foreground mb-2">未找到匹配的项目</p>
      <p class="text-sm text-muted-foreground">尝试使用不同的关键词搜索</p>
    </div>

    <!-- 按服务器分组显示 -->
    <div v-else class="space-y-6">
      <div v-for="group in filteredProjectGroups" :key="group.serverId" class="space-y-3">
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
          <div v-for="project in group.projects" :key="project.id">
            <ProjectCard
              :project="project"
              :server-status="group.serverStatus"
              :backup-task="null"
              @edit="dialog.openUpdate(project)"
              @remove="centerStore.removeProject(project.id)"></ProjectCard>
          </div>
        </div>

        <!-- 列表视图 -->
        <div v-else class="card divide-y divide-border">
          <div v-for="project in group.projects" :key="project.id">
            <ProjectListItem
              :project="project"
              :server-status="group.serverStatus"
              :backup-task="null"
              @edit="dialog.openUpdate(project)"
              @remove="centerStore.removeProject(project.id)"></ProjectListItem>
          </div>
        </div>
      </div>
    </div>
  </div>

  <ProjectDialog
    :open="dialog.state.open"
    :mode="dialog.state.mode"
    :initial-data="dialog.initialData.value"
    :server-name="currentServer?.name"
    @close="dialog.close()"
    @save="handleSave" />
</template>

<script setup lang="ts">
import type { Project } from '@models'
import ProjectListItem from '@/components/project/ProjectListItem.vue'
import {
  FolderOpen as FolderOpenIcon,
  Plus as PlusIcon,
  Search as SearchIcon,
  X as XIcon,
  LayoutGrid as LayoutGridIcon,
  LayoutList as LayoutListIcon,
  Server as ServerIcon,
} from 'lucide-vue-next'

const serverStore = useServerStore()
const projectStore = useProjectStore()
const centerStore = useCenterStore()

const { projects } = storeToRefs(projectStore)
const { currentServer, servers } = storeToRefs(serverStore)

import { useDialog } from '@/composables/useDialog'
import ProjectCard from '@/components/project/ProjectCard.vue'

const dialog = useDialog()

// 视图模式和搜索
const viewMode = useStorage<'grid' | 'list'>('project.view.mode', 'list')

const searchQuery = ref('')

// 按服务器分组项目
const projectGroups = computed(() => {
  const groups = new Map()

  projects.value.forEach((project: Project) => {
    const server = servers.value.find((s: any) => s.id === project.serverId)
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

// 过滤项目组
const filteredProjectGroups = computed(() => {
  if (!searchQuery.value.trim()) {
    return projectGroups.value
  }

  const query = searchQuery.value.toLowerCase()

  return projectGroups.value
    .map((group) => ({
      ...group,
      projects: group.projects.filter(
        (project: any) =>
          project.name.toLowerCase().includes(query) ||
          project.path.toLowerCase().includes(query) ||
          project.dataPath?.toLowerCase().includes(query) ||
          group.serverName.toLowerCase().includes(query),
      ),
    }))
    .filter((group) => group.projects.length > 0)
})

const handleSave = async (project: Project) => {
  if (!serverStore.currentServerId) return

  if (dialog.state.mode === 'create') {
    await centerStore.createProject(project)
  } else if (dialog.state.editingData) {
    await projectStore.updateProject(dialog.state.editingData.id, { ...project })
  }
  dialog.close()
}
</script>
