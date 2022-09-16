import React from 'react'

// MainRoom component
function MainRoom() {
  // Renders this final markup
  return (
    <div
      className="d-flex align-items-center justify-content-start px-0 flex-column"
      style={{ height: '100%', maxHeight: '100vh' }}
    >
      <div className="d-flex" style={{ width: '100vw' }}>
        <section
          style={{
            width: `100%`,
            zIndex: '1',
          }}
          className="d-flex flex-column align-items-center justify-content-start"
        ></section>
      </div>
    </div>
  )
}

export default MainRoom
