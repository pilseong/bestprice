import React from 'react'
import Heading from "@/components/Heading"
import Link from "next/link"
// import { getFeaturedReview, getItems } from "@/lib/reviews"
import Image from "next/image"
import { getTmonItems } from "@/lib/items"

// export const dynamic = 'force-dynamic'
export const revalidate = 600

async function TmonPage() {

  // const featuredReview = await getFeaturedReview()
  const items = await getTmonItems()




  return (
    <>
      <ul className="flex flex-row flex-wrap gap-3 justify-center">
        {
          items && items.map((item, index) => (
            <li key={item.id} className="list-none border w-80 bg-white shadow rounded hover:shadow-xl mb-4">
              <div className="flex justify-around items-center bg-pink-100 text-gray-900">
                <div className="font-exo2"><span className="text-red-900">할인가 {item.price.toLocaleString()}</span><span className="text-sm"> <span className="line-through text-sm">{item.originalPrice.toLocaleString()}</span> ({item.discountRate || '0'}% off)</span></div>
              </div>
              <div className="flex justify-around items-center bg-pink-100 text-gray-900">
                <div><span className="text-gray-700">{item.deliveryFeePolicy !== '' ? '배송비' : ''} {
                  item.deliveryFeePolicy === 'FREE' ? '무료' : item.deliveryFeePolicy === 'CONDITION' ? '조건부 무료' : item.deliveryFeePolicy === '' ? '디지털 구폰' : ''
                } {!(item.deliveryFeePolicy === 'FREE' || item.deliveryFeePolicy === '') ? item.deliveryFee : ''}</span></div>
              </div>
              <a href={`${item.url}`} target="_blank">
                <Image src={`${item.pc3ColImageUrl}`} alt=""
                  width={320} height={180} className="rounded-t" />
                <h2 className="p-2 border-t text-sm font-semibold font-notosanskr">{item.title}</h2>
              </a>
            </li>
          ))
        }
      </ul>
    </>
  )
}

export default TmonPage