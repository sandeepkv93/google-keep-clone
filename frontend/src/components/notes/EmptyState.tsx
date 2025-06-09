export function EmptyState() {
  return (
    <div className="flex flex-col items-center justify-center py-24">
      {/* Light bulb icon */}
      <div className="w-32 h-32 mb-6 relative">
        <svg viewBox="0 0 200 200" className="w-full h-full">
          {/* Outer circle */}
          <circle cx="100" cy="100" r="80" fill="#F8F9FA" stroke="#E8EAED" strokeWidth="2"/>
          
          {/* Light bulb */}
          <path 
            d="M100 40 C120 40, 140 60, 140 80 C140 95, 135 105, 125 115 L125 135 L75 135 L75 115 C65 105, 60 95, 60 80 C60 60, 80 40, 100 40 Z" 
            fill="#9AA0A6"
          />
          
          {/* Base of bulb */}
          <rect x="80" y="140" width="40" height="8" rx="4" fill="#9AA0A6"/>
          <rect x="85" y="152" width="30" height="4" rx="2" fill="#9AA0A6"/>
          
          {/* Light rays */}
          <g stroke="#E8EAED" strokeWidth="3" strokeLinecap="round">
            <line x1="100" y1="20" x2="100" y2="30"/>
            <line x1="140" y1="40" x2="135" y2="45"/>
            <line x1="160" y1="80" x2="150" y2="80"/>
            <line x1="140" y1="120" x2="135" y2="115"/>
            <line x1="60" y1="40" x2="65" y2="45"/>
            <line x1="40" y1="80" x2="50" y2="80"/>
            <line x1="60" y1="120" x2="65" y2="115"/>
          </g>
        </svg>
      </div>
      
      {/* Text */}
      <p className="text-gray-500 text-xl">Notes you add appear here</p>
    </div>
  )
}