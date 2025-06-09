import type { Note } from '@/types/note'
import { NoteCard } from './NoteCard'
import { EmptyState } from './EmptyState'

interface NoteGridProps {
  notes: Note[]
  onNoteUpdate: (note: Note) => void
  onNoteDelete: (id: string) => void
  onNoteClick: (note: Note) => void
}

export function NoteGrid({ notes, onNoteUpdate, onNoteDelete, onNoteClick }: NoteGridProps) {
  // Filter notes
  const activeNotes = notes.filter(note => !note.is_archived && !note.is_deleted)
  const pinnedNotes = activeNotes.filter(note => note.is_pinned)
  const unpinnedNotes = activeNotes.filter(note => !note.is_pinned)

  if (activeNotes.length === 0) {
    return <EmptyState />
  }

  return (
    <div className="space-y-8">
      {/* Pinned Notes Section */}
      {pinnedNotes.length > 0 && (
        <section>
          <h2 className="text-xs font-medium text-gray-600 mb-4 uppercase tracking-wide pl-1">
            Pinned
          </h2>
          <div className="columns-1 sm:columns-2 lg:columns-3 xl:columns-4 2xl:columns-5 gap-4">
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

      {/* Unpinned Notes Section */}
      {unpinnedNotes.length > 0 && (
        <section>
          {pinnedNotes.length > 0 && (
            <h2 className="text-xs font-medium text-gray-600 mb-4 uppercase tracking-wide pl-1">
              Others
            </h2>
          )}
          <div className="columns-1 sm:columns-2 lg:columns-3 xl:columns-4 2xl:columns-5 gap-4">
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