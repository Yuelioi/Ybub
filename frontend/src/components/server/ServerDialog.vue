<template>
  <dialog :open="open" class="dialog">
    <div class="dialog-body dialog-body-lg">
      <div class="dialog-header">
        <h3 class="dialog-title">{{ title }}</h3>
        <button class="dialog-close" @click="$emit('close')">
          <XIcon class="size-5" />
        </button>
      </div>

      <div class="dialog-content space-y-4 p-2">
        <div>
          <label class="label mb-2 block">服务器名称</label>
          <input v-model="formData.name" class="input w-full" placeholder="我的 Ubuntu 服务器" />
        </div>
        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="label mb-2 block">主机地址</label>
            <input v-model="formData.host" class="input w-full" placeholder="192.168.1.100" />
          </div>
          <div>
            <label class="label mb-2 block">SSH 端口</label>
            <input
              v-model.number="formData.port"
              type="number"
              class="input w-full"
              placeholder="22" />
          </div>
        </div>
        <div>
          <label class="label mb-2 block">用户名</label>
          <input v-model="formData.username" class="input w-full" placeholder="root" />
        </div>
        <div>
          <label class="label mb-2 block">密码</label>
          <input
            v-model="formData.password"
            type="password"
            class="input w-full"
            placeholder="******" />
        </div>
        <div>
          <label class="label mb-2 block">密钥文件</label>
          <input v-model="formData.identityFile" type="text" class="input w-full" />
        </div>
      </div>

      <div class="dialog-footer">
        <button class="btn btn-outline" @click="$emit('close')">取消</button>
        <button class="btn btn-primary" @click="handleSave" :disabled="!isValid">确认</button>
      </div>
    </div>
  </dialog>
</template>

<script setup lang="ts">
import { X as XIcon } from 'lucide-vue-next'
import { reactive, computed, watch } from 'vue'
import { ServerStatus, type Server } from '@models'

const props = defineProps<{
  open: boolean
  mode: 'create' | 'update'
  initialData: Server | null
}>()

const emit = defineEmits<{
  close: []
  save: [data: Server]
}>()

const title = computed(() => (props.mode === 'create' ? '添加服务器' : '修改服务器'))

const createDefaultData = (): Server => ({
  id: randomID(),
  name: '',
  host: '',
  port: 22,
  username: '',
  password: '',
  identityFile: '',
  status: ServerStatus.ServerStatusDisconnected,
})

const formData = reactive<Server>(createDefaultData())

const isValid = computed(
  () => !!(formData.name && formData.host && formData.username && formData.port),
)

// 监听对话框打开，初始化表单
watch(
  () => props.open,
  (isOpen) => {
    if (isOpen) {
      if (props.mode === 'update') {
        Object.assign(formData, props.initialData)
      } else {
        Object.assign(formData, createDefaultData())
      }
    }
  },
)

const handleSave = () => {
  if (!isValid.value) return
  emit('save', { ...formData })
}
</script>
