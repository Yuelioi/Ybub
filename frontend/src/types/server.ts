import { Server } from '@models'

export type ServerStatus = 'connected' | 'connecting' | 'disconnected' | 'error'

export type ServerFormData = Omit<Server, 'id' | 'status'>

export type ServerViewMode = 'grid' | 'list'
