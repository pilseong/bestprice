import Link from "next/link"
import React from 'react'

function Navbar() {
  return (
    <nav className="flex flex-col font-notosanskr">
      <div className="text-gray-600 uppercase bg-trust h-14 flex items-center justify-center sm:h-8 text-center sm:justify-start px-5">
        Pilseong.net
      </div>
      <div className="text-center bg-yellow-50">
        <h2 className="text-gray-600 py-3 text-2xl font-bold uppercase">Best Price</h2>
      </div>
      <ul className="hidden sm:block sm:flex sm:justify-around text-center">
        <li className="py-4 text-gray-100 text-sm flex-grow hover:underline font-semibold bg-gmarket">
          <Link href="/">지마켓</Link>
        </li>
        <li className="py-4 text-gray-100 flex-grow text-sm hover:underline bg-red-600 font-semibold">
          <Link href="/reviews">옥션</Link>
        </li>
        <li className="py-4 text-gray-100 flex-grow text-sm hover:underline bg-red-500 font-semibold">
          <Link href="/about">11번가</Link>
        </li>
        <li className="py-4 text-gray-100 flex-grow text-sm hover:underline bg-orange-600 font-semibold">
          <Link href="/about">티몬</Link>
        </li>
      </ul>
    </nav>
  )
}

export default Navbar