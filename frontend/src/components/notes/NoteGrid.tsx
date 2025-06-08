import type { Note } from '@/types/note'
import { NoteCard } from './NoteCard'

interface NoteGridProps {
  notes: Note[]
  onNoteUpdate: (note: Note) => void
  onNoteDelete: (id: string) => void
  onNoteClick: (note: Note) => void
}

export function NoteGrid({ notes, onNoteUpdate, onNoteDelete, onNoteClick }: NoteGridProps) {
  // Separate pinned and unpinned notes
  const pinnedNotes = notes.filter(note => note.is_pinned && !note.is_archived && !note.is_deleted)
  const unpinnedNotes = notes.filter(note => !note.is_pinned && !note.is_archived && !note.is_deleted)

  if (notes.length === 0) {
    return (
      <div className="flex items-center justify-center h-64">
        <p className="text-gray-500 text-lg">No notes yet. Create your first note!</p>
      </div>
    )
  }

  return (
    <div className="space-y-8">
      {pinnedNotes.length > 0 && (
        <section>
          <h2 className="text-sm font-medium text-gray-600 mb-4 uppercase tracking-wide">
            Pinned
          </h2>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
            {pinnedNotes.map((note) => (
              <NoteCard
                key={note.id}
                note={note}
                onUpdate={onNoteUpdate}
                onDelete={onNoteDelete}
                onClick={() => onNoteClick(note)}
              />
            ))}
          </div>
        </section>
      )}

      {unpinnedNotes.length > 0 && (
        <section>
          {pinnedNotes.length > 0 && (
            <h2 className="text-sm font-medium text-gray-600 mb-4 uppercase tracking-wide">
              Others
            </h2>
          )}
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
            {unpinnedNotes.map((note) => (
              <NoteCard
                key={note.id}
                note={note}
                onUpdate={onNoteUpdate}
                onDelete={onNoteDelete}
                onClick={() => onNoteClick(note)}
              />
            ))}
          </div>
        </section>
      )}
    </div>
  )
}