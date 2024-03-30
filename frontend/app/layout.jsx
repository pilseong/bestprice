import React from 'react'

import './globals.css'
import Navbar from "../components/Navbar"
import { orbitron, exo2, notosanskr } from "./fonts"

export const metadata = {
  title: {
    default: 'Best Price',
    template: '%s | Best Price'
  }
}

function RootLayout({ children, }) {
  return (
    <html lang="en" className={`${orbitron.variable} ${exo2.variable} ${notosanskr.variable}`}>
      <body className="flex flex-col min-h-screen">
        <header className=" sticky top-0 shadow-md border-b-4">
          <Navbar />
        </header>
        <main className="container mx-auto py-3 grow">
          {children}
        </main>
        <footer className="text-center text-xs border-t py-4 text-slate-500">
          Enjoy development and let people feel good {' '}<br />
          <div className="text-orange-800 font-semibold pt-1">
            Copyright 2024 Pilseong Heo
          </div>
        </footer>
      </body>
    </html>
  )
}

export default RootLayout