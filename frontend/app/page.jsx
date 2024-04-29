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
          <iframe src="https://ads-partners.coupang.com/widgets.html?id=775787&template=carousel&trackingCode=AF0672510&subId=&width=1600&height=140&tsource="
            width="1600" height="140" frameborder="0" scrolling="no" referrerPolicy="unsafe-url" browsingtopics>
          </iframe>
        </div>
        <div className="flex flex-row justify-center my-2">
          <div>
            <a href="https://link.coupang.com/a/bzJkAZ" target="_blank" referrerpolicy="unsafe-url">
              <Image src="https://img3a.coupangcdn.com/image/affiliate/banner/21d52affcc622aeecbb309b7cbaa7269@2x.jpg" alt="빈슨메시프 티클라 원터치 캐노피 텐트 오페라스위트 + 폴대2p, 그린, 4인용" width="240" height="240" />
            </a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/bzJlLh" target="_blank" referrerpolicy="unsafe-url">
              <Image src="https://img2a.coupangcdn.com/image/affiliate/banner/af424621250bbbe38fbb66648c8e82db@2x.jpg" alt="인스타쉐이드 원터치 그늘막 비치 텐트 원액션 캠핑 파라솔 경량 타프 그늘" width="240" height="240" />
            </a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/bzJprw" target="_blank" referrerpolicy="unsafe-url">
              <Image src="https://img1c.coupangcdn.com/image/affiliate/banner/692bc398c8b56a3eaed005d36c3377ac@2x.jpg" alt="누하스 뉴클래식 안마의자 패브릭 에디션 + 전용러그 세트 방문설치, N-0003HF, 엘더화이트" width="240" height="240" />
            </a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/bzJqb1" target="_blank" referrerpolicy="unsafe-url">
              <Image src="https://img5a.coupangcdn.com/image/affiliate/banner/2277a3bf3c8cb8f3b7e9e0a48e6123e3@2x.jpg" alt="쿠쿠 리네이처 안마의자 방문설치, 그레이 브라운 + 골드, CMS-D10SLGB" width="240" height="240" /></a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/bzJrQI" target="_blank" referrerpolicy="unsafe-url">
              <Image src="https://image4.coupangcdn.com/image/affiliate/banner/e97d47953b25164d9df13ec8dbb9d4ae@2x.jpg" alt="마샬 스탠모어 III 3세대 무선 블루투스 스피커 + 어댑터 세트, 크림 화이트" width="240" height="240" />
            </a>
          </div>
          <div>
            <a href="https://link.coupang.com/a/bzJunp" target="_blank" referrerpolicy="unsafe-url">
              <Image src="https://image5.coupangcdn.com/image/affiliate/banner/9477cdb5f784fab3b1994ef2104828bc@2x.jpg" alt="스탠리 퀜처 H2.0 플로우스테이트 텀블러, 포그, 887ml, 1개" width="240" height="240" />
            </a>
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