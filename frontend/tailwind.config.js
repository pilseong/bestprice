/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        'trust': 'rgb(var(--trust))',
        'gmarket': '#0abb26'
      },
      fontFamily: {
        // sans: ["var(--font-exo2)", "sans-serif"],
        orbitron: ["var(--font-orbitron)", "sans-serif"],
        notosanskr: ["var(--font-notosanskr)", "sans-serif"],
        exo2: ["var(--font-exo2)", "sans-serif"]
      },
    },
  },
  plugins: [],
};