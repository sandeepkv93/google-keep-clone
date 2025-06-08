import { useState } from 'react'
import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Pin, Archive, Trash2, MoreVertical, Palette } from 'lucide-react'
import { Note } from '@/types/note'
import { cn } from '@/lib/utils'

interface NoteCardProps {
  note: Note
  onUpdate: (note: Note) => void
  onDelete: (id: string) => void
  onClick: () => void
}

export function NoteCard({ note, onUpdate, onDelete, onClick }: NoteCardProps) {
  const [isHovered, setIsHovered] = useState(false)

  const handlePin = (e: React.MouseEvent) => {
    e.stopPropagation()
    onUpdate({ ...note, is_pinned: !note.is_pinned })
  }

  const handleArchive = (e: React.MouseEvent) => {
    e.stopPropagation()
    onUpdate({ ...note, is_archived: !note.is_archived })
  }

  const handleDelete = (e: React.MouseEvent) => {
    e.stopPropagation()
    onDelete(note.id)
  }

  return (
    <Card
      className={cn(
        'cursor-pointer transition-all duration-200 hover:shadow-md',
        note.is_pinned && 'ring-2 ring-yellow-400',
        isHovered && 'shadow-lg'
      )}
      style={{ backgroundColor: note.color }}
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
      onClick={onClick}
    >
      <CardContent className='p-4'>
        <div className='flex justify-between items-start mb-2'>
          {note.title && (
            <h3 className='font-medium text-sm text-gray-800 line-clamp-2'>
              {note.title}
            </h3>
          )}
          <Button
            variant='ghost'
            size='sm'
            className={cn(
              'h-6 w-6 p-0 opacity-0 transition-opacity',
              (isHovered || note.is_pinned) && 'opacity-100'
            )}
            onClick={handlePin}
          >
            <Pin className={cn('h-4 w-4', note.is_pinned && 'fill-current')} />
          </Button>
        </div>

        {note.content && (
          <p className='text-sm text-gray-700 line-clamp-4 whitespace-pre-wrap'>
            {note.content}
          </p>
        )}

        {note.labels && note.labels.length > 0 && (
          <div className='flex flex-wrap gap-1 mt-2'>
            {note.labels.map((label) => (
              <span
                key={label.id}
                className='px-2 py-1 text-xs rounded-full bg-gray-200 text-gray-700'
              >
                {label.name}
              </span>
            ))}
          </div>
        )}

        <div
          className={cn(
            'flex justify-between items-center mt-3 opacity-0 transition-opacity',
            isHovered && 'opacity-100'
          )}
        >
          <div className='flex gap-1'>
            <Button variant='ghost' size='sm' className='h-6 w-6 p-0'>
              <Palette className='h-4 w-4' />
            </Button>
          </div>

          <div className='flex gap-1'>
            <Button
              variant='ghost'
              size='sm'
              className='h-6 w-6 p-0'
              onClick={handleArchive}
            >
              <Archive className='h-4 w-4' />
            </Button>
            <Button
              variant='ghost'
              size='sm'
              className='h-6 w-6 p-0'
              onClick={handleDelete}
            >
              <Trash2 className='h-4 w-4' />
            </Button>
            <Button variant='ghost' size='sm' className='h-6 w-6 p-0'>
              <MoreVertical className='h-4 w-4' />
            </Button>
          </div>
        </div>
      </CardContent>
    </Card>
  )
}