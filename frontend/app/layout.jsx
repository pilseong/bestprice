import React from 'react'

import './globals.css'
import Navbar from "../components/Navbar"
import { orbitron, exo2, notosanskr, inter } from "./fonts"
import Script from "next/script"

export const metadata = {
  title: {
    default: '필성 닷넷 대형쇼핑 실시간 베스트셀링 제품만 모아서 보여주는 핫딜 모음',
    description: '가장 많이 팔리는 제품만 모아서 보여주는 핫딜 모음',
    template: '%s | 필성.net'
  }
}

function RootLayout({ children }) {
  return (
    <html lang="en" className={`${orbitron.variable} ${exo2.variable} ${notosanskr.variable} ${inter.variable}`}>
      <body className="flex flex-col min-h-screen bg-gray-200">
        <header className=" sticky top-0 z-0">
          <Navbar />
        </header>
        <main className="container mx-auto py-3 grow">
          {children}
        </main>
        <footer className="text-center text-xs border-t py-4 text-slate-500">
          Enjoy development and let people feel good. {' '} <span className="text-sm text-yellow-900">Powered By Go + NextJS Headless</span><br />
          <div className="text-orange-800 font-semibold pt-1">
            Copyright 2024 Pilseong Heo
          </div>
        </footer>
      </body>
    </html>
  )
}

export default RootLayout