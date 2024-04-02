'use client'
import { usePathname } from 'next/navigation'

import Link from "next/link"
import React from 'react'
import Image from "next/image"

function Navbar() {
  const pathname = usePathname()

  console.log(pathname)

  return (
    <nav className="flex flex-col font-notosanskr">
      <div className="flex justify-center sm:flex sm:justify-between bg-trust items-center px-5 h-14 sm:h-8">
        <div className="text-gray-600 uppercase font-orbitron flex gap-2">
          <Image src="/icon.jpg" width={24} height={24} alt="logo for pilseong.net" />
          Pilseong.net</div>
        <div className="hidden sm:block  font-exo2">Best Price Based On Sales</div>
      </div>
      <div className="text-center bg-slate-50">
        <h2 className="text-gray-600 py-3 text-2xl font-bold uppercase font-orbitron">핫딜 TO 구매욕 Booster</h2>
      </div>
      <ul className="hidden sm:block sm:flex text-center">
        <li className={`py-4 text-gray-100 text-sm w-[25%] hover:underline font-semibold bg-gmarket ${pathname === '/gmarket' ? '' : ''}`}>
          <Link href="/gmarket" className={`${pathname === '/gmarket' ? 'font-extrabold px-4 py-1 bg-green-600 rounded-full' : 'font-normal'}`}>지마켓</Link>
        </li>
        <li className={`py-4 text-gray-100 w-[25%] text-sm hover:underline bg-red-600 font-semibold ${pathname === '/wemakeprice' ? '' : ''}`}>
          <Link href="/wemakeprice" className={`${pathname === '/wemakeprice' ? 'font-extrabold px-4 py-1 bg-red-500 rounded-full' : 'font-normal'}`}>위메프</Link>
        </li>
        <li className={`py-4 text-gray-100 w-[25%] text-sm hover:underline bg-red-500 font-semibold ${pathname === '/11st' ? '' : ''}`}>
          <Link href="/11st" className={`${pathname === '/11st' ? 'font-extrabold px-4 py-1 bg-red-400 rounded-full' : 'font-normal'}`}>11번가</Link>
        </li>
        <li className={`py-4 text-gray-100 w-[25%] text-sm hover:underline bg-orange-600 font-semibold ${pathname === '/tmon' ? '' : ''}`}>
          <Link href="/tmon" className={`${pathname === '/tmon' ? 'font-extrabold px-4 py-1 bg-orange-400 rounded-full' : 'font-normal'}`}>티몬</Link>
        </li>
      </ul>
    </nav>
  )
}

export default Navbar