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
      <div className="mb-10 text-center">
        <div className="flex flex-row justify-center my-2">
          <div>
            <a href="https://link.coupang.com/a/bw5sZT" target="_blank" referrerPolicy="unsafe-url">
              <Image src="https://image9.coupangcdn.com/image/affiliate/banner/bd05a5ee63b542241ed92000d09c61cc@2x.jpg" alt="마샬 워번 III 블루투스 스피커, 크림 화이트 + 3.5MM 어댑터" width="240" height="240" />
            </a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/bw5OIY" target="_blank" referrerPolicy="unsafe-url">
              <Image src="https://image3.coupangcdn.com/image/affiliate/banner/2e46dc8a0e30840ddd8b13acb361500d@2x.jpg" alt="흙대파, 1kg, 5개" width="240" height="240" />
            </a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/bw5QCu" target="_blank" referrerPolicy="unsafe-url">
              <Image src="https://img1a.coupangcdn.com/image/affiliate/banner/acd5444eb0300ea48bb7155a8ab85389@2x.jpg" alt="기타 [하남쭈꾸미]쭈꾸미볶음 500g 3팩, 3개, 350g" width="240" height="240" /></a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/bw5Uqv" target="_blank" referrerPolicy="unsafe-url">
              <Image src="https://image2.coupangcdn.com/image/affiliate/banner/e687b48f9af276ae00ea08ca72cba972@2x.jpg" alt="농심 신라면, 40개" width="240" height="240" /></a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/bw5VDz" target="_blank" referrerPolicy="unsafe-url">
              <Image src="https://img3a.coupangcdn.com/image/affiliate/banner/09241466c55536bb506a09130819e52e@2x.jpg" alt="태국라면 마마 똠냠크리미(MAMA KOONG CREAMY) 55g X 30ea X 1BOX, 30개" width="240" height="240" /></a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/bw5Xor" target="_blank" referrerPolicy="unsafe-url">
              <Image src="https://image1.coupangcdn.com/image/affiliate/banner/2f3878b427d3d0b5cfa9e4a1347fb52b@2x.jpg" alt="롤앤롤 라벤더 화장지 30롤 3겹 고급롤화장지(3겹이상) 30m, 30개입, 2개" width="240" height="240" /></a>
          </div>
        </div>
        <h2 className="text-xs sm:text-base">위의 포스팅은 쿠팡 파트너스 활동의 일환으로, 이에 따른 일정액의 수수료를 제공받습니다.</h2>
      </div>
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