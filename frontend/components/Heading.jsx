import React from 'react'

function Heading({ children }) {
  return (
    <h1 className={`font-bold pb-3 text-2xl font-notosanskr`}>
      {children}
    </h1>
  )
}

export default Heading