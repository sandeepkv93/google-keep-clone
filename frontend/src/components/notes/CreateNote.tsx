import { useState } from 'react'
import type { CreateNoteRequest } from '@/types/note'

interface CreateNoteProps {
  onCreateNote: (noteData: CreateNoteRequest) => void
  isLoading?: boolean
}

export function CreateNote({ onCreateNote, isLoading = false }: CreateNoteProps) {
  const [isExpanded, setIsExpanded] = useState(false)
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')
  const [showColorPicker, setShowColorPicker] = useState(false)
  const [selectedColor, setSelectedColor] = useState('#ffffff')

  const colors = [
    '#ffffff', '#f28b82', '#fbbc04', '#fff475', '#ccff90', '#a7ffeb',
    '#cbf0f8', '#aecbfa', '#d7aefb', '#fdcfe8', '#e6c9a8', '#e8eaed'
  ]

  const handleSubmit = () => {
    if (!title.trim() && !content.trim()) return

    onCreateNote({
      title: title.trim(),
      content: content.trim(),
    })

    // Reset form
    setTitle('')
    setContent('')
    setSelectedColor('#ffffff')
    setIsExpanded(false)
  }

  const handleCancel = () => {
    setTitle('')
    setContent('')
    setSelectedColor('#ffffff')
    setIsExpanded(false)
    setShowColorPicker(false)
  }

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === 'Escape') {
      handleCancel()
    }
  }

  // Collapsed state - "Take a note..." bar
  if (!isExpanded) {
    return (
      <div className="max-w-2xl mx-auto mb-8">
        <div 
          className="bg-white rounded-lg border border-gray-300 shadow-sm hover:shadow-md transition-shadow duration-200 cursor-text p-4"
          onClick={() => setIsExpanded(true)}
        >
          <div className="flex items-center gap-4">
            <span className="text-base text-gray-500">Take a note...</span>
            <div className="ml-auto flex items-center gap-3">
              <button className="p-2 rounded-full hover:bg-gray-100 transition-colors">
                <svg className="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />
                </svg>
              </button>
              <button className="p-2 rounded-full hover:bg-gray-100 transition-colors">
                <svg className="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    )
  }

  // Expanded state - Full note creation form
  return (
    <div className="max-w-2xl mx-auto mb-8 relative">
      <div 
        className="bg-white rounded-lg border border-gray-300 shadow-lg relative"
        style={{ backgroundColor: selectedColor }}
      >
        {/* Title Input */}
        <div className="p-4 pb-2">
          <input
            type="text"
            placeholder="Title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            onKeyDown={handleKeyDown}
            className="w-full text-base font-medium placeholder-gray-500 bg-transparent border-none outline-none"
            autoFocus
          />
        </div>

        {/* Content Input */}
        <div className="px-4 pb-4">
          <textarea
            placeholder="Take a note..."
            value={content}
            onChange={(e) => setContent(e.target.value)}
            onKeyDown={handleKeyDown}
            className="w-full placeholder-gray-500 bg-transparent border-none outline-none resize-none min-h-[80px] text-sm leading-relaxed"
            rows={4}
          />
        </div>

        {/* Toolbar */}
        <div className="flex items-center justify-between p-3 pt-0">
          <div className="flex items-center gap-1">
            {/* Remind me */}
            <button className="p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors" title="Remind me">
              <svg className="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </button>

            {/* Collaborator */}
            <button className="p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors" title="Collaborator">
              <svg className="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
              </svg>
            </button>

            {/* Color Palette */}
            <div className="relative">
              <button 
                className="p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors"
                onClick={() => setShowColorPicker(!showColorPicker)}
                title="Background options"
              >
                <svg className="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zM21 5a2 2 0 00-2-2h-4a2 2 0 00-2 2v12a4 4 0 004 4h4a2 2 0 002-2V5z" />
                </svg>
              </button>

              {/* Color Picker Dropdown */}
              {showColorPicker && (
                <div className="absolute bottom-full left-0 mb-2 bg-white rounded-lg shadow-lg border border-gray-200 p-2 z-10">
                  <div className="grid grid-cols-4 gap-2">
                    {colors.map((color) => (
                      <button
                        key={color}
                        className="w-8 h-8 rounded-full border-2 border-gray-300 hover:border-gray-500 transition-colors relative"
                        style={{ backgroundColor: color }}
                        onClick={() => {
                          setSelectedColor(color)
                          setShowColorPicker(false)
                        }}
                        title={color === '#ffffff' ? 'Default' : color}
                      >
                        {selectedColor === color && (
                          <svg className="w-4 h-4 absolute inset-0 m-auto text-gray-700" fill="currentColor" viewBox="0 0 24 24">
                            <path d="M9 16.2L4.8 12l-1.4 1.4L9 19 21 7l-1.4-1.4L9 16.2z"/>
                          </svg>
                        )}
                      </button>
                    ))}
                  </div>
                </div>
              )}
            </div>

            {/* Add image */}
            <button className="p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors" title="Add image">
              <svg className="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
              </svg>
            </button>

            {/* Archive */}
            <button className="p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors" title="Archive">
              <svg className="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 8l4 4 10-10M3 17l2 2 4-4" />
              </svg>
            </button>

            {/* More */}
            <button className="p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors" title="More">
              <svg className="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 5v.01M12 12v.01M12 19v.01M12 6a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2zm0 7a1 1 0 110-2 1 1 0 010 2z" />
              </svg>
            </button>

            {/* Undo */}
            <button className="p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors" title="Undo">
              <svg className="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 10h10a8 8 0 018 8v2M3 10l6 6m-6-6l6-6" />
              </svg>
            </button>

            {/* Redo */}
            <button className="p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors" title="Redo">
              <svg className="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 10h-10a8 8 0 00-8 8v2m18-10l-6 6m6-6l-6-6" />
              </svg>
            </button>
          </div>

          {/* Close Button */}
          <button
            onClick={handleCancel}
            disabled={isLoading}
            className="px-6 py-1 text-sm font-medium text-gray-600 hover:bg-black hover:bg-opacity-10 rounded transition-colors disabled:opacity-50"
          >
            Close
          </button>
        </div>
      </div>

      {/* Click outside to save/close */}
      {isExpanded && (
        <div 
          className="fixed inset-0 z-[-1]" 
          onClick={() => {
            if (title.trim() || content.trim()) {
              handleSubmit()
            } else {
              handleCancel()
            }
          }}
        />
      )}
    </div>
  )
}