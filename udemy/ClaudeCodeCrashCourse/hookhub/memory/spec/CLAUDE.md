# HookHub Specification

## Project Overview

**HookHub** is a web application for discovering, browsing, and exploring open source Claude Code hooks. It serves as a centralized gallery where developers can find useful hooks to enhance their Claude Code workflows.

### What are Claude Code Hooks?

Claude Code hooks are user-defined shell commands that execute at various points in Claude Code's lifecycle. They provide deterministic control over Claude Code's behavior, enabling automation, validation, notifications, and custom workflows.

### Hook Lifecycle Events

Claude Code provides 8 hook events:

1. **PreToolUse** - Runs before tool calls (can block them)
2. **PermissionRequest** - Runs when a permission dialog is shown
3. **PostToolUse** - Runs after successful tool execution
4. **ToolError** - Runs when a tool fails
5. **SessionStart** - Runs when a new session begins
6. **SessionEnd** - Runs when a session ends
7. **MessageSent** - Runs when user sends a message
8. **ResponseReceived** - Runs when Claude responds

### Common Hook Use Cases

- **Formatting**: Auto-format code after edits (prettier, gofmt, etc.)
- **Notifications**: Alert users when Claude needs input or permission
- **Logging**: Track and audit all executed commands
- **Validation**: Enforce code quality standards and conventions
- **Security**: Block modifications to sensitive files or directories
- **Integration**: Connect Claude Code to external tools and services

## MVP Goals

For the initial release, HookHub will:

- Display a curated collection of open source Claude Code hooks
- Present hooks in an organized, browsable grid layout
- Provide essential information about each hook (name, category, description)
- Link directly to GitHub repositories for installation
- Support basic categorization for easier discovery

**Out of Scope for MVP:**

- User authentication and accounts
- User-submitted hooks
- Ratings, reviews, or comments
- Advanced search and filtering
- Hook installation/download functionality
- Analytics or usage tracking

## Data Model

### Hook Object Structure

```typescript
interface Hook {
  id: string; // Unique identifier (e.g., "auto-prettier")
  name: string; // Display name (e.g., "Auto Prettier")
  description: string; // Brief description of what the hook does
  category: HookCategory; // Primary category
  repoUrl: string; // GitHub repository URL
  author?: string; // GitHub username or organization
  tags?: string[]; // Additional keywords for future search
  hookEvent: HookEvent; // Which lifecycle event it uses
  createdAt?: string; // When added to HookHub (ISO 8601)
}

type HookEvent =
  | "PreToolUse"
  | "PermissionRequest"
  | "PostToolUse"
  | "ToolError"
  | "SessionStart"
  | "SessionEnd"
  | "MessageSent"
  | "ResponseReceived";

type HookCategory =
  | "Formatting"
  | "Notifications"
  | "Logging"
  | "Validation"
  | "Security"
  | "Integration"
  | "Productivity"
  | "Other";
```

### Data Source

**MVP Approach**: Static JSON file in codebase

- File location: `app/data/hooks.json` or `lib/data/hooks.ts`
- Manually curated list of quality hooks
- No external API calls or database
- Simple, fast, no rate limits

**Future**: Could migrate to GitHub API, Airtable, or custom database

## Feature Requirements

### 1. Home Page - Hook Gallery

**Route**: `/` (app/page.tsx)

**Layout**: Responsive grid of hook cards

**Grid Specifications**:

- Desktop: 3 columns
- Tablet: 2 columns
- Mobile: 1 column
- Gap: 24px between cards
- Padding: Container with responsive padding

**Hook Card Components**:
Each card displays:

- Hook name (heading)
- Category badge (colored, top-right or top-left)
- Hook event badge (smaller badge)
- Description (2-3 lines, truncated with ellipsis)
- GitHub icon/link
- Author/organization (if available)

**Card Interactions**:

- Hover state: Subtle elevation/shadow increase
- Click anywhere on card: Opens GitHub repo in new tab
- Focus state: Visible outline for keyboard navigation

### 2. Category System

**Category Badges**:

