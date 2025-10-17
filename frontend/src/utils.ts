import { formatDate } from '@yuelioi/utils'
import cronstrue from 'cronstrue'
import 'cronstrue/locales/zh_CN'

export { formatDate }

export function randomID(): string {
  const charset = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'
  let result = ''
  for (let i = 0; i < 6; i++) {
    const randomIndex = Math.floor(Math.random() * charset.length)
    result += charset[randomIndex]
  }
  return result
}

export function formatCron(expr: string): string {
  const text = cronstrue.toString(expr, { locale: 'zh_CN' })
  return text
}
