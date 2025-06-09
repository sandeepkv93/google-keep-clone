import { useState } from 'react'
import type { Note, CreateNoteRequest } from './types/note'
import './App.css'

// Real Google Keep sample data
const SAMPLE_NOTES: Note[] = [
  {
    id: '1',
    title: 'Ideas for the dog park article',
    content: '',
    color: '#ffffff',
    is_pinned: true,
    is_archived: false,
    is_deleted: false,
    created_at: new Date().toISOString(),
    updated_at: new Date().toISOString(),
    user_id: 'demo-user',
    labels: []
  },
  {
    id: '2',
    title: '',
    content: 'Bring mom lunch on Thursday',
    color: '#fff475',
    is_pinned: false,
    is_archived: false,
    is_deleted: false,
    created_at: new Date().toISOString(),
    updated_at: new Date().toISOString(),
    user_id: 'demo-user',
    labels: []
  },
  {
    id: '3',
    title: 'Books to read',
    content: 'The 7 Habits of Highly Effective People\nEverything Is F*cked\nYour Money or Your Life',
    color: '#a7ffeb',
    is_pinned: false,
    is_archived: false,
    is_deleted: false,
    created_at: new Date().toISOString(),
    updated_at: new Date().toISOString(),
    user_id: 'demo-user',
    labels: []
  }
]

