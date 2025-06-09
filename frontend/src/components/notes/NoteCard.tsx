import { useState } from 'react'
import type { Note } from '@/types/note'

interface NoteCardProps {
  note: Note
  onUpdate: (note: Note) => void
  onDelete: (id: string) => void
  onClick: () => void
}

export function NoteCard({ note, onUpdate, onDelete, onClick }: NoteCardProps) {
  const [isHovered, setIsHovered] = useState(false)
  const [showColorPicker, setShowColorPicker] = useState(false)
  const [showMoreMenu, setShowMoreMenu] = useState(false)

  const colors = [
    { name: 'Default', value: '#ffffff' },
    { name: 'Red', value: '#f28b82' },
    { name: 'Orange', value: '#fbbc04' },
    { name: 'Yellow', value: '#fff475' },
    { name: 'Green', value: '#ccff90' },
    { name: 'Teal', value: '#a7ffeb' },
    { name: 'Blue', value: '#cbf0f8' },
    { name: 'Dark Blue', value: '#aecbfa' },
    { name: 'Purple', value: '#d7aefb' },
    { name: 'Pink', value: '#fdcfe8' },
    { name: 'Brown', value: '#e6c9a8' },
    { name: 'Gray', value: '#e8eaed' }
  ]

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
    setShowMoreMenu(false)
    onDelete(note.id)
  }

  const handleColorChange = (e: React.MouseEvent, color: string) => {
    e.stopPropagation()
    onUpdate({ ...note, color })
    setShowColorPicker(false)
  }

  return (
    <div
      className={`relative bg-white rounded-lg border border-gray-300 cursor-pointer break-inside-avoid mb-4 transition-all duration-200 ${
        isHovered ? 'shadow-lg' : 'shadow-sm hover:shadow-md'
      }`}
      style={{ backgroundColor: note.color }}
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => {
        setIsHovered(false)
        setShowColorPicker(false)
        setShowMoreMenu(false)
      }}
      onClick={onClick}
    >
      {/* Pin Button */}
      <div className="absolute top-2 right-2 z-10">
        <button
          className={`p-2 rounded-full transition-all duration-200 ${
            isHovered || note.is_pinned ? 'opacity-100' : 'opacity-0'
          } hover:bg-black hover:bg-opacity-10`}
          onClick={handlePin}
          title={note.is_pinned ? 'Unpin note' : 'Pin note'}
        >
          <svg 
            className="w-4 h-4 text-gray-600" 
            fill={note.is_pinned ? 'currentColor' : 'none'}
            stroke="currentColor" 
            viewBox="0 0 24 24"
          >
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
          </svg>
        </button>
      </div>

      {/* Note Content */}
      <div className="p-4 pr-12">
        {note.title && (
          <h3 className="font-medium text-gray-900 mb-3 text-base leading-tight">
            {note.title}
          </h3>
        )}
        
        {note.content && (
          <div className="text-gray-700 text-sm leading-relaxed whitespace-pre-wrap">
            {note.content}
          </div>
        )}

        {note.labels && note.labels.length > 0 && (
          <div className="flex flex-wrap gap-1 mt-3">
            {note.labels.map((label) => (
              <span
                key={label.id}
                className="px-2 py-1 text-xs rounded-full bg-gray-200 text-gray-700"
              >
                {label.name}
              </span>
            ))}
          </div>
        )}
      </div>

      {/* Hover Toolbar */}
      <div className={`absolute bottom-0 left-0 right-0 bg-white bg-opacity-95 rounded-b-lg transition-all duration-200 ${
        isHovered ? 'opacity-100 visible' : 'opacity-0 invisible'
      }`}>
        <div className="flex items-center justify-between p-2">
          <div className="flex items-center gap-1">
            {/* Remind me */}
            <button
              className="p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors"
              title="Remind me"
              onClick={(e) => e.stopPropagation()}
            >
              <svg className="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </button>

            {/* Collaborator */}
            <button
              className="p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors"
              title="Collaborator"
              onClick={(e) => e.stopPropagation()}
            >
              <svg className="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
              </svg>
            </button>

            {/* Color Palette */}
            <div className="relative">
              <button
                className="p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors"
                onClick={(e) => {
                  e.stopPropagation()
                  setShowColorPicker(!showColorPicker)
                  setShowMoreMenu(false)
                }}
                title="Background options"
              >
                <svg className="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zM21 5a2 2 0 00-2-2h-4a2 2 0 00-2 2v12a4 4 0 004 4h4a2 2 0 002-2V5z" />
                </svg>
              </button>
              
              {/* Color Picker Dropdown */}
              {showColorPicker && (
                <div className="absolute bottom-full left-0 mb-2 bg-white rounded-lg shadow-lg border border-gray-200 p-2 z-20">
                  <div className="grid grid-cols-4 gap-2">
                    {colors.map((color) => (
                      <button
                        key={color.value}
                        className="w-8 h-8 rounded-full border-2 border-gray-300 hover:border-gray-500 transition-colors relative"
                        style={{ backgroundColor: color.value }}
                        onClick={(e) => handleColorChange(e, color.value)}
                        title={color.name}
                      >
                        {note.color === color.value && (
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
            <button
              className="p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors"
              title="Add image"
              onClick={(e) => e.stopPropagation()}
            >
              <svg className="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
              </svg>
            </button>

            {/* Archive */}
            <button
              onClick={handleArchive}
              className="p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors"
              title="Archive"
            >
              <svg className="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 8l4 4 10-10M3 17l2 2 4-4" />
              </svg>
            </button>
          </div>

          <div className="flex items-center gap-1">
            {/* More */}
            <div className="relative">
              <button
                className="p-2 rounded-full hover:bg-black hover:bg-opacity-10 transition-colors"
                onClick={(e) => {
                  e.stopPropagation()
                  setShowMoreMenu(!showMoreMenu)
                  setShowColorPicker(false)
                }}
                title="More"
              >
                <svg className="w-4 h-4 text-gray-600" fill="currentColor" viewBox="0 0 24 24">
                  <circle cx="12" cy="6" r="2"/>
                  <circle cx="12" cy="12" r="2"/>
                  <circle cx="12" cy="18" r="2"/>
                </svg>
              </button>
              
              {/* More Menu Dropdown */}
              {showMoreMenu && (
                <div className="absolute bottom-full right-0 mb-2 bg-white rounded-lg shadow-lg border border-gray-200 py-2 min-w-[140px] z-20">
                  <button
                    onClick={handleDelete}
                    className="w-full px-4 py-2 text-left text-sm text-gray-700 hover:bg-gray-100 transition-colors"
                  >
                    Delete note
                  </button>
                  <button 
                    onClick={(e) => e.stopPropagation()}
                    className="w-full px-4 py-2 text-left text-sm text-gray-700 hover:bg-gray-100 transition-colors"
                  >
                    Add label
                  </button>
                  <button 
                    onClick={(e) => e.stopPropagation()}
                    className="w-full px-4 py-2 text-left text-sm text-gray-700 hover:bg-gray-100 transition-colors"
                  >
                    Make a copy
                  </button>
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}