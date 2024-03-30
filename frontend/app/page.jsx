import React from 'react'
import Heading from "@/components/Heading"
import Link from "next/link"
// import { getFeaturedReview, getItems } from "@/lib/reviews"
import Image from "next/image"
import { getIt, getItems } from "@/lib/items"

// export const dynamic = 'force-dynamic'
export const revalidate = 3600

async function HomePage() {

  // const featuredReview = await getFeaturedReview()
  const items = await getItems()


  return (
    <>
      <ul className="flex flex-row flex-wrap gap-3 justify-center">
        {
          items && items.map((item, index) => (
            <li key={item.id} className="list-none border w-80 bg-white shadow rounded hover:shadow-xl">
              <a href={`${item.linkUrl.replace('http', 'https')}`} target="_blank">
                <Image src={`https:${item.imageUrl}/300?ver=${item.goodsCode}`} alt=""
                  width={320} height={180} className="rounded-t" />
                <h2 className="p-2 border-t text-sm font-semibold font-notosanskr">{item.goodsName}</h2>
              </a>
            </li>
          ))
        }
      </ul>
    </>
  )
}

export default HomePage