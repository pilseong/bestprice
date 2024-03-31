'use client'
import { usePathname } from 'next/navigation'

import Link from "next/link"
import React from 'react'

function Navbar() {
  const pathname = usePathname()

  console.log(pathname)

  return (
    <nav className="flex flex-col font-notosanskr">
      <div className="flex justify-center sm:flex sm:justify-between bg-trust items-center px-5 h-14 sm:h-8">
        <div className="text-gray-600 uppercase font-orbitron">Pilseong.net</div>
        <div className="hidden sm:block  font-exo2">Best Price Based On Sales</div>
      </div>
      <div className="text-center bg-slate-50">
        <h2 className="text-gray-600 py-3 text-2xl font-bold uppercase font-orbitron">핫딜 TO 구매욕 Booster</h2>
      </div>
      <ul className="hidden sm:block sm:flex sm:justify-around text-center">
        <li className={`py-4 text-gray-100 text-sm flex-grow hover:underline font-semibold bg-gmarket ${pathname === '/gmarket' ? 'text-xl' : 'text-sm'}`}>
          <Link href="/gmarket">지마켓</Link>
        </li>
        <li className="py-4 text-gray-100 flex-grow text-sm hover:underline bg-red-600 font-semibold">
          <Link href="/gmarket">위메프</Link>
        </li>
        <li className="py-4 text-gray-100 flex-grow text-sm hover:underline bg-red-500 font-semibold">
          <Link href="/tmon">11번가</Link>
        </li>
        <li className="py-4 text-gray-100 flex-grow text-sm hover:underline bg-orange-600 font-semibold">
          <Link href="/tmon">티몬</Link>
        </li>
      </ul>
    </nav>
  )
}

export default Navbar