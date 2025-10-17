<template>
  <dialog :open="open" class="dialog">
    <div class="dialog-body dialog-body-lg">
      <div class="dialog-header">
        <h3 v-if="mode == 'create'" class="dialog-title">添加项目到 {{ serverName }}</h3>
        <h3 v-if="mode == 'update'" class="dialog-title">
          修改项目 (<span class="text-primary"> {{ serverName }} </span>)
        </h3>
        <button class="dialog-close" @click="$emit('close')">
          <XIcon class="size-5" />
        </button>
      </div>

      <div class="dialog-content space-y-4 p-2">
        <div>
          <label class="label mb-2 block">项目名称</label>
          <input v-model="project.name" class="input w-full" placeholder="我的 Web 应用" />
        </div>
        <div>
          <label class="label mb-2 block">项目路径 (docker-compose.yml 所在目录)</label>
          <input v-model="project.path" class="input w-full" placeholder="/home/user/my-web-app" />
          <p class="text-xs text-muted-foreground mt-1">Docker 命令会在此目录下执行</p>
        </div>
        <div>
          <label class="label mb-2 block">数据目录 (可选，用于备份)</label>
          <input
            v-model="project.dataPath"
            class="input w-full"
            placeholder="/home/user/my-web-app/data" />
          <p class="text-xs text-muted-foreground mt-1">通常是 docker-compose 中挂载的数据卷目录</p>
        </div>
        <div class="border-t border-border pt-4">
          <div class="flex items-center gap-2 mb-3">
            <input
              v-model="project.enableBackup"
              type="checkbox"
              class="checkbox"
              id="autoBackup" />
            <label for="autoBackup" class="label">启用自动备份</label>
          </div>
          <div v-if="project.enableBackup" class="pl-6 space-y-3">
            <div>
              <label class="label mb-2 block">备份频率</label>
              <select v-model="project.schedule" class="select w-full">
                <option value="0 2 * * *">每天凌晨 2:00</option>
                <option value="0 2 * * 0">每周日凌晨 2:00</option>
                <option value="0 2 1 * *">每月1日凌晨 2:00</option>
                <option value="custom">自定义</option>
              </select>
              <input
                v-if="isCustomSchedule"
                v-model="customSchedule"
                type="text"
                class="input w-full mt-2"
                placeholder="请输入 Cron 表达式，例如: 0 3 * * *" />
            </div>
          </div>
        </div>
      </div>

      <div class="dialog-footer">
        <button class="btn btn-outline" @click="$emit('close')">取消</button>
        <button class="btn btn-primary" @click="save" :disabled="!project.name || !project.path">
          确认
        </button>
      </div>
    </div>
  </dialog>
</template>

<script setup lang="ts">
import { ProjectStatus, type Project } from '@models'
import { X as XIcon } from 'lucide-vue-next'

const props = defineProps<{
  open: boolean
  mode: 'create' | 'update'
  initialData: Project | null
  serverName?: string
}>()

const customSchedule = ref('')

const createDefaultProject = (): Project => ({
  id: randomID(),
  serverId: '',
  name: '',
  path: '',
  status: ProjectStatus.ProjectBackupStatusStopped,
  enableBackup: false,
  schedule: '0 0 2 * * *',
  backupStatus: ProjectStatus.ProjectStatusStopped,
})

const project = reactive<Project>(createDefaultProject())

const isCustomSchedule = computed(() => {
  return project.schedule == 'custom'
})
watch(
  () => props.open,
  (isOpen) => {
    if (isOpen) {
      if (props.mode === 'update') {
        Object.assign(project, props.initialData)
      } else {
        Object.assign(project, createDefaultProject())
      }
    }
  },
)

const emit = defineEmits<{
  close: []
  save: [project: Project]
}>()

const save = () => {
  if (isCustomSchedule.value) {
    project.schedule = customSchedule.value
  }

  if (project.enableBackup) {
    project.backupStatus = ProjectStatus.ProjectBackupStatusRunning
  }

  emit('save', project)
}
</script>
