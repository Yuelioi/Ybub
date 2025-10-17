<template>
  <div class="">
    <TheHeader></TheHeader>

    <div class="flex h-[calc(100vh-64px)]">
      <SideBar />

      <main class="flex-1 overflow-auto p-6">
        <ServerManagementPage v-if="appStore.activeTab === 'servers'" />

        <ProjectManagerPage v-if="appStore.activeTab === 'projects'" />

        <BackupManager v-if="appStore.activeTab === 'backup'" />

        <SshTerminal v-if="appStore.activeTab === 'terminal'" />
      </main>
    </div>

    <CommandOutputDialog
      :open="dockerStore.showCommandOutput"
      :command-output="dockerStore.commandOutput"
      @close="dockerStore.closeCommandOutput" />
  </div>

  <ToastContainer></ToastContainer>
</template>

<script setup lang="ts">
import { ToastContainer } from '@yuelioi/toast'

// 使用 stores
const appStore = useAppStore()
const serverStore = useServerStore()
const projectStore = useProjectStore()
const terminalStore = useTerminalStore()
const dockerStore = useDockerStore()
const backupStore = useBackupStore()

// 表单数据

// 初始化
onMounted(async () => {
  await serverStore.init()
  await projectStore.init()
  await dockerStore.init()
  await backupStore.init()
  terminalStore.init()
})
</script>
