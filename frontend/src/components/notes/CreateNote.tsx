import { useState } from 'react'
import { Card, CardContent } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Plus, Check, X } from 'lucide-react'
import type { CreateNoteRequest } from '@/types/note'

interface CreateNoteProps {
  onCreateNote: (noteData: CreateNoteRequest) => void
  isLoading?: boolean
}

export function CreateNote({ onCreateNote, isLoading = false }: CreateNoteProps) {
  const [isExpanded, setIsExpanded] = useState(false)
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')

  const handleSubmit = () => {
    if (!title.trim() && !content.trim()) return

    onCreateNote({
      title: title.trim(),
      content: content.trim(),
    })

    // Reset form
    setTitle('')
    setContent('')
    setIsExpanded(false)
  }

  const handleCancel = () => {
    setTitle('')
    setContent('')
    setIsExpanded(false)
  }

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === 'Escape') {
      handleCancel()
    } else if (e.key === 'Enter' && e.ctrlKey) {
      handleSubmit()
    }
  }

  if (!isExpanded) {
    return (
      <Card 
        className="cursor-pointer hover:shadow-md transition-shadow duration-200 mb-8"
        onClick={() => setIsExpanded(true)}
      >
        <CardContent className="p-4">
          <div className="flex items-center gap-3 text-gray-500">
            <Plus className="h-5 w-5" />
            <span>Take a note...</span>
          </div>
        </CardContent>
      </Card>
    )
  }

  return (
    <Card className="mb-8 shadow-lg">
      <CardContent className="p-4">
        <div className="space-y-3">
          <input
            type="text"
            placeholder="Title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            onKeyDown={handleKeyDown}
            className="w-full text-lg font-medium placeholder-gray-400 bg-transparent border-none outline-none resize-none"
            autoFocus
          />
          
          <textarea
            placeholder="Take a note..."
            value={content}
            onChange={(e) => setContent(e.target.value)}
            onKeyDown={handleKeyDown}
            className="w-full placeholder-gray-400 bg-transparent border-none outline-none resize-none min-h-[100px]"
            rows={4}
          />
          
          <div className="flex items-center justify-end gap-2 pt-2">
            <Button
              variant="ghost"
              size="sm"
              onClick={handleCancel}
              disabled={isLoading}
            >
              <X className="h-4 w-4" />
            </Button>
            <Button
              variant="default"
              size="sm"
              onClick={handleSubmit}
              disabled={isLoading || (!title.trim() && !content.trim())}
            >
              <Check className="h-4 w-4" />
            </Button>
          </div>
        </div>
      </CardContent>
    </Card>
  )
}