import React from 'react'
import Heading from "@/components/Heading"
import Link from "next/link"
// import { getFeaturedReview, getItems } from "@/lib/reviews"
import Image from "next/image"
import { getElevenStItems, getGmarketItems } from "@/lib/items"

export const dynamic = 'force-dynamic'
// export const revalidate = 600

async function ElevenstPage() {

  // const featuredReview = await getFeaturedReview()
  const items = await getElevenStItems()


  return (
    <>
      <ul className="flex flex-row flex-wrap gap-4 justify-center">
        {
          items && items.map((item, index) => (
            <li key={item.id} className="list-none mb-4 border w-72 bg-white shadow rounded hover:shadow-xl flex flex-col">
              <a href={`${item.linkInfo}`} target="_blank">
                <div>
                  <Image src={`${item.imageLink}`} alt=""
                    width={320} height={180} className="rounded-t" />
                  <h2 className="p-2 border-t text-sm font-notosanskr">{item.title}</h2>
                </div>
              </a>
              <div className="mt-auto">
                <div className="flex justify-around items-center bg-pink-100 text-gray-900 pt-1">
                  <div className="font-exo2"><span className="text-red-900">할인가 {item.discountPrice}</span><span className="text-sm"> <span className="line-through text-sm">{item.originalPrice}</span> ({item.discountRate || '0'}% off)</span></div>
                </div>
                <div className="flex justify-around items-center bg-pink-100 text-gray-900 pb-1">
                  <div><span className="text-gray-700 text-sm font-exo2">배송비 {item.shippingText || "유료"}</span></div>
                </div>
              </div>
            </li>
          ))
        }
      </ul>
    </>
  )
}

export default ElevenstPage