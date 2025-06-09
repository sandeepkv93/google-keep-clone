interface SidebarProps {
  isOpen: boolean
  activeSection: string
  onSectionChange: (section: string) => void
}

export function Sidebar({ isOpen, activeSection, onSectionChange }: SidebarProps) {
  const menuItems = [
    {
      id: 'notes',
      label: 'Notes',
      icon: (
        <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
          <path d="M9 21c0 .5.4 1 1 1h4c.6 0 1-.5 1-1v-1H9v1zm3-19C8.1 2 5 5.1 5 9c0 2.4 1.2 4.5 3 5.7V17h8v-2.3c1.8-1.3 3-3.4 3-5.7 0-3.9-3.1-7-7-7z"/>
        </svg>
      )
    },
    {
      id: 'reminders',
      label: 'Reminders',
      icon: (
        <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
          <path d="M12 2A10 10 0 0 0 2 12a10 10 0 0 0 10 10 10 10 0 0 0 10-10A10 10 0 0 0 12 2zm4.2 14.2L11 13V7h1.5v5.2l4.5 2.7-.8 1.3z"/>
        </svg>
      )
    },
    {
      id: 'labels',
      label: 'Edit labels',
      icon: (
        <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
          <path d="M17.63 5.84C17.27 5.33 16.67 5 16 5L5 5.01C3.9 5.01 3 5.9 3 7v10c0 1.1.9 1.99 2 1.99L16 19c.67 0 1.27-.33 1.63-.84L22 12l-4.37-6.16zM16 17H5V7h11l3.55 5L16 17z"/>
        </svg>
      )
    },
    {
      id: 'archive',
      label: 'Archive',
      icon: (
        <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
          <path d="M20.54 5.23l-1.39-1.68C18.88 3.21 18.47 3 18 3H6c-.47 0-.88.21-1.16.55L3.46 5.23C3.17 5.57 3 6.02 3 6.5V19c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V6.5c0-.48-.17-.93-.46-1.27zM12 17.5L6.5 12H10v-2h4v2h3.5L12 17.5zM5.12 5l.81-1h12l.94 1H5.12z"/>
        </svg>
      )
    },
    {
      id: 'trash',
      label: 'Trash',
      icon: (
        <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
          <path d="M6 19c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2V7H6v12zM19 4h-3.5l-1-1h-5l-1 1H5v2h14V4z"/>
        </svg>
      )
    }
  ]

  return (
    <aside className={`bg-white border-r border-gray-200 transition-all duration-300 ${
      isOpen ? 'w-72' : 'w-0 overflow-hidden'
    }`}>
      <nav className="p-2 mt-2">
        <ul className="space-y-1">
          {menuItems.map((item) => (
            <li key={item.id}>
              <button
                onClick={() => onSectionChange(item.id)}
                className={`w-full flex items-center gap-6 px-6 py-3 rounded-r-full text-left transition-all duration-200 ${
                  activeSection === item.id
                    ? 'bg-yellow-50 text-gray-900 shadow-sm'
                    : 'text-gray-600 hover:bg-gray-50'
                }`}
              >
                <span className={`${
                  activeSection === item.id ? 'text-orange-500' : 'text-gray-600'
                }`}>
                  {item.icon}
                </span>
                <span className="font-medium">{item.label}</span>
              </button>
            </li>
          ))}
        </ul>
      </nav>
    </aside>
  )
}