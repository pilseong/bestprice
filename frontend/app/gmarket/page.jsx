import React from 'react'
import Heading from "@/components/Heading"
import Link from "next/link"
// import { getFeaturedReview, getItems } from "@/lib/reviews"
import Image from "next/image"
import { getGmarketItems } from "@/lib/items"

// export const dynamic = 'force-dynamic'
export const revalidate = 600

async function GmarketPage() {

  // const featuredReview = await getFeaturedReview()
  const items = await getGmarketItems()


  return (
    <>
      <ul className="flex flex-row flex-wrap gap-4 justify-center">
        {
          items && items.map((item, index) => (
            <li key={item.id} className="list-none mb-4 border w-72 bg-white shadow rounded hover:shadow-xl flex flex-col">
              <a href={`${item.linkUrl.replace('http', 'https')}`} target="_blank">
                <div>
                  <Image src={`https:${item.imageUrl}/300?ver=${item.goodsCode}`} alt=""
                    width={320} height={180} className="rounded-t" />
                  <h2 className="p-2 border-t text-sm font-semibold font-notosanskr">{item.goodsName}</h2>
                </div>
              </a>
              <div className="mt-auto">
                <div className="flex justify-around items-center bg-pink-100 text-gray-900 pt-1">
                  <div className="font-exo2"><span className="text-red-900">할인가 {item.discountPrice}</span><span className="text-sm"> <span className="line-through text-sm">{item.price}</span> ({item.discountRate || '0'}% off)</span></div>
                </div>
                <div className="flex justify-around items-center bg-pink-100 text-gray-900 pb-1 text-sm font-exo2">
                  <div><span className="text-gray-700">배송비 {item.deliveryInfo.replace('<br/>', '')}</span></div>
                </div>
              </div>
            </li>
          ))
        }
      </ul>
    </>
  )
}

export default GmarketPage