import type { Note, CreateNoteRequest, UpdateNoteRequest } from '@/types/note'
import { authService } from './auth'

const API_BASE = import.meta.env.VITE_API_URL || 'http://localhost:8080'

export class NotesService {
  private getAuthHeaders(): Record<string, string> {
    const token = authService.getToken()
    return {
      'Content-Type': 'application/json',
      ...(token && { 'Authorization': `Bearer ${token}` }),
    }
  }

  async getNotes(includeArchived = false, includeDeleted = false): Promise<Note[]> {
    const params = new URLSearchParams()
    if (includeArchived) params.set('archived', 'true')
    if (includeDeleted) params.set('deleted', 'true')

    const response = await fetch(`${API_BASE}/notes?${params}`, {
      headers: this.getAuthHeaders(),
    })

    if (!response.ok) {
      throw new Error('Failed to fetch notes')
    }

    return response.json()
  }

  async createNote(noteData: CreateNoteRequest): Promise<Note> {
    const response = await fetch(`${API_BASE}/notes`, {
      method: 'POST',
      headers: this.getAuthHeaders(),
      body: JSON.stringify(noteData),
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.error || 'Failed to create note')
    }

    return response.json()
  }

  async getNoteById(id: string): Promise<Note> {
    const response = await fetch(`${API_BASE}/notes/${id}`, {
      headers: this.getAuthHeaders(),
    })

    if (!response.ok) {
      throw new Error('Failed to fetch note')
    }

    return response.json()
  }

  async updateNote(id: string, noteData: UpdateNoteRequest): Promise<Note> {
    const response = await fetch(`${API_BASE}/notes/${id}`, {
      method: 'PUT',
      headers: this.getAuthHeaders(),
      body: JSON.stringify(noteData),
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.error || 'Failed to update note')
    }

    return response.json()
  }

  async deleteNote(id: string, permanent = false): Promise<void> {
    const params = permanent ? '?permanent=true' : ''
    const response = await fetch(`${API_BASE}/notes/${id}${params}`, {
      method: 'DELETE',
      headers: this.getAuthHeaders(),
    })

    if (!response.ok) {
      throw new Error('Failed to delete note')
    }
  }

  async togglePin(id: string): Promise<Note> {
    const response = await fetch(`${API_BASE}/notes/${id}/pin`, {
      method: 'PATCH',
      headers: this.getAuthHeaders(),
    })

    if (!response.ok) {
      throw new Error('Failed to toggle pin')
    }

    return response.json()
  }

  async toggleArchive(id: string): Promise<Note> {
    const response = await fetch(`${API_BASE}/notes/${id}/archive`, {
      method: 'PATCH',
      headers: this.getAuthHeaders(),
    })

    if (!response.ok) {
      throw new Error('Failed to toggle archive')
    }

    return response.json()
  }

  async updateColor(id: string, color: string): Promise<Note> {
    const response = await fetch(`${API_BASE}/notes/${id}/color`, {
      method: 'PATCH',
      headers: this.getAuthHeaders(),
      body: JSON.stringify({ color }),
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.error || 'Failed to update color')
    }

    return response.json()
  }

  async searchNotes(query: string, limit = 20, page = 0): Promise<Note[]> {
    const params = new URLSearchParams({
      q: query,
      limit: limit.toString(),
      page: page.toString(),
    })

    const response = await fetch(`${API_BASE}/notes/search?${params}`, {
      headers: this.getAuthHeaders(),
    })

    if (!response.ok) {
      throw new Error('Failed to search notes')
    }

    return response.json()
  }

  async getPinnedNotes(): Promise<Note[]> {
    const response = await fetch(`${API_BASE}/notes/pinned`, {
      headers: this.getAuthHeaders(),
    })

    if (!response.ok) {
      throw new Error('Failed to fetch pinned notes')
    }

    return response.json()
  }

  async getArchivedNotes(): Promise<Note[]> {
    const response = await fetch(`${API_BASE}/notes/archived`, {
      headers: this.getAuthHeaders(),
    })

    if (!response.ok) {
      throw new Error('Failed to fetch archived notes')
    }

    return response.json()
  }
}

export const notesService = new NotesService()