function App() {
  const [notes, setNotes] = useState<Note[]>(SAMPLE_NOTES)
  const [isLoading, setIsLoading] = useState(false)
  const [isSidebarOpen, setIsSidebarOpen] = useState(true)
  const [searchQuery, setSearchQuery] = useState('')
  const [isCreateNoteExpanded, setIsCreateNoteExpanded] = useState(false)
  const [newNoteTitle, setNewNoteTitle] = useState('')
  const [newNoteContent, setNewNoteContent] = useState('')

  const handleCreateNote = () => {
    if (!newNoteTitle.trim() && !newNoteContent.trim()) return

    const newNote: Note = {
      id: Date.now().toString(),
      title: newNoteTitle,
      content: newNoteContent,
      color: '#ffffff',
      is_pinned: false,
      is_archived: false,
      is_deleted: false,
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
      user_id: 'demo-user',
      labels: []
    }
    
    setNotes(prev => [newNote, ...prev])
    setNewNoteTitle('')
    setNewNoteContent('')
    setIsCreateNoteExpanded(false)
  }

  const handleNoteUpdate = (updatedNote: Note) => {
    setNotes(prev => prev.map(note => 
      note.id === updatedNote.id ? updatedNote : note
    ))
  }

  const handleNoteDelete = (noteId: string) => {
    setNotes(prev => prev.filter(note => note.id !== noteId))
  }

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === 'Escape') {
      setIsCreateNoteExpanded(false)
      setNewNoteTitle('')
      setNewNoteContent('')
    }
  }

  const activeNotes = notes.filter(note => !note.is_archived && !note.is_deleted)
  const pinnedNotes = activeNotes.filter(note => note.is_pinned)
  const unpinnedNotes = activeNotes.filter(note => !note.is_pinned)

  return (
    <div className="h-screen bg-white flex flex-col">
      {/* Header - EXACT Google Keep sizing: 64px height */}
      <header className="h-16 border-b border-gray-200 flex items-center px-4">
        <div className="flex items-center gap-3">
          {/* Menu Button - 18px icon */}
          <button
            onClick={() => setIsSidebarOpen(!isSidebarOpen)}
            className="p-3 rounded-full hover:bg-gray-100"
          >
            <svg className="w-4.5 h-4.5 text-gray-700" style={{width:'18px', height:'18px'}} fill="currentColor" viewBox="0 0 24 24">
              <path d="M3 18h18v-2H3v2zm0-5h18v-2H3v2zm0-7v2h18V6H3z"/>
            </svg>
          </button>

          {/* Keep Logo & Title - 32px logo */}
          <div className="flex items-center gap-2">
            <svg className="w-8 h-8" style={{width:'32px', height:'32px'}} viewBox="0 0 24 24">
              <path fill="#fbbc04" d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
            </svg>
            <span className="text-xl text-gray-700 font-normal">Keep</span>
          </div>
        </div>

        {/* Search - 16px icon */}
        <div className="flex-1 max-w-2xl mx-6">
          <div className="relative">
            <svg className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-500" style={{width:'16px', height:'16px'}} fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
            <input
              type="text"
              placeholder="Search"
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              className="w-full pl-10 pr-4 py-2.5 bg-gray-100 rounded-lg border-0 focus:bg-white focus:shadow-md focus:outline-none text-sm"
            />
          </div>
        </div>

        {/* Right icons - 18px icons */}
        <div className="flex items-center gap-1">
          <button className="p-3 rounded-full hover:bg-gray-100">
            <svg className="text-gray-700" style={{width:'18px', height:'18px'}} fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
            </svg>
          </button>
          <button className="p-3 rounded-full hover:bg-gray-100">
            <svg className="text-gray-700" style={{width:'18px', height:'18px'}} fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
          </button>
          <div className="w-8 h-8 bg-blue-500 rounded-full flex items-center justify-center">
            <span className="text-white text-sm font-medium">A</span>
          </div>
        </div>
      </header>

      <div className="flex flex-1 overflow-hidden">
        {/* Sidebar - EXACT Google Keep: 280px width */}
        {isSidebarOpen && (
          <aside className="border-r border-gray-200" style={{width:'280px'}}>
            <nav className="p-2">
              <ul className="space-y-1">
                <li>
                  <button className="w-full flex items-center gap-5 px-6 py-3 rounded-r-full bg-orange-50 text-left">
                    <svg className="text-orange-500" style={{width:'20px', height:'20px'}} fill="currentColor" viewBox="0 0 24 24">
                      <path d="M9 21c0 .5.4 1 1 1h4c.6 0 1-.5 1-1v-1H9v1zm3-19C8.1 2 5 5.1 5 9c0 2.4 1.2 4.5 3 5.7V17h8v-2.3c1.8-1.3 3-3.4 3-5.7 0-3.9-3.1-7-7-7z"/>
                    </svg>
                    <span className="text-gray-900 text-sm font-medium">Notes</span>
                  </button>
                </li>
                <li>
                  <button className="w-full flex items-center gap-5 px-6 py-3 rounded-r-full text-left hover:bg-gray-50">
                    <svg className="text-gray-600" style={{width:'20px', height:'20px'}} fill="currentColor" viewBox="0 0 24 24">
                      <path d="M12 2A10 10 0 0 0 2 12a10 10 0 0 0 10 10 10 10 0 0 0 10-10A10 10 0 0 0 12 2zm4.2 14.2L11 13V7h1.5v5.2l4.5 2.7-.8 1.3z"/>
                    </svg>
                    <span className="text-gray-600 text-sm">Reminders</span>
                  </button>
                </li>
                <li>
                  <button className="w-full flex items-center gap-5 px-6 py-3 rounded-r-full text-left hover:bg-gray-50">
                    <svg className="text-gray-600" style={{width:'20px', height:'20px'}} fill="currentColor" viewBox="0 0 24 24">
                      <path d="M17.63 5.84C17.27 5.33 16.67 5 16 5L5 5.01C3.9 5.01 3 5.9 3 7v10c0 1.1.9 1.99 2 1.99L16 19c.67 0 1.27-.33 1.63-.84L22 12l-4.37-6.16z"/>
                    </svg>
                    <span className="text-gray-600 text-sm">Edit labels</span>
                  </button>
                </li>
                <li>
                  <button className="w-full flex items-center gap-5 px-6 py-3 rounded-r-full text-left hover:bg-gray-50">
                    <svg className="text-gray-600" style={{width:'20px', height:'20px'}} fill="currentColor" viewBox="0 0 24 24">
                      <path d="M20.54 5.23l-1.39-1.68C18.88 3.21 18.47 3 18 3H6c-.47 0-.88.21-1.16.55L3.46 5.23C3.17 5.57 3 6.02 3 6.5V19c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V6.5c0-.48-.17-.93-.46-1.27z"/>
                    </svg>
                    <span className="text-gray-600 text-sm">Archive</span>
                  </button>
                </li>
                <li>
                  <button className="w-full flex items-center gap-5 px-6 py-3 rounded-r-full text-left hover:bg-gray-50">
                    <svg className="text-gray-600" style={{width:'20px', height:'20px'}} fill="currentColor" viewBox="0 0 24 24">
                      <path d="M6 19c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2V7H6v12zM19 4h-3.5l-1-1h-5l-1 1H5v2h14V4z"/>
                    </svg>
                    <span className="text-gray-600 text-sm">Trash</span>
                  </button>
                </li>
              </ul>
            </nav>
          </aside>
        )}

        {/* Main Content */}
        <main className="flex-1 overflow-auto">
          <div className="max-w-5xl mx-auto p-6">
            {/* Take a note */}
            <div className="max-w-2xl mx-auto mb-8">
              {!isCreateNoteExpanded ? (
                <div 
                  className="bg-white rounded-lg border border-gray-300 shadow-sm hover:shadow-md transition-shadow cursor-text p-4"
                  onClick={() => setIsCreateNoteExpanded(true)}
                >
                  <div className="flex items-center gap-4">
                    <span className="text-gray-500 text-base">Take a note...</span>
                    <div className="ml-auto flex items-center gap-3">
                      <button className="p-2.5 rounded-full hover:bg-gray-100">
                        <svg className="text-gray-500" style={{width:'16px', height:'16px'}} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                        </svg>
                      </button>
                      <button className="p-2.5 rounded-full hover:bg-gray-100">
                        <svg className="text-gray-500" style={{width:'16px', height:'16px'}} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                        </svg>
                      </button>
                    </div>
                  </div>
                </div>
              ) : (
                <div className="bg-white rounded-lg border border-gray-300 shadow-lg">
                  <div className="p-4">
                    <input
                      type="text"
                      placeholder="Title"
                      value={newNoteTitle}
                      onChange={(e) => setNewNoteTitle(e.target.value)}
                      onKeyDown={handleKeyDown}
                      className="w-full text-base font-medium placeholder-gray-500 bg-transparent border-none outline-none mb-3"
                      autoFocus
                    />
                    <textarea
                      placeholder="Take a note..."
                      value={newNoteContent}
                      onChange={(e) => setNewNoteContent(e.target.value)}
                      onKeyDown={handleKeyDown}
                      className="w-full placeholder-gray-500 bg-transparent border-none outline-none resize-none min-h-[80px] text-sm"
                      rows={4}
                    />
                  </div>
                  <div className="flex items-center justify-between p-3">
                    <div className="flex items-center gap-1">
                      <button className="p-2.5 rounded-full hover:bg-gray-100">
                        <svg className="text-gray-600" style={{width:'14px', height:'14px'}} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                        </svg>
                      </button>
                      <button className="p-2.5 rounded-full hover:bg-gray-100">
                        <svg className="text-gray-600" style={{width:'14px', height:'14px'}} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                        </svg>
                      </button>
                      <button className="p-2.5 rounded-full hover:bg-gray-100">
                        <svg className="text-gray-600" style={{width:'14px', height:'14px'}} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4z" />
                        </svg>
                      </button>
                    </div>
                    <button
                      onClick={handleCreateNote}
                      className="px-6 py-1 text-sm font-medium text-gray-600 hover:bg-gray-100 rounded"
                    >
                      Close
                    </button>
                  </div>
                </div>
              )}
            </div>

            {/* Notes Grid */}
            {activeNotes.length === 0 ? (
              <div className="flex flex-col items-center justify-center py-24">
                <svg className="w-32 h-32 text-gray-300 mb-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1} d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
                </svg>
                <p className="text-gray-500 text-xl">Notes you add appear here</p>
              </div>
            ) : (
              <div className="space-y-8">
                {/* Pinned Notes */}
                {pinnedNotes.length > 0 && (
                  <section>
                    <h2 className="text-xs font-medium text-gray-600 mb-4 uppercase tracking-wide">Pinned</h2>
                    <div className="columns-1 sm:columns-2 lg:columns-3 xl:columns-4 gap-4">
                      {pinnedNotes.map((note) => (
                        <div
                          key={note.id}
                          className="relative bg-white rounded-lg border border-gray-300 shadow-sm hover:shadow-md transition-shadow cursor-pointer break-inside-avoid mb-4 group"
                          style={{ backgroundColor: note.color }}
                        >
                          {/* Pin button - 14px icon */}
                          <div className="absolute top-2 right-2">
                            <button
                              className="p-2 rounded-full opacity-0 group-hover:opacity-100 hover:bg-black hover:bg-opacity-10 transition-all"
                              onClick={(e) => {
                                e.stopPropagation()
                                handleNoteUpdate({ ...note, is_pinned: !note.is_pinned })
                              }}
                            >
                              <svg className="text-gray-600" style={{width:'14px', height:'14px'}} fill="currentColor" viewBox="0 0 24 24">
                                <path d="M14 4l-4 4v7l4-4V4z"/>
                              </svg>
                            </button>
                          </div>

                          <div className="p-4 pr-12">
                            {note.title && (
                              <h3 className="font-medium text-gray-900 mb-3 text-base">{note.title}</h3>
                            )}
                            {note.content && (
                              <div className="text-gray-700 text-sm whitespace-pre-wrap">{note.content}</div>
                            )}
                          </div>

                          {/* Hover toolbar - 14px icons */}
                          <div className="absolute bottom-0 left-0 right-0 bg-white bg-opacity-95 p-2 opacity-0 group-hover:opacity-100 transition-opacity">
                            <div className="flex items-center justify-between">
                              <div className="flex items-center gap-1">
                                <button className="p-2 rounded-full hover:bg-black hover:bg-opacity-10">
                                  <svg className="text-gray-600" style={{width:'14px', height:'14px'}} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                                  </svg>
                                </button>
                                <button className="p-2 rounded-full hover:bg-black hover:bg-opacity-10">
                                  <svg className="text-gray-600" style={{width:'14px', height:'14px'}} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M16 7a4 4 0 11-8 0 4 4 0 018 0z" />
                                  </svg>
                                </button>
                                <button className="p-2 rounded-full hover:bg-black hover:bg-opacity-10">
                                  <svg className="text-gray-600" style={{width:'14px', height:'14px'}} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4z" />
                                  </svg>
                                </button>
                              </div>
                              <button
                                onClick={(e) => {
                                  e.stopPropagation()
                                  handleNoteDelete(note.id)
                                }}
                                className="p-2 rounded-full hover:bg-black hover:bg-opacity-10"
                              >
                                <svg className="text-gray-600" style={{width:'14px', height:'14px'}} fill="currentColor" viewBox="0 0 24 24">
                                  <circle cx="12" cy="6" r="2"/>
                                  <circle cx="12" cy="12" r="2"/>
                                  <circle cx="12" cy="18" r="2"/>
                                </svg>
                              </button>
                            </div>
                          </div>
                        </div>
                      ))}
                    </div>
                  </section>
                )}

                {/* Other Notes */}
                {unpinnedNotes.length > 0 && (
                  <section>
                    {pinnedNotes.length > 0 && (
                      <h2 className="text-xs font-medium text-gray-600 mb-4 uppercase tracking-wide">Others</h2>
                    )}
                    <div className="columns-1 sm:columns-2 lg:columns-3 xl:columns-4 gap-4">
                      {unpinnedNotes.map((note) => (
                        <div
                          key={note.id}
                          className="relative bg-white rounded-lg border border-gray-300 shadow-sm hover:shadow-md transition-shadow cursor-pointer break-inside-avoid mb-4 group"
                          style={{ backgroundColor: note.color }}
                        >
                          {/* Pin button - 14px icon */}
                          <div className="absolute top-2 right-2">
                            <button
                              className="p-2 rounded-full opacity-0 group-hover:opacity-100 hover:bg-black hover:bg-opacity-10 transition-all"
                              onClick={(e) => {
                                e.stopPropagation()
                                handleNoteUpdate({ ...note, is_pinned: !note.is_pinned })
                              }}
                            >
                              <svg className="text-gray-600" style={{width:'14px', height:'14px'}} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z" />
                              </svg>
                            </button>
                          </div>

                          <div className="p-4 pr-12">
                            {note.title && (
                              <h3 className="font-medium text-gray-900 mb-3 text-base">{note.title}</h3>
                            )}
                            {note.content && (
                              <div className="text-gray-700 text-sm whitespace-pre-wrap">{note.content}</div>
                            )}
                          </div>

                          {/* Hover toolbar - 14px icons */}
                          <div className="absolute bottom-0 left-0 right-0 bg-white bg-opacity-95 p-2 opacity-0 group-hover:opacity-100 transition-opacity">
                            <div className="flex items-center justify-between">
                              <div className="flex items-center gap-1">
                                <button className="p-2 rounded-full hover:bg-black hover:bg-opacity-10">
                                  <svg className="text-gray-600" style={{width:'14px', height:'14px'}} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                                  </svg>
                                </button>
                                <button className="p-2 rounded-full hover:bg-black hover:bg-opacity-10">
                                  <svg className="text-gray-600" style={{width:'14px', height:'14px'}} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M16 7a4 4 0 11-8 0 4 4 0 018 0z" />
                                  </svg>
                                </button>
                                <button className="p-2 rounded-full hover:bg-black hover:bg-opacity-10">
                                  <svg className="text-gray-600" style={{width:'14px', height:'14px'}} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4z" />
                                  </svg>
                                </button>
                              </div>
                              <button
                                onClick={(e) => {
                                  e.stopPropagation()
                                  handleNoteDelete(note.id)
                                }}
                                className="p-2 rounded-full hover:bg-black hover:bg-opacity-10"
                              >
                                <svg className="text-gray-600" style={{width:'14px', height:'14px'}} fill="currentColor" viewBox="0 0 24 24">
                                  <circle cx="12" cy="6" r="2"/>
                                  <circle cx="12" cy="12" r="2"/>
                                  <circle cx="12" cy="18" r="2"/>
                                </svg>
                              </button>
                            </div>
                          </div>
                        </div>
                      ))}
                    </div>
                  </section>
                )}
              </div>
            )}
          </div>
        </main>
      </div>
    </div>
  )
}

export default App