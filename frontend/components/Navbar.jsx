'use client'
import { usePathname } from 'next/navigation'

import Link from "next/link"
import React, { useState } from 'react'
import Image from "next/image"
import { Bars3Icon } from '@heroicons/react/20/solid'

function Navbar() {
  const pathname = usePathname()
  const [menuOpen, setMenuOpen] = useState(false)

  console.log(pathname)

  return (
    <nav className="flex flex-col font-notosanskr">
      <div className="flex justify-center sm:flex sm:justify-between bg-trust items-center px-5 h-14 sm:h-8">
        <Link href="/">
          <div className="text-gray-600 uppercase font-inter flex gap-2">
            <Image src="/icon.jpg" width={24} height={24} alt="logo for pilseong.net" />
            필성.net</div>
        </Link>
        <div className="hidden sm:block font-inter">판매량 기준 핫딜 모음</div>
      </div>
      <div className="bg-slate-50 flex items-center justify-between sm:justify-center shadow-md sm:shadow-none">
        <h2 className="text-gray-600 px-4 py-3 text-2xl sm:px-0 font-bold uppercase font-orbitron">
          {pathname !== '/' ? pathname : '핫딜 TO 구매욕 Booster'}
        </h2>
        <div>
          <Bars3Icon onClick={() => setMenuOpen(!menuOpen)} width={24} height={24} className="mx-2 sm:hidden hover:cursor-pointer" />
        </div>
      </div>

      {
        menuOpen && (
          <div className="flex-col text-center ml-auto w-48 sm:hidden bg-slate-50" >
            <div className={`py-4 text-gray-100 text-sm hover:underline font-semibold bg-trust ${pathname === '/gmarket' ? '' : ''}`}>
              <Link href="/" onClick={() => setMenuOpen(false)} className={`${pathname === '/' ? 'font-bold px-4 py-1 bg-green-500 rounded-full' : 'font-normal'}`}>메인 페이지</Link>
            </div>
            <div className={`py-4 text-gray-100 text-sm hover:underline font-semibold bg-gmarket ${pathname === '/gmarket' ? '' : ''}`}>
              <Link href="/gmarket" onClick={() => setMenuOpen(false)} className={`${pathname === '/gmarket' ? 'font-extrabold px-4 py-1 bg-green-600 rounded-full' : 'font-normal'}`}>지마켓</Link>
            </div>
            <div className={`py-4 text-gray-100 text-sm hover:underline bg-red-600 font-semibold ${pathname === '/wemakeprice' ? '' : ''}`}>
              <Link href="/wemakeprice" onClick={() => setMenuOpen(false)} className={`${pathname === '/wemakeprice' ? 'font-extrabold px-4 py-1 bg-red-500 rounded-full' : 'font-normal'}`}>위메프</Link>
            </div>
            <div className={`py-4 text-gray-100 text-sm hover:underline bg-red-500 font-semibold ${pathname === '/11st' ? '' : ''}`}>
              <Link href="/11st" onClick={() => setMenuOpen(false)} className={`${pathname === '/11st' ? 'font-extrabold px-4 py-1 bg-red-400 rounded-full' : 'font-normal'}`}>11번가</Link>
            </div>
            <div className={`py-4 text-gray-100 text-sm hover:underline bg-orange-600 font-semibold ${pathname === '/tmon' ? '' : ''}`}>
              <Link href="/tmon" onClick={() => setMenuOpen(false)} className={`${pathname === '/tmon' ? 'font-extrabold px-4 py-1 bg-orange-400 rounded-full' : 'font-normal'}`}>티몬</Link>
            </div>
          </div>
        )
      }

      <ul className="hidden sm:block sm:flex text-center shadow-md border-b-2">
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
    </nav >
  )
}

export default Navbar