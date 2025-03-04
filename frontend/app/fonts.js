import { Exo_2, Noto_Sans_KR, Orbitron, Inter } from 'next/font/google'

export const orbitron = Orbitron({
  subsets: ['latin'],
  variable: '--font-orbitron'
})

export const notosanskr = Noto_Sans_KR({
  subsets: ['latin'],
  variable: '--font-notosanskr'
})

export const exo2 = Exo_2({
  subsets: ['latin'],
  variable: '--font-exo2'
})

export const inter = Inter({
  subsets: ['latin'],
  variable: '--font-inter'
})