- Color-coded by category type
- Small, pill-shaped design
- Positioned consistently on each card

**Category Colors** (Tailwind classes):

- Formatting: Blue (`bg-blue-100 text-blue-800`)
- Notifications: Yellow (`bg-yellow-100 text-yellow-800`)
- Logging: Gray (`bg-gray-100 text-gray-800`)
- Validation: Green (`bg-green-100 text-green-800`)
- Security: Red (`bg-red-100 text-red-800`)
- Integration: Purple (`bg-purple-100 text-purple-800`)
- Productivity: Indigo (`bg-indigo-100 text-indigo-800`)
- Other: Neutral (`bg-neutral-100 text-neutral-800`)

### 3. Header/Navigation

**Components**:

- Site logo/title: "HookHub"
- Tagline: "Discover Claude Code Hooks"
- Optional: Link to Claude Code documentation
- Optional: GitHub icon linking to HookHub repo

**Design**:

- Clean, minimal header
- Sticky/fixed position optional for MVP
- Responsive typography

### 4. Footer (Optional for MVP)

**Simple footer with**:

- Brief description
- Links to resources (Claude Code docs, GitHub)
- Built with attribution
- Copyright/license info

## Technical Architecture

### Tech Stack

- **Framework**: Next.js 16.1.1 (App Router)
- **React**: 19.2.3
- **TypeScript**: 5.x (strict mode)
- **Styling**: Tailwind CSS v4
- **Fonts**: Geist Sans, Geist Mono

### File Structure

```
hookhub/
├── app/
│   ├── page.tsx                 # Home page (hook gallery)
│   ├── layout.tsx               # Root layout (existing)
│   └── globals.css              # Global styles (existing)
├── components/
│   ├── HookCard.tsx             # Individual hook card component
│   ├── HookGrid.tsx             # Grid container for hooks
│   ├── CategoryBadge.tsx        # Category badge component
│   ├── Header.tsx               # Site header
│   └── Footer.tsx               # Site footer (optional)
├── lib/
│   ├── data/
│   │   └── hooks.ts             # Static hooks data
│   └── types/
│       └── hook.ts              # TypeScript interfaces
├── public/                      # Static assets
└── spec/
    └── CLAUDE.md                # This specification
```

### Component Architecture

**Server Components** (default):

- `app/page.tsx` - Main page, fetches hooks data
- `HookGrid.tsx` - Container component
- `Header.tsx` - Static header
- `Footer.tsx` - Static footer

**Client Components** (if needed):

- `HookCard.tsx` - May need "use client" for hover interactions
- `CategoryBadge.tsx` - Likely server component

### Data Flow

```
hooks.ts (static data)
    ↓
app/page.tsx (Server Component)
    ↓
HookGrid component
    ↓
HookCard components (mapped)
```

## UI/UX Specifications

### Design Principles

- **Clean & Minimal**: Focus on content, not decoration
- **Scannable**: Easy to quickly browse and compare hooks
- **Consistent**: Uniform card design and spacing
- **Accessible**: Proper semantic HTML, ARIA labels, keyboard navigation
- **Responsive**: Mobile-first approach

### Visual Design

**Color Scheme**:

- Use existing Tailwind v4 theme variables
- Support dark mode via `prefers-color-scheme`
- Background: `var(--background)`
- Foreground: `var(--foreground)`
- Cards: White/dark surface with subtle borders

**Typography**:

- Headers: Geist Sans (var(--font-geist-sans))
- Body: Geist Sans
- Code references: Geist Mono (if needed)
- Hook name: text-lg or text-xl, font-semibold
- Description: text-sm or text-base
- Category badge: text-xs, font-medium

**Spacing**:

- Container padding: px-4 md:px-8 lg:px-16
- Vertical spacing: py-8 md:py-12 lg:py-16
- Card padding: p-6
- Grid gap: gap-6

**Cards**:

- Border: border border-gray-200/border-gray-800
- Rounded corners: rounded-lg
- Hover: shadow-md transition
- Background: white/dark surface

### Responsive Breakpoints

