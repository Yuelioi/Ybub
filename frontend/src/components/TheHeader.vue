<template>
  <header class="sticky top-0 z-50 bg-card/50 backdrop-blur-lg h-16 border-b border-border">
    <div class="flex items-center justify-between px-4 h-full">
      <div class="flex items-center gap-2">
        <div
          class="w-8 h-8 rounded-lg bg-gradient-to-br from-primary to-chart-2 flex items-center justify-center">
          <span class="text-white font-bold text-sm">Y</span>
        </div>
        <h1 class="text-lg font-bold pointer-events-none select-none">YBub</h1>
      </div>

      <div class="flex-1 h-full" style="--wails-draggable: drag"></div>

      <div class="flex items-center">
        <ThemeToggle class="btn btn-icon-sm btn-ghost"></ThemeToggle>

        <button class="btn btn-icon-sm btn-ghost">
          <Settings class="size-5"></Settings>
        </button>

        <div class="border mx-2 h-5"></div>

        <!-- 置顶按钮 -->
        <button
          class="btn btn-icon-sm btn-ghost"
          @click="toggleAlwaysOnTop"
          :title="isAlwaysOnTop ? '取消置顶' : '窗口置顶'"
          :class="{ 'bg-accent': isAlwaysOnTop }">
          <PinIcon v-if="!isAlwaysOnTop" class="size-4"></PinIcon>

          <PinOff v-else class="size-4"></PinOff>
        </button>

        <!-- 最小化按钮 -->
        <button class="btn btn-icon-sm btn-ghost" @click="minimizeWindow" title="最小化">
          <svg class="size-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 12h12"></path>
          </svg>
        </button>

        <!-- 最大化/还原按钮 -->
        <button
          class="btn btn-icon-sm btn-ghost"
          @click="toggleMaximize"
          :title="isMaximized ? '还原' : '最大化'">
          <!-- 最大化图标 -->
          <MaximizeIcon v-if="!isMaximized" class="size-4"></MaximizeIcon>
          <!-- 还原图标  -->

          <svg
            v-else
            class="size-4"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24">
            <path
              fill="currentColor"
              d="M16 20v-9H4v9zm2-5v-2h2V4H8v5H6V4q0-.825.588-1.412T8 2h12q.825 0 1.413.588T22 4v9q0 .825-.587 1.413T20 15zM4 22q-.825 0-1.412-.587T2 20v-9q0-.825.588-1.412T4 9h12q.825 0 1.413.588T18 11v9q0 .825-.587 1.413T16 22zm6-6.5" />
          </svg>
        </button>

        <!-- 关闭按钮 -->
        <button
          class="btn btn-icon-sm hover:bg-red-500 hover:text-white"
          @click="closeWindow"
          title="最小化到托盘">
          <XIcon class="size-4"></XIcon>
        </button>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { ThemeToggle } from '@yuelioi/ui'

import { XIcon, PinIcon, PinOff, MaximizeIcon, Settings } from 'lucide-vue-next'

import { Window } from '@wailsio/runtime'

const isMaximized = ref(false)
const isAlwaysOnTop = ref(false)

// 最小化窗口
const minimizeWindow = async () => {
  try {
    await Window.Minimise()
  } catch (error) {
    console.error('Failed to minimize window:', error)
  }
}

// 切换最大化/还原
const toggleMaximize = async () => {
  try {
    if (isMaximized.value) {
      isMaximized.value = false
      await Window.ToggleMaximise()
    } else {
      isMaximized.value = true
      await Window.ToggleMaximise()
    }
  } catch (error) {
    console.error('Failed to toggle maximize:', error)
  }
}

// 切换窗口置顶
const toggleAlwaysOnTop = async () => {
  try {
    isAlwaysOnTop.value = !isAlwaysOnTop.value
    await Window.SetAlwaysOnTop(isAlwaysOnTop.value)
  } catch (error) {
    console.error('Failed to toggle always on top:', error)
    // 如果失败，回滚状态
    isAlwaysOnTop.value = !isAlwaysOnTop.value
  }
}

// 关闭窗口
const closeWindow = async () => {
  try {
    await Window.Hide()
  } catch (error) {
    console.error('Failed to close window:', error)
  }
}

// 检查初始窗口状态
const checkWindowState = async () => {
  try {
    isMaximized.value = await Window.IsMaximised()
    // 如果有获取置顶状态的 API，在这里调用
    // isAlwaysOnTop.value = await Window.IsAlwaysOnTop();
  } catch (error) {
    console.error('Failed to check window state:', error)
  }
}

// 监听窗口状态变化
onMounted(() => {
  checkWindowState()
})
</script>

<style scoped>
/* 确保按钮区域不会被拖拽影响 */
.btn btn-icon-sm {
  --wails-draggable: no-drag;
}
</style>
