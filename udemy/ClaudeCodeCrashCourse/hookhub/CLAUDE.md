# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

HookHub is a Next.js 16.1.1 application using React 19, TypeScript, and Tailwind CSS v4. This is a modern web application bootstrapped with `create-next-app` using the App Router architecture.

## Development Commands

### Start development server
```bash
npm run dev
```
Opens at http://localhost:3000 with hot reload enabled.

### Build for production
```bash
npm run build
```

### Start production server
```bash
npm start
```

### Lint code
```bash
npm run lint
```
Uses ESLint with Next.js configuration for both TypeScript and core web vitals rules.

## Architecture

### Next.js App Router Structure
- Uses Next.js 16 App Router (not Pages Router)
- Entry point: `app/page.tsx` (home page)
- Root layout: `app/layout.tsx` (wraps all pages)
- Global styles: `app/globals.css`

### Styling System
- Tailwind CSS v4 with PostCSS
- Uses `@theme inline` directive for CSS custom properties
- Geist Sans and Geist Mono fonts from Google Fonts with CSS variables:
  - `--font-geist-sans`
  - `--font-geist-mono`
- Theme variables for dark mode support:
  - `--background` and `--foreground` switch based on `prefers-color-scheme`

### TypeScript Configuration
- Target: ES2017
- Module resolution: bundler (Next.js specific)
- Strict mode enabled
- Path alias: `@/*` maps to project root
- JSX mode: `react-jsx` (React 19)

### Key Configuration Files
- `next.config.ts`: Next.js configuration (currently minimal)
- `tsconfig.json`: TypeScript compiler options with Next.js plugin
- `eslint.config.mjs`: ESLint configuration using flat config format
- `postcss.config.mjs`: PostCSS with Tailwind plugin
- Ignores: `.next/`, `out/`, `build/`, `next-env.d.ts`

## File Organization Notes

When adding new features:
- Server Components are the default in App Router
- Client Components require `"use client"` directive
- Place reusable components in a `components/` directory (not yet created)
- API routes go in `app/api/` directory (not yet created)
- Shared utilities can go in `lib/` or `utils/` directory (not yet created)

## Important Conventions

- This project uses the Next.js App Router (app directory), not the Pages Router
- All components in `app/` are Server Components by default unless marked with `"use client"`
- The `layout.tsx` file provides the HTML structure and wraps all pages
- Static assets go in the `public/` directory and are served from the root path
