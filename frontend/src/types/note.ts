export interface Note {
  id: string
  user_id: string
  title: string
  content: string
  color: string
  is_pinned: boolean
  is_archived: boolean
  is_deleted: boolean
  position: number
  created_at: string
  updated_at: string
  labels?: Label[]
  attachments?: Attachment[]
}

export interface Label {
  id: string
  user_id: string
  name: string
  color: string
  created_at: string
  updated_at: string
}

export interface Attachment {
  id: string
  note_id: string
  filename: string
  url: string
  size: number
  mime_type: string
  created_at: string
}

export interface CreateNoteRequest {
  title?: string
  content?: string
  color?: string
  is_pinned?: boolean
}

export interface UpdateNoteRequest {
  title?: string
  content?: string
  color?: string
  is_pinned?: boolean
  is_archived?: boolean
  position?: number
}

export interface NotesState {
  notes: Note[]
  selectedNote: Note | null
  isLoading: boolean
  error: string | null
  searchQuery: string
  filter: 'all' | 'pinned' | 'archived'
}