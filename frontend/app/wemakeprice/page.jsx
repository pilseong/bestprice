import React from 'react'
import Heading from "@/components/Heading"
import Link from "next/link"
// import { getFeaturedReview, getItems } from "@/lib/reviews"
import Image from "next/image"
import { getWeMakePriceItems } from "@/lib/items"

export const dynamic = 'force-dynamic'
export const revalidate = 600

async function WeMakePricePage() {

  // const featuredReview = await getFeaturedReview()
  const items = await getWeMakePriceItems()

  return (
    <>
      <ul className="flex flex-row flex-wrap gap-4 justify-center">
        {
          items && items.map((item, index) => (
            <li key={item.id} className="list-none border w-72 bg-white shadow rounded hover:shadow-xl mb-4 flex flex-col">
              <a href={`https://${item.linkType.toLowerCase() === 'ticket' ? 'ticket' : item.linkType.toLowerCase() === 'tour' ? 'tour.wd' : 'front'}.wemakeprice.com/${(item.linkType.toLowerCase() === 'ticket' || item.linkType.toLowerCase() === 'prod') ? 'product/' : item.linkType.toLowerCase() === 'tour' ? 'wmp/deal.html?dealId=' : item.linkType.toLowerCase() + '/'}${item.linkInfo}`} target="_blank">
                <Image src={`${item.mediumImgUrl}`} alt=""
                  width={320} height={180} className="rounded-t" />
                <h2 className="p-2 border-t text-sm font-notosanskr">{item.dispNm}</h2>
              </a>
              <div className="mt-auto">
                <div className="flex justify-around items-center bg-pink-100 text-gray-900">
                  <div className="font-exo2 pt-1"><span className="text-red-900">할인가 {item.discountPrice !== 0 ? item.discountPrice.toLocaleString() : item.salePrice.toLocaleString()}</span><span className="text-sm"> <span className="line-through text-sm">{item.discountPrice !== 0 ? item.salePrice.toLocaleString() : item.originPrice.toLocaleString()}</span> ({item.discountRate || '0'}% off)</span></div>
                </div>
                <div className="flex justify-around items-center bg-pink-100 text-gray-900">
                  <div><span className="text-gray-700">{item.deliveryFeePolicy !== '' ? '배송비' : ''} {
                    '무료'
                  } {''}</span></div>
                </div>
              </div>
            </li>
          ))
        }
      </ul >
    </>
  )
}

export default WeMakePricePage