- Mobile: < 768px (1 column)
- Tablet: 768px - 1024px (2 columns)
- Desktop: > 1024px (3 columns)

### Accessibility Requirements

- Semantic HTML (`<article>` for cards, `<nav>` for header)
- Proper heading hierarchy (h1 → h2 → h3)
- Alt text for all images/icons
- Keyboard navigation support
- Focus indicators visible
- ARIA labels where needed
- Color contrast ratio ≥ 4.5:1

## Sample Data

Initial curated hooks to include (examples):

```typescript
const sampleHooks: Hook[] = [
  {
    id: "auto-prettier",
    name: "Auto Prettier",
    description:
      "Automatically formats TypeScript and JavaScript files with Prettier after every file edit.",
    category: "Formatting",
    hookEvent: "PostToolUse",
    repoUrl: "https://github.com/example/auto-prettier-hook",
    author: "example",
    tags: ["prettier", "formatting", "javascript", "typescript"],
  },
  {
    id: "commit-validator",
    name: "Commit Message Validator",
    description:
      "Validates commit messages follow conventional commits format before allowing git commits.",
    category: "Validation",
    hookEvent: "PreToolUse",
    repoUrl: "https://github.com/example/commit-validator",
    author: "example",
    tags: ["git", "commits", "validation"],
  },
  // Add 8-12 more quality hooks
];
```

## Implementation Phases

### Phase 1: Core Structure ✓

- [x] Set up Next.js project
- [x] Configure TypeScript and Tailwind
- [x] Create project documentation

### Phase 2: Data Layer

- [ ] Define TypeScript interfaces (`lib/types/hook.ts`)
- [ ] Create static hooks data file (`lib/data/hooks.ts`)
- [ ] Populate with 10-15 curated hooks

### Phase 3: Components

- [ ] Create `HookCard` component
- [ ] Create `CategoryBadge` component
- [ ] Create `HookGrid` component
- [ ] Create `Header` component
- [ ] Create `Footer` component (optional)

### Phase 4: Main Page

- [ ] Update `app/page.tsx` with hook gallery
- [ ] Import and display hooks data
- [ ] Implement responsive grid layout
- [ ] Test responsive behavior

### Phase 5: Styling & Polish

- [ ] Refine card design and interactions
- [ ] Add hover/focus states
- [ ] Ensure dark mode support
- [ ] Optimize typography and spacing
- [ ] Test accessibility

### Phase 6: Testing & Launch

- [ ] Cross-browser testing
- [ ] Mobile device testing
- [ ] Accessibility audit
- [ ] Performance check
- [ ] Deploy to production

## Future Enhancements

**Post-MVP Features** (not for initial release):

- Search functionality (by name, description, tags)
- Filter by category
- Filter by hook event type
- Sort options (newest, popular, alphabetical)
- Hook detail pages with installation instructions
- Copy-paste hook configuration
- User submissions
- Ratings and reviews
- Related hooks suggestions
- Integration with Claude Code API (if available)
- Analytics dashboard
- RSS feed for new hooks

## Success Metrics

**MVP Success Criteria**:

- Displays minimum 10 curated hooks
- Fully responsive on mobile, tablet, desktop
- All external links work correctly
- Passes accessibility audit (WCAG 2.1 AA)
- Page load time < 2 seconds
- Works in Chrome, Firefox, Safari, Edge

## Resources & References

- [Claude Code Hooks Documentation](https://code.claude.com/docs/en/hooks-guide)
- [Claude Code Hooks Reference](https://docs.claude.com/en/docs/claude-code/hooks)
- [GitHub - Claude Code Hooks Mastery](https://github.com/disler/claude-code-hooks-mastery)
- [Complete Guide to Hooks in Claude Code](https://www.eesel.ai/blog/hooks-in-claude-code)
- [Next.js 15 Documentation](https://nextjs.org/docs)
- [Tailwind CSS v4 Documentation](https://tailwindcss.com/docs)

---

**Document Version**: 1.0
**Last Updated**: 2026-01-12
**Author**: Claude Code
**Project**: HookHub MVP
