import { useState } from 'react'

interface HeaderProps {
  onMenuClick: () => void
  searchQuery: string
  onSearchChange: (query: string) => void
}

export function Header({ onMenuClick, searchQuery, onSearchChange }: HeaderProps) {
  const [isSearchFocused, setIsSearchFocused] = useState(false)

  return (
    <header className="h-16 bg-white border-b border-gray-200 flex items-center px-5 relative z-50">
      {/* Left Section */}
      <div className="flex items-center gap-4">
        {/* Menu Button */}
        <button
          onClick={onMenuClick}
          className="p-3 rounded-full hover:bg-gray-100 transition-colors"
          aria-label="Main menu"
        >
          <svg className="w-6 h-6 text-gray-600" fill="currentColor" viewBox="0 0 24 24">
            <path d="M3 18h18v-2H3v2zm0-5h18v-2H3v2zm0-7v2h18V6H3z"/>
          </svg>
        </button>

        {/* Google Keep Logo */}
        <div className="flex items-center gap-2">
          <div className="w-10 h-10 relative">
            {/* Keep Logo - Multi-colored lightbulb */}
            <svg viewBox="0 0 24 24" className="w-10 h-10">
              <path d="M9 21c0 .5.4 1 1 1h4c.6 0 1-.5 1-1v-1H9v1zm3-19C8.1 2 5 5.1 5 9c0 2.4 1.2 4.5 3 5.7V17h8v-2.3c1.8-1.3 3-3.4 3-5.7 0-3.9-3.1-7-7-7z" fill="#fbbc04"/>
              <path d="M12 2C8.1 2 5 5.1 5 9c0 2.4 1.2 4.5 3 5.7V17h8v-2.3c1.8-1.3 3-3.4 3-5.7 0-3.9-3.1-7-7-7z" fill="#ea4335" fillOpacity="0.3"/>
              <path d="M12 2C8.1 2 5 5.1 5 9c0 2.4 1.2 4.5 3 5.7V17h8v-2.3c1.8-1.3 3-3.4 3-5.7 0-3.9-3.1-7-7-7z" fill="#34a853" fillOpacity="0.4"/>
              <path d="M12 2C8.1 2 5 5.1 5 9c0 2.4 1.2 4.5 3 5.7V17h8v-2.3c1.8-1.3 3-3.4 3-5.7 0-3.9-3.1-7-7-7z" fill="#4285f4" fillOpacity="0.5"/>
            </svg>
          </div>
          <span className="text-2xl text-gray-700 font-normal">Keep</span>
        </div>
      </div>

      {/* Search Bar */}
      <div className="flex-1 max-w-3xl mx-6">
        <div className="relative">
          <div className="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
            <svg className="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
          <input
            type="text"
            placeholder="Search"
            value={searchQuery}
            onChange={(e) => onSearchChange(e.target.value)}
            onFocus={() => setIsSearchFocused(true)}
            onBlur={() => setIsSearchFocused(false)}
            className={`w-full pl-12 pr-4 py-2 rounded-lg border-0 focus:outline-none transition-all duration-200 ${
              isSearchFocused 
                ? 'bg-white shadow-md' 
                : 'bg-gray-100 hover:bg-gray-50'
            }`}
          />
        </div>
      </div>

      {/* Right Section */}
      <div className="flex items-center gap-2">
        {/* Refresh Button */}
        <button
          className="p-3 rounded-full hover:bg-gray-100 transition-colors"
          aria-label="Refresh"
        >
          <svg className="w-6 h-6 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
        </button>

        {/* List View Toggle */}
        <button
          className="p-3 rounded-full hover:bg-gray-100 transition-colors"
          aria-label="List view"
        >
          <svg className="w-6 h-6 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
          </svg>
        </button>

        {/* Settings */}
        <button
          className="p-3 rounded-full hover:bg-gray-100 transition-colors"
          aria-label="Settings"
        >
          <svg className="w-6 h-6 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
        </button>

        {/* User Avatar */}
        <button
          className="w-8 h-8 rounded-full bg-blue-500 flex items-center justify-center hover:shadow-md transition-shadow"
          aria-label="Account"
        >
          <span className="text-white text-sm font-medium">A</span>
        </button>
      </div>
    </header>
  )
}