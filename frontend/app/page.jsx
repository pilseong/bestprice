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
        <div className="flex flex-row justify-center my-2 mb-4">
          <div>
            <a href="https://link.coupang.com/a/by1VfJ" target="_blank" referrerPolicy="unsafe-url">
              <Image src="https://image6.coupangcdn.com/image/affiliate/event/promotion/2024/04/19/c35b93b184a5002e01d2b70f4bad5332.png" width="400" height="400" />
            </a>
          </div>
        </div>
        <div className="flex flex-row justify-center my-2">
          <div>
            <a href="https://link.coupang.com/a/by1XSC" target="_blank" referrerPolicy="unsafe-url">
              <Image src="https://image8.coupangcdn.com/image/affiliate/banner/151675f781cf5fddce2e98f84e767984@2x.jpg" alt="[안테나푸드] 설악물냉면세트구성상품 10인분면+육수350g x10봉, 02, 칡냉면+육수10봉" width="240" height="240" />
            </a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/by1YcP" target="_blank" referrerPolicy="unsafe-url">
              <Image src="https://img5c.coupangcdn.com/image/affiliate/banner/cb8235c9bfb644ada75dac34f8dc1143@2x.jpg" alt="[이겼소] 국내산 냉장 마장동 소고기 옵션 육회300g+간장맛소스40g, 1세트" width="240" height="240" />
            </a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/by1YIH" target="_blank" referrerPolicy="unsafe-url">
              <Image src="https://image7.coupangcdn.com/image/affiliate/banner/9eaff3f5be59436c36b11ed7cd19baf1@2x.jpg" alt="농협안심한우 정육 세절 1등급 국거리용 (냉장), 300g, 1팩" width="240" height="240" />
            </a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/by1YZQ" target="_blank" referrerPolicy="unsafe-url">
              <Image src="https://img4c.coupangcdn.com/image/affiliate/banner/ebd25ebf1c8a28aa51b708767f1442c4@2x.jpg" alt="사조참치 살코기 안심따개, 100g, 4개" width="240" height="240" />
            </a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/by1ZrE" target="_blank" referrerPolicy="unsafe-url">
              <Image src="https://image12.coupangcdn.com/image/affiliate/banner/9acd642961df64cca1819b8157f81221@2x.jpg" alt="오뚜기 뿌셔뿌셔 불고기맛, 90g, 4개" width="240" height="240" />
            </a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/by1Z8A" target="_blank" referrerPolicy="unsafe-url"><Image src="https://image2.coupangcdn.com/image/affiliate/banner/7e77922d4315539a17442192451a2903@2x.jpg" alt="립톤 아이스티 분말 복숭아맛, 1.5kg, 1개입, 1개" width="240" height="240" /></a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/by1ZKO" target="_blank" referrerPolicy="unsafe-url">
              <Image src="https://img3c.coupangcdn.com/image/affiliate/banner/a1334a853bd912e024fc9045ba769612@2x.jpg" alt="마샬 엠버튼 엠버튼2 블루투스 스피커 스탠드 받침대 거치대 인테리어, 화이트" width="240" height="240" /></a>
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