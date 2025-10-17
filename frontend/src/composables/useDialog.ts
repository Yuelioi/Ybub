import { reactive, computed } from 'vue'

export function useDialog() {
  const state = reactive({
    open: false,
    mode: 'create' as 'create' | 'update',
    editingData: null as any,
  })

  // 计算初始表单数据
  const initialData = computed(() => {
    if (state.mode === 'update' && state.editingData) {
      return { ...state.editingData }
    }
    return null
  })

  const openCreate = () => {
    state.mode = 'create'
    state.editingData = null
    state.open = true
  }

  const openUpdate = (data: any) => {
    state.mode = 'update'
    state.editingData = data
    state.open = true
  }

  const close = () => {
    state.open = false
    setTimeout(() => {
      state.editingData = null
    }, 300)
  }

  return {
    state,
    initialData,
    openCreate,
    openUpdate,
    close,
  }
}
