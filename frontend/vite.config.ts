import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import tailwindcss from '@tailwindcss/vite'
import path from 'path'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'

import { UIResolver } from '@yuelioi/ui/resolver'

import { resolve } from 'path'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),
    tailwindcss(),
    AutoImport({
      imports: [
        'vue',
        'vue-router',
        'pinia',
        '@vueuse/core',
        {
          '@/utils': ['formatDate', 'randomID', 'formatCron'],
        },
        {
          '@yuelioi/toast': ['toast'],
        },
      ],
      dts: 'src/auto-imports.d.ts', // 自动生成类型声明
      dirs: ['./src/stores', './src/composables', './src/models/', './src/api/'],

      eslintrc: {
        enabled: true, // 生成 ESLint 配置，避免报 no-undef
        filepath: './.eslintrc-auto-import.json',
        globalsPropValue: true,
      },
    }),
    Components({
      dirs: ['src/components', 'src/views'], // 自动导入的组件目录
      extensions: ['vue'],
      deep: true, // 支持子目录
      dts: 'src/components.d.ts', // 生成类型声明文件
      resolvers: [UIResolver()],
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
      '@services': resolve(__dirname, 'bindings', 'ybub', 'services'),
      '@models': resolve(__dirname, 'bindings', 'ybub', 'models'),
    },
  },
  server: {
    watch: {
      ignored: ['**/pkg/**', '!**/stores/**'],
    },
    hmr: {
      host: 'localhost',
      protocol: 'ws',
    },
  },
})
