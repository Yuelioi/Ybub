<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div class="title-deco">
        <h2>服务器连接</h2>
      </div>
      <button class="btn btn-primary" @click="dialog.openCreate()">
        <PlusIcon class="size-4" />
        添加服务器
      </button>
    </div>

    <!-- 搜索和视图切换 -->
    <div class="flex items-center gap-3">
      <div class="relative flex-1">
        <SearchIcon class="absolute left-3 top-1/2 -translate-y-1/2 size-4 text-muted-foreground" />
        <input
          v-model="searchQuery"
          type="text"
          placeholder="搜索服务器名称、主机地址或用户名..."
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
    <div v-if="serverStore.servers.length === 0" class="card p-8 text-center">
      <ServerIcon class="size-12 text-muted-foreground mx-auto mb-4" />
      <p class="text-muted-foreground mb-4">还没有添加服务器</p>
      <button class="btn btn-sm btn-primary" @click="dialog.openCreate()">
        <PlusIcon class="size-4" />
        添加第一个服务器
      </button>
    </div>

    <!-- 搜索无结果 -->
    <div v-else-if="filteredServers.length === 0" class="card p-8 text-center">
      <SearchIcon class="size-12 text-muted-foreground mx-auto mb-4" />
      <p class="text-muted-foreground mb-2">未找到匹配的服务器</p>
      <p class="text-sm text-muted-foreground">尝试使用不同的关键词搜索</p>
    </div>

    <!-- 网格视图 -->
    <div
      v-else-if="viewMode === 'grid'"
      class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <ServerCard
        v-for="server in filteredServers"
        :key="server.id"
        :server="server"
        :project-count="getServerProjectsCount(server.id)"
        @edit="dialog.openUpdate(server)"
        @remove="handleDelete(server.id)"
        @connect="serverStore.connectServer(server.id)"
        @disconnect="serverStore.disconnectServer(server.id)" />
    </div>

    <!-- 列表视图 -->
    <div v-else class="card divide-y divide-border">
      <ServerListItem
        v-for="server in filteredServers"
        :key="server.id"
        :server="server"
        :project-count="getServerProjectsCount(server.id)"
        @edit="dialog.openUpdate(server)"
        @remove="handleDelete(server.id)"
        @connect="serverStore.connectServer(server.id)"
        @disconnect="serverStore.disconnectServer(server.id)" />
    </div>
  </div>

  <!-- 对话框 -->
  <ServerDialog
    :open="dialog.state.open"
    :mode="dialog.state.mode"
    :initial-data="dialog.initialData.value"
    @close="dialog.close()"
    @save="handleSave" />

  <YamiDialog
    :title="'确定要删除这个服务器吗？'"
    v-model="showDeleteDialog"
    @confirm="centerStore.removeServer(selectId)">
    <p>其下的 Docker 项目也会被<span class="text-red-400">删除</span></p>
  </YamiDialog>
</template>

<script setup lang="ts">
import {
  Server as ServerIcon,
  Plus as PlusIcon,
  Search as SearchIcon,
  X as XIcon,
  LayoutGrid as LayoutGridIcon,
  LayoutList as LayoutListIcon,
} from 'lucide-vue-next'
import { ref, computed } from 'vue'
import type { Server } from '@models'
import { useDialog } from '@/composables/useDialog'
import ServerCard from '@/components/server/ServerCard.vue'
import ServerListItem from '@/components/server/ServerListItem.vue'
import ServerDialog from '@/components/server/ServerDialog.vue'
import { YamiDialog } from '@yuelioi/ui'

const serverStore = useServerStore()
const projectStore = useProjectStore()
const centerStore = useCenterStore()

const dialog = useDialog()

const showDeleteDialog = ref(false)
const selectId = ref('')

// 视图模式和搜索
const viewMode = useStorage<'grid' | 'list'>('server.view.mode', 'grid')
const searchQuery = ref('')

// 过滤服务器
const filteredServers = computed(() => {
  if (!searchQuery.value.trim()) {
    return serverStore.servers
  }

  const query = searchQuery.value.toLowerCase()
  return serverStore.servers.filter(
    (server) =>
      server.name.toLowerCase().includes(query) ||
      server.host.toLowerCase().includes(query) ||
      server.username.toLowerCase().includes(query),
  )
})

const getServerProjectsCount = (serverId: string) => {
  return projectStore.getProjectsByServerId(serverId).length
}

const handleSave = async (data: Server) => {
  if (dialog.state.mode === 'create') {
    await centerStore.createServer(data)
  } else if (dialog.state.editingData) {
    await serverStore.updateServer(dialog.state.editingData.id, { ...data })
  }
  dialog.close()
}

const handleDelete = async (id: string) => {
  showDeleteDialog.value = true
  selectId.value = id
}
</script>
