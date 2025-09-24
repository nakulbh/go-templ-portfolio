/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./ui/**/*.{html,js,templ}",
    "./**/*.templ"
  ],
  theme: {
    extend: {
      colors: {
        obsidian: '#1b1c1d',
      },
      fontFamily: {
        sans: ["Gilda Display", "sans-serif"],
        mono: ["JetBrains Mono", "monospace"],
      },
    },
  },
  plugins: [],
}