import React from 'react'
import Heading from "@/components/Heading"
import Link from "next/link"
// import { getFeaturedReview, getItems } from "@/lib/reviews"
import Image from "next/image"
import { getFrontItems } from "@/lib/items"

// export const dynamic = 'force-dynamic'
export const revalidate = 600

async function HomePage() {

  // const featuredReview = await getFeaturedReview()
  const items = await getFrontItems()

  return (
    <>
      <div className="flex flex-row flex-wrap gap-4 justify-center">
        {
          items && items.map((item, index) => (
            <>
              <div key={item.id} className="list-none mb-4 border w-72 bg-white shadow rounded hover:shadow-xl flex flex-col">
                <div className="flex justify-between">
                  <div className={`w-8 text-center ${item.rank < 5 ? 'bg-red-400 text-white' : 'bg-blue-300 text-gray-100'} `}>{item.rank}</div>
                  <div className=
                    {`capitalize pr-2 flex-1 text-right 
                  ${item.shoppingMall === "지마켓" ? 'bg-green-600 text-gray-200' :
                        item.shoppingMall === "위메프" ? 'bg-red-500 text-gray-100' :
                          item.shoppingMall === "11번가" ? 'bg-red-400 text-gray-100' :
                            item.shoppingMall === "티몬" ? 'bg-orange-400 text-gray-100' : ''
                      }`}>{item.shoppingMall}</div>
                </div>
                <a href={item.linkUrl} target="_blank">
                  <div>
                    <Image src={item.imageUrl} alt=""
                      width={320} height={180} className="rounded-t" />
                    <h2 className="p-2 border-t text-sm font-semibold font-notosanskr">{item.title}</h2>
                  </div>
                </a>
                <div className="mt-auto">
                  <div className="flex justify-around items-center bg-pink-100 text-gray-900 pt-1">
                    <div className="font-exo2">
                      <span className="text-red-900">할인가 {item.afterPrice.toLocaleString()}</span>{' '}
                      <span className="text-sm">
                        <span className="line-through text-sm">{item.beforePrice.toLocaleString()}</span> ({item.discountRate}% off)
                      </span>
                    </div>
                  </div>
                  <div className="flex justify-around items-center bg-pink-100 text-gray-900 pb-1 text-sm font-exo2">
                    <div><span className="text-gray-700">{item.shippingInfo}</span></div>
                  </div>
                </div>
              </div>
            </>
          ))
        }
      </div >
    </>
  )
}

export default HomePage